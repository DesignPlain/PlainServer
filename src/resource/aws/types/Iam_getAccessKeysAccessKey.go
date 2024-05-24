package types

type Iam_getAccessKeysAccessKey struct {
	// Access key ID.
	AccessKeyId string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`

	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the access key was created.
	CreateDate string `json:"createDate,omitempty" yaml:"createDate,omitempty"`

	// Access key status. Possible values are `Active` and `Inactive`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
