package mediaconvert

import types "libds/aws/types"

type Queue struct {
	// A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A description of the queue
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A unique identifier describing the queue
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
	PricingPlan string `json:"pricingPlan,omitempty" yaml:"pricingPlan,omitempty"`

	// A detail pricing plan of the  reserved queue. See below.
	ReservationPlanSettings types.Mediaconvert_QueueReservationPlanSettings `json:"reservationPlanSettings,omitempty" yaml:"reservationPlanSettings,omitempty"`
}
