package types

type Gkebackup_RestorePlanRestoreConfigClusterResourceRestoreScope struct {
	/*
	   If True, no cluster-scoped resources will be restored.
	   Mutually exclusive to any other field in `clusterResourceRestoreScope`.
	*/
	NoGroupKinds bool `json:"noGroupKinds,omitempty" yaml:"noGroupKinds,omitempty"`

	/*
	   A list of cluster-scoped resource group kinds to restore from the backup.
	   If specified, only the selected resources will be restored.
	   Mutually exclusive to any other field in the `clusterResourceRestoreScope`.
	   Structure is documented below.
	*/
	SelectedGroupKinds []Gkebackup_RestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKind `json:"selectedGroupKinds,omitempty" yaml:"selectedGroupKinds,omitempty"`

	/*
	   If True, all valid cluster-scoped resources will be restored.
	   Mutually exclusive to any other field in `clusterResourceRestoreScope`.
	*/
	AllGroupKinds bool `json:"allGroupKinds,omitempty" yaml:"allGroupKinds,omitempty"`

	/*
	   A list of cluster-scoped resource group kinds to NOT restore from the backup.
	   If specified, all valid cluster-scoped resources will be restored except
	   for those specified in the list.
	   Mutually exclusive to any other field in `clusterResourceRestoreScope`.
	   Structure is documented below.
	*/
	ExcludedGroupKinds []Gkebackup_RestorePlanRestoreConfigClusterResourceRestoreScopeExcludedGroupKind `json:"excludedGroupKinds,omitempty" yaml:"excludedGroupKinds,omitempty"`
}
