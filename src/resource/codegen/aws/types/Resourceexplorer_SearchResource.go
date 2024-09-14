package types

type Resourceexplorer_SearchResource struct {
	// Amazon Web Services Region in which the resource was created and exists.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Type of the resource.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Amazon Web Service that owns the resource and is responsible for creating and updating it.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// Amazon resource name of resource.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The date and time that the information about this resource property was last updated.
	LastReportedAt string `json:"lastReportedAt,omitempty" yaml:"lastReportedAt,omitempty"`

	// Amazon Web Services account that owns the resource.
	OwningAccountId string `json:"owningAccountId,omitempty" yaml:"owningAccountId,omitempty"`

	// Structure with additional type-specific details about the resource.  See `properties` below.
	Properties []Resourceexplorer_SearchResourceProperty `json:"properties,omitempty" yaml:"properties,omitempty"`
}
