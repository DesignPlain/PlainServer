package firestore

type Database struct {
	/*
	   State of delete protection for the database. When delete protection is enabled, this database cannot be deleted. The
	   default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
	   --Note:-- Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'.
	   Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
	*/
	DeleteProtectionState string `json:"deleteProtectionState,omitempty" yaml:"deleteProtectionState,omitempty"`

	/*
	   The location of the database. Available locations are listed at
	   https://cloud.google.com/firestore/docs/locations.
	*/
	LocationId string `json:"locationId,omitempty" yaml:"locationId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The App Engine integration mode to use for this database.
	   Possible values are: `ENABLED`, `DISABLED`.
	*/
	AppEngineIntegrationMode string `json:"appEngineIntegrationMode,omitempty" yaml:"appEngineIntegrationMode,omitempty"`

	/*
	   The concurrency control mode to use for this database.
	   Possible values are: `OPTIMISTIC`, `PESSIMISTIC`, `OPTIMISTIC_WITH_ENTITY_GROUPS`.
	*/
	ConcurrencyMode string `json:"concurrencyMode,omitempty" yaml:"concurrencyMode,omitempty"`

	/*
	   Deletion behavior for this database. If the deletion policy is 'ABANDON', the database will be removed from Terraform
	   state but not deleted from Google Cloud upon destruction. If the deletion policy is 'DELETE', the database will both be
	   removed from Terraform state and deleted from Google Cloud upon destruction. The default value is 'ABANDON'. See also
	   'delete_protection'.
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	/*
	   The ID to use for the database, which will become the final
	   component of the database's resource name. This value should be 4-63
	   characters. Valid characters are /[a-z][0-9]-/ with first character
	   a letter and the last a letter or a number. Must not be
	   UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	   "(default)" database id is also valid.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Whether to enable the PITR feature on this database.
	   If `POINT_IN_TIME_RECOVERY_ENABLED` is selected, reads are supported on selected versions of the data from within the past 7 days.
	   versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
	   and reads against 1-minute snapshots beyond 1 hour and within 7 days.
	   If `POINT_IN_TIME_RECOVERY_DISABLED` is selected, reads are supported on any version of the data from within the past 1 hour.
	   Default value is `POINT_IN_TIME_RECOVERY_DISABLED`.
	   Possible values are: `POINT_IN_TIME_RECOVERY_ENABLED`, `POINT_IN_TIME_RECOVERY_DISABLED`.
	*/
	PointInTimeRecoveryEnablement string `json:"pointInTimeRecoveryEnablement,omitempty" yaml:"pointInTimeRecoveryEnablement,omitempty"`

	/*
	   The type of the database.
	   See https://cloud.google.com/datastore/docs/firestore-or-datastore
	   for information about how to choose.
	   Possible values are: `FIRESTORE_NATIVE`, `DATASTORE_MODE`.


	   - - -
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
