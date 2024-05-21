package types

type Cloudbuild_getTriggerApprovalConfig struct {
	/*
	   Whether or not approval is needed. If this is set on a build, it will become pending when run,
	   and will need to be explicitly approved to start.
	*/
	ApprovalRequired bool `json:"approvalRequired,omitempty" yaml:"approvalRequired,omitempty"`
}
