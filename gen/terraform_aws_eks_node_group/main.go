package terraform_aws_eks_node_group

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"terraform_aws_eks_node_group.TerraformAwsEksNodeGroup",
		reflect.TypeOf((*TerraformAwsEksNodeGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalTagMap", GoGetter: "AdditionalTagMap"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "afterClusterJoiningUserdata", GoGetter: "AfterClusterJoiningUserdata"},
			_jsii_.MemberProperty{JsiiProperty: "amiImageId", GoGetter: "AmiImageId"},
			_jsii_.MemberProperty{JsiiProperty: "amiReleaseVersion", GoGetter: "AmiReleaseVersion"},
			_jsii_.MemberProperty{JsiiProperty: "amiType", GoGetter: "AmiType"},
			_jsii_.MemberProperty{JsiiProperty: "associateClusterSecurityGroup", GoGetter: "AssociateClusterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "associatedSecurityGroupIds", GoGetter: "AssociatedSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "beforeClusterJoiningUserdata", GoGetter: "BeforeClusterJoiningUserdata"},
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappings", GoGetter: "BlockDeviceMappings"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapAdditionalOptions", GoGetter: "BootstrapAdditionalOptions"},
			_jsii_.MemberProperty{JsiiProperty: "capacityType", GoGetter: "CapacityType"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterAutoscalerEnabled", GoGetter: "ClusterAutoscalerEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "createBeforeDestroy", GoGetter: "CreateBeforeDestroy"},
			_jsii_.MemberProperty{JsiiProperty: "delimiter", GoGetter: "Delimiter"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "descriptorFormats", GoGetter: "DescriptorFormats"},
			_jsii_.MemberProperty{JsiiProperty: "desiredSize", GoGetter: "DesiredSize"},
			_jsii_.MemberProperty{JsiiProperty: "detailedMonitoringEnabled", GoGetter: "DetailedMonitoringEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "ec2SshKeyName", GoGetter: "Ec2SshKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupArnOutput", GoGetter: "EksNodeGroupArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupCbdPetNameOutput", GoGetter: "EksNodeGroupCbdPetNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupIdOutput", GoGetter: "EksNodeGroupIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupLaunchTemplateIdOutput", GoGetter: "EksNodeGroupLaunchTemplateIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupLaunchTemplateNameOutput", GoGetter: "EksNodeGroupLaunchTemplateNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupRemoteAccessSecurityGroupIdOutput", GoGetter: "EksNodeGroupRemoteAccessSecurityGroupIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupResourcesOutput", GoGetter: "EksNodeGroupResourcesOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupRoleArnOutput", GoGetter: "EksNodeGroupRoleArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupRoleNameOutput", GoGetter: "EksNodeGroupRoleNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupStatusOutput", GoGetter: "EksNodeGroupStatusOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupTagsAllOutput", GoGetter: "EksNodeGroupTagsAllOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eksNodeGroupWindowsNoteOutput", GoGetter: "EksNodeGroupWindowsNoteOutput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enclaveEnabled", GoGetter: "EnclaveEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberProperty{JsiiProperty: "idLengthLimit", GoGetter: "IdLengthLimit"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypes", GoGetter: "InstanceTypes"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "kubeletAdditionalOptions", GoGetter: "KubeletAdditionalOptions"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesLabels", GoGetter: "KubernetesLabels"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesTaints", GoGetter: "KubernetesTaints"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesVersion", GoGetter: "KubernetesVersion"},
			_jsii_.MemberProperty{JsiiProperty: "labelKeyCase", GoGetter: "LabelKeyCase"},
			_jsii_.MemberProperty{JsiiProperty: "labelOrder", GoGetter: "LabelOrder"},
			_jsii_.MemberProperty{JsiiProperty: "labelsAsTags", GoGetter: "LabelsAsTags"},
			_jsii_.MemberProperty{JsiiProperty: "labelValueCase", GoGetter: "LabelValueCase"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateId", GoGetter: "LaunchTemplateId"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateVersion", GoGetter: "LaunchTemplateVersion"},
			_jsii_.MemberProperty{JsiiProperty: "maxSize", GoGetter: "MaxSize"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpEndpointEnabled", GoGetter: "MetadataHttpEndpointEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpPutResponseHopLimit", GoGetter: "MetadataHttpPutResponseHopLimit"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHttpTokensRequired", GoGetter: "MetadataHttpTokensRequired"},
			_jsii_.MemberProperty{JsiiProperty: "minSize", GoGetter: "MinSize"},
			_jsii_.MemberProperty{JsiiProperty: "moduleDependsOn", GoGetter: "ModuleDependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeGroupTerraformTimeouts", GoGetter: "NodeGroupTerraformTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "nodeRoleArn", GoGetter: "NodeRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "nodeRoleCniPolicyEnabled", GoGetter: "NodeRoleCniPolicyEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "nodeRolePermissionsBoundary", GoGetter: "NodeRolePermissionsBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "nodeRolePolicyArns", GoGetter: "NodeRolePolicyArns"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "placement", GoGetter: "Placement"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "regexReplaceChars", GoGetter: "RegexReplaceChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "resourcesToTag", GoGetter: "ResourcesToTag"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sshAccessSecurityGroupIds", GoGetter: "SshAccessSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tenant", GoGetter: "Tenant"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updateConfig", GoGetter: "UpdateConfig"},
			_jsii_.MemberProperty{JsiiProperty: "userdataOverrideBase64", GoGetter: "UserdataOverrideBase64"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformAwsEksNodeGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"terraform_aws_eks_node_group.TerraformAwsEksNodeGroupConfig",
		reflect.TypeOf((*TerraformAwsEksNodeGroupConfig)(nil)).Elem(),
	)
}