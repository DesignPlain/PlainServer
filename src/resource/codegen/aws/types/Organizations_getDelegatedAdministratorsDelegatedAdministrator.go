package types

type Organizations_getDelegatedAdministratorsDelegatedAdministrator struct {
	// The friendly name of the delegated administrator's account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The status of the delegated administrator's account in the organization.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The ARN of the delegated administrator's account.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The date when the account was made a delegated administrator.
	DelegationEnabledDate string `json:"delegationEnabledDate,omitempty" yaml:"delegationEnabledDate,omitempty"`

	// The email address that is associated with the delegated administrator's AWS account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// The unique identifier (ID) of the delegated administrator's account.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The method by which the delegated administrator's account joined the organization.
	JoinedMethod string `json:"joinedMethod,omitempty" yaml:"joinedMethod,omitempty"`

	// The date when the delegated administrator's account became a part of the organization.
	JoinedTimestamp string `json:"joinedTimestamp,omitempty" yaml:"joinedTimestamp,omitempty"`
}
