package types

type Sagemaker_FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsd struct {
	// The whole number of dollars in the amount. Valid value range between `0` and `2`.
	Dollars int `json:"dollars,omitempty" yaml:"dollars,omitempty"`

	// Fractions of a cent, in tenths. Valid value range between `0` and `9`.
	TenthFractionsOfACent int `json:"tenthFractionsOfACent,omitempty" yaml:"tenthFractionsOfACent,omitempty"`

	// The fractional portion, in cents, of the amount. Valid value range between `0` and `99`.
	Cents int `json:"cents,omitempty" yaml:"cents,omitempty"`
}
