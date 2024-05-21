package types

type Monitoring_GenericServiceBasicService struct {
	/*
	   Labels that specify the resource that emits the monitoring data
	   which is used for SLO reporting of this `Service`.
	*/
	ServiceLabels map[string]string `json:"serviceLabels,omitempty" yaml:"serviceLabels,omitempty"`

	/*
	   The type of service that this basic service defines, e.g.
	   APP_ENGINE service type
	*/
	ServiceType string `json:"serviceType,omitempty" yaml:"serviceType,omitempty"`
}
