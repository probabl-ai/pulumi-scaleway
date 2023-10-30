// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func GetDocumentdbLoadBalancerEndpoint(ctx *pulumi.Context, args *GetDocumentdbLoadBalancerEndpointArgs, opts ...pulumi.InvokeOption) (*GetDocumentdbLoadBalancerEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDocumentdbLoadBalancerEndpointResult
	err := ctx.Invoke("scaleway:index/getDocumentdbLoadBalancerEndpoint:getDocumentdbLoadBalancerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointArgs struct {
	InstanceId   *string `pulumi:"instanceId"`
	InstanceName *string `pulumi:"instanceName"`
	ProjectId    *string `pulumi:"projectId"`
	Region       *string `pulumi:"region"`
}

// A collection of values returned by getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointResult struct {
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	InstanceId   string `pulumi:"instanceId"`
	InstanceName string `pulumi:"instanceName"`
	Ip           string `pulumi:"ip"`
	Name         string `pulumi:"name"`
	Port         int    `pulumi:"port"`
	ProjectId    string `pulumi:"projectId"`
	Region       string `pulumi:"region"`
}

func GetDocumentdbLoadBalancerEndpointOutput(ctx *pulumi.Context, args GetDocumentdbLoadBalancerEndpointOutputArgs, opts ...pulumi.InvokeOption) GetDocumentdbLoadBalancerEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDocumentdbLoadBalancerEndpointResult, error) {
			args := v.(GetDocumentdbLoadBalancerEndpointArgs)
			r, err := GetDocumentdbLoadBalancerEndpoint(ctx, &args, opts...)
			var s GetDocumentdbLoadBalancerEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDocumentdbLoadBalancerEndpointResultOutput)
}

// A collection of arguments for invoking getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointOutputArgs struct {
	InstanceId   pulumi.StringPtrInput `pulumi:"instanceId"`
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	ProjectId    pulumi.StringPtrInput `pulumi:"projectId"`
	Region       pulumi.StringPtrInput `pulumi:"region"`
}

func (GetDocumentdbLoadBalancerEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentdbLoadBalancerEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointResultOutput struct{ *pulumi.OutputState }

func (GetDocumentdbLoadBalancerEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentdbLoadBalancerEndpointResult)(nil)).Elem()
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) ToGetDocumentdbLoadBalancerEndpointResultOutput() GetDocumentdbLoadBalancerEndpointResultOutput {
	return o
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) ToGetDocumentdbLoadBalancerEndpointResultOutputWithContext(ctx context.Context) GetDocumentdbLoadBalancerEndpointResultOutput {
	return o
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDocumentdbLoadBalancerEndpointResult] {
	return pulumix.Output[GetDocumentdbLoadBalancerEndpointResult]{
		OutputState: o.OutputState,
	}
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDocumentdbLoadBalancerEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Ip }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDocumentdbLoadBalancerEndpointResultOutput{})
}
