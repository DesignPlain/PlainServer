package types

type Codepipeline_PipelineArtifactStoreEncryptionKey struct {
	// The KMS key ARN or ID
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The type of key; currently only `KMS` is supported
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
