package types

type Projects_AccessApprovalSettingsEnrolledService struct {
	/*
	   The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
	   all
	   appengine.googleapis.com
	   bigquery.googleapis.com
	   bigtable.googleapis.com
	   cloudkms.googleapis.com
	   compute.googleapis.com
	   dataflow.googleapis.com
	   iam.googleapis.com
	   pubsub.googleapis.com
	   storage.googleapis.com
	*/
	CloudProduct string `json:"cloudProduct,omitempty" yaml:"cloudProduct,omitempty"`

	/*
	   The enrollment level of the service.
	   Default value is `BLOCK_ALL`.
	   Possible values are: `BLOCK_ALL`.

	   - - -
	*/
	EnrollmentLevel string `json:"enrollmentLevel,omitempty" yaml:"enrollmentLevel,omitempty"`
}
