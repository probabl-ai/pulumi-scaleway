// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Use this data source to get SSH key information based on its ID or name.
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
//			_, err := scaleway.LookupAccountSshKey(ctx, &scaleway.LookupAccountSshKeyArgs{
//				SshKeyId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAccountSshKey(ctx *pulumi.Context, args *LookupAccountSshKeyArgs, opts ...pulumi.InvokeOption) (*LookupAccountSshKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountSshKeyResult
	err := ctx.Invoke("scaleway:index/getAccountSshKey:getAccountSshKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountSshKey.
type LookupAccountSshKeyArgs struct {
	// The SSH key name. Only one of `name` and `sshKeyId` should be specified.
	Name *string `pulumi:"name"`
	// The SSH key id. Only one of `name` and `sshKeyId` should be specified.
	SshKeyId *string `pulumi:"sshKeyId"`
}

// A collection of values returned by getAccountSshKey.
type LookupAccountSshKeyResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Disabled    bool   `pulumi:"disabled"`
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// The ID of the organization the SSH key is associated with.
	OrganizationId string `pulumi:"organizationId"`
	ProjectId      string `pulumi:"projectId"`
	// The SSH public key string
	PublicKey string  `pulumi:"publicKey"`
	SshKeyId  *string `pulumi:"sshKeyId"`
	UpdatedAt string  `pulumi:"updatedAt"`
}

func LookupAccountSshKeyOutput(ctx *pulumi.Context, args LookupAccountSshKeyOutputArgs, opts ...pulumi.InvokeOption) LookupAccountSshKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountSshKeyResult, error) {
			args := v.(LookupAccountSshKeyArgs)
			r, err := LookupAccountSshKey(ctx, &args, opts...)
			var s LookupAccountSshKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountSshKeyResultOutput)
}

// A collection of arguments for invoking getAccountSshKey.
type LookupAccountSshKeyOutputArgs struct {
	// The SSH key name. Only one of `name` and `sshKeyId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The SSH key id. Only one of `name` and `sshKeyId` should be specified.
	SshKeyId pulumi.StringPtrInput `pulumi:"sshKeyId"`
}

func (LookupAccountSshKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountSshKeyArgs)(nil)).Elem()
}

// A collection of values returned by getAccountSshKey.
type LookupAccountSshKeyResultOutput struct{ *pulumi.OutputState }

func (LookupAccountSshKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountSshKeyResult)(nil)).Elem()
}

func (o LookupAccountSshKeyResultOutput) ToLookupAccountSshKeyResultOutput() LookupAccountSshKeyResultOutput {
	return o
}

func (o LookupAccountSshKeyResultOutput) ToLookupAccountSshKeyResultOutputWithContext(ctx context.Context) LookupAccountSshKeyResultOutput {
	return o
}

func (o LookupAccountSshKeyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAccountSshKeyResult] {
	return pulumix.Output[LookupAccountSshKeyResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupAccountSshKeyResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupAccountSshKeyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccountSshKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the organization the SSH key is associated with.
func (o LookupAccountSshKeyResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The SSH public key string
func (o LookupAccountSshKeyResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o LookupAccountSshKeyResultOutput) SshKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) *string { return v.SshKeyId }).(pulumi.StringPtrOutput)
}

func (o LookupAccountSshKeyResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountSshKeyResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountSshKeyResultOutput{})
}
