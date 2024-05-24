package appsync

import types "DesignSphere_Server/src/resource/aws/types"

type Resolver struct {
	// The Caching Config. See Caching Config.
	CachingConfig types.Appsync_ResolverCachingConfig `json:"cachingConfig,omitempty" yaml:"cachingConfig,omitempty"`

	// Data source name.
	DataSource string `json:"dataSource,omitempty" yaml:"dataSource,omitempty"`

	// Field name from the schema defined in the GraphQL API.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate string `json:"requestTemplate,omitempty" yaml:"requestTemplate,omitempty"`

	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig types.Appsync_ResolverSyncConfig `json:"syncConfig,omitempty" yaml:"syncConfig,omitempty"`

	// Type name from the schema defined in the GraphQL API.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// API ID for the GraphQL API.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize int `json:"maxBatchSize,omitempty" yaml:"maxBatchSize,omitempty"`

	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig types.Appsync_ResolverPipelineConfig `json:"pipelineConfig,omitempty" yaml:"pipelineConfig,omitempty"`

	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate string `json:"responseTemplate,omitempty" yaml:"responseTemplate,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime types.Appsync_ResolverRuntime `json:"runtime,omitempty" yaml:"runtime,omitempty"`
}
