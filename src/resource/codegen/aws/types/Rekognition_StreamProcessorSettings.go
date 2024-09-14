package types

type Rekognition_StreamProcessorSettings struct {
	// Label detection settings to use on a streaming video. See `connected_home`.
	ConnectedHome Rekognition_StreamProcessorSettingsConnectedHome `json:"connectedHome,omitempty" yaml:"connectedHome,omitempty"`

	// Input face recognition parameters for an Amazon Rekognition stream processor. See `face_search`.
	FaceSearch Rekognition_StreamProcessorSettingsFaceSearch `json:"faceSearch,omitempty" yaml:"faceSearch,omitempty"`
}
