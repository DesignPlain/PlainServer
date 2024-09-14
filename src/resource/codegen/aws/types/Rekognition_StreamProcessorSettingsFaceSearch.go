package types

type Rekognition_StreamProcessorSettingsFaceSearch struct {
	// ID of a collection that contains faces that you want to search for.
	CollectionId string `json:"collectionId,omitempty" yaml:"collectionId,omitempty"`

	// Minimum face match confidence score that must be met to return a result for a recognized face.
	FaceMatchThreshold float64 `json:"faceMatchThreshold,omitempty" yaml:"faceMatchThreshold,omitempty"`
}
