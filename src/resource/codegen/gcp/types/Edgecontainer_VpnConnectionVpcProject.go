package types

type Edgecontainer_VpnConnectionVpcProject struct {
	// The project of the VPC to connect to. If not specified, it is the same as the cluster project.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
