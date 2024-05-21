package spanner

import types "DesignSphere_Server/src/resource/gcp/types"

type Database struct {
	/*
	   Encryption configuration for the database
	   Structure is documented below.
	*/
	EncryptionConfig types.Spanner_DatabaseEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	/*
	   The instance to create the database on.


	   - - -
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   A unique identifier for the database, which cannot be changed after
	   the instance is created. Values are of the form [a-z][-a-z0-9]-[a-z0-9].
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The retention period for the database. The retention period must be between 1 hour
	   and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
	   the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
	   If this property is used, you must avoid adding new DDL statements to `ddl` that
	   update the database's version_retention_period.
	*/
	VersionRetentionPeriod string `json:"versionRetentionPeriod,omitempty" yaml:"versionRetentionPeriod,omitempty"`

	/*
	   The dialect of the Cloud Spanner Database.
	   If it is not provided, "GOOGLE_STANDARD_SQL" will be used.
	   Possible values are: `GOOGLE_STANDARD_SQL`, `POSTGRESQL`.
	*/
	DatabaseDialect string `json:"databaseDialect,omitempty" yaml:"databaseDialect,omitempty"`

	/*
	   An optional list of DDL statements to run inside the newly created
	   database. Statements can create tables, indexes, etc. These statements
	   execute atomically with the creation of the database: if there is an
	   error in any statement, the database is not created.
	*/
	Ddls []string `json:"ddls,omitempty" yaml:"ddls,omitempty"`

	/*
	   Whether or not to allow the provider to destroy the instance. Unless this field is set to false
	   in state, a `destroy` or `update` that would delete the instance will fail.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	/*
	   Whether drop protection is enabled for this database. Defaults to false. Drop protection is different from the
	   "deletion_protection" attribute in the following ways: (1) "deletion_protection" only protects the database from
	   deletions in Terraform. whereas setting “enableDropProtection” to true protects the database from deletions in all
	   interfaces. (2) Setting "enableDropProtection" to true also prevents the deletion of the parent instance containing the
	   database. "deletion_protection" attribute does not provide protection against the deletion of the parent instance.
	*/
	EnableDropProtection bool `json:"enableDropProtection,omitempty" yaml:"enableDropProtection,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
