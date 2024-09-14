package compute

import types "libds/gcp/types"

type RegionHealthCheck struct {
	/*
	   A so-far unhealthy instance will be marked healthy after this many
	   consecutive successes. The default value is 2.
	*/
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The Region in which the created health check should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	TcpHealthCheck types.Compute_RegionHealthCheckTcpHealthCheck `json:"tcpHealthCheck,omitempty" yaml:"tcpHealthCheck,omitempty"`

	/*
	   How often (in seconds) to send a health check. The default value is 5
	   seconds.
	*/
	CheckIntervalSec int `json:"checkIntervalSec,omitempty" yaml:"checkIntervalSec,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	GrpcHealthCheck types.Compute_RegionHealthCheckGrpcHealthCheck `json:"grpcHealthCheck,omitempty" yaml:"grpcHealthCheck,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	Http2HealthCheck types.Compute_RegionHealthCheckHttp2HealthCheck `json:"http2HealthCheck,omitempty" yaml:"http2HealthCheck,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	HttpHealthCheck types.Compute_RegionHealthCheckHttpHealthCheck `json:"httpHealthCheck,omitempty" yaml:"httpHealthCheck,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	HttpsHealthCheck types.Compute_RegionHealthCheckHttpsHealthCheck `json:"httpsHealthCheck,omitempty" yaml:"httpsHealthCheck,omitempty"`

	/*
	   Configure logging on this health check.
	   Structure is documented below.
	*/
	LogConfig types.Compute_RegionHealthCheckLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

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
	   A nested object resource
	   Structure is documented below.
	*/
	SslHealthCheck types.Compute_RegionHealthCheckSslHealthCheck `json:"sslHealthCheck,omitempty" yaml:"sslHealthCheck,omitempty"`

	/*
	   How long (in seconds) to wait before claiming failure.
	   The default value is 5 seconds.  It is invalid for timeoutSec to have
	   greater value than checkIntervalSec.
	*/
	TimeoutSec int `json:"timeoutSec,omitempty" yaml:"timeoutSec,omitempty"`

	/*
	   A so-far healthy instance will be marked unhealthy after this many
	   consecutive failures. The default value is 2.
	*/
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
