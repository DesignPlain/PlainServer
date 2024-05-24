package codedeploy

import types "DesignSphere_Server/src/resource/aws/types"

type DeploymentGroup struct {
	// Configuration block of the automatic rollback configuration associated with the deployment group (documented below).
	AutoRollbackConfiguration types.Codedeploy_DeploymentGroupAutoRollbackConfiguration `json:"autoRollbackConfiguration,omitempty" yaml:"autoRollbackConfiguration,omitempty"`

	// Autoscaling groups associated with the deployment group.
	AutoscalingGroups []string `json:"autoscalingGroups,omitempty" yaml:"autoscalingGroups,omitempty"`

	// Configuration block of the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer (documented below).
	DeploymentStyle types.Codedeploy_DeploymentGroupDeploymentStyle `json:"deploymentStyle,omitempty" yaml:"deploymentStyle,omitempty"`

	// Tag filters associated with the deployment group. See the AWS docs for details.
	Ec2TagFilters []types.Codedeploy_DeploymentGroupEc2TagFilter `json:"ec2TagFilters,omitempty" yaml:"ec2TagFilters,omitempty"`

	// Configuration block(s) of Tag filters associated with the deployment group, which are also referred to as tag groups (documented below). See the AWS docs for details.
	Ec2TagSets []types.Codedeploy_DeploymentGroupEc2TagSet `json:"ec2TagSets,omitempty" yaml:"ec2TagSets,omitempty"`

	// Configuration block of Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision. Valid values are `UPDATE` and `IGNORE`. Defaults to `UPDATE`.
	OutdatedInstancesStrategy string `json:"outdatedInstancesStrategy,omitempty" yaml:"outdatedInstancesStrategy,omitempty"`

	// The name of the deployment group.
	DeploymentGroupName string `json:"deploymentGroupName,omitempty" yaml:"deploymentGroupName,omitempty"`

	// Single configuration block of the load balancer to use in a blue/green deployment (documented below).
	LoadBalancerInfo types.Codedeploy_DeploymentGroupLoadBalancerInfo `json:"loadBalancerInfo,omitempty" yaml:"loadBalancerInfo,omitempty"`

	// The service role ARN that allows deployments.
	ServiceRoleArn string `json:"serviceRoleArn,omitempty" yaml:"serviceRoleArn,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block(s) of the triggers for the deployment group (documented below).
	TriggerConfigurations []types.Codedeploy_DeploymentGroupTriggerConfiguration `json:"triggerConfigurations,omitempty" yaml:"triggerConfigurations,omitempty"`

	// Configuration block of alarms associated with the deployment group (documented below).
	AlarmConfiguration types.Codedeploy_DeploymentGroupAlarmConfiguration `json:"alarmConfiguration,omitempty" yaml:"alarmConfiguration,omitempty"`

	// The name of the application.
	AppName string `json:"appName,omitempty" yaml:"appName,omitempty"`

	// Configuration block of the blue/green deployment options for a deployment group (documented below).
	BlueGreenDeploymentConfig types.Codedeploy_DeploymentGroupBlueGreenDeploymentConfig `json:"blueGreenDeploymentConfig,omitempty" yaml:"blueGreenDeploymentConfig,omitempty"`

	// Configuration block(s) of the ECS services for a deployment group (documented below).
	EcsService types.Codedeploy_DeploymentGroupEcsService `json:"ecsService,omitempty" yaml:"ecsService,omitempty"`

	// The name of the group's deployment config. The default is "CodeDeployDefault.OneAtATime".
	DeploymentConfigName string `json:"deploymentConfigName,omitempty" yaml:"deploymentConfigName,omitempty"`

	// On premise tag filters associated with the group. See the AWS docs for details.
	OnPremisesInstanceTagFilters []types.Codedeploy_DeploymentGroupOnPremisesInstanceTagFilter `json:"onPremisesInstanceTagFilters,omitempty" yaml:"onPremisesInstanceTagFilters,omitempty"`
}
