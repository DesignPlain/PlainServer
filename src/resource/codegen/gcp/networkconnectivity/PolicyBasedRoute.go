package networkconnectivity

import types "libds/gcp/types"

type PolicyBasedRoute struct {
	// The priority of this policy-based route. Priority is used to break ties in cases where there are more than one matching policy-based routes found. In cases where multiple policy-based routes are matched, the one with the lowest-numbered priority value wins. The default value is 1000. The priority value must be from 1 to 65535, inclusive.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   VM instances to which this policy-based route applies to.
	   Structure is documented below.
	*/
	VirtualMachine types.Networkconnectivity_PolicyBasedRouteVirtualMachine `json:"virtualMachine,omitempty" yaml:"virtualMachine,omitempty"`

	/*
	   The filter to match L4 traffic.
	   Structure is documented below.
	*/
	Filter types.Networkconnectivity_PolicyBasedRouteFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	/*
	   The interconnect attachments that this policy-based route applies to.
	   Structure is documented below.
	*/
	InterconnectAttachment types.Networkconnectivity_PolicyBasedRouteInterconnectAttachment `json:"interconnectAttachment,omitempty" yaml:"interconnectAttachment,omitempty"`

	// The name of the policy based route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The IP address of a global-access-enabled L4 ILB that is the next hop for matching packets.
	NextHopIlbIp string `json:"nextHopIlbIp,omitempty" yaml:"nextHopIlbIp,omitempty"`

	/*
	   Other routes that will be referenced to determine the next hop of the packet.
	   Possible values are: `DEFAULT_ROUTING`.
	*/
	NextHopOtherRoutes string `json:"nextHopOtherRoutes,omitempty" yaml:"nextHopOtherRoutes,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   User-defined labels.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
