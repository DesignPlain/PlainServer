package types

type Licensemanager_getLicenseGrantsFilter struct {
	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_ListReceivedGrants.html#API_ListReceivedGrants_RequestSyntax).
	   For example, if filtering using `ProductSKU`, use:

	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const selected = aws.licensemanager.getLicenseGrants({
	       filters: [{
	           name: "ProductSKU",
	           values: [""],
	       }],
	   });
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   selected = aws.licensemanager.get_license_grants(filters=[aws.licensemanager.GetLicenseGrantsFilterArgs(
	       name="ProductSKU",
	       values=[""],
	   )])
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var selected = Aws.LicenseManager.GetLicenseGrants.Invoke(new()
	       {
	           Filters = new[]
	           {
	               new Aws.LicenseManager.Inputs.GetLicenseGrantsFilterInputArgs
	               {
	                   Name = "ProductSKU",
	                   Values = new[]
	                   {
	                       "",
	                   },
	               },
	           },
	       });

	   });
	   ```
	   ```go
	   package main

	   import (
	   	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/licensemanager"
	   	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	   )

	   func main() {
	   	pulumi.Run(func(ctx -pulumi.Context) error {
	   		_, err := licensemanager.GetLicenseGrants(ctx, &licensemanager.GetLicenseGrantsArgs{
	   			Filters: []licensemanager.GetLicenseGrantsFilter{
	   				{
	   					Name: "ProductSKU",
	   					Values: []string{
	   						"",
	   					},
	   				},
	   			},
	   		}, nil)
	   		if err != nil {
	   			return err
	   		}
	   		return nil
	   	})
	   }
	   ```
	   ```java
	   package generated_program;

	   import com.pulumi.Context;
	   import com.pulumi.Pulumi;
	   import com.pulumi.core.Output;
	   import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
	   import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
	   import java.util.List;
	   import java.util.ArrayList;
	   import java.util.Map;
	   import java.io.File;
	   import java.nio.file.Files;
	   import java.nio.file.Paths;

	   public class App {
	       public static void main(String[] args) {
	           Pulumi.run(App::stack);
	       }

	       public static void stack(Context ctx) {
	           final var selected = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
	               .filters(GetLicenseGrantsFilterArgs.builder()
	                   .name("ProductSKU")
	                   .values("")
	                   .build())
	               .build());

	       }
	   }
	   ```
	   ```yaml
	   variables:
	     selected:
	       fn::invoke:
	         Function: aws:licensemanager:getLicenseGrants
	         Arguments:
	           filters:
	             - name: ProductSKU
	               values:
	                 -
	   ```
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values that are accepted for the given field.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
