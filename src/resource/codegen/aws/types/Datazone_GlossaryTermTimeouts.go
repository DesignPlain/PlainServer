package types

type Datazone_GlossaryTermTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	Create string `json:"create,omitempty" yaml:"create,omitempty"`
}
