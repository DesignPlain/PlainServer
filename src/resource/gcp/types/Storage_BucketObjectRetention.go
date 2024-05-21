package types

type Storage_BucketObjectRetention struct {
	// The retention policy mode. Either `Locked` or `Unlocked`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   The time to retain the object until in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.

	   <a name>
	*/
	RetainUntilTime string `json:"retainUntilTime,omitempty" yaml:"retainUntilTime,omitempty"`
}
