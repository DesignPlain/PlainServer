package types

type Lakeformation_DataLakeSettingsCreateDatabaseDefaultPermission struct {
	// List of permissions that are granted to the principal. Valid values may include `ALL`, `SELECT`, `ALTER`, `DROP`, `DELETE`, `INSERT`, `DESCRIBE`, and `CREATE_TABLE`. For more details, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// Principal who is granted permissions. To enforce metadata and underlying data access control only by IAM on new databases and tables set `principal` to `IAM_ALLOWED_PRINCIPALS` and `permissions` to `["ALL"]`.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
}
