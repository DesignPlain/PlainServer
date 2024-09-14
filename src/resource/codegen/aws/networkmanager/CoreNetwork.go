package networkmanager

type CoreNetwork struct {
	// Sets the base policy document for the core network. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	BasePolicyDocument string `json:"basePolicyDocument,omitempty" yaml:"basePolicyDocument,omitempty"`

	// The base policy created by setting the `create_base_policy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `base_policy_region` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	BasePolicyRegion string `json:"basePolicyRegion,omitempty" yaml:"basePolicyRegion,omitempty"`

	// A list of regions to add to the base policy. The base policy created by setting the `create_base_policy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `base_policy_regions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	BasePolicyRegions []string `json:"basePolicyRegions,omitempty" yaml:"basePolicyRegions,omitempty"`

	/*
	   Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `aws.networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `aws.networkmanager.CoreNetworkPolicyAttachment` resource.

	   ```json
	   {
	   "version": "2021.12",
	   "core-network-configuration": {
	   "asn-ranges": [
	   "64512-65534"
	   ],
	   "vpn-ecmp-support": false,
	   "edge-locations": [
	   {
	   "location": "us-east-1"
	   }
	   ]
	   },
	   "segments": [
	   {
	   "name": "segment",
	   "description": "base-policy",
	   "isolate-attachments": false,
	   "require-attachment-acceptance": false
	   }
	   ]
	   }
	   ```
	*/
	CreateBasePolicy bool `json:"createBasePolicy,omitempty" yaml:"createBasePolicy,omitempty"`

	// Description of the Core Network.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`

	// Key-value tags for the Core Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
