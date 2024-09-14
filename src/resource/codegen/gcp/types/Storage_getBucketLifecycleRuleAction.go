package types

type Storage_getBucketLifecycleRuleAction struct {
	// The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// The type of the action of this Lifecycle Rule. Supported values include: Delete, SetStorageClass and AbortIncompleteMultipartUpload.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
