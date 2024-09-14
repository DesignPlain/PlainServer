package types

type Networkconnectivity_PolicyBasedRouteInterconnectAttachment struct {
	// Cloud region to install this policy-based route on for Interconnect attachments. Use `all` to install it on all Interconnect attachments.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
