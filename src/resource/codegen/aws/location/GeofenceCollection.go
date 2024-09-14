package location

type GeofenceCollection struct {
	// The optional description for the geofence collection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Key-value tags for the geofence collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The name of the geofence collection.

	   The following arguments are optional:
	*/
	CollectionName string `json:"collectionName,omitempty" yaml:"collectionName,omitempty"`
}
