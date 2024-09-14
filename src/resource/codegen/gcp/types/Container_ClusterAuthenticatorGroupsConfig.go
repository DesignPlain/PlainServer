package types

type Container_ClusterAuthenticatorGroupsConfig struct {
	// The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format `gke-security-groups@yourdomain.com`.
	SecurityGroup string `json:"securityGroup,omitempty" yaml:"securityGroup,omitempty"`
}
