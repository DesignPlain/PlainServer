package types

type Vertex_AiEndpointDeployedModelPrivateEndpoint struct {
	/*
	   (Output)
	   Output only. Http(s) path to send explain requests.
	*/
	ExplainHttpUri string `json:"explainHttpUri,omitempty" yaml:"explainHttpUri,omitempty"`

	/*
	   (Output)
	   Output only. Http(s) path to send health check requests.
	*/
	HealthHttpUri string `json:"healthHttpUri,omitempty" yaml:"healthHttpUri,omitempty"`

	/*
	   (Output)
	   Output only. Http(s) path to send prediction requests.
	*/
	PredictHttpUri string `json:"predictHttpUri,omitempty" yaml:"predictHttpUri,omitempty"`

	/*
	   (Output)
	   Output only. The name of the service attachment resource. Populated if private service connect is enabled.
	*/
	ServiceAttachment string `json:"serviceAttachment,omitempty" yaml:"serviceAttachment,omitempty"`
}
