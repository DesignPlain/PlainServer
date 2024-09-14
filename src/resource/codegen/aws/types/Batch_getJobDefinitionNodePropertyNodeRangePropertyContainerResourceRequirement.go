package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerResourceRequirement struct {
	// The type of resource to assign to a container. The supported resources include `GPU`, `MEMORY`, and `VCPU`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The quantity of the specified resource to reserve for the container.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
