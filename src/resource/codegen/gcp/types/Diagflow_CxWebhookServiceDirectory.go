package types

type Diagflow_CxWebhookServiceDirectory struct {
	/*
	   The name of Service Directory service.
	   Structure is documented below.
	*/
	GenericWebService Diagflow_CxWebhookServiceDirectoryGenericWebService `json:"genericWebService,omitempty" yaml:"genericWebService,omitempty"`

	// The name of Service Directory service.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
