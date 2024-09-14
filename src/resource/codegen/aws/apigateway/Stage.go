package apigateway

import types "libds/aws/types"

type Stage struct {
	// Whether a cache cluster is enabled for the stage
	CacheClusterEnabled bool `json:"cacheClusterEnabled,omitempty" yaml:"cacheClusterEnabled,omitempty"`

	// Configuration settings of a canary deployment. See Canary Settings below.
	CanarySettings types.Apigateway_StageCanarySettings `json:"canarySettings,omitempty" yaml:"canarySettings,omitempty"`

	// ID of the deployment that the stage points to
	Deployment string `json:"deployment,omitempty" yaml:"deployment,omitempty"`

	// Description of the stage.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Map that defines the stage variables
	Variables map[string]string `json:"variables,omitempty" yaml:"variables,omitempty"`

	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled bool `json:"xrayTracingEnabled,omitempty" yaml:"xrayTracingEnabled,omitempty"`

	// Enables access logs for the API stage. See Access Log Settings below.
	AccessLogSettings types.Apigateway_StageAccessLogSettings `json:"accessLogSettings,omitempty" yaml:"accessLogSettings,omitempty"`

	// Identifier of a client certificate for the stage.
	ClientCertificateId string `json:"clientCertificateId,omitempty" yaml:"clientCertificateId,omitempty"`

	// Version of the associated API documentation
	DocumentationVersion string `json:"documentationVersion,omitempty" yaml:"documentationVersion,omitempty"`

	// ID of the associated REST API
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Name of the stage
	StageName string `json:"stageName,omitempty" yaml:"stageName,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize string `json:"cacheClusterSize,omitempty" yaml:"cacheClusterSize,omitempty"`
}
