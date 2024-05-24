package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type UserProfile struct {
	// A specifier for the type of value specified in `single_sign_on_user_value`. Currently, the only supported value is `UserName`. If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserIdentifier string `json:"singleSignOnUserIdentifier,omitempty" yaml:"singleSignOnUserIdentifier,omitempty"`

	// The username of the associated AWS Single Sign-On User for this User Profile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserValue string `json:"singleSignOnUserValue,omitempty" yaml:"singleSignOnUserValue,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name for the User Profile.
	UserProfileName string `json:"userProfileName,omitempty" yaml:"userProfileName,omitempty"`

	// The user settings. See User Settings below.
	UserSettings types.Sagemaker_UserProfileUserSettings `json:"userSettings,omitempty" yaml:"userSettings,omitempty"`

	// The ID of the associated Domain.
	DomainId string `json:"domainId,omitempty" yaml:"domainId,omitempty"`
}
