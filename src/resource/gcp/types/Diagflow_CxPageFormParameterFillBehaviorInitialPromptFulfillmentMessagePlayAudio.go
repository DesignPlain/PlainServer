package types

type Diagflow_CxPageFormParameterFillBehaviorInitialPromptFulfillmentMessagePlayAudio struct {
	/*
	   (Output)
	   Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	*/
	AllowPlaybackInterruption bool `json:"allowPlaybackInterruption,omitempty" yaml:"allowPlaybackInterruption,omitempty"`

	// URI of the audio clip. Dialogflow does not impose any validation on this value. It is specific to the client that reads it.
	AudioUri string `json:"audioUri,omitempty" yaml:"audioUri,omitempty"`
}
