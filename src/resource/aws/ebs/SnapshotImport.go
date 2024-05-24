package ebs

import types "DesignSphere_Server/src/resource/aws/types"

type SnapshotImport struct {
	// The client-specific data. Detailed below.
	ClientData types.Ebs_SnapshotImportClientData `json:"clientData,omitempty" yaml:"clientData,omitempty"`

	// Information about the disk container. Detailed below.
	DiskContainer types.Ebs_SnapshotImportDiskContainer `json:"diskContainer,omitempty" yaml:"diskContainer,omitempty"`

	// Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore bool `json:"permanentRestore,omitempty" yaml:"permanentRestore,omitempty"`

	// The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`

	// The description string for the import snapshot task.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
	StorageTier string `json:"storageTier,omitempty" yaml:"storageTier,omitempty"`

	// A map of tags to assign to the snapshot.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays int `json:"temporaryRestoreDays,omitempty" yaml:"temporaryRestoreDays,omitempty"`
}
