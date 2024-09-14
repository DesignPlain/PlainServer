package types

type Workspaces_DirectoryWorkspaceCreationProperties struct {
	// The default organizational unit (OU) for your WorkSpace directories. Should conform `"OU=<value>,DC=<value>,...,DC=<value>"` pattern.
	DefaultOu string `json:"defaultOu,omitempty" yaml:"defaultOu,omitempty"`

	// Indicates whether internet access is enabled for your WorkSpaces.
	EnableInternetAccess bool `json:"enableInternetAccess,omitempty" yaml:"enableInternetAccess,omitempty"`

	// Indicates whether maintenance mode is enabled for your WorkSpaces. For more information, see [WorkSpace Maintenance](https://docs.aws.amazon.com/workspaces/latest/adminguide/workspace-maintenance.html)..
	EnableMaintenanceMode bool `json:"enableMaintenanceMode,omitempty" yaml:"enableMaintenanceMode,omitempty"`

	// Indicates whether users are local administrators of their WorkSpaces.
	UserEnabledAsLocalAdministrator bool `json:"userEnabledAsLocalAdministrator,omitempty" yaml:"userEnabledAsLocalAdministrator,omitempty"`

	// The identifier of your custom security group. Should relate to the same VPC, where workspaces reside in.
	CustomSecurityGroupId string `json:"customSecurityGroupId,omitempty" yaml:"customSecurityGroupId,omitempty"`
}
