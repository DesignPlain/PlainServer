package types

type Eks_getNodeGroupResource struct {
	// List of objects containing information about AutoScaling Groups.
	AutoscalingGroups []Eks_getNodeGroupResourceAutoscalingGroup `json:"autoscalingGroups,omitempty" yaml:"autoscalingGroups,omitempty"`

	// Identifier of the remote access EC2 Security Group.
	RemoteAccessSecurityGroupId string `json:"remoteAccessSecurityGroupId,omitempty" yaml:"remoteAccessSecurityGroupId,omitempty"`
}
