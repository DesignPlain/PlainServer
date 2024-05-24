package identitystore

import types "DesignSphere_Server/src/resource/aws/types"

type User struct {
	// Details about the user's address. At most 1 address is allowed. Detailed below.
	Addresses types.Identitystore_UserAddresses `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	// An alternate name for the user.
	Nickname string `json:"nickname,omitempty" yaml:"nickname,omitempty"`

	// The name that is typically displayed when the user is referenced.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Details about the user's phone number. At most 1 phone number is allowed. Detailed below.
	PhoneNumbers types.Identitystore_UserPhoneNumbers `json:"phoneNumbers,omitempty" yaml:"phoneNumbers,omitempty"`

	// The preferred language of the user.
	PreferredLanguage string `json:"preferredLanguage,omitempty" yaml:"preferredLanguage,omitempty"`

	// The user's time zone.
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`

	// The user type.
	UserType string `json:"userType,omitempty" yaml:"userType,omitempty"`

	// The globally unique identifier for the identity store that this user is in.
	IdentityStoreId string `json:"identityStoreId,omitempty" yaml:"identityStoreId,omitempty"`

	// The user's geographical region or location.
	Locale string `json:"locale,omitempty" yaml:"locale,omitempty"`

	// An URL that may be associated with the user.
	ProfileUrl string `json:"profileUrl,omitempty" yaml:"profileUrl,omitempty"`

	/*
	   A unique string used to identify the user. This value can consist of letters, accented characters, symbols, numbers, and punctuation. This value is specified at the time the user is created and stored as an attribute of the user object in the identity store. The limit is 128 characters.

	   The following arguments are optional:
	*/
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// Details about the user's email. At most 1 email is allowed. Detailed below.
	Emails types.Identitystore_UserEmails `json:"emails,omitempty" yaml:"emails,omitempty"`

	// Details about the user's full name. Detailed below.
	Name types.Identitystore_UserName `json:"name,omitempty" yaml:"name,omitempty"`

	// The user's title.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
