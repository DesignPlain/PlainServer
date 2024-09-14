package types

type Bigquery_JobLoadTimePartitioning struct {
	// Number of milliseconds for which to keep the storage for a partition. A wrapper is used here because 0 is an invalid value.
	ExpirationMs string `json:"expirationMs,omitempty" yaml:"expirationMs,omitempty"`

	/*
	   If not set, the table is partitioned by pseudo column '_PARTITIONTIME'; if set, the table is partitioned by this field.
	   The field must be a top-level TIMESTAMP or DATE field. Its mode must be NULLABLE or REQUIRED.
	   A wrapper is used here because an empty string is an invalid value.
	*/
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   The only type supported is DAY, which will generate one partition per day. Providing an empty string used to cause an error,
	   but in OnePlatform the field will be treated as unset.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
