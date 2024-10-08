package types

type Licensemanager_getReceivedLicensesFilter struct {
	// Set of values that are accepted for the given field.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_ListReceivedLicenses.html#API_ListReceivedLicenses_RequestSyntax).
	   For example, if filtering using `ProductSKU`, use:

	   <!--Start PulumiCodeChooser -->
	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const selected = aws.licensemanager.getReceivedLicenses({
	       filters: [{
	           name: "ProductSKU",
	           values: [""],
	       }],
	   });
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   selected = aws.licensemanager.get_received_licenses(filters=[{
	       "name": "ProductSKU",
	       "values": [""],
	   }])
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var selected = Aws.LicenseManager.GetReceivedLicenses.Invoke(new()
	       {
	           Filters = new[]
	           {
	               new Aws.LicenseManager.Inputs.GetReceivedLicensesFilterInputArgs
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
	   		_, err := licensemanager.GetReceivedLicenses(ctx, &licensemanager.GetReceivedLicensesArgs{
	   			Filters: []licensemanager.GetReceivedLicensesFilter{
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
	   import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
	           final var selected = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
	               .filters(GetReceivedLicensesFilterArgs.builder()
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
	         Function: aws:licensemanager:getReceivedLicenses
	         Arguments:
	           filters:
	             - name: ProductSKU
	               values:
	                 -
	   ```
	   <!--End PulumiCodeChooser -->
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
