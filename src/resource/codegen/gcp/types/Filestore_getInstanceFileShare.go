package types

type Filestore_getInstanceFileShare struct {
	/*
	   The resource name of the backup, in the format
	   projects/{projectId}/locations/{locationId}/backups/{backupId},
	   that this file share has been restored from.
	*/
	SourceBackup string `json:"sourceBackup,omitempty" yaml:"sourceBackup,omitempty"`

	/*
	   File share capacity in GiB. This must be at least 1024 GiB
	   for the standard tier, or 2560 GiB for the premium tier.
	*/
	CapacityGb int `json:"capacityGb,omitempty" yaml:"capacityGb,omitempty"`

	/*
	   The name of a Filestore instance.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Nfs Export Options. There is a limit of 10 export options per file share.
	NfsExportOptions []Filestore_getInstanceFileShareNfsExportOption `json:"nfsExportOptions,omitempty" yaml:"nfsExportOptions,omitempty"`
}
