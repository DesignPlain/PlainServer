package netapp

import types "DesignSphere_Server/src/resource/gcp/types"

type VolumeReplication struct {
	/*
	   Set to false to stop/break the mirror. Stopping the mirror makes the destination volume read-write
	   and act independently from the source volume.
	   Set to true to enable/resume the mirror. WARNING: Resuming a mirror overwrites any changes
	   done to the destination volume with the content of the source volume.
	*/
	ReplicationEnabled bool `json:"replicationEnabled,omitempty" yaml:"replicationEnabled,omitempty"`

	/*
	   Replication resource state is independent of mirror_state. With enough data, it can take many hours for mirror_state to
	   reach MIRRORED. If you want Terraform to wait for the mirror to finish on create/stop/resume operations, set this
	   parameter to true. Default is false.
	*/
	WaitForMirror bool `json:"waitForMirror,omitempty" yaml:"waitForMirror,omitempty"`

	/*
	   A destination volume is created as part of replication creation. The destination volume will not became under Terraform
	   management unless you import it manually. If you delete the replication, this volume will remain. Setting this parameter
	   to true will delete the -current- destination volume when destroying the replication. If you reversed the replication
	   direction, this will be your former source volume! For production use, it is recommended to keep this parameter false to
	   avoid accidental volume deletion. Handle with care. Default is false.
	*/
	DeleteDestinationVolume bool `json:"deleteDestinationVolume,omitempty" yaml:"deleteDestinationVolume,omitempty"`

	// An description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Only replications with mirror_state=MIRRORED can be stopped. A replication in mirror_state=TRANSFERRING
	   currently receives an update and stopping the update might be undesirable. Set this parameter to true
	   to stop anyway. All data transferred to the destination will be discarded and content of destination
	   volume will remain at the state of the last successful update. Default is false.
	*/
	ForceStopping bool `json:"forceStopping,omitempty" yaml:"forceStopping,omitempty"`

	// Name of region for this resource. The resource needs to be created in the region of the destination volume.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The name of the replication. Needs to be unique per location.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Destination volume parameters.
	   Structure is documented below.
	*/
	DestinationVolumeParameters types.Netapp_VolumeReplicationDestinationVolumeParameters `json:"destinationVolumeParameters,omitempty" yaml:"destinationVolumeParameters,omitempty"`

	/*
	   Labels as key value pairs. Example: `{ "owner": "Bob", "department": "finance", "purpose": "testing" }`

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Specifies the replication interval.
	   Possible values are: `EVERY_10_MINUTES`, `HOURLY`, `DAILY`.
	*/
	ReplicationSchedule string `json:"replicationSchedule,omitempty" yaml:"replicationSchedule,omitempty"`

	// The name of the existing source volume.
	VolumeName string `json:"volumeName,omitempty" yaml:"volumeName,omitempty"`
}
