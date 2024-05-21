package firestore

import types "DesignSphere_Server/src/resource/gcp/types"

type Field struct {
	// The id of the collection group to configure.
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`

	// The Firestore database id. Defaults to `"(default)"`.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   The id of the field to configure.


	   - - -
	*/
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   The single field index configuration for this field.
	   Creating an index configuration for this field will override any inherited configuration with the
	   indexes specified. Configuring the index configuration with an empty block disables all indexes on
	   the field.
	   Structure is documented below.
	*/
	IndexConfig types.Firestore_FieldIndexConfig `json:"indexConfig,omitempty" yaml:"indexConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The TTL configuration for this Field. If set to an empty block (i.e. `ttl_config {}`), a TTL policy is configured based on the field. If unset, a TTL policy is not configured (or will be disabled upon updating the resource).
	   Structure is documented below.
	*/
	TtlConfig types.Firestore_FieldTtlConfig `json:"ttlConfig,omitempty" yaml:"ttlConfig,omitempty"`
}
