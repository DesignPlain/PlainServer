package netapp

type VolumeSnapshot struct {
	// Description for the snapshot.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Labels as key value pairs. Example: `{ "owner": "Bob", "department": "finance", "purpose": "testing" }`.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Name of the snapshot location. Snapshots are child resources of volumes and live in the same location.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The name of the snapshot.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the volume to create the snapshot in.
	VolumeName string `json:"volumeName,omitempty" yaml:"volumeName,omitempty"`
}
