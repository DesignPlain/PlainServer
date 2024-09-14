package dns

import types "libds/gcp/types"

type RecordSet struct {
	/*
	   The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
	   data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
	   record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
	   string (e.g. "first255characters\"\"morecharacters").
	*/
	Rrdatas []string `json:"rrdatas,omitempty" yaml:"rrdatas,omitempty"`

	// The time-to-live of this record set (seconds).
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	/*
	   The DNS record set type.

	   - - -
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The name of the zone in which this record set will
	   reside.
	*/
	ManagedZone string `json:"managedZone,omitempty" yaml:"managedZone,omitempty"`

	// The DNS name this record set will apply to.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The configuration for steering traffic based on query.
	   Now you can specify either Weighted Round Robin(WRR) type or Geolocation(GEO) type.
	   Structure is documented below.
	*/
	RoutingPolicy types.Dns_RecordSetRoutingPolicy `json:"routingPolicy,omitempty" yaml:"routingPolicy,omitempty"`
}
