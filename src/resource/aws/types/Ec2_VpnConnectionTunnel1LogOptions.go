package types

type Ec2_VpnConnectionTunnel1LogOptions struct {
	// Options for sending VPN tunnel logs to CloudWatch. See CloudWatch Log Options below for more details.
	CloudwatchLogOptions Ec2_VpnConnectionTunnel1LogOptionsCloudwatchLogOptions `json:"cloudwatchLogOptions,omitempty" yaml:"cloudwatchLogOptions,omitempty"`
}
