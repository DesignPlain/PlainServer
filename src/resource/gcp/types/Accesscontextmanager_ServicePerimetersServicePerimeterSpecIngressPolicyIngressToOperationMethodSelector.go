package types

type Accesscontextmanager_ServicePerimetersServicePerimeterSpecIngressPolicyIngressToOperationMethodSelector struct {
	/*
	   Value for permission should be a valid Cloud IAM permission for the
	   corresponding `serviceName` in `ApiOperation`.
	*/
	Permission string `json:"permission,omitempty" yaml:"permission,omitempty"`

	/*
	   Value for `method` should be a valid method name for the corresponding
	   `serviceName` in `ApiOperation`. If `-` used as value for method,
	   then ALL methods and permissions are allowed.
	*/
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}
