package types

type Codebuild_ProjectCache struct {
	// Type of storage that will be used for the AWS CodeBuild project cache. Valid values: `NO_CACHE`, `LOCAL`, `S3`. Defaults to `NO_CACHE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Location where the AWS CodeBuild project stores cached resources. For type `S3`, the value must be a valid S3 bucket name/prefix.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Specifies settings that AWS CodeBuild uses to store and reuse build dependencies. Valid values:  `LOCAL_SOURCE_CACHE`, `LOCAL_DOCKER_LAYER_CACHE`, `LOCAL_CUSTOM_CACHE`.
	Modes []string `json:"modes,omitempty" yaml:"modes,omitempty"`
}
