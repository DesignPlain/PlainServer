package types

type Accesscontextmanager_AccessLevelsAccessLevelBasicCondition struct {
	/*
	   The request must originate from one of the provided VPC networks in Google Cloud. Cannot specify this field together with `ip_subnetworks`.
	   Structure is documented below.
	*/
	VpcNetworkSources []Accesscontextmanager_AccessLevelsAccessLevelBasicConditionVpcNetworkSource `json:"vpcNetworkSources,omitempty" yaml:"vpcNetworkSources,omitempty"`

	/*
	   Device specific restrictions, all restrictions must hold for
	   the Condition to be true. If not specified, all devices are
	   allowed.
	   Structure is documented below.
	*/
	DevicePolicy Accesscontextmanager_AccessLevelsAccessLevelBasicConditionDevicePolicy `json:"devicePolicy,omitempty" yaml:"devicePolicy,omitempty"`

	/*
	   A list of CIDR block IP subnetwork specification. May be IPv4
	   or IPv6.
	   Note that for a CIDR IP address block, the specified IP address
	   portion must be properly truncated (i.e. all the host bits must
	   be zero) or the input is considered malformed. For example,
	   "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	   for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	   is not. The originating IP of a request must be in one of the
	   listed subnets in order for this Condition to be true.
	   If empty, all IP addresses are allowed.
	*/
	IpSubnetworks []string `json:"ipSubnetworks,omitempty" yaml:"ipSubnetworks,omitempty"`

	/*
	   An allowed list of members (users, service accounts).
	   Using groups is not supported yet.
	   The signed-in user originating the request must be a part of one
	   of the provided members. If not specified, a request may come
	   from any user (logged in/not logged in, not present in any
	   groups, etc.).
	   Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	*/
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   Whether to negate the Condition. If true, the Condition becomes
	   a NAND over its non-empty fields, each field must be false for
	   the Condition overall to be satisfied. Defaults to false.
	*/
	Negate bool `json:"negate,omitempty" yaml:"negate,omitempty"`

	/*
	   The request must originate from one of the provided
	   countries/regions.
	   Format: A valid ISO 3166-1 alpha-2 code.
	*/
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`

	/*
	   A list of other access levels defined in the same Policy,
	   referenced by resource name. Referencing an AccessLevel which
	   does not exist is an error. All access levels listed must be
	   granted for the Condition to be true.
	   Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	*/
	RequiredAccessLevels []string `json:"requiredAccessLevels,omitempty" yaml:"requiredAccessLevels,omitempty"`
}
