package types

type Emr_ClusterBootstrapAction struct {
	// Location of the script to run during a bootstrap action. Can be either a location in Amazon S3 or on a local file system.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// List of command line arguments to pass to the bootstrap action script.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// Name of the bootstrap action.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
