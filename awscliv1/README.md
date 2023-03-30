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
// Example automatically generated from non-compiling source. May contain errors.
// AwsCliLayer bundles the AWS CLI in a lambda layer
import "github.com/aws-samples/dummy/awscdkassetawscliv1"
import "github.com/aws-samples/dummy/awscdklib/awslambda"
import s3_assets "github.com/aws-samples/dummy/awscdklib/awss3assets"
import "github.com/aws-samples/dummy/awscdklib"

var fn lambda.Function

asset := s3_assets.NewAsset(this, jsii.String("layer-asset"), map[string]*string{
	"path": awscdkassetawscliv1.ASSET_FILE,
	"assetHash": awscdklib.FileSystem_fingerprint(awscdkassetawscliv1.LAYER_SOURCE_DIR),
})
fn.addLayers(lambda.NewLayerVersion(this, jsii.String("AwsCliLayer"), map[string]interface{}{
	"code": lambda.Code_fromBucket(asset.bucket, asset.s3ObjectKey),
}))
```

The CLI will be installed under `/opt/awscli/aws`.
