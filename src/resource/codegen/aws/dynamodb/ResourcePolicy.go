package dynamodb

type ResourcePolicy struct {
	// The Amazon Resource Name (ARN) of the DynamoDB resource to which the policy will be attached. The resources you can specify include tables and streams. You can control index permissions using the base table's policy. To specify the same permission level for your table and its indexes, you can provide both the table and index Amazon Resource Name (ARN)s in the Resource field of a given Statement in your policy document. Alternatively, to specify different permissions for your table, indexes, or both, you can define multiple Statement fields in your policy document.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// Set this parameter to true to confirm that you want to remove your permissions to change the policy of this resource in the future.
	ConfirmRemoveSelfResourceAccess bool `json:"confirmRemoveSelfResourceAccess,omitempty" yaml:"confirmRemoveSelfResourceAccess,omitempty"`

	/*
	   n Amazon Web Services resource-based policy document in JSON format. The maximum size supported for a resource-based policy document is 20 KB. DynamoDB counts whitespaces when calculating the size of a policy against this limit. For a full list of all considerations that you should keep in mind while attaching a resource-based policy, see Resource-based policy considerations.

	   The following arguments are optional:
	*/
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
