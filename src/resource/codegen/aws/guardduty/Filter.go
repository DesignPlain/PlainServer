package guardduty

import types "libds/aws/types"

type Filter struct {
	// Represents the criteria to be used in the filter for querying findings. Contains one or more `criterion` blocks, documented below.
	FindingCriteria types.Guardduty_FilterFindingCriteria `json:"findingCriteria,omitempty" yaml:"findingCriteria,omitempty"`

	// The name of your filter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank int `json:"rank,omitempty" yaml:"rank,omitempty"`

	// The tags that you want to add to the Filter resource. A tag consists of a key and a value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the action that is to be applied to the findings that match the filter. Can be one of `ARCHIVE` or `NOOP`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// Description of the filter.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ID of a GuardDuty detector, attached to your account.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`
}
