package sagemaker

type ImageVersion struct {
	// The registry path of the container image on which this image version is based.
	BaseImage string `json:"baseImage,omitempty" yaml:"baseImage,omitempty"`

	// The name of the image. Must be unique to your account.
	ImageName string `json:"imageName,omitempty" yaml:"imageName,omitempty"`
}
