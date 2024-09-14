package types

type Fms_PolicyExcludeMap struct {
	// A list of AWS Organization member Accounts that you want to include for this AWS FMS Policy.
	Accounts []string `json:"accounts,omitempty" yaml:"accounts,omitempty"`

	/*
	   A list of IDs of the AWS Organizational Units that you want to include for this AWS FMS Policy. Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.

	   You can specify inclusions or exclusions, but not both. If you specify an `include_map`, AWS Firewall Manager applies the policy to all accounts specified by the `include_map`, and does not evaluate any `exclude_map` specifications. If you do not specify an `include_map`, then Firewall Manager applies the policy to all accounts except for those specified by the `exclude_map`.
	*/
	Orgunits []string `json:"orgunits,omitempty" yaml:"orgunits,omitempty"`
}
