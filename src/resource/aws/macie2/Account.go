package macie2

type Account struct {
	// Specifies how often to publish updates to policy findings for the account. This includes publishing updates to AWS Security Hub and Amazon EventBridge (formerly called Amazon CloudWatch Events). Valid values are `FIFTEEN_MINUTES`, `ONE_HOUR` or `SIX_HOURS`.
	FindingPublishingFrequency string `json:"findingPublishingFrequency,omitempty" yaml:"findingPublishingFrequency,omitempty"`

	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
