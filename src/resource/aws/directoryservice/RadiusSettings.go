package directoryservice

type RadiusSettings struct {
	// Display label.
	DisplayLabel string `json:"displayLabel,omitempty" yaml:"displayLabel,omitempty"`

	// Required for enabling RADIUS on the directory.
	SharedSecret string `json:"sharedSecret,omitempty" yaml:"sharedSecret,omitempty"`

	// The protocol specified for your RADIUS endpoints. Valid values: `PAP`, `CHAP`, `MS-CHAPv1`, `MS-CHAPv2`.
	AuthenticationProtocol string `json:"authenticationProtocol,omitempty" yaml:"authenticationProtocol,omitempty"`

	// The identifier of the directory for which you want to manager RADIUS settings.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// The port that your RADIUS server is using for communications. Your self-managed network must allow inbound traffic over this port from the AWS Directory Service servers.
	RadiusPort int `json:"radiusPort,omitempty" yaml:"radiusPort,omitempty"`

	// The maximum number of times that communication with the RADIUS server is attempted. Minimum value of `0`. Maximum value of `10`.
	RadiusRetries int `json:"radiusRetries,omitempty" yaml:"radiusRetries,omitempty"`

	// An array of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
	RadiusServers []string `json:"radiusServers,omitempty" yaml:"radiusServers,omitempty"`

	// The amount of time, in seconds, to wait for the RADIUS server to respond. Minimum value of `1`. Maximum value of `50`.
	RadiusTimeout int `json:"radiusTimeout,omitempty" yaml:"radiusTimeout,omitempty"`

	// Not currently used.
	UseSameUsername bool `json:"useSameUsername,omitempty" yaml:"useSameUsername,omitempty"`
}
