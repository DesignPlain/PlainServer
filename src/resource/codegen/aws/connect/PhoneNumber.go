package connect

type PhoneNumber struct {
	// The description of the phone number.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain `+` as part of the country code. Do not specify this argument when importing the resource.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Tags to apply to the Phone Number. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`

	// The type of phone number. Valid Values: `TOLL_FREE` | `DID`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The ISO country code. For a list of Valid values, refer to [PhoneNumberCountryCode](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html#connect-SearchAvailablePhoneNumbers-request-PhoneNumberCountryCode).
	CountryCode string `json:"countryCode,omitempty" yaml:"countryCode,omitempty"`
}
