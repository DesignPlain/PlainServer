package types

type Gkebackup_RestorePlanRestoreConfigSelectedApplications struct {
	/*
	   A list of namespaced Kubernetes resources.
	   Structure is documented below.
	*/
	NamespacedNames []Gkebackup_RestorePlanRestoreConfigSelectedApplicationsNamespacedName `json:"namespacedNames,omitempty" yaml:"namespacedNames,omitempty"`
}
