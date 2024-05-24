package types

type Glue_MLTransformParametersFindMatchesParameters struct {
	// The value selected when tuning your transform for a balance between precision and recall.
	PrecisionRecallTradeOff float64 `json:"precisionRecallTradeOff,omitempty" yaml:"precisionRecallTradeOff,omitempty"`

	// The name of a column that uniquely identifies rows in the source table.
	PrimaryKeyColumnName string `json:"primaryKeyColumnName,omitempty" yaml:"primaryKeyColumnName,omitempty"`

	// The value that is selected when tuning your transform for a balance between accuracy and cost.
	AccuracyCostTradeOff float64 `json:"accuracyCostTradeOff,omitempty" yaml:"accuracyCostTradeOff,omitempty"`

	// The value to switch on or off to force the output to match the provided labels from users.
	EnforceProvidedLabels bool `json:"enforceProvidedLabels,omitempty" yaml:"enforceProvidedLabels,omitempty"`
}
