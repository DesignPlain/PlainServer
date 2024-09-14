package folder

import types "libds/gcp/types"

type AccessApprovalSettings struct {
	/*
	   The asymmetric crypto key version to use for signing approval requests.
	   Empty active_key_version indicates that a Google-managed key should be used for signing.
	   This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.
	*/
	ActiveKeyVersion string `json:"activeKeyVersion,omitempty" yaml:"activeKeyVersion,omitempty"`

	/*
	   A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	   Access requests for the resource given by name against any of these services contained here will be required
	   to have explicit approval. Enrollment can only be done on an all or nothing basis.
	   A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	   Structure is documented below.
	*/
	EnrolledServices []types.Folder_AccessApprovalSettingsEnrolledService `json:"enrolledServices,omitempty" yaml:"enrolledServices,omitempty"`

	// ID of the folder of the access approval settings.
	FolderId string `json:"folderId,omitempty" yaml:"folderId,omitempty"`

	/*
	   A list of email addresses to which notifications relating to approval requests should be sent.
	   Notifications relating to a resource will be sent to all emails in the settings of ancestor
	   resources of that resource. A maximum of 50 email addresses are allowed.
	*/
	NotificationEmails []string `json:"notificationEmails,omitempty" yaml:"notificationEmails,omitempty"`
}
