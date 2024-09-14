package types

type Vmwareengine_getPrivateCloudVcenter struct {
	// Version of the appliance.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Fully qualified domain name of the appliance.
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	// Internal IP address of the appliance.
	InternalIp string `json:"internalIp,omitempty" yaml:"internalIp,omitempty"`

	// State of the appliance. Possible values: ["ACTIVE", "CREATING"]
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
