package types

type Pubsub_getTopicSchemaSetting struct {
	// The encoding of messages validated against schema. Default value: "ENCODING_UNSPECIFIED" Possible values: ["ENCODING_UNSPECIFIED", "JSON", "BINARY"]
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	/*
	   The name of the schema that messages published should be
	   validated against. Format is projects/{project}/schemas/{schema}.
	   The value of this field will be _deleted-schema_
	   if the schema has been deleted.
	*/
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
