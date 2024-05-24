package cloud9

type EnvironmentMembership struct {
	// The ID of the environment that contains the environment member you want to add.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	// The type of environment member permissions you want to associate with this environment member. Allowed values are `read-only` and `read-write` .
	Permissions string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// The Amazon Resource Name (ARN) of the environment member you want to add.
	UserArn string `json:"userArn,omitempty" yaml:"userArn,omitempty"`
}
