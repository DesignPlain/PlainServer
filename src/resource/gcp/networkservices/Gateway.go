package networkservices

type Gateway struct {
	/*
	   A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection.
	   This feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.
	*/
	CertificateUrls []string `json:"certificateUrls,omitempty" yaml:"certificateUrls,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Short name of the Gateway resource to be created.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The relative resource name identifying the VPC network that is using this configuration.
	   For example: `projects/-/global/networks/network-1`.
	   Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Zero or one IPv4-address on which the Gateway will receive the traffic. When no address is provided,
	   an IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'.
	   Gateways of type 'OPEN_MESH' listen on 0.0.0.0.
	*/
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	/*
	   When deleting a gateway of type 'SECURE_WEB_GATEWAY', this boolean option will also delete auto generated router by the gateway creation.
	   If there is no other gateway of type 'SECURE_WEB_GATEWAY' remaining for that region and network it will be deleted.
	*/
	DeleteSwgAutogenRouterOnDestroy bool `json:"deleteSwgAutogenRouterOnDestroy,omitempty" yaml:"deleteSwgAutogenRouterOnDestroy,omitempty"`

	/*
	   Set of label tags associated with the Gateway resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated.
	   If empty, TLS termination is disabled.
	*/
	ServerTlsPolicy string `json:"serverTlsPolicy,omitempty" yaml:"serverTlsPolicy,omitempty"`

	/*
	   The relative resource name identifying the subnetwork in which this SWG is allocated.
	   For example: `projects/-/regions/us-central1/subnetworks/network-1`.
	   Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   The location of the gateway.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Immutable. Scope determines how configuration across multiple Gateway instances are merged.
	   The configuration for multiple Gateway instances with the same scope will be merged as presented as
	   a single coniguration to the proxy/load balancer.
	   Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
	*/
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	/*
	   A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections.
	   For example: `projects/-/locations/-/gatewaySecurityPolicies/swg-policy`.
	   This policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.
	*/
	GatewaySecurityPolicy string `json:"gatewaySecurityPolicy,omitempty" yaml:"gatewaySecurityPolicy,omitempty"`

	/*
	   One or more port numbers (1-65535), on which the Gateway will receive traffic.
	   The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are
	   limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.
	*/
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Immutable. The type of the customer-managed gateway. Possible values are: - OPEN_MESH - SECURE_WEB_GATEWAY.
	   Possible values are: `TYPE_UNSPECIFIED`, `OPEN_MESH`, `SECURE_WEB_GATEWAY`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
