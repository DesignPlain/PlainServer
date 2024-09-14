package redshift

type ClusterIamRoles struct {
	// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
	IamRoleArns []string `json:"iamRoleArns,omitempty" yaml:"iamRoleArns,omitempty"`

	// The name of the Redshift Cluster IAM Roles.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.
	DefaultIamRoleArn string `json:"defaultIamRoleArn,omitempty" yaml:"defaultIamRoleArn,omitempty"`
}
