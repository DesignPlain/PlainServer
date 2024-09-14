package bigtable

import types "libds/gcp/types"

type Table struct {
	// The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A list of predefined keys to split the table on.
	   !> --Warning:-- Modifying the `split_keys` of an existing table will cause the provider
	   to delete/recreate the entire `gcp.bigtable.Table` resource.
	*/
	SplitKeys []string `json:"splitKeys,omitempty" yaml:"splitKeys,omitempty"`

	/*
	   Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.

	   -----
	*/
	ChangeStreamRetention string `json:"changeStreamRetention,omitempty" yaml:"changeStreamRetention,omitempty"`

	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies []types.Bigtable_TableColumnFamily `json:"columnFamilies,omitempty" yaml:"columnFamilies,omitempty"`

	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection string `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// The name of the Bigtable instance.
	InstanceName string `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`
}
