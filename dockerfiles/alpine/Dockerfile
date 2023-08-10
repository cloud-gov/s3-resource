<<<<<<< HEAD
ARG base_image
ARG builder_image=concourse/golang-builder

FROM ${builder_image} as builder
COPY . /go/src/github.com/concourse/s3-resource
<<<<<<< HEAD
=======
=======
FROM concourse/golang-builder as builder
COPY . /go/src/github.com/concourse/s3-resource
WORKDIR /go/src/github.com/concourse/s3-resource
>>>>>>> ad6dcd4 (match upstream route to update to 1.16 instead of our own)
ENV CGO_ENABLED 0
RUN go mod download
RUN go build -o /assets/in github.com/concourse/s3-resource/cmd/in
RUN go build -o /assets/out github.com/concourse/s3-resource/cmd/out
RUN go build -o /assets/check github.com/concourse/s3-resource/cmd/check
>>>>>>> 0ef6a2b (make dockerfiles work with go 1.16, stop building outside of Docker)
WORKDIR /go/src/github.com/concourse/s3-resource
ENV CGO_ENABLED 0
RUN go mod download
RUN go build -o /assets/in ./cmd/in
RUN go build -o /assets/out ./cmd/out
RUN go build -o /assets/check ./cmd/check
RUN set -e; for pkg in $(go list ./...); do \
		go test -o "/tests/$(basename $pkg).test" -c $pkg; \
	done

FROM ${base_image} AS resource
USER root
RUN apt update && apt upgrade -y -o Dpkg::Options::="--force-confdef"
RUN apt update \
      && apt install -y --no-install-recommends \
        tzdata \
        ca-certificates \
        unzip \
        zip \
      && rm -rf /var/lib/apt/lists/*
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource AS tests
ARG S3_TESTING_ACCESS_KEY_ID
ARG S3_TESTING_SECRET_ACCESS_KEY
ARG S3_TESTING_SESSION_TOKEN
ARG S3_TESTING_AWS_ROLE_ARN
ARG S3_VERSIONED_TESTING_BUCKET
ARG S3_TESTING_BUCKET
ARG S3_TESTING_REGION
ARG S3_ENDPOINT
ARG TEST_SESSION_TOKEN
ARG S3_USE_INSTANCE_PROFILE
COPY --from=builder /tests /go-tests
WORKDIR /go-tests
RUN set -e; for test in /go-tests/*.test; do \
		$test; \
	done

FROM resource
