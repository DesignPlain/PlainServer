package cloudfront

import types "libds/aws/types"

type FieldLevelEncryptionProfile struct {
	// An optional comment about the Field Level Encryption Profile.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The encryption entities config block for field-level encryption profiles that contains an attribute `items` which includes the encryption key and field pattern specifications.
	EncryptionEntities types.Cloudfront_FieldLevelEncryptionProfileEncryptionEntities `json:"encryptionEntities,omitempty" yaml:"encryptionEntities,omitempty"`

	// The name of the Field Level Encryption Profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
