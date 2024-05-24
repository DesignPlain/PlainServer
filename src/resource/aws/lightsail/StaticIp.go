package lightsail

type StaticIp struct {
	// The name for the allocated static IP
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
