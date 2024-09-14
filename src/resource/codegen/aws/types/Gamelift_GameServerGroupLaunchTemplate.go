package types

type Gamelift_GameServerGroupLaunchTemplate struct {
	// A readable identifier for an existing EC2 launch template.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The version of the EC2 launch template to use. If none is set, the default is the first version created.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// A unique identifier for an existing EC2 launch template.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
