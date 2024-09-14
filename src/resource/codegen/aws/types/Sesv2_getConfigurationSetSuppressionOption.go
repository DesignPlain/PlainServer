package types

type Sesv2_getConfigurationSetSuppressionOption struct {
	// A list that contains the reasons that email addresses are automatically added to the suppression list for your account.
	SuppressedReasons []string `json:"suppressedReasons,omitempty" yaml:"suppressedReasons,omitempty"`
}
