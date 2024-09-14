package types

type Rekognition_StreamProcessorRegionsOfInterestBoundingBox struct {
	// Width of the bounding box as a ratio of the overall image width.
	Width float64 `json:"width,omitempty" yaml:"width,omitempty"`

	// Height of the bounding box as a ratio of the overall image height.
	Height float64 `json:"height,omitempty" yaml:"height,omitempty"`

	// Left coordinate of the bounding box as a ratio of overall image width.
	Left float64 `json:"left,omitempty" yaml:"left,omitempty"`

	// Top coordinate of the bounding box as a ratio of overall image height.
	Top float64 `json:"top,omitempty" yaml:"top,omitempty"`
}
