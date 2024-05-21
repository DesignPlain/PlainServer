package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig struct {
	// Output only. The name of the Instance Group Manager for this group.
	InstanceGroupManagerName string `json:"instanceGroupManagerName,omitempty" yaml:"instanceGroupManagerName,omitempty"`

	// Output only. The name of the Instance Template used for the Managed Instance Group.
	InstanceTemplateName string `json:"instanceTemplateName,omitempty" yaml:"instanceTemplateName,omitempty"`
}
