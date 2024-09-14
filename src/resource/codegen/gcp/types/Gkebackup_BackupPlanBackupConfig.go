package types

type Gkebackup_BackupPlanBackupConfig struct {
	/*
	   This flag specifies whether volume data should be backed up when PVCs are
	   included in the scope of a Backup.
	*/
	IncludeVolumeData bool `json:"includeVolumeData,omitempty" yaml:"includeVolumeData,omitempty"`

	/*
	   A list of namespaced Kubernetes Resources.
	   Structure is documented below.
	*/
	SelectedApplications Gkebackup_BackupPlanBackupConfigSelectedApplications `json:"selectedApplications,omitempty" yaml:"selectedApplications,omitempty"`

	/*
	   If set, include just the resources in the listed namespaces.
	   Structure is documented below.
	*/
	SelectedNamespaces Gkebackup_BackupPlanBackupConfigSelectedNamespaces `json:"selectedNamespaces,omitempty" yaml:"selectedNamespaces,omitempty"`

	// If True, include all namespaced resources.
	AllNamespaces bool `json:"allNamespaces,omitempty" yaml:"allNamespaces,omitempty"`

	/*
	   This defines a customer managed encryption key that will be used to encrypt the "config"
	   portion (the Kubernetes resources) of Backups created via this plan.
	   Structure is documented below.
	*/
	EncryptionKey Gkebackup_BackupPlanBackupConfigEncryptionKey `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`

	/*
	   This flag specifies whether Kubernetes Secret resources should be included
	   when they fall into the scope of Backups.
	*/
	IncludeSecrets bool `json:"includeSecrets,omitempty" yaml:"includeSecrets,omitempty"`
}
