package networkconnectivity

type Hub struct {
	// An optional description of the hub.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`



	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
