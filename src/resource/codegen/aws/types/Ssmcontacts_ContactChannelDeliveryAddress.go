package types

type Ssmcontacts_ContactChannelDeliveryAddress struct {
	// Details to engage this contact channel. The expected format depends on the contact channel type and is described in the [`ContactChannelAddress` section of the SSM Contacts API Reference](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_SSMContacts_ContactChannelAddress.html).
	SimpleAddress string `json:"simpleAddress,omitempty" yaml:"simpleAddress,omitempty"`
}
