package types

type Codedeploy_DeploymentGroupEc2TagSet struct {
	// Tag filters associated with the deployment group. See the AWS docs for details.
	Ec2TagFilters []Codedeploy_DeploymentGroupEc2TagSetEc2TagFilter `json:"ec2TagFilters,omitempty" yaml:"ec2TagFilters,omitempty"`
}
