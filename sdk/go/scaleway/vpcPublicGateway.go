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

// Creates and manages Scaleway VPC Public Gateway.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/public-gateway).
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
//			_, err := scaleway.NewVpcPublicGateway(ctx, "main", &scaleway.VpcPublicGatewayArgs{
//				Tags: pulumi.StringArray{
//					pulumi.String("demo"),
//					pulumi.String("terraform"),
//				},
//				Type: pulumi.String("VPC-GW-S"),
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
// Public gateway can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/vpcPublicGateway:VpcPublicGateway main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type VpcPublicGateway struct {
	pulumi.CustomResourceState

	// Enable SSH bastion on the gateway.
	BastionEnabled pulumi.BoolPtrOutput `pulumi:"bastionEnabled"`
	// The port on which the SSH bastion will listen.
	BastionPort pulumi.IntOutput `pulumi:"bastionPort"`
	// The date and time of the creation of the public gateway.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Enable SMTP on the gateway.
	EnableSmtp pulumi.BoolOutput `pulumi:"enableSmtp"`
	// attach an existing flexible IP to the gateway.
	IpId pulumi.StringOutput `pulumi:"ipId"`
	// The name of the public gateway. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the public gateway is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Trigger a refresh of the SSH keys on the public gateway by changing this field's value.
	RefreshSshKeys pulumi.StringPtrOutput `pulumi:"refreshSshKeys"`
	// The status of the public gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags associated with the public gateway.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The gateway type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the public gateway.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers pulumi.StringArrayOutput `pulumi:"upstreamDnsServers"`
	// `zone`) The zone in which the public gateway should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGateway registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGateway(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayArgs, opts ...pulumi.ResourceOption) (*VpcPublicGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcPublicGateway
	err := ctx.RegisterResource("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGateway gets an existing VpcPublicGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayState, opts ...pulumi.ResourceOption) (*VpcPublicGateway, error) {
	var resource VpcPublicGateway
	err := ctx.ReadResource("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGateway resources.
type vpcPublicGatewayState struct {
	// Enable SSH bastion on the gateway.
	BastionEnabled *bool `pulumi:"bastionEnabled"`
	// The port on which the SSH bastion will listen.
	BastionPort *int `pulumi:"bastionPort"`
	// The date and time of the creation of the public gateway.
	CreatedAt *string `pulumi:"createdAt"`
	// Enable SMTP on the gateway.
	EnableSmtp *bool `pulumi:"enableSmtp"`
	// attach an existing flexible IP to the gateway.
	IpId *string `pulumi:"ipId"`
	// The name of the public gateway. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the public gateway is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Trigger a refresh of the SSH keys on the public gateway by changing this field's value.
	RefreshSshKeys *string `pulumi:"refreshSshKeys"`
	// The status of the public gateway.
	Status *string `pulumi:"status"`
	// The tags associated with the public gateway.
	Tags []string `pulumi:"tags"`
	// The gateway type.
	Type *string `pulumi:"type"`
	// The date and time of the last update of the public gateway.
	UpdatedAt *string `pulumi:"updatedAt"`
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers []string `pulumi:"upstreamDnsServers"`
	// `zone`) The zone in which the public gateway should be created.
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayState struct {
	// Enable SSH bastion on the gateway.
	BastionEnabled pulumi.BoolPtrInput
	// The port on which the SSH bastion will listen.
	BastionPort pulumi.IntPtrInput
	// The date and time of the creation of the public gateway.
	CreatedAt pulumi.StringPtrInput
	// Enable SMTP on the gateway.
	EnableSmtp pulumi.BoolPtrInput
	// attach an existing flexible IP to the gateway.
	IpId pulumi.StringPtrInput
	// The name of the public gateway. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the public gateway is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId pulumi.StringPtrInput
	// Trigger a refresh of the SSH keys on the public gateway by changing this field's value.
	RefreshSshKeys pulumi.StringPtrInput
	// The status of the public gateway.
	Status pulumi.StringPtrInput
	// The tags associated with the public gateway.
	Tags pulumi.StringArrayInput
	// The gateway type.
	Type pulumi.StringPtrInput
	// The date and time of the last update of the public gateway.
	UpdatedAt pulumi.StringPtrInput
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers pulumi.StringArrayInput
	// `zone`) The zone in which the public gateway should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayState)(nil)).Elem()
}

type vpcPublicGatewayArgs struct {
	// Enable SSH bastion on the gateway.
	BastionEnabled *bool `pulumi:"bastionEnabled"`
	// The port on which the SSH bastion will listen.
	BastionPort *int `pulumi:"bastionPort"`
	// Enable SMTP on the gateway.
	EnableSmtp *bool `pulumi:"enableSmtp"`
	// attach an existing flexible IP to the gateway.
	IpId *string `pulumi:"ipId"`
	// The name of the public gateway. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Trigger a refresh of the SSH keys on the public gateway by changing this field's value.
	RefreshSshKeys *string `pulumi:"refreshSshKeys"`
	// The tags associated with the public gateway.
	Tags []string `pulumi:"tags"`
	// The gateway type.
	Type string `pulumi:"type"`
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers []string `pulumi:"upstreamDnsServers"`
	// `zone`) The zone in which the public gateway should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGateway resource.
type VpcPublicGatewayArgs struct {
	// Enable SSH bastion on the gateway.
	BastionEnabled pulumi.BoolPtrInput
	// The port on which the SSH bastion will listen.
	BastionPort pulumi.IntPtrInput
	// Enable SMTP on the gateway.
	EnableSmtp pulumi.BoolPtrInput
	// attach an existing flexible IP to the gateway.
	IpId pulumi.StringPtrInput
	// The name of the public gateway. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway is associated with.
	ProjectId pulumi.StringPtrInput
	// Trigger a refresh of the SSH keys on the public gateway by changing this field's value.
	RefreshSshKeys pulumi.StringPtrInput
	// The tags associated with the public gateway.
	Tags pulumi.StringArrayInput
	// The gateway type.
	Type pulumi.StringInput
	// override the gateway's default recursive DNS servers, if DNS features are enabled.
	UpstreamDnsServers pulumi.StringArrayInput
	// `zone`) The zone in which the public gateway should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayArgs)(nil)).Elem()
}

type VpcPublicGatewayInput interface {
	pulumi.Input

	ToVpcPublicGatewayOutput() VpcPublicGatewayOutput
	ToVpcPublicGatewayOutputWithContext(ctx context.Context) VpcPublicGatewayOutput
}

func (*VpcPublicGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGateway)(nil)).Elem()
}

func (i *VpcPublicGateway) ToVpcPublicGatewayOutput() VpcPublicGatewayOutput {
	return i.ToVpcPublicGatewayOutputWithContext(context.Background())
}

func (i *VpcPublicGateway) ToVpcPublicGatewayOutputWithContext(ctx context.Context) VpcPublicGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayOutput)
}

// VpcPublicGatewayArrayInput is an input type that accepts VpcPublicGatewayArray and VpcPublicGatewayArrayOutput values.
// You can construct a concrete instance of `VpcPublicGatewayArrayInput` via:
//
//	VpcPublicGatewayArray{ VpcPublicGatewayArgs{...} }
type VpcPublicGatewayArrayInput interface {
	pulumi.Input

	ToVpcPublicGatewayArrayOutput() VpcPublicGatewayArrayOutput
	ToVpcPublicGatewayArrayOutputWithContext(context.Context) VpcPublicGatewayArrayOutput
}

type VpcPublicGatewayArray []VpcPublicGatewayInput

func (VpcPublicGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGateway)(nil)).Elem()
}

func (i VpcPublicGatewayArray) ToVpcPublicGatewayArrayOutput() VpcPublicGatewayArrayOutput {
	return i.ToVpcPublicGatewayArrayOutputWithContext(context.Background())
}

func (i VpcPublicGatewayArray) ToVpcPublicGatewayArrayOutputWithContext(ctx context.Context) VpcPublicGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayArrayOutput)
}

// VpcPublicGatewayMapInput is an input type that accepts VpcPublicGatewayMap and VpcPublicGatewayMapOutput values.
// You can construct a concrete instance of `VpcPublicGatewayMapInput` via:
//
//	VpcPublicGatewayMap{ "key": VpcPublicGatewayArgs{...} }
type VpcPublicGatewayMapInput interface {
	pulumi.Input

	ToVpcPublicGatewayMapOutput() VpcPublicGatewayMapOutput
	ToVpcPublicGatewayMapOutputWithContext(context.Context) VpcPublicGatewayMapOutput
}

type VpcPublicGatewayMap map[string]VpcPublicGatewayInput

func (VpcPublicGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGateway)(nil)).Elem()
}

func (i VpcPublicGatewayMap) ToVpcPublicGatewayMapOutput() VpcPublicGatewayMapOutput {
	return i.ToVpcPublicGatewayMapOutputWithContext(context.Background())
}

func (i VpcPublicGatewayMap) ToVpcPublicGatewayMapOutputWithContext(ctx context.Context) VpcPublicGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayMapOutput)
}

type VpcPublicGatewayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGateway)(nil)).Elem()
}

func (o VpcPublicGatewayOutput) ToVpcPublicGatewayOutput() VpcPublicGatewayOutput {
	return o
}

func (o VpcPublicGatewayOutput) ToVpcPublicGatewayOutputWithContext(ctx context.Context) VpcPublicGatewayOutput {
	return o
}

// Enable SSH bastion on the gateway.
func (o VpcPublicGatewayOutput) BastionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.BoolPtrOutput { return v.BastionEnabled }).(pulumi.BoolPtrOutput)
}

// The port on which the SSH bastion will listen.
func (o VpcPublicGatewayOutput) BastionPort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.IntOutput { return v.BastionPort }).(pulumi.IntOutput)
}

// The date and time of the creation of the public gateway.
func (o VpcPublicGatewayOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Enable SMTP on the gateway.
func (o VpcPublicGatewayOutput) EnableSmtp() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.BoolOutput { return v.EnableSmtp }).(pulumi.BoolOutput)
}

// attach an existing flexible IP to the gateway.
func (o VpcPublicGatewayOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.IpId }).(pulumi.StringOutput)
}

// The name of the public gateway. If not provided it will be randomly generated.
func (o VpcPublicGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the public gateway is associated with.
func (o VpcPublicGatewayOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the public gateway is associated with.
func (o VpcPublicGatewayOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Trigger a refresh of the SSH keys on the public gateway by changing this field's value.
func (o VpcPublicGatewayOutput) RefreshSshKeys() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringPtrOutput { return v.RefreshSshKeys }).(pulumi.StringPtrOutput)
}

// The status of the public gateway.
func (o VpcPublicGatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags associated with the public gateway.
func (o VpcPublicGatewayOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The gateway type.
func (o VpcPublicGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date and time of the last update of the public gateway.
func (o VpcPublicGatewayOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// override the gateway's default recursive DNS servers, if DNS features are enabled.
func (o VpcPublicGatewayOutput) UpstreamDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringArrayOutput { return v.UpstreamDnsServers }).(pulumi.StringArrayOutput)
}

// `zone`) The zone in which the public gateway should be created.
func (o VpcPublicGatewayOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGateway) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcPublicGatewayArrayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGateway)(nil)).Elem()
}

func (o VpcPublicGatewayArrayOutput) ToVpcPublicGatewayArrayOutput() VpcPublicGatewayArrayOutput {
	return o
}

func (o VpcPublicGatewayArrayOutput) ToVpcPublicGatewayArrayOutputWithContext(ctx context.Context) VpcPublicGatewayArrayOutput {
	return o
}

func (o VpcPublicGatewayArrayOutput) Index(i pulumi.IntInput) VpcPublicGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPublicGateway {
		return vs[0].([]*VpcPublicGateway)[vs[1].(int)]
	}).(VpcPublicGatewayOutput)
}

type VpcPublicGatewayMapOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGateway)(nil)).Elem()
}

func (o VpcPublicGatewayMapOutput) ToVpcPublicGatewayMapOutput() VpcPublicGatewayMapOutput {
	return o
}

func (o VpcPublicGatewayMapOutput) ToVpcPublicGatewayMapOutputWithContext(ctx context.Context) VpcPublicGatewayMapOutput {
	return o
}

func (o VpcPublicGatewayMapOutput) MapIndex(k pulumi.StringInput) VpcPublicGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPublicGateway {
		return vs[0].(map[string]*VpcPublicGateway)[vs[1].(string)]
	}).(VpcPublicGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayInput)(nil)).Elem(), &VpcPublicGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayArrayInput)(nil)).Elem(), VpcPublicGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayMapInput)(nil)).Elem(), VpcPublicGatewayMap{})
	pulumi.RegisterOutputType(VpcPublicGatewayOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayArrayOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayMapOutput{})
}
