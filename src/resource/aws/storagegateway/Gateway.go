package storagegateway

import types "DesignSphere_Server/src/resource/aws/types"

type Gateway struct {
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone string `json:"gatewayTimezone,omitempty" yaml:"gatewayTimezone,omitempty"`

	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
	MediumChangerType string `json:"mediumChangerType,omitempty" yaml:"mediumChangerType,omitempty"`

	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings types.Storagegateway_GatewaySmbActiveDirectorySettings `json:"smbActiveDirectorySettings,omitempty" yaml:"smbActiveDirectorySettings,omitempty"`

	// Specifies whether the shares on this gateway appear when listing shares.
	SmbFileShareVisibility bool `json:"smbFileShareVisibility,omitempty" yaml:"smbFileShareVisibility,omitempty"`

	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec int `json:"averageUploadRateLimitInBitsPerSec,omitempty" yaml:"averageUploadRateLimitInBitsPerSec,omitempty"`

	// Name of the gateway.
	GatewayName string `json:"gatewayName,omitempty" yaml:"gatewayName,omitempty"`

	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType string `json:"tapeDriveType,omitempty" yaml:"tapeDriveType,omitempty"`

	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType string `json:"gatewayType,omitempty" yaml:"gatewayType,omitempty"`

	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint string `json:"gatewayVpcEndpoint,omitempty" yaml:"gatewayVpcEndpoint,omitempty"`

	// The gateway's weekly maintenance start time information, including day and time of the week. The maintenance time is the time in your gateway's time zone. More details below.
	MaintenanceStartTime types.Storagegateway_GatewayMaintenanceStartTime `json:"maintenanceStartTime,omitempty" yaml:"maintenanceStartTime,omitempty"`

	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword string `json:"smbGuestPassword,omitempty" yaml:"smbGuestPassword,omitempty"`

	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy string `json:"smbSecurityStrategy,omitempty" yaml:"smbSecurityStrategy,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Gateway activation key during resource creation. Conflicts with `gateway_ip_address`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey string `json:"activationKey,omitempty" yaml:"activationKey,omitempty"`

	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec int `json:"averageDownloadRateLimitInBitsPerSec,omitempty" yaml:"averageDownloadRateLimitInBitsPerSec,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn,omitempty" yaml:"cloudwatchLogGroupArn,omitempty"`

	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress string `json:"gatewayIpAddress,omitempty" yaml:"gatewayIpAddress,omitempty"`
}
