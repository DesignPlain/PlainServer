package dns

import types "DesignSphere_Server/src/resource/gcp/types"

type ManagedZone struct {
	/*
	   Cloud logging configuration
	   Structure is documented below.
	*/
	CloudLoggingConfig types.Dns_ManagedZoneCloudLoggingConfig `json:"cloudLoggingConfig,omitempty" yaml:"cloudLoggingConfig,omitempty"`

	/*
	   DNSSEC configuration
	   Structure is documented below.
	*/
	DnssecConfig types.Dns_ManagedZoneDnssecConfig `json:"dnssecConfig,omitempty" yaml:"dnssecConfig,omitempty"`

	/*
	   The presence for this field indicates that outbound forwarding is enabled
	   for this zone. The value of this field contains the set of destinations
	   to forward to.
	   Structure is documented below.
	*/
	ForwardingConfig types.Dns_ManagedZoneForwardingConfig `json:"forwardingConfig,omitempty" yaml:"forwardingConfig,omitempty"`

	// The DNS name of this managed zone, for instance "example.com.".
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	// Set this true to delete all records in the zone.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	/*
	   User assigned name for this resource.
	   Must be unique within the project.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The presence of this field indicates that DNS Peering is enabled for this
	   zone. The value of this field contains the network to peer with.
	   Structure is documented below.
	*/
	PeeringConfig types.Dns_ManagedZonePeeringConfig `json:"peeringConfig,omitempty" yaml:"peeringConfig,omitempty"`

	/*
	   The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.
	   Structure is documented below.
	*/
	ServiceDirectoryConfig types.Dns_ManagedZoneServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty" yaml:"serviceDirectoryConfig,omitempty"`

	/*
	   The zone's visibility: public zones are exposed to the Internet,
	   while private zones are visible only to Virtual Private Cloud resources.
	   Default value is `public`.
	   Possible values are: `private`, `public`.
	*/
	Visibility string `json:"visibility,omitempty" yaml:"visibility,omitempty"`

	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A set of key/value label pairs to assign to this ManagedZone.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
	   lookup queries using automatically configured records for VPC resources. This only applies
	   to networks listed under `private_visibility_config`.
	*/
	ReverseLookup bool `json:"reverseLookup,omitempty" yaml:"reverseLookup,omitempty"`

	/*
	   For privately visible zones, the set of Virtual Private Cloud
	   resources that the zone is visible from. At least one of `gke_clusters` or `networks` must be specified.
	   Structure is documented below.
	*/
	PrivateVisibilityConfig types.Dns_ManagedZonePrivateVisibilityConfig `json:"privateVisibilityConfig,omitempty" yaml:"privateVisibilityConfig,omitempty"`
}
