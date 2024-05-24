package types

type Cloudtrail_EventDataStoreAdvancedEventSelectorFieldSelector struct {
	// A list of values that excludes events that match the first few characters of the event record field specified as the value of `field`.
	NotStartsWiths []string `json:"notStartsWiths,omitempty" yaml:"notStartsWiths,omitempty"`

	// A list of values that includes events that match the first few characters of the event record field specified as the value of `field`.
	StartsWiths []string `json:"startsWiths,omitempty" yaml:"startsWiths,omitempty"`

	// A list of values that includes events that match the last few characters of the event record field specified as the value of `field`.
	EndsWiths []string `json:"endsWiths,omitempty" yaml:"endsWiths,omitempty"`

	// A list of values that includes events that match the exact value of the event record field specified as the value of `field`. This is the only valid operator that you can use with the `readOnly`, `eventCategory`, and `resources.type` fields.
	Equals []string `json:"equals,omitempty" yaml:"equals,omitempty"`

	// Specifies a field in an event record on which to filter events to be logged. You can specify only the following values: `readOnly`, `eventSource`, `eventName`, `eventCategory`, `resources.type`, `resources.ARN`.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// A list of values that excludes events that match the last few characters of the event record field specified as the value of `field`.
	NotEndsWiths []string `json:"notEndsWiths,omitempty" yaml:"notEndsWiths,omitempty"`

	// A list of values that excludes events that match the exact value of the event record field specified as the value of `field`.
	NotEquals []string `json:"notEquals,omitempty" yaml:"notEquals,omitempty"`
}
