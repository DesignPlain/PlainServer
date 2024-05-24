package types

type Appstream_StackUserSetting struct {
	/*
	   Action that is enabled or disabled.
	   Valid values are `CLIPBOARD_COPY_FROM_LOCAL_DEVICE`,  `CLIPBOARD_COPY_TO_LOCAL_DEVICE`, `FILE_UPLOAD`, `FILE_DOWNLOAD`, `PRINTING_TO_LOCAL_DEVICE`, `DOMAIN_PASSWORD_SIGNIN`, or `DOMAIN_SMART_CARD_SIGNIN`.
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	/*
	   Whether the action is enabled or disabled.
	   Valid values are `ENABLED` or `DISABLED`.
	*/
	Permission string `json:"permission,omitempty" yaml:"permission,omitempty"`
}
