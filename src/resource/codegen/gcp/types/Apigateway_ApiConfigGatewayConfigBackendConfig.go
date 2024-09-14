package types

type Apigateway_ApiConfigGatewayConfigBackendConfig struct {
	/*
	   Google Cloud IAM service account used to sign OIDC tokens for backends that have authentication configured
	   (https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services.configs#backend).
	*/
	GoogleServiceAccount string `json:"googleServiceAccount,omitempty" yaml:"googleServiceAccount,omitempty"`
}
