package types

type Codebuild_ProjectSecondaryArtifact struct {
	// Build output artifact's type. Valid values `CODEPIPELINE`, `NO_ARTIFACTS`, and `S3`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Artifact identifier. Must be the same specified inside the AWS CodeBuild build specification.
	ArtifactIdentifier string `json:"artifactIdentifier,omitempty" yaml:"artifactIdentifier,omitempty"`

	// Name of the project. If `type` is set to `CODEPIPELINE` or `NO_ARTIFACTS`, this value is ignored if specified. If `type` is set to `S3`, this is the name of the output artifact object.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Namespace to use in storing build artifacts. If `type` is set to `CODEPIPELINE` or `NO_ARTIFACTS`, this value is ignored if specified. If `type` is set to `S3`, valid values are `BUILD_ID` or `NONE`.
	NamespaceType string `json:"namespaceType,omitempty" yaml:"namespaceType,omitempty"`

	// Type of build output artifact to create. If `type` is set to `CODEPIPELINE` or `NO_ARTIFACTS`, this value is ignored if specified. If `type` is set to `S3`, valid values are `NONE` or `ZIP`.
	Packaging string `json:"packaging,omitempty" yaml:"packaging,omitempty"`

	// Along with `namespace_type` and `name`, the pattern that AWS CodeBuild uses to name and store the output artifact. If `type` is set to `CODEPIPELINE` or `NO_ARTIFACTS`, this value is ignored if specified. If `type` is set to `S3`, this is the path to the output artifact.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket. By default, only the account that uploads the objects to the bucket has access to these objects. This property allows you to give the bucket owner access to these objects. Valid values are `NONE`, `READ_ONLY`, and `FULL`. The CodeBuild service role must have the `s3:PutBucketAcl` permission. This permission allows CodeBuild to modify the access control list for the bucket.
	BucketOwnerAccess string `json:"bucketOwnerAccess,omitempty" yaml:"bucketOwnerAccess,omitempty"`

	// Whether to disable encrypting output artifacts. If `type` is set to `NO_ARTIFACTS`, this value is ignored. Defaults to `false`.
	EncryptionDisabled bool `json:"encryptionDisabled,omitempty" yaml:"encryptionDisabled,omitempty"`

	// Information about the build output artifact location. If `type` is set to `CODEPIPELINE` or `NO_ARTIFACTS`, this value is ignored if specified. If `type` is set to `S3`, this is the name of the output bucket. If `path` is not specified, `location` can specify the path of the output artifact in the output bucket.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Whether a name specified in the build specification overrides the artifact name.
	OverrideArtifactName bool `json:"overrideArtifactName,omitempty" yaml:"overrideArtifactName,omitempty"`
}
