package types

type Pipes_PipeTargetParametersEventbridgeEventBusParameters struct {
	// List of AWS resources, identified by Amazon Resource Name (ARN), which the event primarily concerns. Any number, including zero, may be present.
	Resources []string `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Source resource of the pipe (typically an ARN).
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// The time stamp of the event, per RFC3339. If no time stamp is provided, the time stamp of the PutEvents call is used. This is the JSON path to the field in the event e.g. $.detail.timestamp
	Time string `json:"time,omitempty" yaml:"time,omitempty"`

	// A free-form string, with a maximum of 128 characters, used to decide what fields to expect in the event detail.
	DetailType string `json:"detailType,omitempty" yaml:"detailType,omitempty"`

	// The URL subdomain of the endpoint. For example, if the URL for Endpoint is https://abcde.veo.endpoints.event.amazonaws.com, then the EndpointId is abcde.veo.
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`
}
