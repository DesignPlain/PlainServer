package compute

import types "libds/gcp/types"

type PacketMirroring struct {
	/*
	   The Region in which the created address should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Specifies the mirrored VPC network. Only packets in this network
	   will be mirrored. All mirrored VMs should have a NIC in the given
	   network. All mirrored subnetworks should belong to the given network.
	   Structure is documented below.
	*/
	Network types.Compute_PacketMirroringNetwork `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A filter for mirrored traffic.  If unset, all traffic is mirrored.
	   Structure is documented below.
	*/
	Filter types.Compute_PacketMirroringFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	/*
	   A means of specifying which resources to mirror.
	   Structure is documented below.
	*/
	MirroredResources types.Compute_PacketMirroringMirroredResources `json:"mirroredResources,omitempty" yaml:"mirroredResources,omitempty"`

	// The name of the packet mirroring rule
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Since only one rule can be active at a time, priority is
	   used to break ties in the case of two rules that apply to
	   the same instances.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	   that will be used as collector for mirrored traffic. The
	   specified forwarding rule must have is_mirroring_collector
	   set to true.
	   Structure is documented below.
	*/
	CollectorIlb types.Compute_PacketMirroringCollectorIlb `json:"collectorIlb,omitempty" yaml:"collectorIlb,omitempty"`

	// A human-readable description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
