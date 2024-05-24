package types

type Vpc_getSecurityGroupRulesFilter struct {
	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroupRules.html).
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Set of values that are accepted for the given field.
	   Security group rule IDs will be selected if any one of the given values match.
	*/
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
