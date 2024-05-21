package types

type Gkebackup_BackupPlanBackupConfigSelectedApplicationsNamespacedName struct {
	// The name of a Kubernetes Resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The namespace of a Kubernetes Resource.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
