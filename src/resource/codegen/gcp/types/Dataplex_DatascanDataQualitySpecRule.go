package types

type Dataplex_DatascanDataQualitySpecRule struct {
	// The dimension a rule belongs to. Results are also aggregated at the dimension level. Supported dimensions are ["COMPLETENESS", "ACCURACY", "CONSISTENCY", "VALIDITY", "UNIQUENESS", "INTEGRITY"]
	Dimension string `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Rows with null values will automatically fail a rule, unless ignoreNull is true. In that case, such null rows are trivially considered passing. Only applicable to ColumnMap rules.
	IgnoreNull bool `json:"ignoreNull,omitempty" yaml:"ignoreNull,omitempty"`

	// ColumnMap rule which evaluates whether each column value is null.
	NonNullExpectation Dataplex_DatascanDataQualitySpecRuleNonNullExpectation `json:"nonNullExpectation,omitempty" yaml:"nonNullExpectation,omitempty"`

	/*
	   ColumnMap rule which evaluates whether each column value is contained by a specified set.
	   Structure is documented below.
	*/
	SetExpectation Dataplex_DatascanDataQualitySpecRuleSetExpectation `json:"setExpectation,omitempty" yaml:"setExpectation,omitempty"`

	// The unnested column which this rule is evaluated against.
	Column string `json:"column,omitempty" yaml:"column,omitempty"`

	/*
	   Table rule which evaluates whether the provided expression is true.
	   Structure is documented below.
	*/
	TableConditionExpectation Dataplex_DatascanDataQualitySpecRuleTableConditionExpectation `json:"tableConditionExpectation,omitempty" yaml:"tableConditionExpectation,omitempty"`

	/*
	   ColumnMap rule which evaluates whether each column value lies between a specified range.
	   Structure is documented below.
	*/
	RangeExpectation Dataplex_DatascanDataQualitySpecRuleRangeExpectation `json:"rangeExpectation,omitempty" yaml:"rangeExpectation,omitempty"`

	/*
	   Table rule which evaluates whether each row passes the specified condition.
	   Structure is documented below.
	*/
	RowConditionExpectation Dataplex_DatascanDataQualitySpecRuleRowConditionExpectation `json:"rowConditionExpectation,omitempty" yaml:"rowConditionExpectation,omitempty"`

	/*
	   ColumnAggregate rule which evaluates whether the column aggregate statistic lies between a specified range.
	   Structure is documented below.
	*/
	StatisticRangeExpectation Dataplex_DatascanDataQualitySpecRuleStatisticRangeExpectation `json:"statisticRangeExpectation,omitempty" yaml:"statisticRangeExpectation,omitempty"`

	// Row-level rule which evaluates whether each column value is unique.
	UniquenessExpectation Dataplex_DatascanDataQualitySpecRuleUniquenessExpectation `json:"uniquenessExpectation,omitempty" yaml:"uniquenessExpectation,omitempty"`

	/*
	   Description of the rule.
	   The maximum length is 1,024 characters.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   ColumnMap rule which evaluates whether each column value matches a specified regex.
	   Structure is documented below.
	*/
	RegexExpectation Dataplex_DatascanDataQualitySpecRuleRegexExpectation `json:"regexExpectation,omitempty" yaml:"regexExpectation,omitempty"`

	// The minimum ratio of passing_rows / total_rows required to pass this rule, with a range of [0.0, 1.0]. 0 indicates default value (i.e. 1.0).
	Threshold float64 `json:"threshold,omitempty" yaml:"threshold,omitempty"`

	/*
	   A mutable name for the rule.
	   The name must contain only letters (a-z, A-Z), numbers (0-9), or hyphens (-).
	   The maximum length is 63 characters.
	   Must start with a letter.
	   Must end with a number or a letter.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
