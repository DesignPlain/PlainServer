package types

type Dataproc_WorkflowTemplatePlacement struct {
	// A selector that chooses target cluster for jobs based on metadata. The selector is evaluated at the time each job is submitted.
	ClusterSelector Dataproc_WorkflowTemplatePlacementClusterSelector `json:"clusterSelector,omitempty" yaml:"clusterSelector,omitempty"`

	// A cluster that is managed by the workflow.
	ManagedCluster Dataproc_WorkflowTemplatePlacementManagedCluster `json:"managedCluster,omitempty" yaml:"managedCluster,omitempty"`
}
