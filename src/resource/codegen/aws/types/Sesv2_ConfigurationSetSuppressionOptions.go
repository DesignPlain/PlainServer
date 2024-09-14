package types

type Sesv2_ConfigurationSetSuppressionOptions struct {
	// A list that contains the reasons that email addresses are automatically added to the suppression list for your account. Valid values: `BOUNCE`, `COMPLAINT`.
	SuppressedReasons []string `json:"suppressedReasons,omitempty" yaml:"suppressedReasons,omitempty"`
}
