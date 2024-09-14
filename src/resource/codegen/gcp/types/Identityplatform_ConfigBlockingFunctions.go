package types

type Identityplatform_ConfigBlockingFunctions struct {
	/*
	   The user credentials to include in the JWT payload that is sent to the registered Blocking Functions.
	   Structure is documented below.
	*/
	ForwardInboundCredentials Identityplatform_ConfigBlockingFunctionsForwardInboundCredentials `json:"forwardInboundCredentials,omitempty" yaml:"forwardInboundCredentials,omitempty"`

	/*
	   Map of Trigger to event type. Key should be one of the supported event types: "beforeCreate", "beforeSignIn".
	   Structure is documented below.
	*/
	Triggers []Identityplatform_ConfigBlockingFunctionsTrigger `json:"triggers,omitempty" yaml:"triggers,omitempty"`
}
