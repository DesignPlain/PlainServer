package types

type Eks_getNodeGroupRemoteAccess struct {
	// EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group.
	Ec2SshKey string `json:"ec2SshKey,omitempty" yaml:"ec2SshKey,omitempty"`

	// Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes.
	SourceSecurityGroupIds []string `json:"sourceSecurityGroupIds,omitempty" yaml:"sourceSecurityGroupIds,omitempty"`
}
