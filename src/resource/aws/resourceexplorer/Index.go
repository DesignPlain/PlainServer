package resourceexplorer

import types "DesignSphere_Server/src/resource/aws/types"

type Index struct {
	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Resourceexplorer_IndexTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The type of the index. Valid values: `AGGREGATOR`, `LOCAL`. To understand the difference between `LOCAL` and `AGGREGATOR`, see the [_AWS Resource Explorer User Guide_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
