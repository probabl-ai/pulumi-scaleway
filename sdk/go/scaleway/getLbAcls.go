// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple Load Balancer ACLs.
//
// For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/reference-content/acls/) or [API reference](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls-get-an-acl).
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
//			_, err := scaleway.GetLbAcls(ctx, &scaleway.GetLbAclsArgs{
//				FrontendId: scaleway_lb_frontend.Frt01.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.GetLbAcls(ctx, &scaleway.GetLbAclsArgs{
//				FrontendId: scaleway_lb_frontend.Frt01.Id,
//				Name:       pulumi.StringRef("tf-acls-datasource"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLbAcls(ctx *pulumi.Context, args *GetLbAclsArgs, opts ...pulumi.InvokeOption) (*GetLbAclsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLbAclsResult
	err := ctx.Invoke("scaleway:index/getLbAcls:getLbAcls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbAcls.
type GetLbAclsArgs struct {
	// The frontend ID this ACL is attached to. ACLs with a matching frontend ID are listed.
	// > **Important:** LB frontend IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
	FrontendId string `pulumi:"frontendId"`
	// The ACL name to filter for. ACLs with a matching name are listed.
	Name      *string `pulumi:"name"`
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which the ACLs exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLbAcls.
type GetLbAclsResult struct {
	// List of retrieved ACLs
	Acls       []GetLbAclsAcl `pulumi:"acls"`
	FrontendId string         `pulumi:"frontendId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Zone           string  `pulumi:"zone"`
}

func GetLbAclsOutput(ctx *pulumi.Context, args GetLbAclsOutputArgs, opts ...pulumi.InvokeOption) GetLbAclsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLbAclsResult, error) {
			args := v.(GetLbAclsArgs)
			r, err := GetLbAcls(ctx, &args, opts...)
			var s GetLbAclsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLbAclsResultOutput)
}

// A collection of arguments for invoking getLbAcls.
type GetLbAclsOutputArgs struct {
	// The frontend ID this ACL is attached to. ACLs with a matching frontend ID are listed.
	// > **Important:** LB frontend IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
	FrontendId pulumi.StringInput `pulumi:"frontendId"`
	// The ACL name to filter for. ACLs with a matching name are listed.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which the ACLs exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetLbAclsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbAclsArgs)(nil)).Elem()
}

// A collection of values returned by getLbAcls.
type GetLbAclsResultOutput struct{ *pulumi.OutputState }

func (GetLbAclsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbAclsResult)(nil)).Elem()
}

func (o GetLbAclsResultOutput) ToGetLbAclsResultOutput() GetLbAclsResultOutput {
	return o
}

func (o GetLbAclsResultOutput) ToGetLbAclsResultOutputWithContext(ctx context.Context) GetLbAclsResultOutput {
	return o
}

// List of retrieved ACLs
func (o GetLbAclsResultOutput) Acls() GetLbAclsAclArrayOutput {
	return o.ApplyT(func(v GetLbAclsResult) []GetLbAclsAcl { return v.Acls }).(GetLbAclsAclArrayOutput)
}

func (o GetLbAclsResultOutput) FrontendId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbAclsResult) string { return v.FrontendId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLbAclsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbAclsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLbAclsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLbAclsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetLbAclsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbAclsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetLbAclsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbAclsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetLbAclsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbAclsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLbAclsResultOutput{})
}
