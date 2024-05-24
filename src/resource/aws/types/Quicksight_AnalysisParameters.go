package types

type Quicksight_AnalysisParameters struct {
	// A list of parameters that have a data type of date-time. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DateTimeParameter.html).
	DateTimeParameters []Quicksight_AnalysisParametersDateTimeParameter `json:"dateTimeParameters,omitempty" yaml:"dateTimeParameters,omitempty"`

	// A list of parameters that have a data type of decimal. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalParameter.html).
	DecimalParameters []Quicksight_AnalysisParametersDecimalParameter `json:"decimalParameters,omitempty" yaml:"decimalParameters,omitempty"`

	// A list of parameters that have a data type of integer. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_IntegerParameter.html).
	IntegerParameters []Quicksight_AnalysisParametersIntegerParameter `json:"integerParameters,omitempty" yaml:"integerParameters,omitempty"`

	// A list of parameters that have a data type of string. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StringParameter.html).
	StringParameters []Quicksight_AnalysisParametersStringParameter `json:"stringParameters,omitempty" yaml:"stringParameters,omitempty"`
}
