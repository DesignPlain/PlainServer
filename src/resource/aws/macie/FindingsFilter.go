package macie

import types "DesignSphere_Server/src/resource/aws/types"

type FindingsFilter struct {
	// The criteria to use to filter findings.
	FindingCriteria types.Macie_FindingsFilterFindingCriteria `json:"findingCriteria,omitempty" yaml:"findingCriteria,omitempty"`

	// A custom name for the filter. The name must contain at least 3 characters and can contain as many as 64 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The position of the filter in the list of saved filters on the Amazon Macie console. This value also determines the order in which the filter is applied to findings, relative to other filters that are also applied to the findings.
	Position int `json:"position,omitempty" yaml:"position,omitempty"`

	// A map of key-value pairs that specifies the tags to associate with the filter.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The action to perform on findings that meet the filter criteria (`finding_criteria`). Valid values are: `ARCHIVE`, suppress (automatically archive) the findings; and, `NOOP`, don't perform any action on the findings.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// A custom description of the filter. The description can contain as many as 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
