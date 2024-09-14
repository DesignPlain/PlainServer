package types

type Dns_RecordSetRoutingPolicy struct {
	/*
	   The configuration for Weighted Round Robin based routing policy.
	   Structure is document below.
	*/
	Wrrs []Dns_RecordSetRoutingPolicyWrr `json:"wrrs,omitempty" yaml:"wrrs,omitempty"`

	// Specifies whether to enable fencing for geo queries.
	EnableGeoFencing bool `json:"enableGeoFencing,omitempty" yaml:"enableGeoFencing,omitempty"`

	/*
	   The configuration for Geolocation based routing policy.
	   Structure is document below.
	*/
	Geos []Dns_RecordSetRoutingPolicyGeo `json:"geos,omitempty" yaml:"geos,omitempty"`

	/*
	   The configuration for a primary-backup policy with global to regional failover. Queries are responded to with the global primary targets, but if none of the primary targets are healthy, then we fallback to a regional failover policy.
	   Structure is document below.
	*/
	PrimaryBackup Dns_RecordSetRoutingPolicyPrimaryBackup `json:"primaryBackup,omitempty" yaml:"primaryBackup,omitempty"`
}
