package types

type Gkebackup_RestorePlanRestoreConfig struct {
	/*
	   A list of selected ProtectedApplications to restore.
	   The listed ProtectedApplications and all the resources
	   to which they refer will be restored.
	   Structure is documented below.
	*/
	SelectedApplications Gkebackup_RestorePlanRestoreConfigSelectedApplications `json:"selectedApplications,omitempty" yaml:"selectedApplications,omitempty"`

	/*
	   A list of selected namespaces to restore from the Backup.
	   The listed Namespaces and all resources contained in them will be restored.
	   Structure is documented below.
	*/
	SelectedNamespaces Gkebackup_RestorePlanRestoreConfigSelectedNamespaces `json:"selectedNamespaces,omitempty" yaml:"selectedNamespaces,omitempty"`

	/*
	   A list of transformation rules to be applied against Kubernetes
	   resources as they are selected for restoration from a Backup.
	   Rules are executed in order defined - this order matters,
	   as changes made by a rule may impact the filtering logic of subsequent
	   rules. An empty list means no transformation will occur.
	   Structure is documented below.
	*/
	TransformationRules []Gkebackup_RestorePlanRestoreConfigTransformationRule `json:"transformationRules,omitempty" yaml:"transformationRules,omitempty"`

	/*
	   Specifies the mechanism to be used to restore volume data.
	   This should be set to a value other than `NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED`
	   if the `namespacedResourceRestoreScope` is anything other than `noNamespaces`.
	   If not specified, it will be treated as `NO_VOLUME_DATA_RESTORATION`.
	   See https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#VolumeDataRestorePolicy
	   for more information on each policy option.
	   Possible values are: `RESTORE_VOLUME_DATA_FROM_BACKUP`, `REUSE_VOLUME_HANDLE_FROM_BACKUP`, `NO_VOLUME_DATA_RESTORATION`.
	*/
	VolumeDataRestorePolicy string `json:"volumeDataRestorePolicy,omitempty" yaml:"volumeDataRestorePolicy,omitempty"`

	/*
	   If True, restore all namespaced resources in the Backup.
	   Setting this field to False will result in an error.
	*/
	AllNamespaces bool `json:"allNamespaces,omitempty" yaml:"allNamespaces,omitempty"`

	/*
	   Defines the behavior for handling the situation where cluster-scoped resources
	   being restored already exist in the target cluster.
	   This MUST be set to a value other than `CLUSTER_RESOURCE_CONFLICT_POLICY_UNSPECIFIED`
	   if `clusterResourceRestoreScope` is anyting other than `noGroupKinds`.
	   See https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#clusterresourceconflictpolicy
	   for more information on each policy option.
	   Possible values are: `USE_EXISTING_VERSION`, `USE_BACKUP_VERSION`.
	*/
	ClusterResourceConflictPolicy string `json:"clusterResourceConflictPolicy,omitempty" yaml:"clusterResourceConflictPolicy,omitempty"`

	/*
	   Identifies the cluster-scoped resources to restore from the Backup.
	   Structure is documented below.
	*/
	ClusterResourceRestoreScope Gkebackup_RestorePlanRestoreConfigClusterResourceRestoreScope `json:"clusterResourceRestoreScope,omitempty" yaml:"clusterResourceRestoreScope,omitempty"`

	/*
	   A list of selected namespaces excluded from restoration.
	   All namespaces except those in this list will be restored.
	   Structure is documented below.
	*/
	ExcludedNamespaces Gkebackup_RestorePlanRestoreConfigExcludedNamespaces `json:"excludedNamespaces,omitempty" yaml:"excludedNamespaces,omitempty"`

	/*
	   Defines the behavior for handling the situation where sets of namespaced resources
	   being restored already exist in the target cluster.
	   This MUST be set to a value other than `NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED`
	   if the `namespacedResourceRestoreScope` is anything other than `noNamespaces`.
	   See https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#namespacedresourcerestoremode
	   for more information on each mode.
	   Possible values are: `DELETE_AND_RESTORE`, `FAIL_ON_CONFLICT`.
	*/
	NamespacedResourceRestoreMode string `json:"namespacedResourceRestoreMode,omitempty" yaml:"namespacedResourceRestoreMode,omitempty"`

	/*
	   Do not restore any namespaced resources if set to "True".
	   Specifying this field to "False" is not allowed.
	*/
	NoNamespaces bool `json:"noNamespaces,omitempty" yaml:"noNamespaces,omitempty"`
}
