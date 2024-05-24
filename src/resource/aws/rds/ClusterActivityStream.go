package rds

type ClusterActivityStream struct {
	// Specifies whether the database activity stream includes engine-native audit fields. This option only applies to an Oracle DB instance. By default, no engine-native audit fields are included. Defaults `false`.
	EngineNativeAuditFieldsIncluded bool `json:"engineNativeAuditFieldsIncluded,omitempty" yaml:"engineNativeAuditFieldsIncluded,omitempty"`

	// The AWS KMS key identifier for encrypting messages in the database activity stream. The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Specifies the mode of the database activity stream. Database events such as a change or access generate an activity stream event. The database session can handle these events either synchronously or asynchronously. One of: `sync`, `async`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// The Amazon Resource Name (ARN) of the DB cluster.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
