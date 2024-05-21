package types

type Diagflow_CxPageEntryFulfillmentMessageText struct {
	/*
	   (Output)
	   Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
	*/
	AllowPlaybackInterruption bool `json:"allowPlaybackInterruption,omitempty" yaml:"allowPlaybackInterruption,omitempty"`

	// A collection of text responses.
	Texts []string `json:"texts,omitempty" yaml:"texts,omitempty"`
}
