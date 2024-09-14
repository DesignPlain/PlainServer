package types

type Ec2_getNetworkInsightsAnalysisExplanationNatGateway struct {
	// ARN of the selected Network Insights Analysis.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	//
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the filter field. Valid values can be found in the EC2 [`DescribeNetworkInsightsAnalyses`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeNetworkInsightsAnalyses.html) API Reference.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
