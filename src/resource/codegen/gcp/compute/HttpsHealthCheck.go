package compute

type HttpsHealthCheck struct {
	/*
	   How often (in seconds) to send a health check. The default value is 5
	   seconds.
	*/
	CheckIntervalSec int `json:"checkIntervalSec,omitempty" yaml:"checkIntervalSec,omitempty"`

	/*
	   A so-far unhealthy instance will be marked healthy after this many
	   consecutive successes. The default value is 2.
	*/
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	/*
	   The value of the host header in the HTTPS health check request. If
	   left empty (default value), the public IP on behalf of which this
	   health check is performed will be used.
	*/
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035.  Specifically, the name must be 1-63 characters long and
	   match the regular expression `a-z?` which means
	   the first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the
	   last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A so-far healthy instance will be marked unhealthy after this many
	   consecutive failures. The default value is 2.
	*/
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The TCP port number for the HTTPS health check request.
	   The default value is 443.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The request path of the HTTPS health check request.
	   The default value is /.
	*/
	RequestPath string `json:"requestPath,omitempty" yaml:"requestPath,omitempty"`

	/*
	   How long (in seconds) to wait before claiming failure.
	   The default value is 5 seconds.  It is invalid for timeoutSec to have
	   greater value than checkIntervalSec.
	*/
	TimeoutSec int `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`
}
