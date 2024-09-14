package types

type Glue_ClassifierGrokClassifier struct {
	// Custom grok patterns used by this classifier.
	CustomPatterns string `json:"customPatterns,omitempty" yaml:"customPatterns,omitempty"`

	// The grok pattern used by this classifier.
	GrokPattern string `json:"grokPattern,omitempty" yaml:"grokPattern,omitempty"`

	// An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, Amazon CloudWatch Logs, and so on.
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`
}
