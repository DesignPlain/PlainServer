package s3

import types "DesignSphere_Server/src/resource/aws/types"

type Bucket struct {
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`

	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a `pulumi preview`. In this case, please make sure you use the verbose/specific version of the policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration types.S3_BucketReplicationConfiguration `json:"replicationConfiguration,omitempty" yaml:"replicationConfiguration,omitempty"`

	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning types.S3_BucketVersioning `json:"versioning,omitempty" yaml:"versioning,omitempty"`

	// A website object (documented below).
	Website types.S3_BucketWebsite `json:"website,omitempty" yaml:"website,omitempty"`

	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus string `json:"accelerationStatus,omitempty" yaml:"accelerationStatus,omitempty"`

	// The name of the bucket. If omitted, this provider will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []types.S3_BucketCorsRule `json:"corsRules,omitempty" yaml:"corsRules,omitempty"`

	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants []types.S3_BucketGrant `json:"grants,omitempty" yaml:"grants,omitempty"`

	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	/*
	   A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)

	   > --NOTE:-- You cannot use `acceleration_status` in `cn-north-1` or `us-gov-west-1`
	*/
	ObjectLockConfiguration types.S3_BucketObjectLockConfiguration `json:"objectLockConfiguration,omitempty" yaml:"objectLockConfiguration,omitempty"`

	/*
	   Specifies who should bear the cost of Amazon S3 data transfer.
	   Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	   the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	   developer guide for more information.
	*/
	RequestPayer string `json:"requestPayer,omitempty" yaml:"requestPayer,omitempty"`

	// A map of tags to assign to the bucket. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings []types.S3_BucketLogging `json:"loggings,omitempty" yaml:"loggings,omitempty"`

	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be lowercase and less than or equal to 37 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are -not- recoverable.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules []types.S3_BucketLifecycleRule `json:"lifecycleRules,omitempty" yaml:"lifecycleRules,omitempty"`

	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration types.S3_BucketServerSideEncryptionConfiguration `json:"serverSideEncryptionConfiguration,omitempty" yaml:"serverSideEncryptionConfiguration,omitempty"`

	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain string `json:"websiteDomain,omitempty" yaml:"websiteDomain,omitempty"`

	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint string `json:"websiteEndpoint,omitempty" yaml:"websiteEndpoint,omitempty"`

	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, and `log-delivery-write`. Defaults to `private`.  Conflicts with `grant`.
	Acl string `json:"acl,omitempty" yaml:"acl,omitempty"`
}
