package types

type Networkservices_HttpRouteRuleActionRequestHeaderModifier struct {
	// Add the headers with given map where key is the name of the header, value is the value of the header.
	Add map[string]string `json:"add,omitempty" yaml:"add,omitempty"`

	// Remove headers (matching by header names) specified in the list.
	Removes []string `json:"removes,omitempty" yaml:"removes,omitempty"`

	// Completely overwrite/replace the headers with given map where key is the name of the header, value is the value of the header.
	Set map[string]string `json:"set,omitempty" yaml:"set,omitempty"`
}
