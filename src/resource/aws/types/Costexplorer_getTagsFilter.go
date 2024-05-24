package types

type Costexplorer_getTagsFilter struct {
	// Tags that match your request.
	Tags Costexplorer_getTagsFilterTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_getTagsFilterAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_getTagsFilterCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_getTagsFilterDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Return results that match both `Dimension` object.
	Not Costexplorer_getTagsFilterNot `json:"not,omitempty" yaml:"not,omitempty"`

	// Return results that match both `Dimension` object.
	Ors []Costexplorer_getTagsFilterOr `json:"ors,omitempty" yaml:"ors,omitempty"`
}
