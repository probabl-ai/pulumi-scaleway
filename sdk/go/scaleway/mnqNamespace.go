// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates and manages Scaleway Messaging and queuing Namespace.
// For further information please check
// our [documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-index)
//
// > NOTE: This resource refers to the old version of the MNQ API. You should use new resources dedicated to your protocol. SQS, NATS.
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewMnqNamespace(ctx, "main", &scaleway.MnqNamespaceArgs{
//				Protocol: pulumi.String("nats"),
//				Region:   pulumi.String("fr-par"),
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
// Namespaces can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/mnqNamespace:MnqNamespace main fr-par/11111111111111111111111111111111
//
// ```
type MnqNamespace struct {
	pulumi.CustomResourceState

	// The date and time the Namespace was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The endpoint of the service matching the Namespace protocol.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The unique name of the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// `projectId`) The ID of the project the
	// namespace is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The protocol to apply on your namespace. Please check our
	// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// `region`). The region
	// in which the namespace should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The date and time the Namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewMnqNamespace registers a new resource with the given unique name, arguments, and options.
func NewMnqNamespace(ctx *pulumi.Context,
	name string, args *MnqNamespaceArgs, opts ...pulumi.ResourceOption) (*MnqNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqNamespace
	err := ctx.RegisterResource("scaleway:index/mnqNamespace:MnqNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqNamespace gets an existing MnqNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqNamespaceState, opts ...pulumi.ResourceOption) (*MnqNamespace, error) {
	var resource MnqNamespace
	err := ctx.ReadResource("scaleway:index/mnqNamespace:MnqNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqNamespace resources.
type mnqNamespaceState struct {
	// The date and time the Namespace was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The endpoint of the service matching the Namespace protocol.
	Endpoint *string `pulumi:"endpoint"`
	// The unique name of the namespace.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the
	// namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The protocol to apply on your namespace. Please check our
	// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
	Protocol *string `pulumi:"protocol"`
	// `region`). The region
	// in which the namespace should be created.
	Region *string `pulumi:"region"`
	// The date and time the Namespace was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type MnqNamespaceState struct {
	// The date and time the Namespace was created.
	CreatedAt pulumi.StringPtrInput
	// The endpoint of the service matching the Namespace protocol.
	Endpoint pulumi.StringPtrInput
	// The unique name of the namespace.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the
	// namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// The protocol to apply on your namespace. Please check our
	// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
	Protocol pulumi.StringPtrInput
	// `region`). The region
	// in which the namespace should be created.
	Region pulumi.StringPtrInput
	// The date and time the Namespace was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (MnqNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqNamespaceState)(nil)).Elem()
}

type mnqNamespaceArgs struct {
	// The unique name of the namespace.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the
	// namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The protocol to apply on your namespace. Please check our
	// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
	Protocol string `pulumi:"protocol"`
	// `region`). The region
	// in which the namespace should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a MnqNamespace resource.
type MnqNamespaceArgs struct {
	// The unique name of the namespace.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the
	// namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// The protocol to apply on your namespace. Please check our
	// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
	Protocol pulumi.StringInput
	// `region`). The region
	// in which the namespace should be created.
	Region pulumi.StringPtrInput
}

func (MnqNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqNamespaceArgs)(nil)).Elem()
}

type MnqNamespaceInput interface {
	pulumi.Input

	ToMnqNamespaceOutput() MnqNamespaceOutput
	ToMnqNamespaceOutputWithContext(ctx context.Context) MnqNamespaceOutput
}

func (*MnqNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqNamespace)(nil)).Elem()
}

func (i *MnqNamespace) ToMnqNamespaceOutput() MnqNamespaceOutput {
	return i.ToMnqNamespaceOutputWithContext(context.Background())
}

func (i *MnqNamespace) ToMnqNamespaceOutputWithContext(ctx context.Context) MnqNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqNamespaceOutput)
}

func (i *MnqNamespace) ToOutput(ctx context.Context) pulumix.Output[*MnqNamespace] {
	return pulumix.Output[*MnqNamespace]{
		OutputState: i.ToMnqNamespaceOutputWithContext(ctx).OutputState,
	}
}

// MnqNamespaceArrayInput is an input type that accepts MnqNamespaceArray and MnqNamespaceArrayOutput values.
// You can construct a concrete instance of `MnqNamespaceArrayInput` via:
//
//	MnqNamespaceArray{ MnqNamespaceArgs{...} }
type MnqNamespaceArrayInput interface {
	pulumi.Input

	ToMnqNamespaceArrayOutput() MnqNamespaceArrayOutput
	ToMnqNamespaceArrayOutputWithContext(context.Context) MnqNamespaceArrayOutput
}

type MnqNamespaceArray []MnqNamespaceInput

func (MnqNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqNamespace)(nil)).Elem()
}

func (i MnqNamespaceArray) ToMnqNamespaceArrayOutput() MnqNamespaceArrayOutput {
	return i.ToMnqNamespaceArrayOutputWithContext(context.Background())
}

func (i MnqNamespaceArray) ToMnqNamespaceArrayOutputWithContext(ctx context.Context) MnqNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqNamespaceArrayOutput)
}

func (i MnqNamespaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*MnqNamespace] {
	return pulumix.Output[[]*MnqNamespace]{
		OutputState: i.ToMnqNamespaceArrayOutputWithContext(ctx).OutputState,
	}
}

// MnqNamespaceMapInput is an input type that accepts MnqNamespaceMap and MnqNamespaceMapOutput values.
// You can construct a concrete instance of `MnqNamespaceMapInput` via:
//
//	MnqNamespaceMap{ "key": MnqNamespaceArgs{...} }
type MnqNamespaceMapInput interface {
	pulumi.Input

	ToMnqNamespaceMapOutput() MnqNamespaceMapOutput
	ToMnqNamespaceMapOutputWithContext(context.Context) MnqNamespaceMapOutput
}

type MnqNamespaceMap map[string]MnqNamespaceInput

func (MnqNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqNamespace)(nil)).Elem()
}

func (i MnqNamespaceMap) ToMnqNamespaceMapOutput() MnqNamespaceMapOutput {
	return i.ToMnqNamespaceMapOutputWithContext(context.Background())
}

func (i MnqNamespaceMap) ToMnqNamespaceMapOutputWithContext(ctx context.Context) MnqNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqNamespaceMapOutput)
}

func (i MnqNamespaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MnqNamespace] {
	return pulumix.Output[map[string]*MnqNamespace]{
		OutputState: i.ToMnqNamespaceMapOutputWithContext(ctx).OutputState,
	}
}

type MnqNamespaceOutput struct{ *pulumi.OutputState }

func (MnqNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqNamespace)(nil)).Elem()
}

func (o MnqNamespaceOutput) ToMnqNamespaceOutput() MnqNamespaceOutput {
	return o
}

func (o MnqNamespaceOutput) ToMnqNamespaceOutputWithContext(ctx context.Context) MnqNamespaceOutput {
	return o
}

func (o MnqNamespaceOutput) ToOutput(ctx context.Context) pulumix.Output[*MnqNamespace] {
	return pulumix.Output[*MnqNamespace]{
		OutputState: o.OutputState,
	}
}

// The date and time the Namespace was created.
func (o MnqNamespaceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNamespace) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The endpoint of the service matching the Namespace protocol.
func (o MnqNamespaceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNamespace) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The unique name of the namespace.
func (o MnqNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the
// namespace is associated with.
func (o MnqNamespaceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNamespace) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The protocol to apply on your namespace. Please check our
// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
func (o MnqNamespaceOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNamespace) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// `region`). The region
// in which the namespace should be created.
func (o MnqNamespaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNamespace) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The date and time the Namespace was updated.
func (o MnqNamespaceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqNamespace) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type MnqNamespaceArrayOutput struct{ *pulumi.OutputState }

func (MnqNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqNamespace)(nil)).Elem()
}

func (o MnqNamespaceArrayOutput) ToMnqNamespaceArrayOutput() MnqNamespaceArrayOutput {
	return o
}

func (o MnqNamespaceArrayOutput) ToMnqNamespaceArrayOutputWithContext(ctx context.Context) MnqNamespaceArrayOutput {
	return o
}

func (o MnqNamespaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MnqNamespace] {
	return pulumix.Output[[]*MnqNamespace]{
		OutputState: o.OutputState,
	}
}

func (o MnqNamespaceArrayOutput) Index(i pulumi.IntInput) MnqNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqNamespace {
		return vs[0].([]*MnqNamespace)[vs[1].(int)]
	}).(MnqNamespaceOutput)
}

type MnqNamespaceMapOutput struct{ *pulumi.OutputState }

func (MnqNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqNamespace)(nil)).Elem()
}

func (o MnqNamespaceMapOutput) ToMnqNamespaceMapOutput() MnqNamespaceMapOutput {
	return o
}

func (o MnqNamespaceMapOutput) ToMnqNamespaceMapOutputWithContext(ctx context.Context) MnqNamespaceMapOutput {
	return o
}

func (o MnqNamespaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MnqNamespace] {
	return pulumix.Output[map[string]*MnqNamespace]{
		OutputState: o.OutputState,
	}
}

func (o MnqNamespaceMapOutput) MapIndex(k pulumi.StringInput) MnqNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqNamespace {
		return vs[0].(map[string]*MnqNamespace)[vs[1].(string)]
	}).(MnqNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqNamespaceInput)(nil)).Elem(), &MnqNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqNamespaceArrayInput)(nil)).Elem(), MnqNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqNamespaceMapInput)(nil)).Elem(), MnqNamespaceMap{})
	pulumi.RegisterOutputType(MnqNamespaceOutput{})
	pulumi.RegisterOutputType(MnqNamespaceArrayOutput{})
	pulumi.RegisterOutputType(MnqNamespaceMapOutput{})
}
