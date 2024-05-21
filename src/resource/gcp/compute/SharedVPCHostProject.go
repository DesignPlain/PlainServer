package compute

type SharedVPCHostProject struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
