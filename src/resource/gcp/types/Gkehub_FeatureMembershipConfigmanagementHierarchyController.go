package types

type Gkehub_FeatureMembershipConfigmanagementHierarchyController struct {
	// Whether hierarchical resource quota is enabled in this cluster.
	EnableHierarchicalResourceQuota bool `json:"enableHierarchicalResourceQuota,omitempty" yaml:"enableHierarchicalResourceQuota,omitempty"`

	// Whether pod tree labels are enabled in this cluster.
	EnablePodTreeLabels bool `json:"enablePodTreeLabels,omitempty" yaml:"enablePodTreeLabels,omitempty"`

	// Whether Hierarchy Controller is enabled in this cluster.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
