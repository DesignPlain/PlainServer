package datastore

import types "libds/gcp/types"

type DataStoreIndex struct {
	/*
	   Policy for including ancestors in the index.
	   Default value is `NONE`.
	   Possible values are: `NONE`, `ALL_ANCESTORS`.
	*/
	Ancestor string `json:"ancestor,omitempty" yaml:"ancestor,omitempty"`

	/*
	   The entity kind which the index applies to.


	   - - -
	*/
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   An ordered list of properties to index on.
	   Structure is documented below.
	*/
	Properties []types.Datastore_DataStoreIndexProperty `json:"properties,omitempty" yaml:"properties,omitempty"`
}
