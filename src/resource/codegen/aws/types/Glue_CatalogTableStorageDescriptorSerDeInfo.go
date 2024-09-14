package types

type Glue_CatalogTableStorageDescriptorSerDeInfo struct {
	// Name of the SerDe.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Usually the class that implements the SerDe. An example is `org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe`.
	SerializationLibrary string `json:"serializationLibrary,omitempty" yaml:"serializationLibrary,omitempty"`
}
