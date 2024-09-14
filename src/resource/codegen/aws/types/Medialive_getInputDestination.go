package types

type Medialive_getInputDestination struct {
	//
	Ip string `json:"ip,omitempty" yaml:"ip,omitempty"`

	//
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	//
	Vpcs []Medialive_getInputDestinationVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`
}
