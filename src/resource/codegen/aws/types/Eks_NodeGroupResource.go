package types

type Eks_NodeGroupResource struct {
	// List of objects containing information about AutoScaling Groups.
	AutoscalingGroups []Eks_NodeGroupResourceAutoscalingGroup `json:"autoscalingGroups,omitempty" yaml:"autoscalingGroups,omitempty"`

	// Identifier of the remote access EC2 Security Group.
	RemoteAccessSecurityGroupId string `json:"remoteAccessSecurityGroupId,omitempty" yaml:"remoteAccessSecurityGroupId,omitempty"`
}
