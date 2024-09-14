package s3

import types "libds/aws/types"

type ObjectCopy struct {
	// Specifies server-side encryption of the object in S3. Valid values are `AES256` and `aws:kms`.
	ServerSideEncryption string `json:"serverSideEncryption,omitempty" yaml:"serverSideEncryption,omitempty"`

	// Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	CustomerKeyMd5 string `json:"customerKeyMd5,omitempty" yaml:"customerKeyMd5,omitempty"`

	// Specifies the AWS KMS Key ARN to use for object encryption. This value is a fully qualified --ARN-- of the KMS Key. If using `aws.kms.Key`, use the exported `arn` attribute: `kms_key_id = aws_kms_key.foo.arn`
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Copies the object if it has been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CopyIfModifiedSince string `json:"copyIfModifiedSince,omitempty" yaml:"copyIfModifiedSince,omitempty"`

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	CustomerAlgorithm string `json:"customerAlgorithm,omitempty" yaml:"customerAlgorithm,omitempty"`

	// The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
	ObjectLockLegalHoldStatus string `json:"objectLockLegalHoldStatus,omitempty" yaml:"objectLockLegalHoldStatus,omitempty"`

	// Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Valid values are `COPY` and `REPLACE`.
	TaggingDirective string `json:"taggingDirective,omitempty" yaml:"taggingDirective,omitempty"`

	//
	BucketKeyEnabled bool `json:"bucketKeyEnabled,omitempty" yaml:"bucketKeyEnabled,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding string `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`

	// Account id of the expected destination bucket owner. If the destination bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Date and time at which the object is no longer cacheable, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	Expires string `json:"expires,omitempty" yaml:"expires,omitempty"`

	// Specifies the AWS KMS Encryption Context to use for object encryption. The value is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
	KmsEncryptionContext string `json:"kmsEncryptionContext,omitempty" yaml:"kmsEncryptionContext,omitempty"`

	/*
	   Specifies the source object for the copy operation. You specify the value in one of two formats. For objects not accessed through an access point, specify the name of the source bucket and the key of the source object, separated by a slash (`/`). For example, `testbucket/test1.json`. For objects accessed through access points, specify the ARN of the object as accessed through the access point, in the format `arn:aws:s3:<Region>:<account-id>:accesspoint/<access-point-name>/object/<key>`. For example, `arn:aws:s3:us-west-2:9999912999:accesspoint/my-access-point/object/testbucket/test1.json`.

	   The following arguments are optional:
	*/
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// Specifies the algorithm to use when decrypting the source object (for example, AES256).
	SourceCustomerAlgorithm string `json:"sourceCustomerAlgorithm,omitempty" yaml:"sourceCustomerAlgorithm,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	SourceCustomerKeyMd5 string `json:"sourceCustomerKeyMd5,omitempty" yaml:"sourceCustomerKeyMd5,omitempty"`

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	CopyIfNoneMatch string `json:"copyIfNoneMatch,omitempty" yaml:"copyIfNoneMatch,omitempty"`

	// Copies the object if it hasn't been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CopyIfUnmodifiedSince string `json:"copyIfUnmodifiedSince,omitempty" yaml:"copyIfUnmodifiedSince,omitempty"`

	// Specifies the desired [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html#AmazonS3-CopyObject-request-header-StorageClass) for the object. Defaults to `STANDARD`.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// Standard MIME type describing the format of the object data, e.g., `application/octet-stream`. All Valid MIME Types are valid for this input.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object. The encryption key provided in this header must be one that was used when the source object was created.
	SourceCustomerKey string `json:"sourceCustomerKey,omitempty" yaml:"sourceCustomerKey,omitempty"`

	// Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition string `json:"contentDisposition,omitempty" yaml:"contentDisposition,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	ContentLanguage string `json:"contentLanguage,omitempty" yaml:"contentLanguage,omitempty"`

	// Indicates the algorithm used to create the checksum for the object. If a value is specified and the object is encrypted with KMS, you must have permission to use the `kms:Decrypt` action. Valid values: `CRC32`, `CRC32C`, `SHA1`, `SHA256`.
	ChecksumAlgorithm string `json:"checksumAlgorithm,omitempty" yaml:"checksumAlgorithm,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Allow the object to be deleted by removing any legal hold on any object version. Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// Configuration block for header grants. Documented below. Conflicts with `acl`.
	Grants []types.S3_ObjectCopyGrant `json:"grants,omitempty" yaml:"grants,omitempty"`

	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. For information about downloading objects from requester pays buckets, see Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html) in the Amazon S3 Developer Guide. If included, the only valid value is `requester`.
	RequestPayer string `json:"requestPayer,omitempty" yaml:"requestPayer,omitempty"`

	// [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `authenticated-read`, `aws-exec-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Conflicts with `grant`.
	Acl string `json:"acl,omitempty" yaml:"acl,omitempty"`

	// Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl string `json:"cacheControl,omitempty" yaml:"cacheControl,omitempty"`

	// Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
	ObjectLockRetainUntilDate string `json:"objectLockRetainUntilDate,omitempty" yaml:"objectLockRetainUntilDate,omitempty"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data. This value is used to store the object and then it is discarded; Amazon S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the x-amz-server-side-encryption-customer-algorithm header.
	CustomerKey string `json:"customerKey,omitempty" yaml:"customerKey,omitempty"`

	// Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Valid values are `COPY` and `REPLACE`.
	MetadataDirective string `json:"metadataDirective,omitempty" yaml:"metadataDirective,omitempty"`

	// Account id of the expected source bucket owner. If the source bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedSourceBucketOwner string `json:"expectedSourceBucketOwner,omitempty" yaml:"expectedSourceBucketOwner,omitempty"`

	// Name of the object once it is in the bucket.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	ObjectLockMode string `json:"objectLockMode,omitempty" yaml:"objectLockMode,omitempty"`

	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect string `json:"websiteRedirect,omitempty" yaml:"websiteRedirect,omitempty"`

	// Name of the bucket to put the file in.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopyIfMatch string `json:"copyIfMatch,omitempty" yaml:"copyIfMatch,omitempty"`
}
