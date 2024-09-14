package resourceexplorer

import types "libds/aws/types"

type View struct {
	// The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
	DefaultView bool `json:"defaultView,omitempty" yaml:"defaultView,omitempty"`

	// Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
	Filters types.Resourceexplorer_ViewFilters `json:"filters,omitempty" yaml:"filters,omitempty"`

	// Optional fields to be included in search results from this view. See Included Properties below for more details.
	IncludedProperties []types.Resourceexplorer_ViewIncludedProperty `json:"includedProperties,omitempty" yaml:"includedProperties,omitempty"`
}
