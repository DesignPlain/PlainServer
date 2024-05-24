package types

type Codecatalyst_DevEnvironmentIdes struct {
	// The name of the IDE. Valid values include Cloud9, IntelliJ, PyCharm, GoLand, and VSCode.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A link to the IDE runtime image. This parameter is not required if the name is VSCode. Values of the runtime can be for example public.ecr.aws/jetbrains/py,public.ecr.aws/jetbrains/go
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`
}
