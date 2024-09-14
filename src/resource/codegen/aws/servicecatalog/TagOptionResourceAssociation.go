package servicecatalog

type TagOptionResourceAssociation struct {
	// Resource identifier.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Tag Option identifier.
	TagOptionId string `json:"tagOptionId,omitempty" yaml:"tagOptionId,omitempty"`
}
