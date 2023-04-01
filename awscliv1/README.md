# Asset with AWS CLI v1

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Stable](https://img.shields.io/badge/cdk--constructs-stable-success.svg?style=for-the-badge)

---
<!--END STABILITY BANNER-->

This module bundles the AWS CLI v1 as a local asset. It exposes
constants `ASSET_FILE` and `LAYER_SOURCE_DIR` that can be consumed
via the CDK `Asset` construct.

Any Lambda Function that uses uses this asset must use a Python 3.x
runtime.

Usage:

```go
// AwsCliLayer bundles the AWS CLI in a lambda layer
import "github.com/aws-samples/dummy/awscdkassetawscliv1"
import "github.com/aws/aws-cdk-go/awscdk"
import s3_assets "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/awscdklib"

var fn function

asset := s3_assets.NewAsset(this, jsii.String("layer-asset"), &assetProps{
	path: awscdkassetawscliv1.ASSET_FILE,
	assetHash: *awscdklib.FileSystem_Fingerprint(*awscdkassetawscliv1.LAYER_SOURCE_DIR),
})
fn.addLayers(lambda.NewLayerVersion(this, jsii.String("AwsCliLayer"), &layerVersionProps{
	code: lambda.code_FromBucket(asset.bucket, asset.s3ObjectKey),
}))
```

The CLI will be installed under `/opt/awscli/aws`.
