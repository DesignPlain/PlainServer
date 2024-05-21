package types

type Clouddeploy_TargetRun struct {
	// Required. The location where the Cloud Run Service should be located. Format is `projects/{project}/locations/{location}`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
