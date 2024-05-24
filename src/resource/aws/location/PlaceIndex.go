package location

import types "DesignSphere_Server/src/resource/aws/types"

type PlaceIndex struct {
	// Specifies the geospatial data provider for the new place index.
	DataSource string `json:"dataSource,omitempty" yaml:"dataSource,omitempty"`

	// Configuration block with the data storage option chosen for requesting Places. Detailed below.
	DataSourceConfiguration types.Location_PlaceIndexDataSourceConfiguration `json:"dataSourceConfiguration,omitempty" yaml:"dataSourceConfiguration,omitempty"`

	// The optional description for the place index resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name of the place index resource.

	   The following arguments are optional:
	*/
	IndexName string `json:"indexName,omitempty" yaml:"indexName,omitempty"`

	// Key-value tags for the place index. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
