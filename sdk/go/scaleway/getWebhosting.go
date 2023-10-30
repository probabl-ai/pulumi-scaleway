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

// Gets information about a webhosting.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.LookupWebhosting(ctx, &scaleway.LookupWebhostingArgs{
//				Domain: pulumi.StringRef("foobar.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupWebhosting(ctx, &scaleway.LookupWebhostingArgs{
//				WebhostingId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupWebhosting(ctx *pulumi.Context, args *LookupWebhostingArgs, opts ...pulumi.InvokeOption) (*LookupWebhostingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebhostingResult
	err := ctx.Invoke("scaleway:index/getWebhosting:getWebhosting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebhosting.
type LookupWebhostingArgs struct {
	// The hosting domain name. Only one of `domain` and `webhostingId` should be specified.
	Domain *string `pulumi:"domain"`
	// The ID of the organization the hosting is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the hosting is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The hosting id. Only one of `domain` and `webhostingId` should be specified.
	WebhostingId *string `pulumi:"webhostingId"`
}

// A collection of values returned by getWebhosting.
type LookupWebhostingResult struct {
	CpanelUrls []GetWebhostingCpanelUrl `pulumi:"cpanelUrls"`
	CreatedAt  string                   `pulumi:"createdAt"`
	DnsStatus  string                   `pulumi:"dnsStatus"`
	Domain     *string                  `pulumi:"domain"`
	Email      string                   `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                `pulumi:"id"`
	OfferId          string                `pulumi:"offerId"`
	OfferName        string                `pulumi:"offerName"`
	OptionIds        []string              `pulumi:"optionIds"`
	Options          []GetWebhostingOption `pulumi:"options"`
	OrganizationId   string                `pulumi:"organizationId"`
	PlatformHostname string                `pulumi:"platformHostname"`
	PlatformNumber   int                   `pulumi:"platformNumber"`
	ProjectId        *string               `pulumi:"projectId"`
	Region           string                `pulumi:"region"`
	Status           string                `pulumi:"status"`
	Tags             []string              `pulumi:"tags"`
	UpdatedAt        string                `pulumi:"updatedAt"`
	Username         string                `pulumi:"username"`
	WebhostingId     *string               `pulumi:"webhostingId"`
}

func LookupWebhostingOutput(ctx *pulumi.Context, args LookupWebhostingOutputArgs, opts ...pulumi.InvokeOption) LookupWebhostingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebhostingResult, error) {
			args := v.(LookupWebhostingArgs)
			r, err := LookupWebhosting(ctx, &args, opts...)
			var s LookupWebhostingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebhostingResultOutput)
}

// A collection of arguments for invoking getWebhosting.
type LookupWebhostingOutputArgs struct {
	// The hosting domain name. Only one of `domain` and `webhostingId` should be specified.
	Domain pulumi.StringPtrInput `pulumi:"domain"`
	// The ID of the organization the hosting is associated with.
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the hosting is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The hosting id. Only one of `domain` and `webhostingId` should be specified.
	WebhostingId pulumi.StringPtrInput `pulumi:"webhostingId"`
}

func (LookupWebhostingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhostingArgs)(nil)).Elem()
}

// A collection of values returned by getWebhosting.
type LookupWebhostingResultOutput struct{ *pulumi.OutputState }

func (LookupWebhostingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhostingResult)(nil)).Elem()
}

func (o LookupWebhostingResultOutput) ToLookupWebhostingResultOutput() LookupWebhostingResultOutput {
	return o
}

func (o LookupWebhostingResultOutput) ToLookupWebhostingResultOutputWithContext(ctx context.Context) LookupWebhostingResultOutput {
	return o
}

func (o LookupWebhostingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebhostingResult] {
	return pulumix.Output[LookupWebhostingResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupWebhostingResultOutput) CpanelUrls() GetWebhostingCpanelUrlArrayOutput {
	return o.ApplyT(func(v LookupWebhostingResult) []GetWebhostingCpanelUrl { return v.CpanelUrls }).(GetWebhostingCpanelUrlArrayOutput)
}

func (o LookupWebhostingResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) DnsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.DnsStatus }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhostingResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o LookupWebhostingResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.Email }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupWebhostingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.OfferId }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) OfferName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.OfferName }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) OptionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebhostingResult) []string { return v.OptionIds }).(pulumi.StringArrayOutput)
}

func (o LookupWebhostingResultOutput) Options() GetWebhostingOptionArrayOutput {
	return o.ApplyT(func(v LookupWebhostingResult) []GetWebhostingOption { return v.Options }).(GetWebhostingOptionArrayOutput)
}

func (o LookupWebhostingResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) PlatformHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.PlatformHostname }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) PlatformNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebhostingResult) int { return v.PlatformNumber }).(pulumi.IntOutput)
}

func (o LookupWebhostingResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhostingResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupWebhostingResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebhostingResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupWebhostingResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhostingResult) string { return v.Username }).(pulumi.StringOutput)
}

func (o LookupWebhostingResultOutput) WebhostingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhostingResult) *string { return v.WebhostingId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebhostingResultOutput{})
}
