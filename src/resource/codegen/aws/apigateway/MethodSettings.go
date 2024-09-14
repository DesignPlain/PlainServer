package apigateway

import types "libds/aws/types"

type MethodSettings struct {
	// Name of the stage
	StageName string `json:"stageName,omitempty" yaml:"stageName,omitempty"`

	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `-/-` for overriding all methods in the stage. Ensure to trim any leading forward slashes in the path (e.g., `trimprefix(aws_api_gateway_resource.example.path, "/")`).
	MethodPath string `json:"methodPath,omitempty" yaml:"methodPath,omitempty"`

	// ID of the REST API
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Settings block, see below.
	Settings types.Apigateway_MethodSettingsSettings `json:"settings,omitempty" yaml:"settings,omitempty"`
}
