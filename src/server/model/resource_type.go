package model

import (
	"strings"

	"libds/aws"
	. "libds/ds_base"
	"libds/gcp"
)

// Get resource URI
// Example: `gcp:compute:Instance`
func GetURI(resourceType ResourceType) string {
	var provider_name string
	var res_uri string

	if resourceType < RESOURCE_TYPE_LIMIT {
		provider_name = "gcp"
		res_uri = gcp.Get_GCP_String(resourceType)
	} else {
		provider_name = "aws"
		res_uri = aws.Get_AWS_String(resourceType)
	}

	res_parts := strings.Split(res_uri, "_")
	res_family := strings.ToLower(res_parts[0])

	res_uri = res_parts[1]

	charArray := []rune(res_uri)
	res_uri = strings.ToUpper(string(charArray[0])) + string(charArray[1:])

	return provider_name + ":" + res_family + ":" + res_uri
}

func GetString(resourceType ResourceType) string {
	var res_str string

	if resourceType < RESOURCE_TYPE_LIMIT {
		res_str = gcp.Get_GCP_String(resourceType)
	} else {
		res_str = aws.Get_AWS_String(resourceType)
	}

	return res_str
}
