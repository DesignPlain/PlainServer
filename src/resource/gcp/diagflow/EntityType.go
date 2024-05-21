package diagflow

import types "DesignSphere_Server/src/resource/gcp/types"

type EntityType struct {
	/*
	   Indicates the kind of entity type.
	   - KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	   - KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	   types can contain references to other entity types (with or without aliases).
	   - KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	   Possible values are: `KIND_MAP`, `KIND_LIST`, `KIND_REGEXP`.


	   - - -
	*/
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of this entity type to be displayed on the console.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction bool `json:"enableFuzzyExtraction,omitempty" yaml:"enableFuzzyExtraction,omitempty"`

	/*
	   The collection of entity entries associated with the entity type.
	   Structure is documented below.
	*/
	Entities []types.Diagflow_EntityTypeEntity `json:"entities,omitempty" yaml:"entities,omitempty"`
}
