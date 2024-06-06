// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about an RDB instance. For further information see our [developers website](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
func LookupDatabaseInstance(ctx *pulumi.Context, args *LookupDatabaseInstanceArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseInstanceResult
	err := ctx.Invoke("scaleway:index/getDatabaseInstance:getDatabaseInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseInstance.
type LookupDatabaseInstanceArgs struct {
	// The RDB instance ID.
	// Only one of `name` and `instanceId` should be specified.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the RDB instance.
	// Only one of `name` and `instanceId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the RDB instance exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getDatabaseInstance.
type LookupDatabaseInstanceResult struct {
	BackupSameRegion        bool   `pulumi:"backupSameRegion"`
	BackupScheduleFrequency int    `pulumi:"backupScheduleFrequency"`
	BackupScheduleRetention int    `pulumi:"backupScheduleRetention"`
	Certificate             string `pulumi:"certificate"`
	DisableBackup           bool   `pulumi:"disableBackup"`
	EndpointIp              string `pulumi:"endpointIp"`
	EndpointPort            int    `pulumi:"endpointPort"`
	Engine                  string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                              `pulumi:"id"`
	InitSettings    map[string]string                   `pulumi:"initSettings"`
	InstanceId      *string                             `pulumi:"instanceId"`
	IsHaCluster     bool                                `pulumi:"isHaCluster"`
	LoadBalancers   []GetDatabaseInstanceLoadBalancer   `pulumi:"loadBalancers"`
	LogsPolicies    []GetDatabaseInstanceLogsPolicy     `pulumi:"logsPolicies"`
	Name            *string                             `pulumi:"name"`
	NodeType        string                              `pulumi:"nodeType"`
	OrganizationId  string                              `pulumi:"organizationId"`
	Password        string                              `pulumi:"password"`
	PrivateNetworks []GetDatabaseInstancePrivateNetwork `pulumi:"privateNetworks"`
	ProjectId       *string                             `pulumi:"projectId"`
	ReadReplicas    []GetDatabaseInstanceReadReplica    `pulumi:"readReplicas"`
	Region          *string                             `pulumi:"region"`
	Settings        map[string]string                   `pulumi:"settings"`
	Tags            []string                            `pulumi:"tags"`
	UserName        string                              `pulumi:"userName"`
	VolumeSizeInGb  int                                 `pulumi:"volumeSizeInGb"`
	VolumeType      string                              `pulumi:"volumeType"`
}

func LookupDatabaseInstanceOutput(ctx *pulumi.Context, args LookupDatabaseInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseInstanceResult, error) {
			args := v.(LookupDatabaseInstanceArgs)
			r, err := LookupDatabaseInstance(ctx, &args, opts...)
			var s LookupDatabaseInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseInstanceResultOutput)
}

// A collection of arguments for invoking getDatabaseInstance.
type LookupDatabaseInstanceOutputArgs struct {
	// The RDB instance ID.
	// Only one of `name` and `instanceId` should be specified.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The name of the RDB instance.
	// Only one of `name` and `instanceId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the RDB instance exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupDatabaseInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseInstance.
type LookupDatabaseInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseInstanceResult)(nil)).Elem()
}

func (o LookupDatabaseInstanceResultOutput) ToLookupDatabaseInstanceResultOutput() LookupDatabaseInstanceResultOutput {
	return o
}

func (o LookupDatabaseInstanceResultOutput) ToLookupDatabaseInstanceResultOutputWithContext(ctx context.Context) LookupDatabaseInstanceResultOutput {
	return o
}

func (o LookupDatabaseInstanceResultOutput) BackupSameRegion() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) bool { return v.BackupSameRegion }).(pulumi.BoolOutput)
}

func (o LookupDatabaseInstanceResultOutput) BackupScheduleFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) int { return v.BackupScheduleFrequency }).(pulumi.IntOutput)
}

func (o LookupDatabaseInstanceResultOutput) BackupScheduleRetention() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) int { return v.BackupScheduleRetention }).(pulumi.IntOutput)
}

func (o LookupDatabaseInstanceResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) DisableBackup() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) bool { return v.DisableBackup }).(pulumi.BoolOutput)
}

func (o LookupDatabaseInstanceResultOutput) EndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.EndpointIp }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) EndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) int { return v.EndpointPort }).(pulumi.IntOutput)
}

func (o LookupDatabaseInstanceResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabaseInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) InitSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) map[string]string { return v.InitSettings }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseInstanceResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseInstanceResultOutput) IsHaCluster() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) bool { return v.IsHaCluster }).(pulumi.BoolOutput)
}

func (o LookupDatabaseInstanceResultOutput) LoadBalancers() GetDatabaseInstanceLoadBalancerArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceLoadBalancer { return v.LoadBalancers }).(GetDatabaseInstanceLoadBalancerArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) LogsPolicies() GetDatabaseInstanceLogsPolicyArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceLogsPolicy { return v.LogsPolicies }).(GetDatabaseInstanceLogsPolicyArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseInstanceResultOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) PrivateNetworks() GetDatabaseInstancePrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstancePrivateNetwork { return v.PrivateNetworks }).(GetDatabaseInstancePrivateNetworkArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseInstanceResultOutput) ReadReplicas() GetDatabaseInstanceReadReplicaArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []GetDatabaseInstanceReadReplica { return v.ReadReplicas }).(GetDatabaseInstanceReadReplicaArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseInstanceResultOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) map[string]string { return v.Settings }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseInstanceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupDatabaseInstanceResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupDatabaseInstanceResultOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) int { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

func (o LookupDatabaseInstanceResultOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseInstanceResult) string { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseInstanceResultOutput{})
}
