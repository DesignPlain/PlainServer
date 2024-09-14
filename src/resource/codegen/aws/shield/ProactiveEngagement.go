package shield

import types "libds/aws/types"

type ProactiveEngagement struct {
	// One or more emergency contacts. You must provide at least one phone number in the emergency contact list. See `emergency_contacts`.
	EmergencyContacts []types.Shield_ProactiveEngagementEmergencyContact `json:"emergencyContacts,omitempty" yaml:"emergencyContacts,omitempty"`

	// Boolean value indicating if Proactive Engagement should be enabled or not.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
