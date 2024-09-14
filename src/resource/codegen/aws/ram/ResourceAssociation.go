package ram

type ResourceAssociation struct {
	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// Amazon Resource Name (ARN) of the RAM Resource Share.
	ResourceShareArn string `json:"resourceShareArn,omitempty" yaml:"resourceShareArn,omitempty"`
}
