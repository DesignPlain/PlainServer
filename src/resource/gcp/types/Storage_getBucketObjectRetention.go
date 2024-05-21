package types

type Storage_getBucketObjectRetention struct {
	// The object retention mode. Supported values include: "Unlocked", "Locked".
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// Time in RFC 3339 (e.g. 2030-01-01T02:03:04Z) until which object retention protects this object.
	RetainUntilTime string `json:"retainUntilTime,omitempty" yaml:"retainUntilTime,omitempty"`
}
