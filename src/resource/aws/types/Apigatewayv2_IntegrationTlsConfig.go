package types

type Apigatewayv2_IntegrationTlsConfig struct {
	// If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.
	ServerNameToVerify string `json:"serverNameToVerify,omitempty" yaml:"serverNameToVerify,omitempty"`
}
