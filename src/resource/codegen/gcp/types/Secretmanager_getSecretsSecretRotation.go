package types

type Secretmanager_getSecretsSecretRotation struct {
	// Timestamp in UTC at which the Secret is scheduled to rotate.
	NextRotationTime string `json:"nextRotationTime,omitempty" yaml:"nextRotationTime,omitempty"`

	// The Duration between rotation notifications.
	RotationPeriod string `json:"rotationPeriod,omitempty" yaml:"rotationPeriod,omitempty"`
}
