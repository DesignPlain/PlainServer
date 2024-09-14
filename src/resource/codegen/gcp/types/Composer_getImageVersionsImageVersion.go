package types

type Composer_getImageVersionsImageVersion struct {
	// The string identifier of the image version, in the form: "composer-x.y.z-airflow-a.b.c"
	ImageVersionId string `json:"imageVersionId,omitempty" yaml:"imageVersionId,omitempty"`

	// Supported python versions for this image version
	SupportedPythonVersions []string `json:"supportedPythonVersions,omitempty" yaml:"supportedPythonVersions,omitempty"`
}
