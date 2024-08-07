package appstream

import types "DesignSphere_Server/src/resource/aws/types"

type Fleet struct {
	// Amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the `disconnect_timeout_in_seconds` time interval begins. Defaults to 60 seconds.
	IdleDisconnectTimeoutInSeconds int `json:"idleDisconnectTimeoutInSeconds,omitempty" yaml:"idleDisconnectTimeoutInSeconds,omitempty"`

	// ARN of the public, private, or shared image to use.
	ImageArn string `json:"imageArn,omitempty" yaml:"imageArn,omitempty"`

	// Instance type to use when launching fleet instances.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	/*
	   Unique name for the fleet.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for the desired capacity of the fleet. See below.
	ComputeCapacity types.Appstream_FleetComputeCapacity `json:"computeCapacity,omitempty" yaml:"computeCapacity,omitempty"`

	// Amount of time that a streaming session remains active after users disconnect.
	DisconnectTimeoutInSeconds int `json:"disconnectTimeoutInSeconds,omitempty" yaml:"disconnectTimeoutInSeconds,omitempty"`

	// Human-readable friendly name for the AppStream fleet.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Fleet type. Valid values are: `ON_DEMAND`, `ALWAYS_ON`
	FleetType string `json:"fleetType,omitempty" yaml:"fleetType,omitempty"`

	// ARN of the IAM role to apply to the fleet.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// Name of the image used to create the fleet.
	ImageName string `json:"imageName,omitempty" yaml:"imageName,omitempty"`

	// Configuration block for the name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain. See below.
	DomainJoinInfo types.Appstream_FleetDomainJoinInfo `json:"domainJoinInfo,omitempty" yaml:"domainJoinInfo,omitempty"`

	// Enables or disables default internet access for the fleet.
	EnableDefaultInternetAccess bool `json:"enableDefaultInternetAccess,omitempty" yaml:"enableDefaultInternetAccess,omitempty"`

	// AppStream 2.0 view that is displayed to your users when they stream from the fleet. When `APP` is specified, only the windows of applications opened by users display. When `DESKTOP` is specified, the standard desktop that is provided by the operating system displays. If not specified, defaults to `APP`.
	StreamView string `json:"streamView,omitempty" yaml:"streamView,omitempty"`

	// Map of tags to attach to AppStream instances.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for the VPC configuration for the image builder. See below.
	VpcConfig types.Appstream_FleetVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Description to display.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Maximum amount of time that a streaming session can remain active, in seconds.
	MaxUserDurationInSeconds int `json:"maxUserDurationInSeconds,omitempty" yaml:"maxUserDurationInSeconds,omitempty"`
}
