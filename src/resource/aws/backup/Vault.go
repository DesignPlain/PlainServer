package backup

type Vault struct {
	// A boolean that indicates that all recovery points stored in the vault are deleted so that the vault can be destroyed without error.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Name of the backup vault to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Metadata that you can assign to help organize the resources that you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
