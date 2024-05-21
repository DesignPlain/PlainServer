package types

type Apigee_SharedflowMetaData struct {
	// Time at which the API proxy was created, in milliseconds since epoch.
	CreatedAt string `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`

	// Time at which the API proxy was most recently modified, in milliseconds since epoch.
	LastModifiedAt string `json:"lastModifiedAt,omitempty" yaml:"lastModifiedAt,omitempty"`

	// The type of entity described
	SubType string `json:"subType,omitempty" yaml:"subType,omitempty"`
}
