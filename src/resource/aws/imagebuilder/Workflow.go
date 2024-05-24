package imagebuilder

type Workflow struct {
	// Description of the workflow.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the workflow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// S3 URI with data of the workflow. Exactly one of `data` and `uri` can be specified.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	/*
	   Version of the workflow.

	   The following arguments are optional:
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Change description of the workflow.
	ChangeDescription string `json:"changeDescription,omitempty" yaml:"changeDescription,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key used to encrypt the workflow.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Key-value map of resource tags for the workflow. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Type of the workflow. Valid values: `BUILD`, `TEST`, `DISTRIBUTION`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Inline YAML string with data of the workflow. Exactly one of `data` and `uri` can be specified.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`
}
