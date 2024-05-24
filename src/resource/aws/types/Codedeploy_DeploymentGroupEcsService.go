package types

type Codedeploy_DeploymentGroupEcsService struct {
	// The name of the ECS cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// The name of the ECS service.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
}
