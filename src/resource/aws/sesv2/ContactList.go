package sesv2

import types "DesignSphere_Server/src/resource/aws/types"

type ContactList struct {
	// Configuration block(s) with topic for the contact list. Detailed below.
	Topics []types.Sesv2_ContactListTopic `json:"topics,omitempty" yaml:"topics,omitempty"`

	/*
	   Name of the contact list.

	   The following arguments are optional:
	*/
	ContactListName string `json:"contactListName,omitempty" yaml:"contactListName,omitempty"`

	// Description of what the contact list is about.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Key-value map of resource tags for the contact list. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
