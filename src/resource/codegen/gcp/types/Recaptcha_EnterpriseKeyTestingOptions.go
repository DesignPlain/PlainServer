package types

type Recaptcha_EnterpriseKeyTestingOptions struct {
	// For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if UNSOLVABLE_CHALLENGE. Possible values: TESTING_CHALLENGE_UNSPECIFIED, NOCAPTCHA, UNSOLVABLE_CHALLENGE
	TestingChallenge string `json:"testingChallenge,omitempty" yaml:"testingChallenge,omitempty"`

	// All assessments for this Key will return this score. Must be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive.
	TestingScore float64 `json:"testingScore,omitempty" yaml:"testingScore,omitempty"`
}
