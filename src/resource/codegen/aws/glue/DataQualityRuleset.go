package glue

import types "libds/aws/types"

type DataQualityRuleset struct {
	// Description of the data quality ruleset.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the data quality ruleset.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A Data Quality Definition Language (DQDL) ruleset. For more information, see the AWS Glue developer guide.
	Ruleset string `json:"ruleset,omitempty" yaml:"ruleset,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A Configuration block specifying a target table associated with the data quality ruleset. See `target_table` below.
	TargetTable types.Glue_DataQualityRulesetTargetTable `json:"targetTable,omitempty" yaml:"targetTable,omitempty"`
}
