package essentialcontacts

type Contact struct {
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []string `json:"notificationCategorySubscriptions,omitempty" yaml:"notificationCategorySubscriptions,omitempty"`

	/*
	   The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// The email address to send notifications to. This does not need to be a Google account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
	LanguageTag string `json:"languageTag,omitempty" yaml:"languageTag,omitempty"`
}
