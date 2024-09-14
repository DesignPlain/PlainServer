package types

type Dataproc_WorkflowTemplatePlacementClusterSelector struct {
	// Required. The cluster labels. Cluster must have all labels to match.
	ClusterLabels map[string]string `json:"clusterLabels,omitempty" yaml:"clusterLabels,omitempty"`

	// The zone where workflow process executes. This parameter does not affect the selection of the cluster. If unspecified, the zone of the first cluster matching the selector is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
