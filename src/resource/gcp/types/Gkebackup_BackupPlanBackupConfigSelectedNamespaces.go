package types

type Gkebackup_BackupPlanBackupConfigSelectedNamespaces struct {
	// A list of Kubernetes Namespaces.
	Namespaces []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}
