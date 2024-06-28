// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Private Network.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.LookupVpcPrivateNetwork(ctx, &scaleway.LookupVpcPrivateNetworkArgs{
//				Name: pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupVpcPrivateNetwork(ctx, &scaleway.LookupVpcPrivateNetworkArgs{
//				Name:  pulumi.StringRef("foobar"),
//				VpcId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupVpcPrivateNetwork(ctx, &scaleway.LookupVpcPrivateNetworkArgs{
//				PrivateNetworkId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVpcPrivateNetwork(ctx *pulumi.Context, args *LookupVpcPrivateNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVpcPrivateNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcPrivateNetworkResult
	err := ctx.Invoke("scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPrivateNetwork.
type LookupVpcPrivateNetworkArgs struct {
	// Name of the Private Network. Cannot be used with `privateNetworkId`.
	Name *string `pulumi:"name"`
	// ID of the Private Network. Cannot be used with `name` or `vpcId`.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The ID of the Project the Private Network is associated with.
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// ID of the VPC the Private Network is in. Cannot be used with `privateNetworkId`.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getVpcPrivateNetwork.
type LookupVpcPrivateNetworkResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IPv4 subnet associated with the Private Network.
	Ipv4Subnets []GetVpcPrivateNetworkIpv4Subnet `pulumi:"ipv4Subnets"`
	// The IPv6 subnets associated with the Private Network.
	Ipv6Subnets      []GetVpcPrivateNetworkIpv6Subnet `pulumi:"ipv6Subnets"`
	IsRegional       bool                             `pulumi:"isRegional"`
	Name             *string                          `pulumi:"name"`
	OrganizationId   string                           `pulumi:"organizationId"`
	PrivateNetworkId *string                          `pulumi:"privateNetworkId"`
	ProjectId        *string                          `pulumi:"projectId"`
	Region           *string                          `pulumi:"region"`
	Tags             []string                         `pulumi:"tags"`
	UpdatedAt        string                           `pulumi:"updatedAt"`
	VpcId            *string                          `pulumi:"vpcId"`
	Zone             string                           `pulumi:"zone"`
}

func LookupVpcPrivateNetworkOutput(ctx *pulumi.Context, args LookupVpcPrivateNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPrivateNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcPrivateNetworkResult, error) {
			args := v.(LookupVpcPrivateNetworkArgs)
			r, err := LookupVpcPrivateNetwork(ctx, &args, opts...)
			var s LookupVpcPrivateNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcPrivateNetworkResultOutput)
}

// A collection of arguments for invoking getVpcPrivateNetwork.
type LookupVpcPrivateNetworkOutputArgs struct {
	// Name of the Private Network. Cannot be used with `privateNetworkId`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the Private Network. Cannot be used with `name` or `vpcId`.
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
	// The ID of the Project the Private Network is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	Region    pulumi.StringPtrInput `pulumi:"region"`
	// ID of the VPC the Private Network is in. Cannot be used with `privateNetworkId`.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupVpcPrivateNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPrivateNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPrivateNetwork.
type LookupVpcPrivateNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPrivateNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPrivateNetworkResult)(nil)).Elem()
}

func (o LookupVpcPrivateNetworkResultOutput) ToLookupVpcPrivateNetworkResultOutput() LookupVpcPrivateNetworkResultOutput {
	return o
}

func (o LookupVpcPrivateNetworkResultOutput) ToLookupVpcPrivateNetworkResultOutputWithContext(ctx context.Context) LookupVpcPrivateNetworkResultOutput {
	return o
}

func (o LookupVpcPrivateNetworkResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPrivateNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IPv4 subnet associated with the Private Network.
func (o LookupVpcPrivateNetworkResultOutput) Ipv4Subnets() GetVpcPrivateNetworkIpv4SubnetArrayOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) []GetVpcPrivateNetworkIpv4Subnet { return v.Ipv4Subnets }).(GetVpcPrivateNetworkIpv4SubnetArrayOutput)
}

// The IPv6 subnets associated with the Private Network.
func (o LookupVpcPrivateNetworkResultOutput) Ipv6Subnets() GetVpcPrivateNetworkIpv6SubnetArrayOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) []GetVpcPrivateNetworkIpv6Subnet { return v.Ipv6Subnets }).(GetVpcPrivateNetworkIpv6SubnetArrayOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) IsRegional() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) bool { return v.IsRegional }).(pulumi.BoolOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPrivateNetworkResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPrivateNetworkResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPrivateNetworkResultOutput{})
}
