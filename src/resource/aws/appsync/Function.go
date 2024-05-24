package appsync

import types "DesignSphere_Server/src/resource/aws/types"

type Function struct {
	// ID of the associated AppSync API.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Function name. The function name does not have to be unique.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate string `json:"requestMappingTemplate,omitempty" yaml:"requestMappingTemplate,omitempty"`

	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig types.Appsync_FunctionSyncConfig `json:"syncConfig,omitempty" yaml:"syncConfig,omitempty"`

	// Function data source name.
	DataSource string `json:"dataSource,omitempty" yaml:"dataSource,omitempty"`

	// Function description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Version of the request mapping template. Currently the supported value is `2018-05-29`. Does not apply when specifying `code`.
	FunctionVersion string `json:"functionVersion,omitempty" yaml:"functionVersion,omitempty"`

	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize int `json:"maxBatchSize,omitempty" yaml:"maxBatchSize,omitempty"`

	// Function response mapping template.
	ResponseMappingTemplate string `json:"responseMappingTemplate,omitempty" yaml:"responseMappingTemplate,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime types.Appsync_FunctionRuntime `json:"runtime,omitempty" yaml:"runtime,omitempty"`
}
