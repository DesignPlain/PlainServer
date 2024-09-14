package types

type Workspaces_getDirectoryWorkspaceCreationProperty struct {
	// Indicates whether users are local administrators of their WorkSpaces.
	UserEnabledAsLocalAdministrator bool `json:"userEnabledAsLocalAdministrator,omitempty" yaml:"userEnabledAsLocalAdministrator,omitempty"`

	// The identifier of your custom security group. Should relate to the same VPC, where workspaces reside in.
	CustomSecurityGroupId string `json:"customSecurityGroupId,omitempty" yaml:"customSecurityGroupId,omitempty"`

	// The default organizational unit (OU) for your WorkSpace directories.
	DefaultOu string `json:"defaultOu,omitempty" yaml:"defaultOu,omitempty"`

	// Indicates whether internet access is enabled for your WorkSpaces.
	EnableInternetAccess bool `json:"enableInternetAccess,omitempty" yaml:"enableInternetAccess,omitempty"`

	// Indicates whether maintenance mode is enabled for your WorkSpaces. For more information, see [WorkSpace Maintenance](https://docs.aws.amazon.com/workspaces/latest/adminguide/workspace-maintenance.html).
	EnableMaintenanceMode bool `json:"enableMaintenanceMode,omitempty" yaml:"enableMaintenanceMode,omitempty"`
}
