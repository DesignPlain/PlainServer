package appsync

import types "DesignSphere_Server/src/resource/aws/types"

type GraphQLApi struct {
	// Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection/).
	IntrospectionConfig string `json:"introspectionConfig,omitempty" yaml:"introspectionConfig,omitempty"`

	// Nested argument containing logging configuration. Defined below.
	LogConfig types.Appsync_GraphQLApiLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

	// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`

	// The maximum number of resolvers that can be invoked in a single request. The default value is `0` (or unspecified), which will set the limit to `10000`. When specified, the limit value can be between `1` and `10000`. This field will produce a limit error if the operation falls out of bounds.
	ResolverCountLimit int `json:"resolverCountLimit,omitempty" yaml:"resolverCountLimit,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig types.Appsync_GraphQLApiUserPoolConfig `json:"userPoolConfig,omitempty" yaml:"userPoolConfig,omitempty"`

	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled bool `json:"xrayEnabled,omitempty" yaml:"xrayEnabled,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig types.Appsync_GraphQLApiLambdaAuthorizerConfig `json:"lambdaAuthorizerConfig,omitempty" yaml:"lambdaAuthorizerConfig,omitempty"`

	// User-supplied name for the GraphqlApi.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`

	// Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
	Visibility string `json:"visibility,omitempty" yaml:"visibility,omitempty"`

	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders []types.Appsync_GraphQLApiAdditionalAuthenticationProvider `json:"additionalAuthenticationProviders,omitempty" yaml:"additionalAuthenticationProviders,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig types.Appsync_GraphQLApiOpenidConnectConfig `json:"openidConnectConfig,omitempty" yaml:"openidConnectConfig,omitempty"`

	/*
	   The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds.

	   Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
	*/
	QueryDepthLimit int `json:"queryDepthLimit,omitempty" yaml:"queryDepthLimit,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
