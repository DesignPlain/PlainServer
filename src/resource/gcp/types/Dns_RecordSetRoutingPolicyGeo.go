package types

type Dns_RecordSetRoutingPolicyGeo struct {
	/*
	   For A and AAAA types only. The list of targets to be health checked. These can be specified along with `rrdatas` within this item.
	   Structure is document below.
	*/
	HealthCheckedTargets Dns_RecordSetRoutingPolicyGeoHealthCheckedTargets `json:"healthCheckedTargets,omitempty" yaml:"healthCheckedTargets,omitempty"`

	// The location name defined in Google Cloud.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Same as `rrdatas` above.
	Rrdatas []string `json:"rrdatas,omitempty" yaml:"rrdatas,omitempty"`
}
