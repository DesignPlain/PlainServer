package types

type Eks_NodeGroupRemoteAccess struct {
	// EC2 Key Pair name that provides access for remote communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify `source_security_group_ids` when you create an EKS Node Group, either port 3389 for Windows, or port 22 for all other operating systems is opened on the worker nodes to the Internet (0.0.0.0/0). For Windows nodes, this will allow you to use RDP, for all others this allows you to SSH into the worker nodes.
	Ec2SshKey string `json:"ec2SshKey,omitempty" yaml:"ec2SshKey,omitempty"`

	// Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify `ec2_ssh_key`, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
	SourceSecurityGroupIds []string `json:"sourceSecurityGroupIds,omitempty" yaml:"sourceSecurityGroupIds,omitempty"`
}
