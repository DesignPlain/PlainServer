package types

type Ec2_VpnConnectionTunnel2LogOptionsCloudwatchLogOptions struct {
	// Enable or disable VPN tunnel logging feature. The default is `false`.
	LogEnabled bool `json:"logEnabled,omitempty" yaml:"logEnabled,omitempty"`

	// The Amazon Resource Name (ARN) of the CloudWatch log group to send logs to.
	LogGroupArn string `json:"logGroupArn,omitempty" yaml:"logGroupArn,omitempty"`

	// Set log format. Default format is json. Possible values are: `json` and `text`. The default is `json`.
	LogOutputFormat string `json:"logOutputFormat,omitempty" yaml:"logOutputFormat,omitempty"`
}
