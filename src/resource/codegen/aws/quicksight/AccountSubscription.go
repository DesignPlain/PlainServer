package quicksight

type AccountSubscription struct {
	// A 10-digit phone number for the author of the Amazon QuickSight account to use for future communications. This field is required if `ENTERPPRISE_AND_Q` is the selected edition of the new Amazon QuickSight account.
	ContactNumber string `json:"contactNumber,omitempty" yaml:"contactNumber,omitempty"`

	// Last name of the author of the Amazon QuickSight account to use for future communications. This field is required if `ENTERPPRISE_AND_Q` is the selected edition of the new Amazon QuickSight account.
	LastName string `json:"lastName,omitempty" yaml:"lastName,omitempty"`

	// Realm of the Active Directory that is associated with your Amazon QuickSight account.
	Realm string `json:"realm,omitempty" yaml:"realm,omitempty"`

	// AWS account ID hosting the QuickSight account. Default to provider account.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// Active Directory ID that is associated with your Amazon QuickSight account.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// Edition of Amazon QuickSight that you want your account to have. Currently, you can choose from `STANDARD`, `ENTERPRISE` or `ENTERPRISE_AND_Q`.
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`

	// Name of your Amazon QuickSight account. This name is unique over all of AWS, and it appears only when users sign in.
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`

	// Method that you want to use to authenticate your Amazon QuickSight account. Currently, the valid values for this parameter are `IAM_AND_QUICKSIGHT`, `IAM_ONLY`, `IAM_IDENTITY_CENTER`, and `ACTIVE_DIRECTORY`.
	AuthenticationMethod string `json:"authenticationMethod,omitempty" yaml:"authenticationMethod,omitempty"`

	// Author group associated with your Active Directory.
	AuthorGroups []string `json:"authorGroups,omitempty" yaml:"authorGroups,omitempty"`

	// First name of the author of the Amazon QuickSight account to use for future communications. This field is required if `ENTERPPRISE_AND_Q` is the selected edition of the new Amazon QuickSight account.
	FirstName string `json:"firstName,omitempty" yaml:"firstName,omitempty"`

	// Reader group associated with your Active Direcrtory.
	ReaderGroups []string `json:"readerGroups,omitempty" yaml:"readerGroups,omitempty"`

	// Name of your Active Directory. This field is required if `ACTIVE_DIRECTORY` is the selected authentication method of the new Amazon QuickSight account.
	ActiveDirectoryName string `json:"activeDirectoryName,omitempty" yaml:"activeDirectoryName,omitempty"`

	// Admin group associated with your Active Directory. This field is required if `ACTIVE_DIRECTORY` is the selected authentication method of the new Amazon QuickSight account.
	AdminGroups []string `json:"adminGroups,omitempty" yaml:"adminGroups,omitempty"`

	// Email address of the author of the Amazon QuickSight account to use for future communications. This field is required if `ENTERPPRISE_AND_Q` is the selected edition of the new Amazon QuickSight account.
	EmailAddress string `json:"emailAddress,omitempty" yaml:"emailAddress,omitempty"`

	// The Amazon Resource Name (ARN) for the IAM Identity Center instance.
	IamIdentityCenterInstanceArn string `json:"iamIdentityCenterInstanceArn,omitempty" yaml:"iamIdentityCenterInstanceArn,omitempty"`

	/*
	   Email address that you want Amazon QuickSight to send notifications to regarding your Amazon QuickSight account or Amazon QuickSight subscription.

	   The following arguments are optional:
	*/
	NotificationEmail string `json:"notificationEmail,omitempty" yaml:"notificationEmail,omitempty"`
}
