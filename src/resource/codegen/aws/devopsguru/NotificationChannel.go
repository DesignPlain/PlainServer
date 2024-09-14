package devopsguru

import types "libds/aws/types"

type NotificationChannel struct {
	// Filter configurations for the Amazon SNS notification topic. See the `filters` argument reference below.
	Filters types.Devopsguru_NotificationChannelFilters `json:"filters,omitempty" yaml:"filters,omitempty"`

	/*
	   SNS noficiation channel configurations. See the `sns` argument reference below.

	   The following arguments are optional:
	*/
	Sns types.Devopsguru_NotificationChannelSns `json:"sns,omitempty" yaml:"sns,omitempty"`
}
