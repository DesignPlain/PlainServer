package types

type Vertex_AiFeatureOnlineStoreDedicatedServingEndpoint struct {
	/*
	   Private service connect config.
	   Structure is documented below.
	*/
	PrivateServiceConnectConfig Vertex_AiFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty" yaml:"privateServiceConnectConfig,omitempty"`

	/*
	   (Output)
	   Domain name to use for this FeatureOnlineStore
	*/
	PublicEndpointDomainName string `json:"publicEndpointDomainName,omitempty" yaml:"publicEndpointDomainName,omitempty"`

	/*
	   (Output)
	   Name of the service attachment resource. Applicable only if private service connect is enabled and after FeatureViewSync is created.
	*/
	ServiceAttachment string `json:"serviceAttachment,omitempty" yaml:"serviceAttachment,omitempty"`
}
