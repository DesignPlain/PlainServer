package types

type Gkeonprem_BareMetalAdminClusterNetworkConfigIslandModeCidr struct {
	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges. This field cannot be changed after creation.
	PodAddressCidrBlocks []string `json:"podAddressCidrBlocks,omitempty" yaml:"podAddressCidrBlocks,omitempty"`

	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges. This field cannot be changed after creation.
	ServiceAddressCidrBlocks []string `json:"serviceAddressCidrBlocks,omitempty" yaml:"serviceAddressCidrBlocks,omitempty"`
}
