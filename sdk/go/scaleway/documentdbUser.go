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

// Creates and manages Scaleway Database DocumentDB Users.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
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
//			dbPassword, err := random.NewRandomPassword(ctx, "dbPassword", &random.RandomPasswordArgs{
//				Length:  pulumi.Int(16),
//				Special: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDocumentdbUser(ctx, "dbAdmin", &scaleway.DocumentdbUserArgs{
//				InstanceId: pulumi.String("11111111-1111-1111-1111-111111111111"),
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
// Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/documentdbUser:DocumentdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
// ```
type DocumentdbUser struct {
	pulumi.CustomResourceState

	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Grant admin permissions to the Database User.
	IsAdmin pulumi.BoolPtrOutput `pulumi:"isAdmin"`
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name pulumi.StringOutput `pulumi:"name"`
	// Database User password.
	Password pulumi.StringOutput `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewDocumentdbUser registers a new resource with the given unique name, arguments, and options.
func NewDocumentdbUser(ctx *pulumi.Context,
	name string, args *DocumentdbUserArgs, opts ...pulumi.ResourceOption) (*DocumentdbUser, error) {
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
	var resource DocumentdbUser
	err := ctx.RegisterResource("scaleway:index/documentdbUser:DocumentdbUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentdbUser gets an existing DocumentdbUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentdbUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentdbUserState, opts ...pulumi.ResourceOption) (*DocumentdbUser, error) {
	var resource DocumentdbUser
	err := ctx.ReadResource("scaleway:index/documentdbUser:DocumentdbUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentdbUser resources.
type documentdbUserState struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId *string `pulumi:"instanceId"`
	// Grant admin permissions to the Database User.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name *string `pulumi:"name"`
	// Database User password.
	Password *string `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region *string `pulumi:"region"`
}

type DocumentdbUserState struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId pulumi.StringPtrInput
	// Grant admin permissions to the Database User.
	IsAdmin pulumi.BoolPtrInput
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name pulumi.StringPtrInput
	// Database User password.
	Password pulumi.StringPtrInput
	// The Scaleway region this resource resides in.
	Region pulumi.StringPtrInput
}

func (DocumentdbUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbUserState)(nil)).Elem()
}

type documentdbUserArgs struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId string `pulumi:"instanceId"`
	// Grant admin permissions to the Database User.
	IsAdmin *bool `pulumi:"isAdmin"`
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name *string `pulumi:"name"`
	// Database User password.
	Password string `pulumi:"password"`
	// The Scaleway region this resource resides in.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DocumentdbUser resource.
type DocumentdbUserArgs struct {
	// UUID of the documentDB instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database User.
	InstanceId pulumi.StringInput
	// Grant admin permissions to the Database User.
	IsAdmin pulumi.BoolPtrInput
	// Database Username.
	//
	// > **Important:** Updates to `name` will recreate the Database User.
	Name pulumi.StringPtrInput
	// Database User password.
	Password pulumi.StringInput
	// The Scaleway region this resource resides in.
	Region pulumi.StringPtrInput
}

func (DocumentdbUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbUserArgs)(nil)).Elem()
}

type DocumentdbUserInput interface {
	pulumi.Input

	ToDocumentdbUserOutput() DocumentdbUserOutput
	ToDocumentdbUserOutputWithContext(ctx context.Context) DocumentdbUserOutput
}

func (*DocumentdbUser) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbUser)(nil)).Elem()
}

func (i *DocumentdbUser) ToDocumentdbUserOutput() DocumentdbUserOutput {
	return i.ToDocumentdbUserOutputWithContext(context.Background())
}

func (i *DocumentdbUser) ToDocumentdbUserOutputWithContext(ctx context.Context) DocumentdbUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbUserOutput)
}

// DocumentdbUserArrayInput is an input type that accepts DocumentdbUserArray and DocumentdbUserArrayOutput values.
// You can construct a concrete instance of `DocumentdbUserArrayInput` via:
//
//	DocumentdbUserArray{ DocumentdbUserArgs{...} }
type DocumentdbUserArrayInput interface {
	pulumi.Input

	ToDocumentdbUserArrayOutput() DocumentdbUserArrayOutput
	ToDocumentdbUserArrayOutputWithContext(context.Context) DocumentdbUserArrayOutput
}

type DocumentdbUserArray []DocumentdbUserInput

func (DocumentdbUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbUser)(nil)).Elem()
}

func (i DocumentdbUserArray) ToDocumentdbUserArrayOutput() DocumentdbUserArrayOutput {
	return i.ToDocumentdbUserArrayOutputWithContext(context.Background())
}

func (i DocumentdbUserArray) ToDocumentdbUserArrayOutputWithContext(ctx context.Context) DocumentdbUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbUserArrayOutput)
}

// DocumentdbUserMapInput is an input type that accepts DocumentdbUserMap and DocumentdbUserMapOutput values.
// You can construct a concrete instance of `DocumentdbUserMapInput` via:
//
//	DocumentdbUserMap{ "key": DocumentdbUserArgs{...} }
type DocumentdbUserMapInput interface {
	pulumi.Input

	ToDocumentdbUserMapOutput() DocumentdbUserMapOutput
	ToDocumentdbUserMapOutputWithContext(context.Context) DocumentdbUserMapOutput
}

type DocumentdbUserMap map[string]DocumentdbUserInput

func (DocumentdbUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbUser)(nil)).Elem()
}

func (i DocumentdbUserMap) ToDocumentdbUserMapOutput() DocumentdbUserMapOutput {
	return i.ToDocumentdbUserMapOutputWithContext(context.Background())
}

func (i DocumentdbUserMap) ToDocumentdbUserMapOutputWithContext(ctx context.Context) DocumentdbUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbUserMapOutput)
}

type DocumentdbUserOutput struct{ *pulumi.OutputState }

func (DocumentdbUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbUser)(nil)).Elem()
}

func (o DocumentdbUserOutput) ToDocumentdbUserOutput() DocumentdbUserOutput {
	return o
}

func (o DocumentdbUserOutput) ToDocumentdbUserOutputWithContext(ctx context.Context) DocumentdbUserOutput {
	return o
}

// UUID of the documentDB instance.
//
// > **Important:** Updates to `instanceId` will recreate the Database User.
func (o DocumentdbUserOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbUser) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Grant admin permissions to the Database User.
func (o DocumentdbUserOutput) IsAdmin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DocumentdbUser) pulumi.BoolPtrOutput { return v.IsAdmin }).(pulumi.BoolPtrOutput)
}

// Database Username.
//
// > **Important:** Updates to `name` will recreate the Database User.
func (o DocumentdbUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Database User password.
func (o DocumentdbUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The Scaleway region this resource resides in.
func (o DocumentdbUserOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbUser) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type DocumentdbUserArrayOutput struct{ *pulumi.OutputState }

func (DocumentdbUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbUser)(nil)).Elem()
}

func (o DocumentdbUserArrayOutput) ToDocumentdbUserArrayOutput() DocumentdbUserArrayOutput {
	return o
}

func (o DocumentdbUserArrayOutput) ToDocumentdbUserArrayOutputWithContext(ctx context.Context) DocumentdbUserArrayOutput {
	return o
}

func (o DocumentdbUserArrayOutput) Index(i pulumi.IntInput) DocumentdbUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentdbUser {
		return vs[0].([]*DocumentdbUser)[vs[1].(int)]
	}).(DocumentdbUserOutput)
}

type DocumentdbUserMapOutput struct{ *pulumi.OutputState }

func (DocumentdbUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbUser)(nil)).Elem()
}

func (o DocumentdbUserMapOutput) ToDocumentdbUserMapOutput() DocumentdbUserMapOutput {
	return o
}

func (o DocumentdbUserMapOutput) ToDocumentdbUserMapOutputWithContext(ctx context.Context) DocumentdbUserMapOutput {
	return o
}

func (o DocumentdbUserMapOutput) MapIndex(k pulumi.StringInput) DocumentdbUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentdbUser {
		return vs[0].(map[string]*DocumentdbUser)[vs[1].(string)]
	}).(DocumentdbUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbUserInput)(nil)).Elem(), &DocumentdbUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbUserArrayInput)(nil)).Elem(), DocumentdbUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbUserMapInput)(nil)).Elem(), DocumentdbUserMap{})
	pulumi.RegisterOutputType(DocumentdbUserOutput{})
	pulumi.RegisterOutputType(DocumentdbUserArrayOutput{})
	pulumi.RegisterOutputType(DocumentdbUserMapOutput{})
}
