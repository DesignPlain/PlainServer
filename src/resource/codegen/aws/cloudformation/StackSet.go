package cloudformation

import types "libds/aws/types"

type StackSet struct {
	// Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the `SERVICE_MANAGED` permission model.
	AutoDeployment types.Cloudformation_StackSetAutoDeployment `json:"autoDeployment,omitempty" yaml:"autoDeployment,omitempty"`

	// Description of the StackSet.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole` when using the `SELF_MANAGED` permission model. This should not be defined when using the `SERVICE_MANAGED` permission model.
	ExecutionRoleName string `json:"executionRoleName,omitempty" yaml:"executionRoleName,omitempty"`

	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set update.
	OperationPreferences types.Cloudformation_StackSetOperationPreferences `json:"operationPreferences,omitempty" yaml:"operationPreferences,omitempty"`

	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
	TemplateBody string `json:"templateBody,omitempty" yaml:"templateBody,omitempty"`

	// Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the `SELF_MANAGED` permission model.
	AdministrationRoleArn string `json:"administrationRoleArn,omitempty" yaml:"administrationRoleArn,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
	CallAs string `json:"callAs,omitempty" yaml:"callAs,omitempty"`

	// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
	Capabilities []string `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`

	// Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
	ManagedExecution types.Cloudformation_StackSetManagedExecution `json:"managedExecution,omitempty" yaml:"managedExecution,omitempty"`

	// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Describes how the IAM roles required for your StackSet are created. Valid values: `SELF_MANAGED` (default), `SERVICE_MANAGED`.
	PermissionModel string `json:"permissionModel,omitempty" yaml:"permissionModel,omitempty"`

	// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
	TemplateUrl string `json:"templateUrl,omitempty" yaml:"templateUrl,omitempty"`
}
