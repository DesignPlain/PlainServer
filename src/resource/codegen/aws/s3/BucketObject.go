package s3

type BucketObject struct {
	// Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Triggers updates when the value changes. This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"` (see `source_hash` instead).
	Etag string `json:"etag,omitempty" yaml:"etag,omitempty"`

	// Whether to allow the object to be deleted by removing any legal hold on any object version. Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// ARN of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the `aws.kms.Key` resource, use the `arn` attribute. If referencing the `aws.kms.Alias` data source or resource, use the `target_key_arn` attribute. The provider will only perform drift detection if a configuration value is provided.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// [Legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
	ObjectLockLegalHoldStatus string `json:"objectLockLegalHoldStatus,omitempty" yaml:"objectLockLegalHoldStatus,omitempty"`

	// Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 string `json:"contentBase64,omitempty" yaml:"contentBase64,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	ContentLanguage string `json:"contentLanguage,omitempty" yaml:"contentLanguage,omitempty"`

	// Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
	ObjectLockRetainUntilDate string `json:"objectLockRetainUntilDate,omitempty" yaml:"objectLockRetainUntilDate,omitempty"`

	// Triggers updates like `etag` but useful to address `etag` encryption limitations.
	SourceHash string `json:"sourceHash,omitempty" yaml:"sourceHash,omitempty"`

	// [Storage Class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html#AmazonS3-PutObject-request-header-StorageClass) for the object. Defaults to "`STANDARD`".
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Defaults to `private`.
	Acl string `json:"acl,omitempty" yaml:"acl,omitempty"`

	// Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
	BucketKeyEnabled bool `json:"bucketKeyEnabled,omitempty" yaml:"bucketKeyEnabled,omitempty"`

	// Caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl string `json:"cacheControl,omitempty" yaml:"cacheControl,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// Presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition string `json:"contentDisposition,omitempty" yaml:"contentDisposition,omitempty"`

	// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	ObjectLockMode string `json:"objectLockMode,omitempty" yaml:"objectLockMode,omitempty"`

	// Path to a file that will be read and uploaded as raw bytes for the object content.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   Target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).

	   If no content is provided through `source`, `content` or `content_base64`, then the object will be empty.
	*/
	WebsiteRedirect string `json:"websiteRedirect,omitempty" yaml:"websiteRedirect,omitempty"`

	// Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding string `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`

	/*
	   Name of the object once it is in the bucket.

	   The following arguments are optional:
	*/
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
	ServerSideEncryption string `json:"serverSideEncryption,omitempty" yaml:"serverSideEncryption,omitempty"`
}
