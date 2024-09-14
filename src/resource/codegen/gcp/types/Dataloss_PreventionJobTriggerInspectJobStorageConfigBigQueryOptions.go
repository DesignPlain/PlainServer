package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigBigQueryOptions struct {
	/*
	   Set of files to scan.
	   Structure is documented below.
	*/
	TableReference Dataloss_PreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference `json:"tableReference,omitempty" yaml:"tableReference,omitempty"`

	/*
	   References to fields excluded from scanning.
	   This allows you to skip inspection of entire columns which you know have no findings.
	   Structure is documented below.
	*/
	ExcludedFields []Dataloss_PreventionJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedField `json:"excludedFields,omitempty" yaml:"excludedFields,omitempty"`

	/*
	   Specifies the BigQuery fields that will be returned with findings.
	   If not specified, no identifying fields will be returned for findings.
	   Structure is documented below.
	*/
	IdentifyingFields []Dataloss_PreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingField `json:"identifyingFields,omitempty" yaml:"identifyingFields,omitempty"`

	/*
	   Limit scanning only to these fields.
	   Structure is documented below.
	*/
	IncludedFields []Dataloss_PreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedField `json:"includedFields,omitempty" yaml:"includedFields,omitempty"`

	/*
	   Max number of rows to scan. If the table has more rows than this value, the rest of the rows are omitted.
	   If not set, or if set to 0, all rows will be scanned. Only one of rowsLimit and rowsLimitPercent can be
	   specified. Cannot be used in conjunction with TimespanConfig.
	*/
	RowsLimit int `json:"rowsLimit,omitempty" yaml:"rowsLimit,omitempty"`

	/*
	   Max percentage of rows to scan. The rest are omitted. The number of rows scanned is rounded down.
	   Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of
	   rowsLimit and rowsLimitPercent can be specified. Cannot be used in conjunction with TimespanConfig.
	*/
	RowsLimitPercent int `json:"rowsLimitPercent,omitempty" yaml:"rowsLimitPercent,omitempty"`

	/*
	   How to sample rows if not all rows are scanned. Meaningful only when used in conjunction with either
	   rowsLimit or rowsLimitPercent. If not specified, rows are scanned in the order BigQuery reads them.
	   Default value is `TOP`.
	   Possible values are: `TOP`, `RANDOM_START`.
	*/
	SampleMethod string `json:"sampleMethod,omitempty" yaml:"sampleMethod,omitempty"`
}
