# syntax=docker/dockerfile:1.4

ARG base_image=ubuntu:22.04
ARG builder_image=concourse/golang-builder

FROM ${builder_image} AS builder
COPY . /go/src/github.com/concourse/s3-resource
WORKDIR /go/src/github.com/concourse/s3-resource
ENV CGO_ENABLED=0
ENV AWS_USE_FIPS_ENDPOINT=true
RUN go mod download
RUN go build -o /assets/in github.com/concourse/s3-resource/cmd/in
RUN go build -o /assets/out github.com/concourse/s3-resource/cmd/out
RUN go build -o /assets/check github.com/concourse/s3-resource/cmd/check
RUN set -e; for pkg in $(go list ./...); do \
  go test -o "/tests/$(basename $pkg).test" -c $pkg; \
  done

FROM ${base_image} AS resource
RUN apt-get update \
  && apt-get upgrade -y -o Dpkg::Options::="--force-confdef" \
  && apt-get install -y --no-install-recommends \
  tzdata \
  ca-certificates \
  unzip \
  zip \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

RUN groupadd -r -g 1001 s3user && \
    useradd -r -u 1001 -g s3user -s /sbin/nologin -c "S3 Resource User" s3user

USER s3user
ENV AWS_USE_FIPS_ENDPOINT=true

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD test -x /opt/resource/check && test -x /opt/resource/in && test -x /opt/resource/out || exit 1

FROM resource AS tests
ARG S3_VERSIONED_TESTING_BUCKET
ARG S3_TESTING_BUCKET
ARG S3_TESTING_REGION
ARG S3_ENDPOINT
ENV AWS_USE_FIPS_ENDPOINT=true

COPY --from=builder /tests /go-tests
WORKDIR /go-tests

# Run tests with secrets mounted via BuildKit
# Secrets are exposed as files in /run/secrets/* and never stored in image layers
RUN --mount=type=secret,id=s3_access_key_id \
    --mount=type=secret,id=s3_secret_access_key \
    --mount=type=secret,id=s3_session_token,required=false \
    --mount=type=secret,id=s3_role_arn,required=false \
    --mount=type=secret,id=test_session_token,required=false \
    set -e; \
    export S3_TESTING_ACCESS_KEY_ID=$(cat /run/secrets/s3_access_key_id 2>/dev/null || echo ""); \
    export S3_TESTING_SECRET_ACCESS_KEY=$(cat /run/secrets/s3_secret_access_key 2>/dev/null || echo ""); \
    export S3_TESTING_SESSION_TOKEN=$(cat /run/secrets/s3_session_token 2>/dev/null || echo ""); \
    export S3_TESTING_AWS_ROLE_ARN=$(cat /run/secrets/s3_role_arn 2>/dev/null || echo ""); \
    export TEST_SESSION_TOKEN=$(cat /run/secrets/test_session_token 2>/dev/null || echo ""); \
    for test in /go-tests/*.test; do \
      $test; \
    done

FROM resource
ENV AWS_USE_FIPS_ENDPOINT=true
