package types

type Gamelift_FleetRuntimeConfigurationServerProcess struct {
	// Location of the server executable in a game build. All game builds are installed on instances at the root : for Windows instances `C:\game`, and for Linux instances `/local/game`.
	LaunchPath string `json:"launchPath,omitempty" yaml:"launchPath,omitempty"`

	// Optional list of parameters to pass to the server executable on launch.
	Parameters string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Number of server processes using this configuration to run concurrently on an instance.
	ConcurrentExecutions int `json:"concurrentExecutions,omitempty" yaml:"concurrentExecutions,omitempty"`
}
