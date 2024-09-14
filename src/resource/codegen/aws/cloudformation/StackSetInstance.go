package cloudformation

import types "libds/aws/types"

type StackSetInstance struct {
	// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
	CallAs string `json:"callAs,omitempty" yaml:"callAs,omitempty"`

	// AWS Organizations accounts to which StackSets deploys. StackSets doesn't deploy stack instances to the organization management account, even if the organization management account is in your organization or in an OU in your organization. Drift detection is not possible for this argument. See deployment_targets below.
	DeploymentTargets types.Cloudformation_StackSetInstanceDeploymentTargets `json:"deploymentTargets,omitempty" yaml:"deploymentTargets,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set operation.
	OperationPreferences types.Cloudformation_StackSetInstanceOperationPreferences `json:"operationPreferences,omitempty" yaml:"operationPreferences,omitempty"`

	// Key-value map of input parameters to override from the StackSet for this Instance.
	ParameterOverrides map[string]string `json:"parameterOverrides,omitempty" yaml:"parameterOverrides,omitempty"`

	// Target AWS Region to create a Stack based on the StackSet. Defaults to current region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// During resource destroy, remove Instance from StackSet while keeping the Stack and its associated resources. Must be enabled in the state _before_ destroy operation to take effect. You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to `false`.
	RetainStack bool `json:"retainStack,omitempty" yaml:"retainStack,omitempty"`

	// Name of the StackSet.
	StackSetName string `json:"stackSetName,omitempty" yaml:"stackSetName,omitempty"`
}
