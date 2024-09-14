package types

type Rekognition_StreamProcessorRegionsOfInterest struct {
	// Box representing a region of interest on screen. Only 1 per region is allowed. See `bounding_box`.
	BoundingBox Rekognition_StreamProcessorRegionsOfInterestBoundingBox `json:"boundingBox,omitempty" yaml:"boundingBox,omitempty"`

	// Shape made up of up to 10 Point objects to define a region of interest. See `polygon`.
	Polygons []Rekognition_StreamProcessorRegionsOfInterestPolygon `json:"polygons,omitempty" yaml:"polygons,omitempty"`
}
