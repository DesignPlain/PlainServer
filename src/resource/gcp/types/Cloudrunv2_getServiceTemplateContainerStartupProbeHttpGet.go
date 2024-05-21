package types

type Cloudrunv2_getServiceTemplateContainerStartupProbeHttpGet struct {
	// Custom headers to set in the request. HTTP allows repeated headers.
	HttpHeaders []Cloudrunv2_getServiceTemplateContainerStartupProbeHttpGetHttpHeader `json:"httpHeaders,omitempty" yaml:"httpHeaders,omitempty"`

	// Path to access on the HTTP server. Defaults to '/'.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   Port number to access on the container. Must be in the range 1 to 65535.
	   If not specified, defaults to the same value as container.ports[0].containerPort.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
