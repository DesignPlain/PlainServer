package types

type Costexplorer_getTagsFilterOr struct {
	// Configuration block for the filter that's based on `CostCategory` values. See `cost_category` block below for details.
	CostCategory Costexplorer_getTagsFilterOrCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See `dimension` block below for details.
	Dimension Costexplorer_getTagsFilterOrDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Tags that match your request.
	Tags Costexplorer_getTagsFilterOrTags `json:"tags,omitempty" yaml:"tags,omitempty"`
}
