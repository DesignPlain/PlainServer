package types

type Firestore_FieldIndexConfig struct {
	/*
	   The indexes to configure on the field. Order or array contains must be specified.
	   Structure is documented below.
	*/
	Indexes []Firestore_FieldIndexConfigIndex `json:"indexes,omitempty" yaml:"indexes,omitempty"`
}
