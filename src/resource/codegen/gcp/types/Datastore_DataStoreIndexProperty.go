package types

type Datastore_DataStoreIndexProperty struct {
	/*
	   The direction the index should optimize for sorting.
	   Possible values are: `ASCENDING`, `DESCENDING`.
	*/
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	// The property name to index.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
