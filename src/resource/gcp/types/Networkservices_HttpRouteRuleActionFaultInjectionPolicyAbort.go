package types

type Networkservices_HttpRouteRuleActionFaultInjectionPolicyAbort struct {
	// The HTTP status code used to abort the request.
	HttpStatus int `json:"httpStatus,omitempty" yaml:"httpStatus,omitempty"`

	// The percentage of traffic which will be aborted.
	Percentage int `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
