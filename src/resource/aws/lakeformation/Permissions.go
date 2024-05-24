package lakeformation

import types "DesignSphere_Server/src/resource/aws/types"

type Permissions struct {
	/*
	   Configuration block for a table with columns resource. Detailed below.

	   The following arguments are optional:
	*/
	TableWithColumns types.Lakeformation_PermissionsTableWithColumns `json:"tableWithColumns,omitempty" yaml:"tableWithColumns,omitempty"`

	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
	CatalogResource bool `json:"catalogResource,omitempty" yaml:"catalogResource,omitempty"`

	// Configuration block for an LF-tag resource. Detailed below.
	LfTag types.Lakeformation_PermissionsLfTag `json:"lfTag,omitempty" yaml:"lfTag,omitempty"`

	/*
	   Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).

	   > --NOTE:-- We highly recommend that the `principal` _NOT_ be a Lake Formation administrator (granted using `aws.lakeformation.DataLakeSettings`). The entity (e.g., IAM role) running the deployment will most likely need to be a Lake Formation administrator. As such, the entity will have implicit permissions and does not need permissions granted through this resource.

	   One of the following is required:
	*/
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// Configuration block for a table resource. Detailed below.
	Table types.Lakeformation_PermissionsTable `json:"table,omitempty" yaml:"table,omitempty"`

	// Configuration block for a data location resource. Detailed below.
	DataLocation types.Lakeformation_PermissionsDataLocation `json:"dataLocation,omitempty" yaml:"dataLocation,omitempty"`

	// Configuration block for a database resource. Detailed below.
	Database types.Lakeformation_PermissionsDatabase `json:"database,omitempty" yaml:"database,omitempty"`

	// Configuration block for an LF-tag policy resource. Detailed below.
	LfTagPolicy types.Lakeformation_PermissionsLfTagPolicy `json:"lfTagPolicy,omitempty" yaml:"lfTagPolicy,omitempty"`

	// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// Subset of `permissions` which the principal can pass.
	PermissionsWithGrantOptions []string `json:"permissionsWithGrantOptions,omitempty" yaml:"permissionsWithGrantOptions,omitempty"`
}
