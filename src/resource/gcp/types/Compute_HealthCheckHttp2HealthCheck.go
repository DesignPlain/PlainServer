package types

type Compute_HealthCheckHttp2HealthCheck struct {
	/*
	   Port name as defined in InstanceGroup#NamedPort#name. If both port and
	   port_name are defined, port takes precedence.
	*/
	PortName string `json:"portName,omitempty" yaml:"portName,omitempty"`

	/*
	   Specifies how port is selected for health checking, can be one of the
	   following values:
	*/
	PortSpecification string `json:"portSpecification,omitempty" yaml:"portSpecification,omitempty"`

	/*
	   Specifies the type of proxy header to append before sending data to the
	   backend.
	   Default value is `NONE`.
	   Possible values are: `NONE`, `PROXY_V1`.
	*/
	ProxyHeader string `json:"proxyHeader,omitempty" yaml:"proxyHeader,omitempty"`

	/*
	   The request path of the HTTP health check request.
	   The default value is /.
	*/
	RequestPath string `json:"requestPath,omitempty" yaml:"requestPath,omitempty"`

	/*
	   The bytes to match against the beginning of the response data. If left empty
	   (the default value), any response will indicate health. The response data
	   can only be ASCII.
	*/
	Response string `json:"response,omitempty" yaml:"response,omitempty"`

	/*
	   The value of the host header in the HTTP health check request.
	   If left empty (default value), the public IP on behalf of which this health
	   check is performed will be used.
	*/
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   The port number for the health check request.
	   Must be specified if portName and portSpecification are not set
	   or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
