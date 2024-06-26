package types

type Firestore_FieldIndexConfigIndex struct {
	/*
	   Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
	   be specified.
	   Possible values are: `CONTAINS`.
	*/
	ArrayConfig string `json:"arrayConfig,omitempty" yaml:"arrayConfig,omitempty"`

	/*
	   Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=, !=.
	   Only one of `order` and `arrayConfig` can be specified.
	   Possible values are: `ASCENDING`, `DESCENDING`.
	*/
	Order string `json:"order,omitempty" yaml:"order,omitempty"`

	/*
	   The scope at which a query is run. Collection scoped queries require you specify
	   the collection at query time. Collection group scope allows queries across all
	   collections with the same id.
	   Default value is `COLLECTION`.
	   Possible values are: `COLLECTION`, `COLLECTION_GROUP`.
	*/
	QueryScope string `json:"queryScope,omitempty" yaml:"queryScope,omitempty"`
}
