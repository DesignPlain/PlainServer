package types

type Eventarc_TriggerMatchingCriteria struct {
	// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.
	Attribute string `json:"attribute,omitempty" yaml:"attribute,omitempty"`

	// Optional. The operator used for matching the events with the value of the filter. If not specified, only events that have an exact key-value pair specified in the filter are matched. The only allowed value is `match-path-pattern`.
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	/*
	   Required. The value for the attribute. See https://cloud.google.com/eventarc/docs/creating-triggers#trigger-gcloud for available values.

	   - - -
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
