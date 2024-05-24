package types

type Costexplorer_getTagsFilterNot struct {
	// Tags that match your request.
	Tags Costexplorer_getTagsFilterNotTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_getTagsFilterNotCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_getTagsFilterNotDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`
}
