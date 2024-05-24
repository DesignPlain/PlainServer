package imagebuilder

type Component struct {
	// Change description of the component.
	ChangeDescription string `json:"changeDescription,omitempty" yaml:"changeDescription,omitempty"`

	// Whether to retain the old version when the resource is destroyed or replacement is necessary. Defaults to `false`.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// Key-value map of resource tags for the component. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Version of the component.

	   The following attributes are optional:
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Inline YAML string with data of the component. Exactly one of `data` and `uri` can be specified. the provider will only perform drift detection of its value when present in a configuration.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	// Description of the component.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key used to encrypt the component.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Name of the component.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Platform of the component.
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`

	// Set of Operating Systems (OS) supported by the component.
	SupportedOsVersions []string `json:"supportedOsVersions,omitempty" yaml:"supportedOsVersions,omitempty"`

	/*
	   S3 URI with data of the component. Exactly one of `data` and `uri` can be specified.

	   > --NOTE:-- Updating `data` or `uri` requires specifying a new `version`. This causes replacement of the resource. The `skip_destroy` argument can be used to retain the old version.
	*/
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
