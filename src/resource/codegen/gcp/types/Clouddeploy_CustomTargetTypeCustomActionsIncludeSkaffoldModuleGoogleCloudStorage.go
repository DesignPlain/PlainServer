package types

type Clouddeploy_CustomTargetTypeCustomActionsIncludeSkaffoldModuleGoogleCloudStorage struct {
	// Relative path from the source to the Skaffold file.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Cloud Storage source paths to copy recursively. For example, providing `gs://my-bucket/dir/configs/-` will result in Skaffold copying all files within the `dir/configs` directory in the bucket `my-bucket`.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
