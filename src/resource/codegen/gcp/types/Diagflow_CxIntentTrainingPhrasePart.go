package types

type Diagflow_CxIntentTrainingPhrasePart struct {
	// The parameter used to annotate this part of the training phrase. This field is required for annotated parts of the training phrase.
	ParameterId string `json:"parameterId,omitempty" yaml:"parameterId,omitempty"`

	// The text for this part.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}
