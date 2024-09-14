package dataexchange

type DataSet struct {
	// The type of asset that is added to a data set. Valid values are: `S3_SNAPSHOT`, `REDSHIFT_DATA_SHARE`, and `API_GATEWAY_API`.
	AssetType string `json:"assetType,omitempty" yaml:"assetType,omitempty"`

	// A description for the data set.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the data set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
