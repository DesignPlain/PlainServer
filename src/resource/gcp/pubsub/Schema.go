package pubsub

type Schema struct {
	/*
	   The type of the schema definition
	   Default value is `TYPE_UNSPECIFIED`.
	   Possible values are: `TYPE_UNSPECIFIED`, `PROTOCOL_BUFFER`, `AVRO`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The definition of the schema.
	   This should contain a string representing the full definition of the schema
	   that is a valid schema definition of the type specified in type. Changes
	   to the definition commit new [schema revisions](https://cloud.google.com/pubsub/docs/commit-schema-revision).
	   A schema can only have up to 20 revisions, so updates that fail with an
	   error indicating that the limit has been reached require manually
	   [deleting old revisions](https://cloud.google.com/pubsub/docs/delete-schema-revision).
	*/
	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`

	/*
	   The ID to use for the schema, which will become the final component of the schema's resource name.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
