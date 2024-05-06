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

// Creates and manages Scaleway VPC Public Gateway DHCP.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-c05544).
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
//			_, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "main", &scaleway.VpcPublicGatewayDhcpArgs{
//				Subnet: pulumi.String("192.168.1.0/24"),
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
// Public gateway DHCP config can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type VpcPublicGatewayDhcp struct {
	pulumi.CustomResourceState

	// The IP address of the public gateway DHCP config.
	Address pulumi.StringOutput `pulumi:"address"`
	// The date and time of the creation of the public gateway DHCP config.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	DnsLocalName pulumi.StringOutput `pulumi:"dnsLocalName"`
	// Additional DNS search paths
	DnsSearches pulumi.StringArrayOutput `pulumi:"dnsSearches"`
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	DnsServersOverrides pulumi.StringArrayOutput `pulumi:"dnsServersOverrides"`
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
	EnableDynamic pulumi.BoolOutput `pulumi:"enableDynamic"`
	// The organization ID the public gateway DHCP config is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	PoolHigh pulumi.StringOutput `pulumi:"poolHigh"`
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	PoolLow pulumi.StringOutput `pulumi:"poolLow"`
	// `projectId`) The ID of the project the public gateway DHCP config is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
	PushDefaultRoute pulumi.BoolOutput `pulumi:"pushDefaultRoute"`
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
	PushDnsServer pulumi.BoolOutput `pulumi:"pushDnsServer"`
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
	RebindTimer pulumi.IntOutput `pulumi:"rebindTimer"`
	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
	RenewTimer pulumi.IntOutput `pulumi:"renewTimer"`
	// The subnet to associate with the public gateway DHCP config.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// The date and time of the last update of the public gateway DHCP config.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	ValidLifetime pulumi.IntOutput `pulumi:"validLifetime"`
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGatewayDhcp registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGatewayDhcp(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayDhcpArgs, opts ...pulumi.ResourceOption) (*VpcPublicGatewayDhcp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcPublicGatewayDhcp
	err := ctx.RegisterResource("scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGatewayDhcp gets an existing VpcPublicGatewayDhcp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGatewayDhcp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayDhcpState, opts ...pulumi.ResourceOption) (*VpcPublicGatewayDhcp, error) {
	var resource VpcPublicGatewayDhcp
	err := ctx.ReadResource("scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGatewayDhcp resources.
type vpcPublicGatewayDhcpState struct {
	// The IP address of the public gateway DHCP config.
	Address *string `pulumi:"address"`
	// The date and time of the creation of the public gateway DHCP config.
	CreatedAt *string `pulumi:"createdAt"`
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	DnsLocalName *string `pulumi:"dnsLocalName"`
	// Additional DNS search paths
	DnsSearches []string `pulumi:"dnsSearches"`
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	DnsServersOverrides []string `pulumi:"dnsServersOverrides"`
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
	EnableDynamic *bool `pulumi:"enableDynamic"`
	// The organization ID the public gateway DHCP config is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	PoolHigh *string `pulumi:"poolHigh"`
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	PoolLow *string `pulumi:"poolLow"`
	// `projectId`) The ID of the project the public gateway DHCP config is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
	PushDefaultRoute *bool `pulumi:"pushDefaultRoute"`
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
	PushDnsServer *bool `pulumi:"pushDnsServer"`
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
	RebindTimer *int `pulumi:"rebindTimer"`
	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
	RenewTimer *int `pulumi:"renewTimer"`
	// The subnet to associate with the public gateway DHCP config.
	Subnet *string `pulumi:"subnet"`
	// The date and time of the last update of the public gateway DHCP config.
	UpdatedAt *string `pulumi:"updatedAt"`
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	ValidLifetime *int `pulumi:"validLifetime"`
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayDhcpState struct {
	// The IP address of the public gateway DHCP config.
	Address pulumi.StringPtrInput
	// The date and time of the creation of the public gateway DHCP config.
	CreatedAt pulumi.StringPtrInput
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	DnsLocalName pulumi.StringPtrInput
	// Additional DNS search paths
	DnsSearches pulumi.StringArrayInput
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	DnsServersOverrides pulumi.StringArrayInput
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
	EnableDynamic pulumi.BoolPtrInput
	// The organization ID the public gateway DHCP config is associated with.
	OrganizationId pulumi.StringPtrInput
	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	PoolHigh pulumi.StringPtrInput
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	PoolLow pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway DHCP config is associated with.
	ProjectId pulumi.StringPtrInput
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
	PushDefaultRoute pulumi.BoolPtrInput
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
	PushDnsServer pulumi.BoolPtrInput
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
	RebindTimer pulumi.IntPtrInput
	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
	RenewTimer pulumi.IntPtrInput
	// The subnet to associate with the public gateway DHCP config.
	Subnet pulumi.StringPtrInput
	// The date and time of the last update of the public gateway DHCP config.
	UpdatedAt pulumi.StringPtrInput
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	ValidLifetime pulumi.IntPtrInput
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayDhcpState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayDhcpState)(nil)).Elem()
}

type vpcPublicGatewayDhcpArgs struct {
	// The IP address of the public gateway DHCP config.
	Address *string `pulumi:"address"`
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	DnsLocalName *string `pulumi:"dnsLocalName"`
	// Additional DNS search paths
	DnsSearches []string `pulumi:"dnsSearches"`
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	DnsServersOverrides []string `pulumi:"dnsServersOverrides"`
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
	EnableDynamic *bool `pulumi:"enableDynamic"`
	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	PoolHigh *string `pulumi:"poolHigh"`
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	PoolLow *string `pulumi:"poolLow"`
	// `projectId`) The ID of the project the public gateway DHCP config is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
	PushDefaultRoute *bool `pulumi:"pushDefaultRoute"`
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
	PushDnsServer *bool `pulumi:"pushDnsServer"`
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
	RebindTimer *int `pulumi:"rebindTimer"`
	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
	RenewTimer *int `pulumi:"renewTimer"`
	// The subnet to associate with the public gateway DHCP config.
	Subnet string `pulumi:"subnet"`
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	ValidLifetime *int `pulumi:"validLifetime"`
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGatewayDhcp resource.
type VpcPublicGatewayDhcpArgs struct {
	// The IP address of the public gateway DHCP config.
	Address pulumi.StringPtrInput
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	DnsLocalName pulumi.StringPtrInput
	// Additional DNS search paths
	DnsSearches pulumi.StringArrayInput
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	DnsServersOverrides pulumi.StringArrayInput
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
	EnableDynamic pulumi.BoolPtrInput
	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	PoolHigh pulumi.StringPtrInput
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	PoolLow pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway DHCP config is associated with.
	ProjectId pulumi.StringPtrInput
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
	PushDefaultRoute pulumi.BoolPtrInput
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
	PushDnsServer pulumi.BoolPtrInput
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
	RebindTimer pulumi.IntPtrInput
	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
	RenewTimer pulumi.IntPtrInput
	// The subnet to associate with the public gateway DHCP config.
	Subnet pulumi.StringInput
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	ValidLifetime pulumi.IntPtrInput
	// `zone`) The zone in which the public gateway DHCP config should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayDhcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayDhcpArgs)(nil)).Elem()
}

type VpcPublicGatewayDhcpInput interface {
	pulumi.Input

	ToVpcPublicGatewayDhcpOutput() VpcPublicGatewayDhcpOutput
	ToVpcPublicGatewayDhcpOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpOutput
}

func (*VpcPublicGatewayDhcp) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayDhcp)(nil)).Elem()
}

func (i *VpcPublicGatewayDhcp) ToVpcPublicGatewayDhcpOutput() VpcPublicGatewayDhcpOutput {
	return i.ToVpcPublicGatewayDhcpOutputWithContext(context.Background())
}

func (i *VpcPublicGatewayDhcp) ToVpcPublicGatewayDhcpOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayDhcpOutput)
}

// VpcPublicGatewayDhcpArrayInput is an input type that accepts VpcPublicGatewayDhcpArray and VpcPublicGatewayDhcpArrayOutput values.
// You can construct a concrete instance of `VpcPublicGatewayDhcpArrayInput` via:
//
//	VpcPublicGatewayDhcpArray{ VpcPublicGatewayDhcpArgs{...} }
type VpcPublicGatewayDhcpArrayInput interface {
	pulumi.Input

	ToVpcPublicGatewayDhcpArrayOutput() VpcPublicGatewayDhcpArrayOutput
	ToVpcPublicGatewayDhcpArrayOutputWithContext(context.Context) VpcPublicGatewayDhcpArrayOutput
}

type VpcPublicGatewayDhcpArray []VpcPublicGatewayDhcpInput

func (VpcPublicGatewayDhcpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayDhcp)(nil)).Elem()
}

func (i VpcPublicGatewayDhcpArray) ToVpcPublicGatewayDhcpArrayOutput() VpcPublicGatewayDhcpArrayOutput {
	return i.ToVpcPublicGatewayDhcpArrayOutputWithContext(context.Background())
}

func (i VpcPublicGatewayDhcpArray) ToVpcPublicGatewayDhcpArrayOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayDhcpArrayOutput)
}

// VpcPublicGatewayDhcpMapInput is an input type that accepts VpcPublicGatewayDhcpMap and VpcPublicGatewayDhcpMapOutput values.
// You can construct a concrete instance of `VpcPublicGatewayDhcpMapInput` via:
//
//	VpcPublicGatewayDhcpMap{ "key": VpcPublicGatewayDhcpArgs{...} }
type VpcPublicGatewayDhcpMapInput interface {
	pulumi.Input

	ToVpcPublicGatewayDhcpMapOutput() VpcPublicGatewayDhcpMapOutput
	ToVpcPublicGatewayDhcpMapOutputWithContext(context.Context) VpcPublicGatewayDhcpMapOutput
}

type VpcPublicGatewayDhcpMap map[string]VpcPublicGatewayDhcpInput

func (VpcPublicGatewayDhcpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayDhcp)(nil)).Elem()
}

func (i VpcPublicGatewayDhcpMap) ToVpcPublicGatewayDhcpMapOutput() VpcPublicGatewayDhcpMapOutput {
	return i.ToVpcPublicGatewayDhcpMapOutputWithContext(context.Background())
}

func (i VpcPublicGatewayDhcpMap) ToVpcPublicGatewayDhcpMapOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayDhcpMapOutput)
}

type VpcPublicGatewayDhcpOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayDhcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayDhcp)(nil)).Elem()
}

func (o VpcPublicGatewayDhcpOutput) ToVpcPublicGatewayDhcpOutput() VpcPublicGatewayDhcpOutput {
	return o
}

func (o VpcPublicGatewayDhcpOutput) ToVpcPublicGatewayDhcpOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpOutput {
	return o
}

// The IP address of the public gateway DHCP config.
func (o VpcPublicGatewayDhcpOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The date and time of the creation of the public gateway DHCP config.
func (o VpcPublicGatewayDhcpOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
func (o VpcPublicGatewayDhcpOutput) DnsLocalName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.DnsLocalName }).(pulumi.StringOutput)
}

// Additional DNS search paths
func (o VpcPublicGatewayDhcpOutput) DnsSearches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringArrayOutput { return v.DnsSearches }).(pulumi.StringArrayOutput)
}

// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
func (o VpcPublicGatewayDhcpOutput) DnsServersOverrides() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringArrayOutput { return v.DnsServersOverrides }).(pulumi.StringArrayOutput)
}

// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
func (o VpcPublicGatewayDhcpOutput) EnableDynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.BoolOutput { return v.EnableDynamic }).(pulumi.BoolOutput)
}

// The organization ID the public gateway DHCP config is associated with.
func (o VpcPublicGatewayDhcpOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
func (o VpcPublicGatewayDhcpOutput) PoolHigh() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.PoolHigh }).(pulumi.StringOutput)
}

// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
func (o VpcPublicGatewayDhcpOutput) PoolLow() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.PoolLow }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the public gateway DHCP config is associated with.
func (o VpcPublicGatewayDhcpOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
func (o VpcPublicGatewayDhcpOutput) PushDefaultRoute() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.BoolOutput { return v.PushDefaultRoute }).(pulumi.BoolOutput)
}

// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
func (o VpcPublicGatewayDhcpOutput) PushDnsServer() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.BoolOutput { return v.PushDnsServer }).(pulumi.BoolOutput)
}

// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
func (o VpcPublicGatewayDhcpOutput) RebindTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.IntOutput { return v.RebindTimer }).(pulumi.IntOutput)
}

// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
func (o VpcPublicGatewayDhcpOutput) RenewTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.IntOutput { return v.RenewTimer }).(pulumi.IntOutput)
}

// The subnet to associate with the public gateway DHCP config.
func (o VpcPublicGatewayDhcpOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// The date and time of the last update of the public gateway DHCP config.
func (o VpcPublicGatewayDhcpOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
func (o VpcPublicGatewayDhcpOutput) ValidLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.IntOutput { return v.ValidLifetime }).(pulumi.IntOutput)
}

// `zone`) The zone in which the public gateway DHCP config should be created.
func (o VpcPublicGatewayDhcpOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayDhcp) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcPublicGatewayDhcpArrayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayDhcpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayDhcp)(nil)).Elem()
}

func (o VpcPublicGatewayDhcpArrayOutput) ToVpcPublicGatewayDhcpArrayOutput() VpcPublicGatewayDhcpArrayOutput {
	return o
}

func (o VpcPublicGatewayDhcpArrayOutput) ToVpcPublicGatewayDhcpArrayOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpArrayOutput {
	return o
}

func (o VpcPublicGatewayDhcpArrayOutput) Index(i pulumi.IntInput) VpcPublicGatewayDhcpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPublicGatewayDhcp {
		return vs[0].([]*VpcPublicGatewayDhcp)[vs[1].(int)]
	}).(VpcPublicGatewayDhcpOutput)
}

type VpcPublicGatewayDhcpMapOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayDhcpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayDhcp)(nil)).Elem()
}

func (o VpcPublicGatewayDhcpMapOutput) ToVpcPublicGatewayDhcpMapOutput() VpcPublicGatewayDhcpMapOutput {
	return o
}

func (o VpcPublicGatewayDhcpMapOutput) ToVpcPublicGatewayDhcpMapOutputWithContext(ctx context.Context) VpcPublicGatewayDhcpMapOutput {
	return o
}

func (o VpcPublicGatewayDhcpMapOutput) MapIndex(k pulumi.StringInput) VpcPublicGatewayDhcpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPublicGatewayDhcp {
		return vs[0].(map[string]*VpcPublicGatewayDhcp)[vs[1].(string)]
	}).(VpcPublicGatewayDhcpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayDhcpInput)(nil)).Elem(), &VpcPublicGatewayDhcp{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayDhcpArrayInput)(nil)).Elem(), VpcPublicGatewayDhcpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayDhcpMapInput)(nil)).Elem(), VpcPublicGatewayDhcpMap{})
	pulumi.RegisterOutputType(VpcPublicGatewayDhcpOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayDhcpArrayOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayDhcpMapOutput{})
}
