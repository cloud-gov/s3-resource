# S3 Resource

Versions objects in an S3 bucket, by pattern-matching filenames to identify
version numbers.

## Source Configuration

* `bucket`: *Required.* The name of the bucket.

* `access_key_id`: *Optional.* The AWS access key to use when accessing the
  bucket. If empty, the resource will try to retrieve credentials from
  [AWS instance profile](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html)

* `secret_access_key`: *Optional.* The AWS secret key to use when accessing
  the bucket. If empty, the resource will try to retrieve credentials from
  [AWS instance profile](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html)

* `session_token`: *Optional.* The AWS STS session token to use when
  accessing the bucket.

* `aws_role_arn`: *Optional.* The AWS role ARN to be assumed by the user
  identified by `access_key_id` and `secret_access_key`.

* `region_name`: *Optional.* The region the bucket is in. Defaults to
  `us-east-1`.

* `private`: *Optional.* Indicates that the bucket is private, so that any
  URLs provided are signed.

* `cloudfront_url`: *Optional.* The URL (scheme and domain) of your CloudFront
  distribution that is fronting this bucket (e.g
  `https://d5yxxxxx.cloudfront.net`).  This will affect `in` but not `check`
  and `put`. `in` will ignore the `bucket` name setting, exclusively using the
  `cloudfront_url`.  When configuring CloudFront with versioned buckets, set
  `Query String Forwarding and Caching` to `Forward all, cache based on all` to
  ensure S3 calls succeed.

* `endpoint`: *Optional.* Custom endpoint for using S3 compatible provider.

* `disable_ssl`: *Optional.* Disable SSL for the endpoint, useful for S3
  compatible providers without SSL.

* `skip_ssl_verification`: *Optional.* Skip SSL verification for S3 endpoint. Useful for S3 compatible providers using self-signed SSL certificates.

* `skip_download`: *Optional.* Skip downloading object from S3. Useful only trigger the pipeline without using the object.

* `server_side_encryption`: *Optional.* An encryption algorithm to use when
  storing objects in S3.

* `sse_kms_key_id`: *Optional.* The ID of the AWS KMS master encryption key
  used for the object.

* `use_v2_signing`: *Optional.* Use signature v2 signing, useful for S3 compatible providers that do not support v4.

* `disable_multipart`: *Optional.* Disable Multipart Upload. useful for S3 compatible providers that do not support multipart upload.

### File Names

One of the following two options must be specified:

* `regexp`: *Optional.* The forward-slash (`/`) delimited sequence of patterns to
  match against the sub-directories and filenames of the objects stored within
  the S3 bucket. The first grouped match is used to extract the version, or if
  a group is explicitly named `version`, that group is used. At least one
  capture group must be specified, with parentheses.

  The version extracted from this pattern is used to version the resource.
  Semantic versions, or just numbers, are supported. Accordingly, full regular
  expressions are supported, to specify the capture groups.

  The full `regexp` will be matched against the S3 objects as if it was anchored
  on both ends, even if you don't specify `^` and `$` explicitly.

* `versioned_file`: *Optional* If you enable versioning for your S3 bucket then
  you can keep the file name the same and upload new versions of your file
  without resorting to version numbers. This property is the path to the file
  in your S3 bucket.

### Initial state

If no resource versions exist you can set up this resource to emit an initial version with a specified content. This won't create a real resource in S3 but only create an initial version for Concourse. The resource file will be created as usual when you `get` a resource with an initial version.

You can define one of the following two options:

* `initial_path`: *Optional.* Must be used with the `regexp` option. You should set this to the file path containing the initial version which would match the given regexp. E.g. if `regexp` is `file/build-(.*).zip`, then `initial_path` might be `file/build-0.0.0.zip`. The resource version will be `0.0.0` in this case.

* `initial_version`: *Optional.* Must be used with the `versioned_file` option. This will be the resource version.

By default the resource file will be created with no content when `get` runs. You can set the content by using one of the following options:

* `initial_content_text`: *Optional.* Initial content as a string.

* `initial_content_binary`: *Optional.* You can pass binary content as a base64 encoded string.

## Behavior

### `check`: Extract versions from the bucket.

Objects will be found via the pattern configured by `regexp`. The versions
will be used to order them (using [semver](http://semver.org/)). Each
object's filename is the resulting version.


### `in`: Fetch an object from the bucket.

Places the following files in the destination:

* `(filename)`: The file fetched from the bucket.

* `url`: A file containing the URL of the object. If `private` is true, this
  URL will be signed.

* `version`: The version identified in the file name.

* `tags.json`: The object's tags represented as a JSON object. Only written if `download_tags` is set to true.

#### Parameters

* `skip_download`: *Optional.* Skip downloading object from S3. Same parameter as source configuration but used to define/override by get. Value needs to be a true/false string.

* `unpack`: *Optional.* If true and the file is an archive (tar, gzipped tar, other gzipped file, or zip), unpack the file. Gzipped tarballs will be both ungzipped and untarred. It is ignored when `get` is running on the initial version.

* `download_tags`: *Optional.* Write object tags to `tags.json`. Value needs to be a true/false string.

### `out`: Upload an object to the bucket.

Given a file specified by `file`, upload it to the S3 bucket. If `regexp` is
specified, the new file will be uploaded to the directory that the regex
searches in. If `versioned_file` is specified, the new file will be uploaded as
a new version of that file.

#### Parameters

* `file`: *Required.* Path to the file to upload, provided by an output of a task.
  If multiple files are matched by the glob, an error is raised. The file which
  matches will be placed into the directory structure on S3 as defined in `regexp`
  in the resource definition. The matching syntax is bash glob expansion, so
  no capture groups, etc.

* `acl`: *Optional.*  [Canned Acl](http://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
  for the uploaded object.

* `content_type`: *Optional.* MIME [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17)
  describing the contents of the uploaded object

## Example Configuration

### Resource

When the file has the version name in the filename

``` yaml
- name: release
  type: s3
  source:
    bucket: releases
    regexp: directory_on_s3/release-(.*).tgz
    access_key_id: ACCESS-KEY
    secret_access_key: SECRET
```

or

When the file is being [versioned by s3](http://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html)

``` yaml
- name: release
  type: s3
  source:
    bucket: releases
    versioned_file: directory_on_s3/release.tgz
    access_key_id: ACCESS-KEY
    secret_access_key: SECRET
```

### Plan

``` yaml
- get: release
```

``` yaml
- put: release
  params:
    file: path/to/release-*.tgz
    acl: public-read
```

## Required IAM Permissions

### Non-versioned Buckets

The bucket itself (e.g. `"arn:aws:s3:::your-bucket"`):
* `s3:ListBucket`

The objects in the bucket (e.g. `"arn:aws:s3:::your-bucket/*"`):
* `s3:PutObject`
* `s3:PutObjectAcl`
* `s3:GetObject`
* `s3:GetObjectTagging` (if using the `download_tags` option)

### Versioned Buckets

Everything above and...

The bucket itself (e.g. `"arn:aws:s3:::your-bucket"`):
* `s3:ListBucketVersions`
* `s3:GetBucketVersioning`

The objects in the bucket (e.g. `"arn:aws:s3:::your-bucket/*"`):
* `s3:GetObjectVersion`
* `s3:PutObjectVersionAcl`
* `s3:GetObjectVersionTagging` (if using the `download_tags` option)

## Development

### Prerequisites

* Go is *required* - version 1.13 is tested; earlier versions may also
  work.
* docker is *required* - version 17.06.x is tested; earlier versions may also
  work.

### Running the tests

The tests have been embedded with the `Dockerfile`; ensuring that the testing
environment is consistent across any `docker` enabled platform. When the docker
image builds, the test are run inside the docker container, on failure they
will stop the build.

Run the tests with the following commands for both `alpine` and `ubuntu` images:

```sh
<<<<<<< HEAD
docker build -t s3-resource --target tests --build-arg base_image=paketobuildpacks/run-jammy-base:latest .
 .
=======
docker build -t s3-resource -f dockerfiles/alpine/Dockerfile .
docker build -t s3-resource -f dockerfiles/ubuntu/Dockerfile .
>>>>>>> 01d768b (Catchup (#3))
```

#### Integration tests

The integration requires two AWS S3 buckets, one without versioning and another
with. The `docker build` step requires setting `--build-args` so the
integration will run.

Run the tests with the following command:

```sh
<<<<<<< HEAD
docker build . -t s3-resource --target tests \
=======
docker build . -t s3-resource -f dockerfiles/alpine/Dockerfile \
  --build-arg S3_TESTING_ACCESS_KEY_ID="access-key" \
  --build-arg S3_TESTING_SECRET_ACCESS_KEY="some-secret" \
  --build-arg S3_TESTING_BUCKET="bucket-non-versioned" \
  --build-arg S3_VERSIONED_TESTING_BUCKET="bucket-versioned" \
  --build-arg S3_TESTING_REGION="us-east-1" \
  --build-arg S3_ENDPOINT="https://s3.amazonaws.com"

docker build . -t s3-resource -f dockerfiles/ubuntu/Dockerfile \
>>>>>>> 01d768b (Catchup (#3))
  --build-arg S3_TESTING_ACCESS_KEY_ID="access-key" \
  --build-arg S3_TESTING_SECRET_ACCESS_KEY="some-secret" \
  --build-arg S3_TESTING_BUCKET="bucket-non-versioned" \
  --build-arg S3_VERSIONED_TESTING_BUCKET="bucket-versioned" \
  --build-arg S3_TESTING_REGION="us-east-1" \
  --build-arg S3_ENDPOINT="https://s3.amazonaws.com"
```

##### Speeding up integration tests by skipping large file upload

One of the integration tests uploads a large file (>40GB) and so can be slow.
It can be skipped by adding the following option when running the tests:
```
  --build-arg S3_TESTING_NO_LARGE_UPLOAD=true
```

##### Integration tests using role assumption

If `S3_TESTING_AWS_ROLE_ARN` is set to a role ARN, this role will be assumed for accessing
the S3 bucket during integration tests. The whole integration test suite runs either
completely using role assumption or completely by direct access via the credentials.

##### Required IAM permissions

In addition to the required permissions above, the `s3:PutObjectTagging` permission is required to run integration tests.

### Contributing

Please make all pull requests to the `master` branch and ensure tests pass
locally.
