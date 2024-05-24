package types

type Cloudtrail_TrailEventSelector struct {
	// A set of event sources to exclude. Valid values include: `kms.amazonaws.com` and `rdsdata.amazonaws.com`. `include_management_events` must be set to`true` to allow this.
	ExcludeManagementEventSources []string `json:"excludeManagementEventSources,omitempty" yaml:"excludeManagementEventSources,omitempty"`

	// Whether to include management events for your trail. Defaults to `true`.
	IncludeManagementEvents bool `json:"includeManagementEvents,omitempty" yaml:"includeManagementEvents,omitempty"`

	// Type of events to log. Valid values are `ReadOnly`, `WriteOnly`, `All`. Default value is `All`.
	ReadWriteType string `json:"readWriteType,omitempty" yaml:"readWriteType,omitempty"`

	// Configuration block for data events. See details below.
	DataResources []Cloudtrail_TrailEventSelectorDataResource `json:"dataResources,omitempty" yaml:"dataResources,omitempty"`
}
