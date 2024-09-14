package types

type Datastream_StreamDestinationConfigGcsDestinationConfigJsonFileFormat struct {
	/*
	   Compression of the loaded JSON file.
	   Possible values are: `NO_COMPRESSION`, `GZIP`.
	*/
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	/*
	   The schema file format along JSON data files.
	   Possible values are: `NO_SCHEMA_FILE`, `AVRO_SCHEMA_FILE`.
	*/
	SchemaFileFormat string `json:"schemaFileFormat,omitempty" yaml:"schemaFileFormat,omitempty"`
}
