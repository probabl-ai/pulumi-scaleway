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

// Creates and manages database users.
// For more information refer to [the API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/).
//
// ## Example Usage
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := scaleway.NewDatabaseInstance(ctx, "main", &scaleway.DatabaseInstanceArgs{
//				NodeType:      pulumi.String("DB-DEV-S"),
//				Engine:        pulumi.String("PostgreSQL-15"),
//				IsHaCluster:   pulumi.Bool(true),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//			})
//			if err != nil {
//				return err
//			}
//			dbPassword, err := random.NewRandomPassword(ctx, "dbPassword", &random.RandomPasswordArgs{
//				Length:  pulumi.Int(16),
//				Special: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDatabaseUser(ctx, "dbAdmin", &scaleway.DatabaseUserArgs{
//				InstanceId: main.ID(),
//				Password:   dbPassword.Result,
//				IsAdmin:    pulumi.Bool(true),
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
// Database users can be imported using `{region}/{instance_id}/{user_name}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/databaseUser:DatabaseUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
// ```
type DatabaseUser struct {
	pulumi.CustomResourceState

	// UUID of the Database Instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the database user.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Grant admin permissions to the database user.
	IsAdmin pulumi.BoolPtrOutput `pulumi:"isAdmin"`
	// database user name.
	//
	// > **Important:** Updates to `name` will recreate the database user.
	Name pulumi.StringOutput `pulumi:"name"`
	// database user password.
	Password pulumi.StringOutput `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewDatabaseUser registers a new resource with the given unique name, arguments, and options.
func NewDatabaseUser(ctx *pulumi.Context,
	name string, args *DatabaseUserArgs, opts ...pulumi.ResourceOption) (*DatabaseUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseUser
	err := ctx.RegisterResource("scaleway:index/databaseUser:DatabaseUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseUser gets an existing DatabaseUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseUserState, opts ...pulumi.ResourceOption) (*DatabaseUser, error) {
	var resource DatabaseUser
	err := ctx.ReadResource("scaleway:index/databaseUser:DatabaseUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseUser resources.
type databaseUserState struct {
	// UUID of the Database Instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the database user.
	InstanceId *string `pulumi:"instanceId"`
	// Grant admin permissions to the database user.
	IsAdmin *bool `pulumi:"isAdmin"`
	// database user name.
	//
	// > **Important:** Updates to `name` will recreate the database user.
	Name *string `pulumi:"name"`
	// database user password.
	Password *string `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region *string `pulumi:"region"`
}

type DatabaseUserState struct {
	// UUID of the Database Instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the database user.
	InstanceId pulumi.StringPtrInput
	// Grant admin permissions to the database user.
	IsAdmin pulumi.BoolPtrInput
	// database user name.
	//
	// > **Important:** Updates to `name` will recreate the database user.
	Name pulumi.StringPtrInput
	// database user password.
	Password pulumi.StringPtrInput
	// The Scaleway region this resource resides in.
	Region pulumi.StringPtrInput
}

func (DatabaseUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseUserState)(nil)).Elem()
}

type databaseUserArgs struct {
	// UUID of the Database Instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the database user.
	InstanceId string `pulumi:"instanceId"`
	// Grant admin permissions to the database user.
	IsAdmin *bool `pulumi:"isAdmin"`
	// database user name.
	//
	// > **Important:** Updates to `name` will recreate the database user.
	Name *string `pulumi:"name"`
	// database user password.
	Password string `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DatabaseUser resource.
type DatabaseUserArgs struct {
	// UUID of the Database Instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the database user.
	InstanceId pulumi.StringInput
	// Grant admin permissions to the database user.
	IsAdmin pulumi.BoolPtrInput
	// database user name.
	//
	// > **Important:** Updates to `name` will recreate the database user.
	Name pulumi.StringPtrInput
	// database user password.
	Password pulumi.StringInput
	// The Scaleway region this resource resides in.
	Region pulumi.StringPtrInput
}

func (DatabaseUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseUserArgs)(nil)).Elem()
}

type DatabaseUserInput interface {
	pulumi.Input

	ToDatabaseUserOutput() DatabaseUserOutput
	ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput
}

func (*DatabaseUser) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseUser)(nil)).Elem()
}

func (i *DatabaseUser) ToDatabaseUserOutput() DatabaseUserOutput {
	return i.ToDatabaseUserOutputWithContext(context.Background())
}

func (i *DatabaseUser) ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserOutput)
}

// DatabaseUserArrayInput is an input type that accepts DatabaseUserArray and DatabaseUserArrayOutput values.
// You can construct a concrete instance of `DatabaseUserArrayInput` via:
//
//	DatabaseUserArray{ DatabaseUserArgs{...} }
type DatabaseUserArrayInput interface {
	pulumi.Input

	ToDatabaseUserArrayOutput() DatabaseUserArrayOutput
	ToDatabaseUserArrayOutputWithContext(context.Context) DatabaseUserArrayOutput
}

type DatabaseUserArray []DatabaseUserInput

func (DatabaseUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseUser)(nil)).Elem()
}

func (i DatabaseUserArray) ToDatabaseUserArrayOutput() DatabaseUserArrayOutput {
	return i.ToDatabaseUserArrayOutputWithContext(context.Background())
}

func (i DatabaseUserArray) ToDatabaseUserArrayOutputWithContext(ctx context.Context) DatabaseUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserArrayOutput)
}

// DatabaseUserMapInput is an input type that accepts DatabaseUserMap and DatabaseUserMapOutput values.
// You can construct a concrete instance of `DatabaseUserMapInput` via:
//
//	DatabaseUserMap{ "key": DatabaseUserArgs{...} }
type DatabaseUserMapInput interface {
	pulumi.Input

	ToDatabaseUserMapOutput() DatabaseUserMapOutput
	ToDatabaseUserMapOutputWithContext(context.Context) DatabaseUserMapOutput
}

type DatabaseUserMap map[string]DatabaseUserInput

func (DatabaseUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseUser)(nil)).Elem()
}

func (i DatabaseUserMap) ToDatabaseUserMapOutput() DatabaseUserMapOutput {
	return i.ToDatabaseUserMapOutputWithContext(context.Background())
}

func (i DatabaseUserMap) ToDatabaseUserMapOutputWithContext(ctx context.Context) DatabaseUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserMapOutput)
}

type DatabaseUserOutput struct{ *pulumi.OutputState }

func (DatabaseUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserOutput) ToDatabaseUserOutput() DatabaseUserOutput {
	return o
}

func (o DatabaseUserOutput) ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput {
	return o
}

// UUID of the Database Instance.
//
// > **Important:** Updates to `instanceId` will recreate the database user.
func (o DatabaseUserOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Grant admin permissions to the database user.
func (o DatabaseUserOutput) IsAdmin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.BoolPtrOutput { return v.IsAdmin }).(pulumi.BoolPtrOutput)
}

// database user name.
//
// > **Important:** Updates to `name` will recreate the database user.
func (o DatabaseUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// database user password.
func (o DatabaseUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The Scaleway region this resource resides in.
func (o DatabaseUserOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type DatabaseUserArrayOutput struct{ *pulumi.OutputState }

func (DatabaseUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserArrayOutput) ToDatabaseUserArrayOutput() DatabaseUserArrayOutput {
	return o
}

func (o DatabaseUserArrayOutput) ToDatabaseUserArrayOutputWithContext(ctx context.Context) DatabaseUserArrayOutput {
	return o
}

func (o DatabaseUserArrayOutput) Index(i pulumi.IntInput) DatabaseUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseUser {
		return vs[0].([]*DatabaseUser)[vs[1].(int)]
	}).(DatabaseUserOutput)
}

type DatabaseUserMapOutput struct{ *pulumi.OutputState }

func (DatabaseUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserMapOutput) ToDatabaseUserMapOutput() DatabaseUserMapOutput {
	return o
}

func (o DatabaseUserMapOutput) ToDatabaseUserMapOutputWithContext(ctx context.Context) DatabaseUserMapOutput {
	return o
}

func (o DatabaseUserMapOutput) MapIndex(k pulumi.StringInput) DatabaseUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseUser {
		return vs[0].(map[string]*DatabaseUser)[vs[1].(string)]
	}).(DatabaseUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserInput)(nil)).Elem(), &DatabaseUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserArrayInput)(nil)).Elem(), DatabaseUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserMapInput)(nil)).Elem(), DatabaseUserMap{})
	pulumi.RegisterOutputType(DatabaseUserOutput{})
	pulumi.RegisterOutputType(DatabaseUserArrayOutput{})
	pulumi.RegisterOutputType(DatabaseUserMapOutput{})
}
