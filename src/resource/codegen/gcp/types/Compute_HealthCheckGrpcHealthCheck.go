package types

type Compute_HealthCheckGrpcHealthCheck struct {
	/*
	   The port number for the health check request.
	   Must be specified if portName and portSpecification are not set
	   or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

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
	   The gRPC service name for the health check.
	   The value of grpcServiceName has the following meanings by convention:
	   - Empty serviceName means the overall status of all services at the backend.
	   - Non-empty serviceName means the health of that gRPC service, as defined by the owner of the service.
	   The grpcServiceName can only be ASCII.
	*/
	GrpcServiceName string `json:"grpcServiceName,omitempty" yaml:"grpcServiceName,omitempty"`
}
