package types

type Firebaserules_RulesetMetadata struct {
	// Services that this ruleset has declarations for (e.g., "cloud.firestore"). There may be 0+ of these.
	Services []string `json:"services,omitempty" yaml:"services,omitempty"`
}
