package route53recoverycontrol

type RoutingControl struct {
	// ARN of the cluster in which this routing control will reside.
	ClusterArn string `json:"clusterArn,omitempty" yaml:"clusterArn,omitempty"`

	// ARN of the control panel in which this routing control will reside.
	ControlPanelArn string `json:"controlPanelArn,omitempty" yaml:"controlPanelArn,omitempty"`

	/*
	   The name describing the routing control.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
