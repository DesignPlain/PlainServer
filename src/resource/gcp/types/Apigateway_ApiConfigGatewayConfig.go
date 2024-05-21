package types



type Apigateway_ApiConfigGatewayConfig struct {
	/*
	   Backend settings that are applied to all backends of the Gateway.
	   Structure is documented below.
	*/
	BackendConfig Apigateway_ApiConfigGatewayConfigBackendConfig `json:"backendConfig,omitempty" yaml:"backendConfig,omitempty"`
}
