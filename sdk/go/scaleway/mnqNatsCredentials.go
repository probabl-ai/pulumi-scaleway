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

// Creates and manages Scaleway Messaging and queuing Nats Credentials.
// For further information please check
// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/)
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainMnqNatsAccount, err := scaleway.NewMnqNatsAccount(ctx, "mainMnqNatsAccount", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewMnqNatsCredentials(ctx, "mainMnqNatsCredentials", &scaleway.MnqNatsCredentialsArgs{
//				AccountId: mainMnqNatsAccount.ID(),
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
// Namespaces can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/mnqNatsCredentials:MnqNatsCredentials main fr-par/11111111111111111111111111111111
// ```
type MnqNatsCredentials struct {
	pulumi.CustomResourceState

	// The ID of the nats account the credentials are generated from
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The content of the credentials file.
	File pulumi.StringOutput `pulumi:"file"`
	// The unique name of the nats credentials.
	Name pulumi.StringOutput `pulumi:"name"`
	// `region`). The region
	// in which the account exists.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewMnqNatsCredentials registers a new resource with the given unique name, arguments, and options.
func NewMnqNatsCredentials(ctx *pulumi.Context,
	name string, args *MnqNatsCredentialsArgs, opts ...pulumi.ResourceOption) (*MnqNatsCredentials, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqNatsCredentials
	err := ctx.RegisterResource("scaleway:index/mnqNatsCredentials:MnqNatsCredentials", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqNatsCredentials gets an existing MnqNatsCredentials resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqNatsCredentials(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqNatsCredentialsState, opts ...pulumi.ResourceOption) (*MnqNatsCredentials, error) {
	var resource MnqNatsCredentials
	err := ctx.ReadResource("scaleway:index/mnqNatsCredentials:MnqNatsCredentials", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqNatsCredentials resources.
type mnqNatsCredentialsState struct {
	// The ID of the nats account the credentials are generated from
	AccountId *string `pulumi:"accountId"`
	// The content of the credentials file.
	File *string `pulumi:"file"`
	// The unique name of the nats credentials.
	Name *string `pulumi:"name"`
	// `region`). The region
	// in which the account exists.
	Region *string `pulumi:"region"`
}

type MnqNatsCredentialsState struct {
	// The ID of the nats account the credentials are generated from
	AccountId pulumi.StringPtrInput
	// The content of the credentials file.
	File pulumi.StringPtrInput
	// The unique name of the nats credentials.
	Name pulumi.StringPtrInput
	// `region`). The region
	// in which the account exists.
	Region pulumi.StringPtrInput
}

func (MnqNatsCredentialsState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqNatsCredentialsState)(nil)).Elem()
}

type mnqNatsCredentialsArgs struct {
	// The ID of the nats account the credentials are generated from
	AccountId string `pulumi:"accountId"`
	// The unique name of the nats credentials.
	Name *string `pulumi:"name"`
	// `region`). The region
	// in which the account exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a MnqNatsCredentials resource.
type MnqNatsCredentialsArgs struct {
	// The ID of the nats account the credentials are generated from
	AccountId pulumi.StringInput
	// The unique name of the nats credentials.
	Name pulumi.StringPtrInput
	// `region`). The region
	// in which the account exists.
	Region pulumi.StringPtrInput
}

func (MnqNatsCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqNatsCredentialsArgs)(nil)).Elem()
}

type MnqNatsCredentialsInput interface {
	pulumi.Input

	ToMnqNatsCredentialsOutput() MnqNatsCredentialsOutput
	ToMnqNatsCredentialsOutputWithContext(ctx context.Context) MnqNatsCredentialsOutput
}

func (*MnqNatsCredentials) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqNatsCredentials)(nil)).Elem()
}

func (i *MnqNatsCredentials) ToMnqNatsCredentialsOutput() MnqNatsCredentialsOutput {
	return i.ToMnqNatsCredentialsOutputWithContext(context.Background())
}

func (i *MnqNatsCredentials) ToMnqNatsCredentialsOutputWithContext(ctx context.Context) MnqNatsCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqNatsCredentialsOutput)
}

// MnqNatsCredentialsArrayInput is an input type that accepts MnqNatsCredentialsArray and MnqNatsCredentialsArrayOutput values.
// You can construct a concrete instance of `MnqNatsCredentialsArrayInput` via:
//
//	MnqNatsCredentialsArray{ MnqNatsCredentialsArgs{...} }
type MnqNatsCredentialsArrayInput interface {
	pulumi.Input

	ToMnqNatsCredentialsArrayOutput() MnqNatsCredentialsArrayOutput
	ToMnqNatsCredentialsArrayOutputWithContext(context.Context) MnqNatsCredentialsArrayOutput
}

type MnqNatsCredentialsArray []MnqNatsCredentialsInput

func (MnqNatsCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqNatsCredentials)(nil)).Elem()
}

func (i MnqNatsCredentialsArray) ToMnqNatsCredentialsArrayOutput() MnqNatsCredentialsArrayOutput {
	return i.ToMnqNatsCredentialsArrayOutputWithContext(context.Background())
}

func (i MnqNatsCredentialsArray) ToMnqNatsCredentialsArrayOutputWithContext(ctx context.Context) MnqNatsCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqNatsCredentialsArrayOutput)
}

// MnqNatsCredentialsMapInput is an input type that accepts MnqNatsCredentialsMap and MnqNatsCredentialsMapOutput values.
// You can construct a concrete instance of `MnqNatsCredentialsMapInput` via:
//
//	MnqNatsCredentialsMap{ "key": MnqNatsCredentialsArgs{...} }
type MnqNatsCredentialsMapInput interface {
	pulumi.Input

	ToMnqNatsCredentialsMapOutput() MnqNatsCredentialsMapOutput
	ToMnqNatsCredentialsMapOutputWithContext(context.Context) MnqNatsCredentialsMapOutput
}

type MnqNatsCredentialsMap map[string]MnqNatsCredentialsInput

func (MnqNatsCredentialsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqNatsCredentials)(nil)).Elem()
}

func (i MnqNatsCredentialsMap) ToMnqNatsCredentialsMapOutput() MnqNatsCredentialsMapOutput {
	return i.ToMnqNatsCredentialsMapOutputWithContext(context.Background())
}

func (i MnqNatsCredentialsMap) ToMnqNatsCredentialsMapOutputWithContext(ctx context.Context) MnqNatsCredentialsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqNatsCredentialsMapOutput)
}

type MnqNatsCredentialsOutput struct{ *pulumi.OutputState }

func (MnqNatsCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqNatsCredentials)(nil)).Elem()
}

func (o MnqNatsCredentialsOutput) ToMnqNatsCredentialsOutput() MnqNatsCredentialsOutput {
	return o
}

func (o MnqNatsCredentialsOutput) ToMnqNatsCredentialsOutputWithContext(ctx context.Context) MnqNatsCredentialsOutput {
	return o
}

// The ID of the nats account the credentials are generated from
func (o MnqNatsCredentialsOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNatsCredentials) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The content of the credentials file.
func (o MnqNatsCredentialsOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNatsCredentials) pulumi.StringOutput { return v.File }).(pulumi.StringOutput)
}

// The unique name of the nats credentials.
func (o MnqNatsCredentialsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNatsCredentials) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `region`). The region
// in which the account exists.
func (o MnqNatsCredentialsOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNatsCredentials) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type MnqNatsCredentialsArrayOutput struct{ *pulumi.OutputState }

func (MnqNatsCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqNatsCredentials)(nil)).Elem()
}

func (o MnqNatsCredentialsArrayOutput) ToMnqNatsCredentialsArrayOutput() MnqNatsCredentialsArrayOutput {
	return o
}

func (o MnqNatsCredentialsArrayOutput) ToMnqNatsCredentialsArrayOutputWithContext(ctx context.Context) MnqNatsCredentialsArrayOutput {
	return o
}

func (o MnqNatsCredentialsArrayOutput) Index(i pulumi.IntInput) MnqNatsCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqNatsCredentials {
		return vs[0].([]*MnqNatsCredentials)[vs[1].(int)]
	}).(MnqNatsCredentialsOutput)
}

type MnqNatsCredentialsMapOutput struct{ *pulumi.OutputState }

func (MnqNatsCredentialsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqNatsCredentials)(nil)).Elem()
}

func (o MnqNatsCredentialsMapOutput) ToMnqNatsCredentialsMapOutput() MnqNatsCredentialsMapOutput {
	return o
}

func (o MnqNatsCredentialsMapOutput) ToMnqNatsCredentialsMapOutputWithContext(ctx context.Context) MnqNatsCredentialsMapOutput {
	return o
}

func (o MnqNatsCredentialsMapOutput) MapIndex(k pulumi.StringInput) MnqNatsCredentialsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqNatsCredentials {
		return vs[0].(map[string]*MnqNatsCredentials)[vs[1].(string)]
	}).(MnqNatsCredentialsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqNatsCredentialsInput)(nil)).Elem(), &MnqNatsCredentials{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqNatsCredentialsArrayInput)(nil)).Elem(), MnqNatsCredentialsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqNatsCredentialsMapInput)(nil)).Elem(), MnqNatsCredentialsMap{})
	pulumi.RegisterOutputType(MnqNatsCredentialsOutput{})
	pulumi.RegisterOutputType(MnqNatsCredentialsArrayOutput{})
	pulumi.RegisterOutputType(MnqNatsCredentialsMapOutput{})
}
