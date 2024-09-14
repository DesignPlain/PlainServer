package mskconnect

type WorkerConfiguration struct {
	// A summary description of the worker configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the worker configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.

	   The following arguments are optional:
	*/
	PropertiesFileContent string `json:"propertiesFileContent,omitempty" yaml:"propertiesFileContent,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
