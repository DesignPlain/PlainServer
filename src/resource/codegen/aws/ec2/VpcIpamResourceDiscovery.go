package ec2

import types "libds/aws/types"

type VpcIpamResourceDiscovery struct {
	// A description for the IPAM Resource Discovery.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. --You must set your provider block region as an operating_region.--
	OperatingRegions []types.Ec2_VpcIpamResourceDiscoveryOperatingRegion `json:"operatingRegions,omitempty" yaml:"operatingRegions,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
