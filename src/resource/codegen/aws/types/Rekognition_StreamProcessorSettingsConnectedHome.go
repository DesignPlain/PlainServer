package types

type Rekognition_StreamProcessorSettingsConnectedHome struct {
	// Specifies what you want to detect in the video, such as people, packages, or pets. The current valid labels you can include in this list are: `PERSON`, `PET`, `PACKAGE`, and `ALL`.
	Labels []string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Minimum confidence required to label an object in the video.
	MinConfidence float64 `json:"minConfidence,omitempty" yaml:"minConfidence,omitempty"`
}
