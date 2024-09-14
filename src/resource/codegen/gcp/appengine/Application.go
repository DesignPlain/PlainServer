package appengine

import types "libds/gcp/types"

type Application struct {
	// Settings for enabling Cloud Identity Aware Proxy
	Iap types.Appengine_ApplicationIap `json:"iap,omitempty" yaml:"iap,omitempty"`

	/*
	   The [location](https://cloud.google.com/appengine/docs/locations)
	   to serve the app from.
	*/
	LocationId string `json:"locationId,omitempty" yaml:"locationId,omitempty"`

	/*
	   The project ID to create the application under.
	   ~>--NOTE:-- GCP only accepts project ID, not project number. If you are using number,
	   you may get a "Permission denied" error.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The serving status of the app.
	ServingStatus string `json:"servingStatus,omitempty" yaml:"servingStatus,omitempty"`

	// The domain to authenticate users with when using App Engine's User API.
	AuthDomain string `json:"authDomain,omitempty" yaml:"authDomain,omitempty"`

	/*
	   The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	   Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
	   instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted by the provider, but will be rejected by the API.
	   To create a Cloud Firestore database without creating an App Engine application, use the
	   `gcp.firestore.Database`
	   resource instead.
	*/
	DatabaseType string `json:"databaseType,omitempty" yaml:"databaseType,omitempty"`

	// A block of optional settings to configure specific App Engine features:
	FeatureSettings types.Appengine_ApplicationFeatureSettings `json:"featureSettings,omitempty" yaml:"featureSettings,omitempty"`
}
