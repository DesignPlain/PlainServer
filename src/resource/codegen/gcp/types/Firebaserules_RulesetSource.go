package types

type Firebaserules_RulesetSource struct {
	// `File` set constituting the `Source` bundle.
	Files []Firebaserules_RulesetSourceFile `json:"files,omitempty" yaml:"files,omitempty"`

	// `Language` of the `Source` bundle. If unspecified, the language will default to `FIREBASE_RULES`. Possible values: LANGUAGE_UNSPECIFIED, FIREBASE_RULES, EVENT_FLOW_TRIGGERS
	Language string `json:"language,omitempty" yaml:"language,omitempty"`
}
