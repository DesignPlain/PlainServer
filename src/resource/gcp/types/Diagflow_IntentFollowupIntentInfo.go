package types

type Diagflow_IntentFollowupIntentInfo struct {
	/*
	   The unique identifier of the followup intent.
	   Format: projects/<Project ID>/agent/intents/<Intent ID>.
	*/
	FollowupIntentName string `json:"followupIntentName,omitempty" yaml:"followupIntentName,omitempty"`

	/*
	   The unique identifier of the parent intent in the chain of followup intents.
	   Format: projects/<Project ID>/agent/intents/<Intent ID>.
	*/
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty" yaml:"parentFollowupIntentName,omitempty"`
}
