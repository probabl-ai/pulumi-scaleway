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

// Gets information about DocumentDB database. More on our official [site](https://www.scaleway.com/en/developers/api/document_db/)
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
//			_, err := scaleway.LookupDocumentdbDatabase(ctx, &scaleway.LookupDocumentdbDatabaseArgs{
//				InstanceId: "11111111-1111-1111-1111-111111111111",
//				Name:       pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDocumentdbDatabase(ctx *pulumi.Context, args *LookupDocumentdbDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDocumentdbDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDocumentdbDatabaseResult
	err := ctx.Invoke("scaleway:index/getDocumentdbDatabase:getDocumentdbDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocumentdbDatabase.
type LookupDocumentdbDatabaseArgs struct {
	// The DocumentDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// The name of the DocumentDB instance.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// A collection of values returned by getDocumentdbDatabase.
type LookupDocumentdbDatabaseResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed bool    `pulumi:"managed"`
	Name    *string `pulumi:"name"`
	// The name of the owner of the database.
	Owner     string  `pulumi:"owner"`
	ProjectId string  `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// Size of the database (in bytes).
	Size string `pulumi:"size"`
}

func LookupDocumentdbDatabaseOutput(ctx *pulumi.Context, args LookupDocumentdbDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDocumentdbDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDocumentdbDatabaseResult, error) {
			args := v.(LookupDocumentdbDatabaseArgs)
			r, err := LookupDocumentdbDatabase(ctx, &args, opts...)
			var s LookupDocumentdbDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDocumentdbDatabaseResultOutput)
}

// A collection of arguments for invoking getDocumentdbDatabase.
type LookupDocumentdbDatabaseOutputArgs struct {
	// The DocumentDB instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The name of the DocumentDB instance.
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupDocumentdbDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentdbDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getDocumentdbDatabase.
type LookupDocumentdbDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDocumentdbDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentdbDatabaseResult)(nil)).Elem()
}

func (o LookupDocumentdbDatabaseResultOutput) ToLookupDocumentdbDatabaseResultOutput() LookupDocumentdbDatabaseResultOutput {
	return o
}

func (o LookupDocumentdbDatabaseResultOutput) ToLookupDocumentdbDatabaseResultOutputWithContext(ctx context.Context) LookupDocumentdbDatabaseResultOutput {
	return o
}

func (o LookupDocumentdbDatabaseResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDocumentdbDatabaseResult] {
	return pulumix.Output[LookupDocumentdbDatabaseResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDocumentdbDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDocumentdbDatabaseResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the database is managed or not.
func (o LookupDocumentdbDatabaseResultOutput) Managed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) bool { return v.Managed }).(pulumi.BoolOutput)
}

func (o LookupDocumentdbDatabaseResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The name of the owner of the database.
func (o LookupDocumentdbDatabaseResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupDocumentdbDatabaseResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupDocumentdbDatabaseResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// Size of the database (in bytes).
func (o LookupDocumentdbDatabaseResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentdbDatabaseResult) string { return v.Size }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentdbDatabaseResultOutput{})
}
