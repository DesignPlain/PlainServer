package workspaces

import types "DesignSphere_Server/src/resource/aws/types"

type Directory struct {
	// The directory identifier for registration in WorkSpaces service.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// The identifiers of the IP access control groups associated with the directory.
	IpGroupIds []string `json:"ipGroupIds,omitempty" yaml:"ipGroupIds,omitempty"`

	// Permissions to enable or disable self-service capabilities. Defined below.
	SelfServicePermissions types.Workspaces_DirectorySelfServicePermissions `json:"selfServicePermissions,omitempty" yaml:"selfServicePermissions,omitempty"`

	// The identifiers of the subnets where the directory resides.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags assigned to the WorkSpaces directory. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies which devices and operating systems users can use to access their WorkSpaces. Defined below.
	WorkspaceAccessProperties types.Workspaces_DirectoryWorkspaceAccessProperties `json:"workspaceAccessProperties,omitempty" yaml:"workspaceAccessProperties,omitempty"`

	// Default properties that are used for creating WorkSpaces. Defined below.
	WorkspaceCreationProperties types.Workspaces_DirectoryWorkspaceCreationProperties `json:"workspaceCreationProperties,omitempty" yaml:"workspaceCreationProperties,omitempty"`
}
