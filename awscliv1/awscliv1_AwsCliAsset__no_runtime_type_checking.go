//go:build no_runtime_type_checking

// An Asset construct that contains the AWS CLI, for use in Lambda Layers
package awscliv1

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCliAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	return nil
}

func (a *jsiiProxy_AwsCliAsset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateAwsCliAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAwsCliAssetParameters(scope constructs.Construct, id *string, options *awss3assets.AssetOptions) error {
	return nil
}

