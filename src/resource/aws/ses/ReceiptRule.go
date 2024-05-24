package ses

import types "DesignSphere_Server/src/resource/aws/types"

type ReceiptRule struct {
	// The name of the rule to place this rule after
	After string `json:"after,omitempty" yaml:"after,omitempty"`

	// If true, the rule will be enabled
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the rule
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If true, incoming emails will be scanned for spam and viruses
	ScanEnabled bool `json:"scanEnabled,omitempty" yaml:"scanEnabled,omitempty"`

	// A list of WorkMail Action blocks. Documented below.
	WorkmailActions []types.Ses_ReceiptRuleWorkmailAction `json:"workmailActions,omitempty" yaml:"workmailActions,omitempty"`

	// A list of Stop Action blocks. Documented below.
	StopActions []types.Ses_ReceiptRuleStopAction `json:"stopActions,omitempty" yaml:"stopActions,omitempty"`

	// A list of email addresses
	Recipients []string `json:"recipients,omitempty" yaml:"recipients,omitempty"`

	// A list of SNS Action blocks. Documented below.
	SnsActions []types.Ses_ReceiptRuleSnsAction `json:"snsActions,omitempty" yaml:"snsActions,omitempty"`

	// A list of Add Header Action blocks. Documented below.
	AddHeaderActions []types.Ses_ReceiptRuleAddHeaderAction `json:"addHeaderActions,omitempty" yaml:"addHeaderActions,omitempty"`

	// A list of Bounce Action blocks. Documented below.
	BounceActions []types.Ses_ReceiptRuleBounceAction `json:"bounceActions,omitempty" yaml:"bounceActions,omitempty"`

	// A list of Lambda Action blocks. Documented below.
	LambdaActions []types.Ses_ReceiptRuleLambdaAction `json:"lambdaActions,omitempty" yaml:"lambdaActions,omitempty"`

	// The name of the rule set
	RuleSetName string `json:"ruleSetName,omitempty" yaml:"ruleSetName,omitempty"`

	// A list of S3 Action blocks. Documented below.
	S3Actions []types.Ses_ReceiptRuleS3Action `json:"s3Actions,omitempty" yaml:"s3Actions,omitempty"`

	// `Require` or `Optional`
	TlsPolicy string `json:"tlsPolicy,omitempty" yaml:"tlsPolicy,omitempty"`
}
