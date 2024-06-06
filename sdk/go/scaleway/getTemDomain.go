// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a transactional email domain.
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
//			_, err := scaleway.LookupTemDomain(ctx, &scaleway.LookupTemDomainArgs{
//				DomainId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTemDomain(ctx *pulumi.Context, args *LookupTemDomainArgs, opts ...pulumi.InvokeOption) (*LookupTemDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTemDomainResult
	err := ctx.Invoke("scaleway:index/getTemDomain:getTemDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemDomain.
type LookupTemDomainArgs struct {
	// The domain id.
	// Only one of `name` and `domainId` should be specified.
	DomainId *string `pulumi:"domainId"`
	// The domain name.
	// Only one of `name` and `domainId` should be specified.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the domain exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getTemDomain.
type LookupTemDomainResult struct {
	AcceptTos   bool    `pulumi:"acceptTos"`
	CreatedAt   string  `pulumi:"createdAt"`
	DkimConfig  string  `pulumi:"dkimConfig"`
	DmarcConfig string  `pulumi:"dmarcConfig"`
	DmarcName   string  `pulumi:"dmarcName"`
	DomainId    *string `pulumi:"domainId"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string                   `pulumi:"id"`
	LastError            string                   `pulumi:"lastError"`
	LastValidAt          string                   `pulumi:"lastValidAt"`
	MxBlackhole          string                   `pulumi:"mxBlackhole"`
	Name                 *string                  `pulumi:"name"`
	NextCheckAt          string                   `pulumi:"nextCheckAt"`
	ProjectId            *string                  `pulumi:"projectId"`
	Region               *string                  `pulumi:"region"`
	Reputations          []GetTemDomainReputation `pulumi:"reputations"`
	RevokedAt            string                   `pulumi:"revokedAt"`
	SmtpHost             string                   `pulumi:"smtpHost"`
	SmtpPort             int                      `pulumi:"smtpPort"`
	SmtpPortAlternative  int                      `pulumi:"smtpPortAlternative"`
	SmtpPortUnsecure     int                      `pulumi:"smtpPortUnsecure"`
	SmtpsAuthUser        string                   `pulumi:"smtpsAuthUser"`
	SmtpsPort            int                      `pulumi:"smtpsPort"`
	SmtpsPortAlternative int                      `pulumi:"smtpsPortAlternative"`
	SpfConfig            string                   `pulumi:"spfConfig"`
	Status               string                   `pulumi:"status"`
}

func LookupTemDomainOutput(ctx *pulumi.Context, args LookupTemDomainOutputArgs, opts ...pulumi.InvokeOption) LookupTemDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemDomainResult, error) {
			args := v.(LookupTemDomainArgs)
			r, err := LookupTemDomain(ctx, &args, opts...)
			var s LookupTemDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemDomainResultOutput)
}

// A collection of arguments for invoking getTemDomain.
type LookupTemDomainOutputArgs struct {
	// The domain id.
	// Only one of `name` and `domainId` should be specified.
	DomainId pulumi.StringPtrInput `pulumi:"domainId"`
	// The domain name.
	// Only one of `name` and `domainId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the domain exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupTemDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemDomainArgs)(nil)).Elem()
}

// A collection of values returned by getTemDomain.
type LookupTemDomainResultOutput struct{ *pulumi.OutputState }

func (LookupTemDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemDomainResult)(nil)).Elem()
}

func (o LookupTemDomainResultOutput) ToLookupTemDomainResultOutput() LookupTemDomainResultOutput {
	return o
}

func (o LookupTemDomainResultOutput) ToLookupTemDomainResultOutputWithContext(ctx context.Context) LookupTemDomainResultOutput {
	return o
}

func (o LookupTemDomainResultOutput) AcceptTos() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTemDomainResult) bool { return v.AcceptTos }).(pulumi.BoolOutput)
}

func (o LookupTemDomainResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) DkimConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.DkimConfig }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) DmarcConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.DmarcConfig }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) DmarcName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.DmarcName }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemDomainResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTemDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) LastError() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.LastError }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) LastValidAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.LastValidAt }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) MxBlackhole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.MxBlackhole }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemDomainResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupTemDomainResultOutput) NextCheckAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.NextCheckAt }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemDomainResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupTemDomainResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemDomainResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupTemDomainResultOutput) Reputations() GetTemDomainReputationArrayOutput {
	return o.ApplyT(func(v LookupTemDomainResult) []GetTemDomainReputation { return v.Reputations }).(GetTemDomainReputationArrayOutput)
}

func (o LookupTemDomainResultOutput) RevokedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.RevokedAt }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) SmtpHost() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.SmtpHost }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) SmtpPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTemDomainResult) int { return v.SmtpPort }).(pulumi.IntOutput)
}

func (o LookupTemDomainResultOutput) SmtpPortAlternative() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTemDomainResult) int { return v.SmtpPortAlternative }).(pulumi.IntOutput)
}

func (o LookupTemDomainResultOutput) SmtpPortUnsecure() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTemDomainResult) int { return v.SmtpPortUnsecure }).(pulumi.IntOutput)
}

func (o LookupTemDomainResultOutput) SmtpsAuthUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.SmtpsAuthUser }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) SmtpsPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTemDomainResult) int { return v.SmtpsPort }).(pulumi.IntOutput)
}

func (o LookupTemDomainResultOutput) SmtpsPortAlternative() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTemDomainResult) int { return v.SmtpsPortAlternative }).(pulumi.IntOutput)
}

func (o LookupTemDomainResultOutput) SpfConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.SpfConfig }).(pulumi.StringOutput)
}

func (o LookupTemDomainResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemDomainResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemDomainResultOutput{})
}
