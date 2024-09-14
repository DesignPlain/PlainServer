package apigateway

import types "libds/gcp/types"

type ApiConfig struct {
	/*
	   The API to attach the config to.


	   - - -
	*/
	Api string `json:"api,omitempty" yaml:"api,omitempty"`

	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId string `json:"apiConfigId,omitempty" yaml:"apiConfigId,omitempty"`

	/*
	   Creates a unique name beginning with the
	   specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
	*/
	ApiConfigIdPrefix string `json:"apiConfigIdPrefix,omitempty" yaml:"apiConfigIdPrefix,omitempty"`

	// A user-visible name for the API.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Immutable. Gateway specific configuration.
	   If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	   Structure is documented below.
	*/
	GatewayConfig types.Apigateway_ApiConfigGatewayConfig `json:"gatewayConfig,omitempty" yaml:"gatewayConfig,omitempty"`

	/*
	   gRPC service definition files. If specified, openapiDocuments must not be included.
	   Structure is documented below.
	*/
	GrpcServices []types.Apigateway_ApiConfigGrpcService `json:"grpcServices,omitempty" yaml:"grpcServices,omitempty"`

	/*
	   Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	   If multiple files are specified, the files are merged with the following rules: - All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. - Repeated fields are concatenated. - Singular embedded messages are merged using these rules for nested fields.
	   Structure is documented below.
	*/
	ManagedServiceConfigs []types.Apigateway_ApiConfigManagedServiceConfig `json:"managedServiceConfigs,omitempty" yaml:"managedServiceConfigs,omitempty"`

	/*
	   OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	   Structure is documented below.
	*/
	OpenapiDocuments []types.Apigateway_ApiConfigOpenapiDocument `json:"openapiDocuments,omitempty" yaml:"openapiDocuments,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
