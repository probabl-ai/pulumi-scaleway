// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a public gateway IP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := scaleway.NewVpcPublicGatewayIP(ctx, "main", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_ = scaleway.LookupVpcPublicGatewayIPOutput(ctx, GetVpcPublicGatewayIPOutputArgs{
// 			IpId: main.ID(),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func LookupVpcPublicGatewayIP(ctx *pulumi.Context, args *LookupVpcPublicGatewayIPArgs, opts ...pulumi.InvokeOption) (*LookupVpcPublicGatewayIPResult, error) {
	var rv LookupVpcPublicGatewayIPResult
	err := ctx.Invoke("scaleway:index/getVpcPublicGatewayIP:getVpcPublicGatewayIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPublicGatewayIP.
type LookupVpcPublicGatewayIPArgs struct {
	IpId *string `pulumi:"ipId"`
}

// A collection of values returned by getVpcPublicGatewayIP.
type LookupVpcPublicGatewayIPResult struct {
	Address   string `pulumi:"address"`
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	IpId           *string  `pulumi:"ipId"`
	OrganizationId string   `pulumi:"organizationId"`
	ProjectId      string   `pulumi:"projectId"`
	Reverse        string   `pulumi:"reverse"`
	Tags           []string `pulumi:"tags"`
	UpdatedAt      string   `pulumi:"updatedAt"`
	Zone           string   `pulumi:"zone"`
}

func LookupVpcPublicGatewayIPOutput(ctx *pulumi.Context, args LookupVpcPublicGatewayIPOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPublicGatewayIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcPublicGatewayIPResult, error) {
			args := v.(LookupVpcPublicGatewayIPArgs)
			r, err := LookupVpcPublicGatewayIP(ctx, &args, opts...)
			return *r, err
		}).(LookupVpcPublicGatewayIPResultOutput)
}

// A collection of arguments for invoking getVpcPublicGatewayIP.
type LookupVpcPublicGatewayIPOutputArgs struct {
	IpId pulumi.StringPtrInput `pulumi:"ipId"`
}

func (LookupVpcPublicGatewayIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayIPArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPublicGatewayIP.
type LookupVpcPublicGatewayIPResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPublicGatewayIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayIPResult)(nil)).Elem()
}

func (o LookupVpcPublicGatewayIPResultOutput) ToLookupVpcPublicGatewayIPResultOutput() LookupVpcPublicGatewayIPResultOutput {
	return o
}

func (o LookupVpcPublicGatewayIPResultOutput) ToLookupVpcPublicGatewayIPResultOutputWithContext(ctx context.Context) LookupVpcPublicGatewayIPResultOutput {
	return o
}

func (o LookupVpcPublicGatewayIPResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPublicGatewayIPResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) IpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) *string { return v.IpId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.Reverse }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayIPResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayIPResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPublicGatewayIPResultOutput{})
}
