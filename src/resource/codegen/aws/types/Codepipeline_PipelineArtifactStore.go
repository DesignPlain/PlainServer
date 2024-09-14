package types

type Codepipeline_PipelineArtifactStore struct {
	// The location where AWS CodePipeline stores artifacts for a pipeline; currently only `S3` is supported.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The type of the artifact store, such as Amazon S3
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An `encryption_key` block is documented below.
	EncryptionKey Codepipeline_PipelineArtifactStoreEncryptionKey `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`
}
