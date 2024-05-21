package netapp

type BackupVault struct {
	/*
	   Labels as key value pairs. Example: `{ "owner": "Bob", "department": "finance", "purpose": "testing" }`.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Location (region) of the backup vault.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The resource name of the backup vault. Needs to be unique per location.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
