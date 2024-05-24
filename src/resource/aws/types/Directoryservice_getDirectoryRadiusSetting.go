package types

type Directoryservice_getDirectoryRadiusSetting struct {
	// The protocol specified for your RADIUS endpoints.
	AuthenticationProtocol string `json:"authenticationProtocol,omitempty" yaml:"authenticationProtocol,omitempty"`

	// Display label.
	DisplayLabel string `json:"displayLabel,omitempty" yaml:"displayLabel,omitempty"`

	// Port that your RADIUS server is using for communications.
	RadiusPort int `json:"radiusPort,omitempty" yaml:"radiusPort,omitempty"`

	// Maximum number of times that communication with the RADIUS server is attempted.
	RadiusRetries int `json:"radiusRetries,omitempty" yaml:"radiusRetries,omitempty"`

	// Set of strings that contains the fully qualified domain name (FQDN) or IP addresses of the RADIUS server endpoints, or the FQDN or IP addresses of your RADIUS server load balancer.
	RadiusServers []string `json:"radiusServers,omitempty" yaml:"radiusServers,omitempty"`

	// Amount of time, in seconds, to wait for the RADIUS server to respond.
	RadiusTimeout int `json:"radiusTimeout,omitempty" yaml:"radiusTimeout,omitempty"`

	// Not currently used.
	UseSameUsername bool `json:"useSameUsername,omitempty" yaml:"useSameUsername,omitempty"`
}
