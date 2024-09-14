package types

type Gkebackup_BackupPlanBackupConfigSelectedApplications struct {
	/*
	   A list of namespaced Kubernetes resources.
	   Structure is documented below.
	*/
	NamespacedNames []Gkebackup_BackupPlanBackupConfigSelectedApplicationsNamespacedName `json:"namespacedNames,omitempty" yaml:"namespacedNames,omitempty"`
}
