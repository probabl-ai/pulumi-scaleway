// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about the privilege on RDB database.
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
//			_, err := scaleway.LookupDatabasePrivilege(ctx, &scaleway.LookupDatabasePrivilegeArgs{
//				DatabaseName: "my-database",
//				InstanceId:   "11111111-1111-111111111111",
//				UserName:     "my-user",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDatabasePrivilege(ctx *pulumi.Context, args *LookupDatabasePrivilegeArgs, opts ...pulumi.InvokeOption) (*LookupDatabasePrivilegeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabasePrivilegeResult
	err := ctx.Invoke("scaleway:index/getDatabasePrivilege:getDatabasePrivilege", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabasePrivilege.
type LookupDatabasePrivilegeArgs struct {
	// The database name.
	DatabaseName string `pulumi:"databaseName"`
	// The RDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// The user name.
	UserName string `pulumi:"userName"`
}

// A collection of values returned by getDatabasePrivilege.
type LookupDatabasePrivilegeResult struct {
	DatabaseName string `pulumi:"databaseName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The permission for this user on the database. Possible values are `readonly`, `readwrite`, `all`
	// , `custom` and `none`.
	Permission string  `pulumi:"permission"`
	Region     *string `pulumi:"region"`
	UserName   string  `pulumi:"userName"`
}

func LookupDatabasePrivilegeOutput(ctx *pulumi.Context, args LookupDatabasePrivilegeOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasePrivilegeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabasePrivilegeResult, error) {
			args := v.(LookupDatabasePrivilegeArgs)
			r, err := LookupDatabasePrivilege(ctx, &args, opts...)
			var s LookupDatabasePrivilegeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabasePrivilegeResultOutput)
}

// A collection of arguments for invoking getDatabasePrivilege.
type LookupDatabasePrivilegeOutputArgs struct {
	// The database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The RDB instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The user name.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (LookupDatabasePrivilegeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrivilegeArgs)(nil)).Elem()
}

// A collection of values returned by getDatabasePrivilege.
type LookupDatabasePrivilegeResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasePrivilegeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrivilegeResult)(nil)).Elem()
}

func (o LookupDatabasePrivilegeResultOutput) ToLookupDatabasePrivilegeResultOutput() LookupDatabasePrivilegeResultOutput {
	return o
}

func (o LookupDatabasePrivilegeResultOutput) ToLookupDatabasePrivilegeResultOutputWithContext(ctx context.Context) LookupDatabasePrivilegeResultOutput {
	return o
}

func (o LookupDatabasePrivilegeResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabasePrivilegeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabasePrivilegeResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The permission for this user on the database. Possible values are `readonly`, `readwrite`, `all`
// , `custom` and `none`.
func (o LookupDatabasePrivilegeResultOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.Permission }).(pulumi.StringOutput)
}

func (o LookupDatabasePrivilegeResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupDatabasePrivilegeResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasePrivilegeResultOutput{})
}
