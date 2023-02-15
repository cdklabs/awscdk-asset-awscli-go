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
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

asset := s3_assets.NewAsset(this, jsii.String("layer-asset"), &AssetProps{
	Path: awscdkassetawscliv1.ASSET_FILE,
	AssetHash: awscdk.FileSystem_Fingerprint(*awscdkassetawscliv1.LAYER_SOURCE_DIR),
})
fn.AddLayers(lambda.NewLayerVersion(this, jsii.String("AwsCliLayer"), &LayerVersionProps{
	Code: lambda.Code_FromBucket(asset.Bucket, asset.S3ObjectKey),
}))
```

The CLI will be installed under `/opt/awscli/aws`.
