package types

type Cloudrunv2_getServiceTemplateContainerLivenessProbeGrpc struct {
	/*
	   Port number to access on the container. Number must be in the range 1 to 65535.
	   If not specified, defaults to the same value as container.ports[0].containerPort.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The name of the service to place in the gRPC HealthCheckRequest
	   (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
	   If this is not specified, the default behavior is defined by gRPC.
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
