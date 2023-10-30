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

// Creates and manages Scaleway Instance Private NICs. For more information, see
// [the documentation](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea).
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
//			_, err := scaleway.NewInstancePrivateNic(ctx, "pnic01", &scaleway.InstancePrivateNicArgs{
//				PrivateNetworkId: pulumi.String("fr-par-1/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
//				ServerId:         pulumi.String("fr-par-1/11111111-1111-1111-1111-111111111111"),
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
// ### With zone
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
//			pn01, err := scaleway.NewVpcPrivateNetwork(ctx, "pn01", &scaleway.VpcPrivateNetworkArgs{
//				Zone: pulumi.String("fr-par-2"),
//			})
//			if err != nil {
//				return err
//			}
//			base, err := scaleway.NewInstanceServer(ctx, "base", &scaleway.InstanceServerArgs{
//				Image: pulumi.String("ubuntu_jammy"),
//				Type:  pulumi.String("DEV1-S"),
//				Zone:  pn01.Zone,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewInstancePrivateNic(ctx, "pnic01", &scaleway.InstancePrivateNicArgs{
//				ServerId:         base.ID(),
//				PrivateNetworkId: pn01.ID(),
//				Zone:             pn01.Zone,
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
// Private NICs can be imported using the `{zone}/{server_id}/{private_nic_id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/instancePrivateNic:InstancePrivateNic pnic01 fr-par-1/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
//
// ```
type InstancePrivateNic struct {
	pulumi.CustomResourceState

	// IPAM ip list, should be for internal use only
	IpIds pulumi.StringArrayOutput `pulumi:"ipIds"`
	// The MAC address of the private NIC.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The ID of the private network attached to.
	PrivateNetworkId pulumi.StringOutput `pulumi:"privateNetworkId"`
	// The ID of the server associated with.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The tags associated with the private NIC.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// `zone`) The zone in which the server must be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstancePrivateNic registers a new resource with the given unique name, arguments, and options.
func NewInstancePrivateNic(ctx *pulumi.Context,
	name string, args *InstancePrivateNicArgs, opts ...pulumi.ResourceOption) (*InstancePrivateNic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateNetworkId'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstancePrivateNic
	err := ctx.RegisterResource("scaleway:index/instancePrivateNic:InstancePrivateNic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstancePrivateNic gets an existing InstancePrivateNic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstancePrivateNic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstancePrivateNicState, opts ...pulumi.ResourceOption) (*InstancePrivateNic, error) {
	var resource InstancePrivateNic
	err := ctx.ReadResource("scaleway:index/instancePrivateNic:InstancePrivateNic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstancePrivateNic resources.
type instancePrivateNicState struct {
	// IPAM ip list, should be for internal use only
	IpIds []string `pulumi:"ipIds"`
	// The MAC address of the private NIC.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network attached to.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The ID of the server associated with.
	ServerId *string `pulumi:"serverId"`
	// The tags associated with the private NIC.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the server must be created.
	Zone *string `pulumi:"zone"`
}

type InstancePrivateNicState struct {
	// IPAM ip list, should be for internal use only
	IpIds pulumi.StringArrayInput
	// The MAC address of the private NIC.
	MacAddress pulumi.StringPtrInput
	// The ID of the private network attached to.
	PrivateNetworkId pulumi.StringPtrInput
	// The ID of the server associated with.
	ServerId pulumi.StringPtrInput
	// The tags associated with the private NIC.
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the server must be created.
	Zone pulumi.StringPtrInput
}

func (InstancePrivateNicState) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePrivateNicState)(nil)).Elem()
}

type instancePrivateNicArgs struct {
	// IPAM ip list, should be for internal use only
	IpIds []string `pulumi:"ipIds"`
	// The ID of the private network attached to.
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// The ID of the server associated with.
	ServerId string `pulumi:"serverId"`
	// The tags associated with the private NIC.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the server must be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstancePrivateNic resource.
type InstancePrivateNicArgs struct {
	// IPAM ip list, should be for internal use only
	IpIds pulumi.StringArrayInput
	// The ID of the private network attached to.
	PrivateNetworkId pulumi.StringInput
	// The ID of the server associated with.
	ServerId pulumi.StringInput
	// The tags associated with the private NIC.
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the server must be created.
	Zone pulumi.StringPtrInput
}

func (InstancePrivateNicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePrivateNicArgs)(nil)).Elem()
}

type InstancePrivateNicInput interface {
	pulumi.Input

	ToInstancePrivateNicOutput() InstancePrivateNicOutput
	ToInstancePrivateNicOutputWithContext(ctx context.Context) InstancePrivateNicOutput
}

func (*InstancePrivateNic) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePrivateNic)(nil)).Elem()
}

func (i *InstancePrivateNic) ToInstancePrivateNicOutput() InstancePrivateNicOutput {
	return i.ToInstancePrivateNicOutputWithContext(context.Background())
}

func (i *InstancePrivateNic) ToInstancePrivateNicOutputWithContext(ctx context.Context) InstancePrivateNicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePrivateNicOutput)
}

func (i *InstancePrivateNic) ToOutput(ctx context.Context) pulumix.Output[*InstancePrivateNic] {
	return pulumix.Output[*InstancePrivateNic]{
		OutputState: i.ToInstancePrivateNicOutputWithContext(ctx).OutputState,
	}
}

// InstancePrivateNicArrayInput is an input type that accepts InstancePrivateNicArray and InstancePrivateNicArrayOutput values.
// You can construct a concrete instance of `InstancePrivateNicArrayInput` via:
//
//	InstancePrivateNicArray{ InstancePrivateNicArgs{...} }
type InstancePrivateNicArrayInput interface {
	pulumi.Input

	ToInstancePrivateNicArrayOutput() InstancePrivateNicArrayOutput
	ToInstancePrivateNicArrayOutputWithContext(context.Context) InstancePrivateNicArrayOutput
}

type InstancePrivateNicArray []InstancePrivateNicInput

func (InstancePrivateNicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancePrivateNic)(nil)).Elem()
}

func (i InstancePrivateNicArray) ToInstancePrivateNicArrayOutput() InstancePrivateNicArrayOutput {
	return i.ToInstancePrivateNicArrayOutputWithContext(context.Background())
}

func (i InstancePrivateNicArray) ToInstancePrivateNicArrayOutputWithContext(ctx context.Context) InstancePrivateNicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePrivateNicArrayOutput)
}

func (i InstancePrivateNicArray) ToOutput(ctx context.Context) pulumix.Output[[]*InstancePrivateNic] {
	return pulumix.Output[[]*InstancePrivateNic]{
		OutputState: i.ToInstancePrivateNicArrayOutputWithContext(ctx).OutputState,
	}
}

// InstancePrivateNicMapInput is an input type that accepts InstancePrivateNicMap and InstancePrivateNicMapOutput values.
// You can construct a concrete instance of `InstancePrivateNicMapInput` via:
//
//	InstancePrivateNicMap{ "key": InstancePrivateNicArgs{...} }
type InstancePrivateNicMapInput interface {
	pulumi.Input

	ToInstancePrivateNicMapOutput() InstancePrivateNicMapOutput
	ToInstancePrivateNicMapOutputWithContext(context.Context) InstancePrivateNicMapOutput
}

type InstancePrivateNicMap map[string]InstancePrivateNicInput

func (InstancePrivateNicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancePrivateNic)(nil)).Elem()
}

func (i InstancePrivateNicMap) ToInstancePrivateNicMapOutput() InstancePrivateNicMapOutput {
	return i.ToInstancePrivateNicMapOutputWithContext(context.Background())
}

func (i InstancePrivateNicMap) ToInstancePrivateNicMapOutputWithContext(ctx context.Context) InstancePrivateNicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePrivateNicMapOutput)
}

func (i InstancePrivateNicMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*InstancePrivateNic] {
	return pulumix.Output[map[string]*InstancePrivateNic]{
		OutputState: i.ToInstancePrivateNicMapOutputWithContext(ctx).OutputState,
	}
}

type InstancePrivateNicOutput struct{ *pulumi.OutputState }

func (InstancePrivateNicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePrivateNic)(nil)).Elem()
}

func (o InstancePrivateNicOutput) ToInstancePrivateNicOutput() InstancePrivateNicOutput {
	return o
}

func (o InstancePrivateNicOutput) ToInstancePrivateNicOutputWithContext(ctx context.Context) InstancePrivateNicOutput {
	return o
}

func (o InstancePrivateNicOutput) ToOutput(ctx context.Context) pulumix.Output[*InstancePrivateNic] {
	return pulumix.Output[*InstancePrivateNic]{
		OutputState: o.OutputState,
	}
}

// IPAM ip list, should be for internal use only
func (o InstancePrivateNicOutput) IpIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePrivateNic) pulumi.StringArrayOutput { return v.IpIds }).(pulumi.StringArrayOutput)
}

// The MAC address of the private NIC.
func (o InstancePrivateNicOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePrivateNic) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// The ID of the private network attached to.
func (o InstancePrivateNicOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePrivateNic) pulumi.StringOutput { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// The ID of the server associated with.
func (o InstancePrivateNicOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePrivateNic) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// The tags associated with the private NIC.
func (o InstancePrivateNicOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstancePrivateNic) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// `zone`) The zone in which the server must be created.
func (o InstancePrivateNicOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePrivateNic) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstancePrivateNicArrayOutput struct{ *pulumi.OutputState }

func (InstancePrivateNicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancePrivateNic)(nil)).Elem()
}

func (o InstancePrivateNicArrayOutput) ToInstancePrivateNicArrayOutput() InstancePrivateNicArrayOutput {
	return o
}

func (o InstancePrivateNicArrayOutput) ToInstancePrivateNicArrayOutputWithContext(ctx context.Context) InstancePrivateNicArrayOutput {
	return o
}

func (o InstancePrivateNicArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*InstancePrivateNic] {
	return pulumix.Output[[]*InstancePrivateNic]{
		OutputState: o.OutputState,
	}
}

func (o InstancePrivateNicArrayOutput) Index(i pulumi.IntInput) InstancePrivateNicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstancePrivateNic {
		return vs[0].([]*InstancePrivateNic)[vs[1].(int)]
	}).(InstancePrivateNicOutput)
}

type InstancePrivateNicMapOutput struct{ *pulumi.OutputState }

func (InstancePrivateNicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancePrivateNic)(nil)).Elem()
}

func (o InstancePrivateNicMapOutput) ToInstancePrivateNicMapOutput() InstancePrivateNicMapOutput {
	return o
}

func (o InstancePrivateNicMapOutput) ToInstancePrivateNicMapOutputWithContext(ctx context.Context) InstancePrivateNicMapOutput {
	return o
}

func (o InstancePrivateNicMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*InstancePrivateNic] {
	return pulumix.Output[map[string]*InstancePrivateNic]{
		OutputState: o.OutputState,
	}
}

func (o InstancePrivateNicMapOutput) MapIndex(k pulumi.StringInput) InstancePrivateNicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstancePrivateNic {
		return vs[0].(map[string]*InstancePrivateNic)[vs[1].(string)]
	}).(InstancePrivateNicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePrivateNicInput)(nil)).Elem(), &InstancePrivateNic{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePrivateNicArrayInput)(nil)).Elem(), InstancePrivateNicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePrivateNicMapInput)(nil)).Elem(), InstancePrivateNicMap{})
	pulumi.RegisterOutputType(InstancePrivateNicOutput{})
	pulumi.RegisterOutputType(InstancePrivateNicArrayOutput{})
	pulumi.RegisterOutputType(InstancePrivateNicMapOutput{})
}
