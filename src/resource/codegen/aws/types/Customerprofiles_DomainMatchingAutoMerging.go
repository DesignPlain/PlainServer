package types

type Customerprofiles_DomainMatchingAutoMerging struct {
	// The flag that enables the auto-merging of duplicate profiles.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	MinAllowedConfidenceScoreForMerging float64 `json:"minAllowedConfidenceScoreForMerging,omitempty" yaml:"minAllowedConfidenceScoreForMerging,omitempty"`

	// A block that specifies how the auto-merging process should resolve conflicts between different profiles. Documented below.
	ConflictResolution Customerprofiles_DomainMatchingAutoMergingConflictResolution `json:"conflictResolution,omitempty" yaml:"conflictResolution,omitempty"`

	/*
	   A block that specifies a list of matching attributes that represent matching criteria. If two profiles meet at least one of the requirements in the matching attributes list, they will be merged. Documented below.
	   - `min_allowed_confidence_score_for_merging ` - (Optional) A number between 0 and 1 that represents the minimum confidence score required for profiles within a matching group to be merged during the auto-merge process. A higher score means higher similarity required to merge profiles.
	*/
	Consolidation Customerprofiles_DomainMatchingAutoMergingConsolidation `json:"consolidation,omitempty" yaml:"consolidation,omitempty"`
}
