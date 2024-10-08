package types

type Auditmanager_ControlControlMappingSourceSourceKeyword struct {
	// Input method for the keyword. Valid values are `INPUT_TEXT`, `SELECT_FROM_LIST`, or `UPLOAD_FILE`.
	KeywordInputType string `json:"keywordInputType,omitempty" yaml:"keywordInputType,omitempty"`

	// The value of the keyword that's used when mapping a control data source. For example, this can be a CloudTrail event name, a rule name for Config, a Security Hub control, or the name of an Amazon Web Services API call. See the [Audit Manager supported control data sources documentation](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources.html) for more information.
	KeywordValue string `json:"keywordValue,omitempty" yaml:"keywordValue,omitempty"`
}
