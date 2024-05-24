package amp

type AlertManagerDefinition struct {
	// ID of the prometheus workspace the alert manager definition should be linked to
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`

	// the alert manager definition that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html).
	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`
}
