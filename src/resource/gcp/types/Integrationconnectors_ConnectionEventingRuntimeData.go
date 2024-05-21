package types

type Integrationconnectors_ConnectionEventingRuntimeData struct {
	// Events listener endpoint. The value will populated after provisioning the events listener.
	EventsListenerEndpoint string `json:"eventsListenerEndpoint,omitempty" yaml:"eventsListenerEndpoint,omitempty"`

	/*
	   (Output)
	   Current status of eventing.
	   Structure is documented below.
	*/
	Statuses []Integrationconnectors_ConnectionEventingRuntimeDataStatus `json:"statuses,omitempty" yaml:"statuses,omitempty"`
}
