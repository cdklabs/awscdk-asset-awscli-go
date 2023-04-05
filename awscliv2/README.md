# Asset with AWS CLI v2

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

---


> This library is currently under development. Do not use!

<!--END STABILITY BANNER-->

This module exports a single class called `AwsCliAsset` which is an `s3_assets.Asset` that bundles the AWS CLI v2.

Usage:

```go
// Example automatically generated from non-compiling source. May contain errors.
// AwsCliLayer bundles the AWS CLI in a lambda layer
import "github.com/aws-samples/dummy/awscdkassetawscliv2"
import "github.com/aws/aws-cdk-go/awscdk"
import s3_assets "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

awscli := awscdkassetawscliv2.NewAwsCliAsset(this, jsii.String("AwsCliCode"))
fn.AddLayers(lambda.NewLayerVersion(this, jsii.String("AwsCliLayer"), &LayerVersionProps{
	Code: lambda.Code_FromBucket(awscli.bucket, awscli.s3ObjectKey),
}))
```

The CLI will be installed under `/opt/awscli/aws`.
