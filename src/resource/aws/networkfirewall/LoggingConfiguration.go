package networkfirewall

import types "DesignSphere_Server/src/resource/aws/types"

type LoggingConfiguration struct {
	// The Amazon Resource Name (ARN) of the Network Firewall firewall.
	FirewallArn string `json:"firewallArn,omitempty" yaml:"firewallArn,omitempty"`

	// A configuration block describing how AWS Network Firewall performs logging for a firewall. See Logging Configuration below for details.
	LoggingConfiguration types.Networkfirewall_LoggingConfigurationLoggingConfiguration `json:"loggingConfiguration,omitempty" yaml:"loggingConfiguration,omitempty"`
}
