package alloydb

import types "libds/gcp/types"

type Backup struct {
	/*
	   The location where the alloydb backup should reside.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The ID of the alloydb backup.
	BackupId string `json:"backupId,omitempty" yaml:"backupId,omitempty"`

	// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// User-provided description of the backup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
	   Structure is documented below.
	*/
	EncryptionConfig types.Alloydb_BackupEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	/*
	   User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// User-settable and human-readable display name for the Backup.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The backup type, which suggests the trigger for the backup.
	   Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
