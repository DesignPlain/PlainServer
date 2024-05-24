package codecatalyst

import types "DesignSphere_Server/src/resource/aws/types"

type DevEnvironment struct {
	/*
	   The Amazon EC2 instace type to use for the Dev Environment. Valid values include dev.standard1.small,dev.standard1.medium,dev.standard1.large,dev.standard1.xlarge

	   The following arguments are optional:
	*/
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Information about the amount of storage allocated to the Dev Environment.
	PersistentStorage types.Codecatalyst_DevEnvironmentPersistentStorage `json:"persistentStorage,omitempty" yaml:"persistentStorage,omitempty"`

	// The name of the project in the space.
	ProjectName string `json:"projectName,omitempty" yaml:"projectName,omitempty"`

	// The source repository that contains the branch to clone into the Dev Environment.
	Repositories []types.Codecatalyst_DevEnvironmentRepository `json:"repositories,omitempty" yaml:"repositories,omitempty"`

	// The name of the space.
	SpaceName string `json:"spaceName,omitempty" yaml:"spaceName,omitempty"`

	//
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// Information about the integrated development environment (IDE) configured for a Dev Environment.
	Ides types.Codecatalyst_DevEnvironmentIdes `json:"ides,omitempty" yaml:"ides,omitempty"`

	// The amount of time the Dev Environment will run without any activity detected before stopping, in minutes. Only whole integers are allowed. Dev Environments consume compute minutes when running.
	InactivityTimeoutMinutes int `json:"inactivityTimeoutMinutes,omitempty" yaml:"inactivityTimeoutMinutes,omitempty"`
}
