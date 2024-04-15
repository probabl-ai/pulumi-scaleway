// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple Flexible IPs.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := scaleway.GetFlexibleIps(ctx, &scaleway.GetFlexibleIpsArgs{
//				Tags: []string{
//					"a tag",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			myOffer, err := scaleway.GetBaremetalOffer(ctx, &scaleway.GetBaremetalOfferArgs{
//				Name: pulumi.StringRef("EM-B112X-SSD"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			base, err := scaleway.NewBaremetalServer(ctx, "base", &scaleway.BaremetalServerArgs{
//				Offer:                  pulumi.String(myOffer.OfferId),
//				InstallConfigAfterward: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewFlexibleIp(ctx, "first", &scaleway.FlexibleIpArgs{
//				ServerId: base.ID(),
//				Tags: pulumi.StringArray{
//					pulumi.String("foo"),
//					pulumi.String("first"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewFlexibleIp(ctx, "second", &scaleway.FlexibleIpArgs{
//				ServerId: base.ID(),
//				Tags: pulumi.StringArray{
//					pulumi.String("foo"),
//					pulumi.String("second"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.GetFlexibleIpsOutput(ctx, scaleway.GetFlexibleIpsOutputArgs{
//				ServerIds: pulumi.StringArray{
//					base.ID(),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetFlexibleIps(ctx *pulumi.Context, args *GetFlexibleIpsArgs, opts ...pulumi.InvokeOption) (*GetFlexibleIpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFlexibleIpsResult
	err := ctx.Invoke("scaleway:index/getFlexibleIps:getFlexibleIps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlexibleIps.
type GetFlexibleIpsArgs struct {
	// (Defaults to provider `projectId`) The ID of the project the IP is in.
	ProjectId *string `pulumi:"projectId"`
	// List of server IDs used as filter. IPs with these exact server IDs are listed.
	ServerIds []string `pulumi:"serverIds"`
	// List of tags used as filter. IPs with these exact tags are listed.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which IPs exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getFlexibleIps.
type GetFlexibleIpsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of found flexible IPS
	Ips []GetFlexibleIpsIp `pulumi:"ips"`
	// (Defaults to provider `organizationId`) The ID of the organization the IP is in.
	OrganizationId string `pulumi:"organizationId"`
	// (Defaults to provider `projectId`) The ID of the project the IP is in.
	ProjectId string   `pulumi:"projectId"`
	ServerIds []string `pulumi:"serverIds"`
	// The list of tags which are attached to the flexible IP.
	Tags []string `pulumi:"tags"`
	// (Defaults to provider `zone`) The zone in which the MAC address exist.
	Zone string `pulumi:"zone"`
}

func GetFlexibleIpsOutput(ctx *pulumi.Context, args GetFlexibleIpsOutputArgs, opts ...pulumi.InvokeOption) GetFlexibleIpsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFlexibleIpsResult, error) {
			args := v.(GetFlexibleIpsArgs)
			r, err := GetFlexibleIps(ctx, &args, opts...)
			var s GetFlexibleIpsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFlexibleIpsResultOutput)
}

// A collection of arguments for invoking getFlexibleIps.
type GetFlexibleIpsOutputArgs struct {
	// (Defaults to provider `projectId`) The ID of the project the IP is in.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// List of server IDs used as filter. IPs with these exact server IDs are listed.
	ServerIds pulumi.StringArrayInput `pulumi:"serverIds"`
	// List of tags used as filter. IPs with these exact tags are listed.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// `zone`) The zone in which IPs exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetFlexibleIpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlexibleIpsArgs)(nil)).Elem()
}

// A collection of values returned by getFlexibleIps.
type GetFlexibleIpsResultOutput struct{ *pulumi.OutputState }

func (GetFlexibleIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlexibleIpsResult)(nil)).Elem()
}

func (o GetFlexibleIpsResultOutput) ToGetFlexibleIpsResultOutput() GetFlexibleIpsResultOutput {
	return o
}

func (o GetFlexibleIpsResultOutput) ToGetFlexibleIpsResultOutputWithContext(ctx context.Context) GetFlexibleIpsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlexibleIpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlexibleIpsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of found flexible IPS
func (o GetFlexibleIpsResultOutput) Ips() GetFlexibleIpsIpArrayOutput {
	return o.ApplyT(func(v GetFlexibleIpsResult) []GetFlexibleIpsIp { return v.Ips }).(GetFlexibleIpsIpArrayOutput)
}

// (Defaults to provider `organizationId`) The ID of the organization the IP is in.
func (o GetFlexibleIpsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlexibleIpsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// (Defaults to provider `projectId`) The ID of the project the IP is in.
func (o GetFlexibleIpsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlexibleIpsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetFlexibleIpsResultOutput) ServerIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlexibleIpsResult) []string { return v.ServerIds }).(pulumi.StringArrayOutput)
}

// The list of tags which are attached to the flexible IP.
func (o GetFlexibleIpsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlexibleIpsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// (Defaults to provider `zone`) The zone in which the MAC address exist.
func (o GetFlexibleIpsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlexibleIpsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlexibleIpsResultOutput{})
}
