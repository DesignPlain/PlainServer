package types

type Codeguruprofiler_getProfilingGroupProfilingStatus struct {
	//
	LatestAgentOrchestratedAt string `json:"latestAgentOrchestratedAt,omitempty" yaml:"latestAgentOrchestratedAt,omitempty"`

	//
	LatestAgentProfileReportedAt string `json:"latestAgentProfileReportedAt,omitempty" yaml:"latestAgentProfileReportedAt,omitempty"`

	//
	LatestAggregatedProfiles []string `json:"latestAggregatedProfiles,omitempty" yaml:"latestAggregatedProfiles,omitempty"`
}
