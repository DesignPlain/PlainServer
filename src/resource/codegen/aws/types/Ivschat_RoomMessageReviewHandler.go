package types

type Ivschat_RoomMessageReviewHandler struct {
	/*
	   The fallback behavior (whether the message
	   is allowed or denied) if the handler does not return a valid response,
	   encounters an error, or times out. Valid values: `ALLOW`, `DENY`.
	*/
	FallbackResult string `json:"fallbackResult,omitempty" yaml:"fallbackResult,omitempty"`

	// ARN of the lambda message review handler function.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
