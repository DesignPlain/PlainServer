package cloudwatch

type QueryDefinition struct {
	// Specific log groups to use with the query.
	LogGroupNames []string `json:"logGroupNames,omitempty" yaml:"logGroupNames,omitempty"`

	// The name of the query.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	QueryString string `json:"queryString,omitempty" yaml:"queryString,omitempty"`
}
