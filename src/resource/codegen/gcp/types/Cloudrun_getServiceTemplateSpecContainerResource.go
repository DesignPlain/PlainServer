package types

type Cloudrun_getServiceTemplateSpecContainerResource struct {
	/*
	   Limits describes the maximum amount of compute resources allowed.
	   The values of the map is string form of the 'quantity' k8s type:
	   https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	*/
	Limits map[string]string `json:"limits,omitempty" yaml:"limits,omitempty"`

	/*
	   Requests describes the minimum amount of compute resources required.
	   If Requests is omitted for a container, it defaults to Limits if that is
	   explicitly specified, otherwise to an implementation-defined value.
	   The values of the map is string form of the 'quantity' k8s type:
	   https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	*/
	Requests map[string]string `json:"requests,omitempty" yaml:"requests,omitempty"`
}
