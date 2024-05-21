package types

type Beyondcorp_getAppConnectionGateway struct {
	// AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.
	AppGateway string `json:"appGateway,omitempty" yaml:"appGateway,omitempty"`

	// Ingress port reserved on the gateways for this AppConnection, if not specified or zero, the default port is 19443.
	IngressPort int `json:"ingressPort,omitempty" yaml:"ingressPort,omitempty"`

	/*
	   The type of hosting used by the gateway. Refer to
	   https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	   for a list of possible values.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Server-defined URI for this resource.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
