package ssmcontacts

import types "libds/aws/types"

type ContactChannel struct {
	// Amazon Resource Name (ARN) of the AWS SSM Contact that the contact channel belongs to.
	ContactId string `json:"contactId,omitempty" yaml:"contactId,omitempty"`

	// Block that contains contact engagement details. See details below.
	DeliveryAddress types.Ssmcontacts_ContactChannelDeliveryAddress `json:"deliveryAddress,omitempty" yaml:"deliveryAddress,omitempty"`

	// Name of the contact channel. Must be between 1 and 255 characters, and may contain alphanumerics, underscores (`_`), hyphens (`-`), periods (`.`), and spaces.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type of the contact channel. One of `SMS`, `VOICE` or `EMAIL`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
