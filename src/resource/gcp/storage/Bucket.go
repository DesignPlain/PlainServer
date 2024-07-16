package storage

import types "DesignSphere_Server/src/resource/gcp/types"

type Bucket struct {
	// The bucket's [Autoclass](https://cloud.google.com/storage/docs/autoclass) configuration.  Structure is documented below.
	Autoclass types.Storage_BucketAutoclass `json:"autoclass,omitempty" yaml:"autoclass,omitempty"`

	/*
	   When deleting a bucket, this
	   boolean option will delete all contained objects. If you try to delete a
	   bucket that contains objects, the provider will fail that run.
	*/
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// The recovery point objective for cross-region replication of the bucket. Applicable only for dual and multi-region buckets. `"DEFAULT"` sets default replication. `"ASYNC_TURBO"` value enables turbo replication, valid for dual-region buckets only. See [Turbo Replication](https://cloud.google.com/storage/docs/managing-turbo-replication) for more information. If rpo is not specified at bucket creation, it defaults to `"DEFAULT"` for dual and multi-region buckets. --NOTE-- If used with single-region bucket, It will throw an error.
	//Rpo string `json:"rpo,omitempty" yaml:"rpo,omitempty"`

	// Configuration if the bucket acts as a website. Structure is documented below.
	Website types.Storage_BucketWebsite `json:"website,omitempty" yaml:"website,omitempty"`

	// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors []types.Storage_BucketCor `json:"cors,omitempty" yaml:"cors,omitempty"`

	// A map of key/value label pairs to assign to the bucket.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of the bucket.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Prevents public access to a bucket. Acceptable values are "inherited" or "enforced". If "inherited", the bucket uses [public access prevention](https://cloud.google.com/storage/docs/public-access-prevention). only if the bucket is subject to the public access prevention organization policy constraint. Defaults to "inherited".
	PublicAccessPrevention string `json:"publicAccessPrevention,omitempty" yaml:"publicAccessPrevention,omitempty"`

	// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
	RequesterPays bool `json:"requesterPays,omitempty" yaml:"requesterPays,omitempty"`

	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy types.Storage_BucketRetentionPolicy `json:"retentionPolicy,omitempty" yaml:"retentionPolicy,omitempty"`

	// The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. Structure is documented below.
	CustomPlacementConfig types.Storage_BucketCustomPlacementConfig `json:"customPlacementConfig,omitempty" yaml:"customPlacementConfig,omitempty"`

	// Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
	DefaultEventBasedHold bool `json:"defaultEventBasedHold,omitempty" yaml:"defaultEventBasedHold,omitempty"`

	// Enables [object retention](https://cloud.google.com/storage/docs/object-lock) on a storage bucket.
	//EnableObjectRetention bool `json:"enableObjectRetention,omitempty" yaml:"enableObjectRetention,omitempty"`

	// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.  Structure is documented below.
	Versioning types.Storage_BucketVersioning `json:"versioning,omitempty" yaml:"versioning,omitempty"`

	// The bucket's encryption configuration. Structure is documented below.
	Encryption types.Storage_BucketEncryption `json:"encryption,omitempty" yaml:"encryption,omitempty"`

	// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRules []types.Storage_BucketLifecycleRule `json:"lifecycleRules,omitempty" yaml:"lifecycleRules,omitempty"`

	/*
	   The [GCS location](https://cloud.google.com/storage/docs/bucket-locations).

	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration. Structure is documented below.
	Logging types.Storage_BucketLogging `json:"logging,omitempty" yaml:"logging,omitempty"`

	// Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
	UniformBucketLevelAccess bool `json:"uniformBucketLevelAccess,omitempty" yaml:"uniformBucketLevelAccess,omitempty"`
}
