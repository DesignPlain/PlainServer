package types

type Diagflow_CxFlowEventHandlerTriggerFulfillmentMessageOutputAudioText struct {
	/*
	   (Output)
	   Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	*/
	AllowPlaybackInterruption bool `json:"allowPlaybackInterruption,omitempty" yaml:"allowPlaybackInterruption,omitempty"`

	// The SSML text to be synthesized. For more information, see SSML.
	Ssml string `json:"ssml,omitempty" yaml:"ssml,omitempty"`

	// The raw text to be synthesized.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}
