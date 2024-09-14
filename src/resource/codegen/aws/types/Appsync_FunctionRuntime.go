package types

type Appsync_FunctionRuntime struct {
	// The name of the runtime to use. Currently, the only allowed value is `APPSYNC_JS`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The version of the runtime to use. Currently, the only allowed version is `1.0.0`.
	RuntimeVersion string `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`
}
