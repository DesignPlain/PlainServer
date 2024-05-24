package route53recoverycontrol

type ControlPanel struct {
	// ARN of the cluster in which this control panel will reside.
	ClusterArn string `json:"clusterArn,omitempty" yaml:"clusterArn,omitempty"`

	// Name describing the control panel.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
