package types

type Cloudrun_getServiceTemplateSpecContainerEnvValueFromSecretKeyRef struct {
	/*
	   A Cloud Secret Manager secret version. Must be 'latest' for the latest
	   version or an integer for a specific version.
	*/
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The name of the Cloud Run Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
