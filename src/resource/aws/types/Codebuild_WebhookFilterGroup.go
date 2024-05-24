package types

type Codebuild_WebhookFilterGroup struct {
	// A webhook filter for the group. Filter blocks are documented below.
	Filters []Codebuild_WebhookFilterGroupFilter `json:"filters,omitempty" yaml:"filters,omitempty"`
}
