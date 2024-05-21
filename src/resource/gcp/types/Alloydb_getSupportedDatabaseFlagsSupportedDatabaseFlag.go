package types

type Alloydb_getSupportedDatabaseFlagsSupportedDatabaseFlag struct {
	// The name of the database flag, e.g. "max_allowed_packets". The is a possibly key for the Instance.database_flags map field.
	FlagName string `json:"flagName,omitempty" yaml:"flagName,omitempty"`

	// Restriction on `INTEGER` type value. Specifies the minimum value and the maximum value that can be specified, if applicable.
	IntegerRestrictions Alloydb_getSupportedDatabaseFlagsSupportedDatabaseFlagIntegerRestrictions `json:"integerRestrictions,omitempty" yaml:"integerRestrictions,omitempty"`

	// The name of the flag resource, following Google Cloud conventions, e.g.: - projects/{project}/locations/{location}/flags/{flag} This field currently has no semantic meaning.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether setting or updating this flag on an Instance requires a database restart. If a flag that requires database restart is set, the backend will automatically restart the database (making sure to satisfy any availability SLO's).
	RequiresDbRestart bool `json:"requiresDbRestart,omitempty" yaml:"requiresDbRestart,omitempty"`

	// Restriction on `STRING` type value. The list of allowed values, if bounded. This field will be empty if there is a unbounded number of allowed values.
	StringRestrictions Alloydb_getSupportedDatabaseFlagsSupportedDatabaseFlagStringRestrictions `json:"stringRestrictions,omitempty" yaml:"stringRestrictions,omitempty"`

	// Major database engine versions for which this flag is supported. The supported values are `POSTGRES_14` and `DATABASE_VERSION_UNSPECIFIED`.
	SupportedDbVersions []string `json:"supportedDbVersions,omitempty" yaml:"supportedDbVersions,omitempty"`

	// ValueType describes the semantic type of the value that the flag accepts. Regardless of the ValueType, the Instance.database_flags field accepts the stringified version of the value, i.e. "20" or "3.14". The supported values are `VALUE_TYPE_UNSPECIFIED`, `STRING`, `INTEGER`, `FLOAT` and `NONE`.
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`

	// Whether the database flag accepts multiple values. If true, a comma-separated list of stringified values may be specified.
	AcceptsMultipleValues bool `json:"acceptsMultipleValues,omitempty" yaml:"acceptsMultipleValues,omitempty"`
}
