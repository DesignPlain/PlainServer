package types

type Glue_PartitionStorageDescriptorSerDeInfo struct {
	// Usually the class that implements the SerDe. An example is: org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe.
	SerializationLibrary string `json:"serializationLibrary,omitempty" yaml:"serializationLibrary,omitempty"`

	// Name of the SerDe.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
