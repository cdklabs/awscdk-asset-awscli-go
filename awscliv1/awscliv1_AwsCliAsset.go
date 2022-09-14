// An Asset construct that contains the AWS CLI, for use in Lambda Layers
package awscliv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/awscdk-asset-awscli-go/awscliv1/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/awscdk-asset-awscli-go/awscliv1/internal"
)

// A CDK Asset construct that contains the AWS CLI.
type AwsCliAsset interface {
	awss3assets.Asset
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	AssetHash() *string
	// The path to the asset, relative to the current Cloud Assembly.
	//
	// If asset staging is disabled, this will just be the original path.
	// If asset staging is enabled it will be the staged path.
	AssetPath() *string
	// The S3 bucket in which this asset resides.
	Bucket() awss3.IBucket
	// Attribute which represents the S3 HTTP URL of this asset.
	//
	// Example:
	//   https://s3.us-west-1.amazonaws.com/bucket/key
	//
	HttpUrl() *string
	// Indicates if this asset is a single file.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	IsFile() *bool
	// Indicates if this asset is a zip archive.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	IsZipArchive() *bool
	// The tree node.
	Node() constructs.Node
	// Attribute that represents the name of the bucket this asset exists in.
	S3BucketName() *string
	// Attribute which represents the S3 object key of this asset.
	S3ObjectKey() *string
	// Attribute which represents the S3 URL of this asset.
	//
	// Example:
	//   s3://bucket/key
	//
	S3ObjectUrl() *string
	// Adds CloudFormation template metadata to the specified resource with information that indicates which resource property is mapped to this local asset.
	//
	// This can be used by tools such as SAM CLI to provide local
	// experience such as local invocation and debugging of Lambda functions.
	//
	// Asset metadata will only be included if the stack is synthesized with the
	// "aws:cdk:enable-asset-metadata" context key defined, which is the default
	// behavior when synthesizing via the CDK Toolkit.
	// See: https://github.com/aws/aws-cdk/issues/1432
	//
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	// Grants read permissions to the principal on the assets bucket.
	GrantRead(grantee awsiam.IGrantable)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AwsCliAsset
type jsiiProxy_AwsCliAsset struct {
	internal.Type__awss3assetsAsset
}

func (j *jsiiProxy_AwsCliAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) AssetPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) HttpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) IsFile() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) IsZipArchive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isZipArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) S3ObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCliAsset) S3ObjectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectUrl",
		&returns,
	)
	return returns
}


func NewAwsCliAsset(scope constructs.Construct, id *string, options *awss3assets.AssetOptions) AwsCliAsset {
	_init_.Initialize()

	j := jsiiProxy_AwsCliAsset{}

	_jsii_.Create(
		"@aws-cdk/asset-awscli-v1.AwsCliAsset",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

func NewAwsCliAsset_Override(a AwsCliAsset, scope constructs.Construct, id *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/asset-awscli-v1.AwsCliAsset",
		[]interface{}{scope, id, options},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AwsCliAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/asset-awscli-v1.AwsCliAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCliAsset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	_jsii_.InvokeVoid(
		a,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

func (a *jsiiProxy_AwsCliAsset) GrantRead(grantee awsiam.IGrantable) {
	_jsii_.InvokeVoid(
		a,
		"grantRead",
		[]interface{}{grantee},
	)
}

func (a *jsiiProxy_AwsCliAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

