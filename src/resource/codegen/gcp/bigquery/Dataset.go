package bigquery

import types "libds/gcp/types"

type Dataset struct {
	/*
	   An array of objects that define dataset access for one or more entities.
	   Structure is documented below.
	*/
	Accesses []types.Bigquery_DatasetAccess `json:"accesses,omitempty" yaml:"accesses,omitempty"`

	// Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
	MaxTimeTravelHours string `json:"maxTimeTravelHours,omitempty" yaml:"maxTimeTravelHours,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A unique ID for this dataset, without the project name. The ID
	   must contain only letters (a-z, A-Z), numbers (0-9), or
	   underscores (_). The maximum length is 1,024 characters.


	   - - -
	*/
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	/*
	   The default partition expiration for all partitioned tables in
	   the dataset, in milliseconds.

	   Once this property is set, all newly-created partitioned tables in
	   the dataset will have an `expirationMs` property in the `timePartitioning`
	   settings set to this value, and changing the value will only
	   affect new tables, not existing ones. The storage in a partition will
	   have an expiration time of its partition time plus this value.
	   Setting this property overrides the use of `defaultTableExpirationMs`
	   for partitioned tables: only one of `defaultTableExpirationMs` and
	   `defaultPartitionExpirationMs` will be used for any new partitioned
	   table. If you provide an explicit `timePartitioning.expirationMs` when
	   creating or updating a partitioned table, that value takes precedence
	   over the default partition expiration time indicated by this property.
	*/
	DefaultPartitionExpirationMs int `json:"defaultPartitionExpirationMs,omitempty" yaml:"defaultPartitionExpirationMs,omitempty"`

	// A descriptive name for the dataset
	FriendlyName string `json:"friendlyName,omitempty" yaml:"friendlyName,omitempty"`

	/*
	   The labels associated with this dataset. You can use these to
	   organize and group your datasets.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Defines the default collation specification of future tables created
	   in the dataset. If a table is created in this dataset without table-level
	   default collation, then the table inherits the dataset default collation,
	   which is applied to the string fields that do not have explicit collation
	   specified. A change to this field affects only tables created afterwards,
	   and does not alter the existing tables.
	   The following values are supported:
	   - 'und:ci': undetermined locale, case insensitive.
	   - '': empty string. Default to case-sensitive behavior.
	*/
	DefaultCollation string `json:"defaultCollation,omitempty" yaml:"defaultCollation,omitempty"`

	/*
	   Information about the external metadata storage where the dataset is defined.
	   Structure is documented below.
	*/
	ExternalDatasetReference types.Bigquery_DatasetExternalDatasetReference `json:"externalDatasetReference,omitempty" yaml:"externalDatasetReference,omitempty"`

	/*
	   TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
	   By default, this is FALSE, which means the dataset and its table names are
	   case-sensitive. This field does not affect routine references.
	*/
	IsCaseInsensitive bool `json:"isCaseInsensitive,omitempty" yaml:"isCaseInsensitive,omitempty"`

	/*
	   Specifies the storage billing model for the dataset.
	   Set this flag value to LOGICAL to use logical bytes for storage billing,
	   or to PHYSICAL to use physical bytes instead.
	   LOGICAL is the default if this flag isn't specified.
	*/
	StorageBillingModel string `json:"storageBillingModel,omitempty" yaml:"storageBillingModel,omitempty"`

	/*
	   The geographic location where the dataset should reside.
	   See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).

	   There are two types of locations, regional or multi-regional. A regional
	   location is a specific geographic place, such as Tokyo, and a multi-regional
	   location is a large geographic area, such as the United States, that
	   contains at least two geographic places.

	   The default value is multi-regional location `US`.
	   Changing this forces a new resource to be created.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The default encryption key for all tables in the dataset. Once this property is set,
	   all newly-created partitioned tables in the dataset will have encryption key set to
	   this value, unless table creation request (or query) overrides the key.
	   Structure is documented below.
	*/
	DefaultEncryptionConfiguration types.Bigquery_DatasetDefaultEncryptionConfiguration `json:"defaultEncryptionConfiguration,omitempty" yaml:"defaultEncryptionConfiguration,omitempty"`

	/*
	   The default lifetime of all tables in the dataset, in milliseconds.
	   The minimum value is 3600000 milliseconds (one hour).

	   Once this property is set, all newly-created tables in the dataset
	   will have an `expirationTime` property set to the creation time plus
	   the value in this property, and changing the value will only affect
	   new tables, not existing ones. When the `expirationTime` for a given
	   table is reached, that table will be deleted automatically.
	   If a table's `expirationTime` is modified or removed before the
	   table expires, or if you provide an explicit `expirationTime` when
	   creating a table, that value takes precedence over the default
	   expiration time indicated by this property.
	*/
	DefaultTableExpirationMs int `json:"defaultTableExpirationMs,omitempty" yaml:"defaultTableExpirationMs,omitempty"`

	/*
	   If set to `true`, delete all the tables in the
	   dataset when destroying the resource; otherwise,
	   destroying the resource will fail if tables are present.
	*/
	DeleteContentsOnDestroy bool `json:"deleteContentsOnDestroy,omitempty" yaml:"deleteContentsOnDestroy,omitempty"`

	// A user-friendly description of the dataset
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
