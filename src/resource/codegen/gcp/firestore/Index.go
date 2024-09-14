package firestore

import types "libds/gcp/types"

type Index struct {
	/*
	   The API scope at which a query is run.
	   Default value is `ANY_API`.
	   Possible values are: `ANY_API`, `DATASTORE_MODE_API`.
	*/
	ApiScope string `json:"apiScope,omitempty" yaml:"apiScope,omitempty"`

	// The collection being indexed.
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`

	// The Firestore database id. Defaults to `"(default)"`.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   The fields supported by this index. The last field entry is always for
	   the field path `__name__`. If, on creation, `__name__` was not
	   specified as the last field, it will be added automatically with the
	   same direction as that of the last field defined. If the final field
	   in a composite index is not directional, the `__name__` will be
	   ordered `"ASCENDING"` (unless explicitly specified otherwise).
	   Structure is documented below.
	*/
	Fields []types.Firestore_IndexField `json:"fields,omitempty" yaml:"fields,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The scope at which a query is run.
	   Default value is `COLLECTION`.
	   Possible values are: `COLLECTION`, `COLLECTION_GROUP`, `COLLECTION_RECURSIVE`.
	*/
	QueryScope string `json:"queryScope,omitempty" yaml:"queryScope,omitempty"`
}
