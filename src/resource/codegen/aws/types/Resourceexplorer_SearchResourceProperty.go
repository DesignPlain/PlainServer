package types

type Resourceexplorer_SearchResourceProperty struct {
	// Details about this property. The content of this field is a JSON object that varies based on the resource type.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	// The date and time that the information about this resource property was last updated.
	LastReportedAt string `json:"lastReportedAt,omitempty" yaml:"lastReportedAt,omitempty"`

	// Name of this property of the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
