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

// Creates and manages Scaleway IAM SSH Keys.
// For more information,
// see [the documentation](https://www.scaleway.com/en/developers/api/iam/#ssh-keys-d8ccd4).
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
//			_, err := scaleway.NewIamSshKey(ctx, "main", &scaleway.IamSshKeyArgs{
//				PublicKey: pulumi.String("<YOUR-PUBLIC-SSH-KEY>"),
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
// SSH keys can be imported using the `id`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/iamSshKey:IamSshKey main 11111111-1111-1111-1111-111111111111
// ```
type IamSshKey struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the SSH key.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The SSH key status.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The fingerprint of the iam SSH key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name of the SSH key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the organization the SSH key is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the SSH key is
	// associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The public SSH key to be added.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The date and time of the last update of the SSH key.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIamSshKey registers a new resource with the given unique name, arguments, and options.
func NewIamSshKey(ctx *pulumi.Context,
	name string, args *IamSshKeyArgs, opts ...pulumi.ResourceOption) (*IamSshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamSshKey
	err := ctx.RegisterResource("scaleway:index/iamSshKey:IamSshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamSshKey gets an existing IamSshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamSshKeyState, opts ...pulumi.ResourceOption) (*IamSshKey, error) {
	var resource IamSshKey
	err := ctx.ReadResource("scaleway:index/iamSshKey:IamSshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamSshKey resources.
type iamSshKeyState struct {
	// The date and time of the creation of the SSH key.
	CreatedAt *string `pulumi:"createdAt"`
	// The SSH key status.
	Disabled *bool `pulumi:"disabled"`
	// The fingerprint of the iam SSH key.
	Fingerprint *string `pulumi:"fingerprint"`
	// The name of the SSH key.
	Name *string `pulumi:"name"`
	// The ID of the organization the SSH key is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the SSH key is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// The public SSH key to be added.
	PublicKey *string `pulumi:"publicKey"`
	// The date and time of the last update of the SSH key.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IamSshKeyState struct {
	// The date and time of the creation of the SSH key.
	CreatedAt pulumi.StringPtrInput
	// The SSH key status.
	Disabled pulumi.BoolPtrInput
	// The fingerprint of the iam SSH key.
	Fingerprint pulumi.StringPtrInput
	// The name of the SSH key.
	Name pulumi.StringPtrInput
	// The ID of the organization the SSH key is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the SSH key is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// The public SSH key to be added.
	PublicKey pulumi.StringPtrInput
	// The date and time of the last update of the SSH key.
	UpdatedAt pulumi.StringPtrInput
}

func (IamSshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamSshKeyState)(nil)).Elem()
}

type iamSshKeyArgs struct {
	// The SSH key status.
	Disabled *bool `pulumi:"disabled"`
	// The name of the SSH key.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the SSH key is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// The public SSH key to be added.
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a IamSshKey resource.
type IamSshKeyArgs struct {
	// The SSH key status.
	Disabled pulumi.BoolPtrInput
	// The name of the SSH key.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the SSH key is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// The public SSH key to be added.
	PublicKey pulumi.StringInput
}

func (IamSshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamSshKeyArgs)(nil)).Elem()
}

type IamSshKeyInput interface {
	pulumi.Input

	ToIamSshKeyOutput() IamSshKeyOutput
	ToIamSshKeyOutputWithContext(ctx context.Context) IamSshKeyOutput
}

func (*IamSshKey) ElementType() reflect.Type {
	return reflect.TypeOf((**IamSshKey)(nil)).Elem()
}

func (i *IamSshKey) ToIamSshKeyOutput() IamSshKeyOutput {
	return i.ToIamSshKeyOutputWithContext(context.Background())
}

func (i *IamSshKey) ToIamSshKeyOutputWithContext(ctx context.Context) IamSshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamSshKeyOutput)
}

// IamSshKeyArrayInput is an input type that accepts IamSshKeyArray and IamSshKeyArrayOutput values.
// You can construct a concrete instance of `IamSshKeyArrayInput` via:
//
//	IamSshKeyArray{ IamSshKeyArgs{...} }
type IamSshKeyArrayInput interface {
	pulumi.Input

	ToIamSshKeyArrayOutput() IamSshKeyArrayOutput
	ToIamSshKeyArrayOutputWithContext(context.Context) IamSshKeyArrayOutput
}

type IamSshKeyArray []IamSshKeyInput

func (IamSshKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamSshKey)(nil)).Elem()
}

func (i IamSshKeyArray) ToIamSshKeyArrayOutput() IamSshKeyArrayOutput {
	return i.ToIamSshKeyArrayOutputWithContext(context.Background())
}

func (i IamSshKeyArray) ToIamSshKeyArrayOutputWithContext(ctx context.Context) IamSshKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamSshKeyArrayOutput)
}

// IamSshKeyMapInput is an input type that accepts IamSshKeyMap and IamSshKeyMapOutput values.
// You can construct a concrete instance of `IamSshKeyMapInput` via:
//
//	IamSshKeyMap{ "key": IamSshKeyArgs{...} }
type IamSshKeyMapInput interface {
	pulumi.Input

	ToIamSshKeyMapOutput() IamSshKeyMapOutput
	ToIamSshKeyMapOutputWithContext(context.Context) IamSshKeyMapOutput
}

type IamSshKeyMap map[string]IamSshKeyInput

func (IamSshKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamSshKey)(nil)).Elem()
}

func (i IamSshKeyMap) ToIamSshKeyMapOutput() IamSshKeyMapOutput {
	return i.ToIamSshKeyMapOutputWithContext(context.Background())
}

func (i IamSshKeyMap) ToIamSshKeyMapOutputWithContext(ctx context.Context) IamSshKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamSshKeyMapOutput)
}

type IamSshKeyOutput struct{ *pulumi.OutputState }

func (IamSshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamSshKey)(nil)).Elem()
}

func (o IamSshKeyOutput) ToIamSshKeyOutput() IamSshKeyOutput {
	return o
}

func (o IamSshKeyOutput) ToIamSshKeyOutputWithContext(ctx context.Context) IamSshKeyOutput {
	return o
}

// The date and time of the creation of the SSH key.
func (o IamSshKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The SSH key status.
func (o IamSshKeyOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The fingerprint of the iam SSH key.
func (o IamSshKeyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The name of the SSH key.
func (o IamSshKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the organization the SSH key is associated with.
func (o IamSshKeyOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the SSH key is
// associated with.
func (o IamSshKeyOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The public SSH key to be added.
func (o IamSshKeyOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// The date and time of the last update of the SSH key.
func (o IamSshKeyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamSshKey) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IamSshKeyArrayOutput struct{ *pulumi.OutputState }

func (IamSshKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamSshKey)(nil)).Elem()
}

func (o IamSshKeyArrayOutput) ToIamSshKeyArrayOutput() IamSshKeyArrayOutput {
	return o
}

func (o IamSshKeyArrayOutput) ToIamSshKeyArrayOutputWithContext(ctx context.Context) IamSshKeyArrayOutput {
	return o
}

func (o IamSshKeyArrayOutput) Index(i pulumi.IntInput) IamSshKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamSshKey {
		return vs[0].([]*IamSshKey)[vs[1].(int)]
	}).(IamSshKeyOutput)
}

type IamSshKeyMapOutput struct{ *pulumi.OutputState }

func (IamSshKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamSshKey)(nil)).Elem()
}

func (o IamSshKeyMapOutput) ToIamSshKeyMapOutput() IamSshKeyMapOutput {
	return o
}

func (o IamSshKeyMapOutput) ToIamSshKeyMapOutputWithContext(ctx context.Context) IamSshKeyMapOutput {
	return o
}

func (o IamSshKeyMapOutput) MapIndex(k pulumi.StringInput) IamSshKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamSshKey {
		return vs[0].(map[string]*IamSshKey)[vs[1].(string)]
	}).(IamSshKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamSshKeyInput)(nil)).Elem(), &IamSshKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamSshKeyArrayInput)(nil)).Elem(), IamSshKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamSshKeyMapInput)(nil)).Elem(), IamSshKeyMap{})
	pulumi.RegisterOutputType(IamSshKeyOutput{})
	pulumi.RegisterOutputType(IamSshKeyArrayOutput{})
	pulumi.RegisterOutputType(IamSshKeyMapOutput{})
}
