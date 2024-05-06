// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Create and manage Scaleway RDB database privilege.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#user-and-permissions).
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
//			mainDatabaseInstance, err := scaleway.NewDatabaseInstance(ctx, "mainDatabaseInstance", &scaleway.DatabaseInstanceArgs{
//				NodeType:      pulumi.String("DB-DEV-S"),
//				Engine:        pulumi.String("PostgreSQL-11"),
//				IsHaCluster:   pulumi.Bool(true),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//			})
//			if err != nil {
//				return err
//			}
//			mainDatabase, err := scaleway.NewDatabase(ctx, "mainDatabase", &scaleway.DatabaseArgs{
//				InstanceId: mainDatabaseInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainDatabaseUser, err := scaleway.NewDatabaseUser(ctx, "mainDatabaseUser", &scaleway.DatabaseUserArgs{
//				InstanceId: mainDatabaseInstance.ID(),
//				Password:   pulumi.String("thiZ_is_v&ry_s3cret"),
//				IsAdmin:    pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDatabasePrivilege(ctx, "mainDatabasePrivilege", &scaleway.DatabasePrivilegeArgs{
//				InstanceId:   mainDatabaseInstance.ID(),
//				UserName:     mainDatabaseUser.Name,
//				DatabaseName: mainDatabase.Name,
//				Permission:   pulumi.String("all"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/databasePrivilege:DatabasePrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
// ```
type DatabasePrivilege struct {
	pulumi.CustomResourceState

	// Name of the database (e.g. `my-db-name`).
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// UUID of the rdb instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission pulumi.StringOutput `pulumi:"permission"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// Name of the user (e.g. `my-db-user`).
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewDatabasePrivilege registers a new resource with the given unique name, arguments, and options.
func NewDatabasePrivilege(ctx *pulumi.Context,
	name string, args *DatabasePrivilegeArgs, opts ...pulumi.ResourceOption) (*DatabasePrivilege, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabasePrivilege
	err := ctx.RegisterResource("scaleway:index/databasePrivilege:DatabasePrivilege", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasePrivilege gets an existing DatabasePrivilege resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasePrivilege(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasePrivilegeState, opts ...pulumi.ResourceOption) (*DatabasePrivilege, error) {
	var resource DatabasePrivilege
	err := ctx.ReadResource("scaleway:index/databasePrivilege:DatabasePrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasePrivilege resources.
type databasePrivilegeState struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName *string `pulumi:"databaseName"`
	// UUID of the rdb instance.
	InstanceId *string `pulumi:"instanceId"`
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission *string `pulumi:"permission"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Name of the user (e.g. `my-db-user`).
	UserName *string `pulumi:"userName"`
}

type DatabasePrivilegeState struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName pulumi.StringPtrInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringPtrInput
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Name of the user (e.g. `my-db-user`).
	UserName pulumi.StringPtrInput
}

func (DatabasePrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrivilegeState)(nil)).Elem()
}

type databasePrivilegeArgs struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName string `pulumi:"databaseName"`
	// UUID of the rdb instance.
	InstanceId string `pulumi:"instanceId"`
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission string `pulumi:"permission"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Name of the user (e.g. `my-db-user`).
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a DatabasePrivilege resource.
type DatabasePrivilegeArgs struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName pulumi.StringInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringInput
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission pulumi.StringInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Name of the user (e.g. `my-db-user`).
	UserName pulumi.StringInput
}

func (DatabasePrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrivilegeArgs)(nil)).Elem()
}

type DatabasePrivilegeInput interface {
	pulumi.Input

	ToDatabasePrivilegeOutput() DatabasePrivilegeOutput
	ToDatabasePrivilegeOutputWithContext(ctx context.Context) DatabasePrivilegeOutput
}

func (*DatabasePrivilege) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrivilege)(nil)).Elem()
}

func (i *DatabasePrivilege) ToDatabasePrivilegeOutput() DatabasePrivilegeOutput {
	return i.ToDatabasePrivilegeOutputWithContext(context.Background())
}

func (i *DatabasePrivilege) ToDatabasePrivilegeOutputWithContext(ctx context.Context) DatabasePrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrivilegeOutput)
}

// DatabasePrivilegeArrayInput is an input type that accepts DatabasePrivilegeArray and DatabasePrivilegeArrayOutput values.
// You can construct a concrete instance of `DatabasePrivilegeArrayInput` via:
//
//	DatabasePrivilegeArray{ DatabasePrivilegeArgs{...} }
type DatabasePrivilegeArrayInput interface {
	pulumi.Input

	ToDatabasePrivilegeArrayOutput() DatabasePrivilegeArrayOutput
	ToDatabasePrivilegeArrayOutputWithContext(context.Context) DatabasePrivilegeArrayOutput
}

type DatabasePrivilegeArray []DatabasePrivilegeInput

func (DatabasePrivilegeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabasePrivilege)(nil)).Elem()
}

func (i DatabasePrivilegeArray) ToDatabasePrivilegeArrayOutput() DatabasePrivilegeArrayOutput {
	return i.ToDatabasePrivilegeArrayOutputWithContext(context.Background())
}

func (i DatabasePrivilegeArray) ToDatabasePrivilegeArrayOutputWithContext(ctx context.Context) DatabasePrivilegeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrivilegeArrayOutput)
}

// DatabasePrivilegeMapInput is an input type that accepts DatabasePrivilegeMap and DatabasePrivilegeMapOutput values.
// You can construct a concrete instance of `DatabasePrivilegeMapInput` via:
//
//	DatabasePrivilegeMap{ "key": DatabasePrivilegeArgs{...} }
type DatabasePrivilegeMapInput interface {
	pulumi.Input

	ToDatabasePrivilegeMapOutput() DatabasePrivilegeMapOutput
	ToDatabasePrivilegeMapOutputWithContext(context.Context) DatabasePrivilegeMapOutput
}

type DatabasePrivilegeMap map[string]DatabasePrivilegeInput

func (DatabasePrivilegeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabasePrivilege)(nil)).Elem()
}

func (i DatabasePrivilegeMap) ToDatabasePrivilegeMapOutput() DatabasePrivilegeMapOutput {
	return i.ToDatabasePrivilegeMapOutputWithContext(context.Background())
}

func (i DatabasePrivilegeMap) ToDatabasePrivilegeMapOutputWithContext(ctx context.Context) DatabasePrivilegeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrivilegeMapOutput)
}

type DatabasePrivilegeOutput struct{ *pulumi.OutputState }

func (DatabasePrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrivilege)(nil)).Elem()
}

func (o DatabasePrivilegeOutput) ToDatabasePrivilegeOutput() DatabasePrivilegeOutput {
	return o
}

func (o DatabasePrivilegeOutput) ToDatabasePrivilegeOutputWithContext(ctx context.Context) DatabasePrivilegeOutput {
	return o
}

// Name of the database (e.g. `my-db-name`).
func (o DatabasePrivilegeOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrivilege) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// UUID of the rdb instance.
func (o DatabasePrivilegeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrivilege) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
func (o DatabasePrivilegeOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrivilege) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o DatabasePrivilegeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrivilege) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Name of the user (e.g. `my-db-user`).
func (o DatabasePrivilegeOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrivilege) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type DatabasePrivilegeArrayOutput struct{ *pulumi.OutputState }

func (DatabasePrivilegeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabasePrivilege)(nil)).Elem()
}

func (o DatabasePrivilegeArrayOutput) ToDatabasePrivilegeArrayOutput() DatabasePrivilegeArrayOutput {
	return o
}

func (o DatabasePrivilegeArrayOutput) ToDatabasePrivilegeArrayOutputWithContext(ctx context.Context) DatabasePrivilegeArrayOutput {
	return o
}

func (o DatabasePrivilegeArrayOutput) Index(i pulumi.IntInput) DatabasePrivilegeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabasePrivilege {
		return vs[0].([]*DatabasePrivilege)[vs[1].(int)]
	}).(DatabasePrivilegeOutput)
}

type DatabasePrivilegeMapOutput struct{ *pulumi.OutputState }

func (DatabasePrivilegeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabasePrivilege)(nil)).Elem()
}

func (o DatabasePrivilegeMapOutput) ToDatabasePrivilegeMapOutput() DatabasePrivilegeMapOutput {
	return o
}

func (o DatabasePrivilegeMapOutput) ToDatabasePrivilegeMapOutputWithContext(ctx context.Context) DatabasePrivilegeMapOutput {
	return o
}

func (o DatabasePrivilegeMapOutput) MapIndex(k pulumi.StringInput) DatabasePrivilegeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabasePrivilege {
		return vs[0].(map[string]*DatabasePrivilege)[vs[1].(string)]
	}).(DatabasePrivilegeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrivilegeInput)(nil)).Elem(), &DatabasePrivilege{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrivilegeArrayInput)(nil)).Elem(), DatabasePrivilegeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrivilegeMapInput)(nil)).Elem(), DatabasePrivilegeMap{})
	pulumi.RegisterOutputType(DatabasePrivilegeOutput{})
	pulumi.RegisterOutputType(DatabasePrivilegeArrayOutput{})
	pulumi.RegisterOutputType(DatabasePrivilegeMapOutput{})
}
