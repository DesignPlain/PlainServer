package projects

type UsageExportBucket struct {
	// The bucket to store reports in.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// A prefix for the reports, for instance, the project name.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
