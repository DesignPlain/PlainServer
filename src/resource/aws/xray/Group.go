package xray

import types "DesignSphere_Server/src/resource/aws/types"

type Group struct {
	// The filter expression defining criteria by which to group traces. more info can be found in official [docs](https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html).
	FilterExpression string `json:"filterExpression,omitempty" yaml:"filterExpression,omitempty"`

	// The name of the group.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// Configuration options for enabling insights.
	InsightsConfiguration types.Xray_GroupInsightsConfiguration `json:"insightsConfiguration,omitempty" yaml:"insightsConfiguration,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
