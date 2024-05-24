package ssm

import types "DesignSphere_Server/src/resource/aws/types"

type MaintenanceWindowTarget struct {
	// The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	/*
	   The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	   (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	*/
	Targets []types.Ssm_MaintenanceWindowTargetTarget `json:"targets,omitempty" yaml:"targets,omitempty"`

	// The Id of the maintenance window to register the target with.
	WindowId string `json:"windowId,omitempty" yaml:"windowId,omitempty"`

	// The description of the maintenance window target.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the maintenance window target.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation string `json:"ownerInformation,omitempty" yaml:"ownerInformation,omitempty"`
}
