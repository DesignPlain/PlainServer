package types

type Ec2_getEipsFilter struct {
	// Set of values that are accepted for the given field. An Elastic IP will be selected if any one of the given values matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
