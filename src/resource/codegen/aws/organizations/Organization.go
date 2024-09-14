package organizations

type Organization struct {
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet string `json:"featureSet,omitempty" yaml:"featureSet,omitempty"`

	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. Some services do not support enablement via this endpoint, see [warning in aws docs](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnableAWSServiceAccess.html).
	AwsServiceAccessPrincipals []string `json:"awsServiceAccessPrincipals,omitempty" yaml:"awsServiceAccessPrincipals,omitempty"`

	// List of Organizations policy types to enable in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g., `AISERVICES_OPT_OUT_POLICY`, `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes []string `json:"enabledPolicyTypes,omitempty" yaml:"enabledPolicyTypes,omitempty"`
}
