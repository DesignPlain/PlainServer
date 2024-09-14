package types

type Quicksight_DataSetColumnGroupGeoSpatialColumnGroup struct {
	// Country code. Valid values are `US`.
	CountryCode string `json:"countryCode,omitempty" yaml:"countryCode,omitempty"`

	// A display name for the hierarchy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Columns in this hierarchy.
	Columns []string `json:"columns,omitempty" yaml:"columns,omitempty"`
}
