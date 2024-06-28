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

// Creates and manages Scaleway Load Balancers.
//
// For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/concepts/#load-balancers) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-list-load-balancers).
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
//			main, err := scaleway.NewLoadbalancerIp(ctx, "main", &scaleway.LoadbalancerIpArgs{
//				Zone: pulumi.String("fr-par-1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewLoadbalancer(ctx, "base", &scaleway.LoadbalancerArgs{
//				IpIds: pulumi.StringArray{
//					main.ID(),
//				},
//				Zone: main.Zone,
//				Type: pulumi.String("LB-S"),
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
// ### Private LB
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
//			_, err := scaleway.NewLoadbalancer(ctx, "base", &scaleway.LoadbalancerArgs{
//				AssignFlexibleIp: pulumi.Bool(false),
//				Type:             pulumi.String("LB-S"),
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
// ### With IPv6
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
//			v4, err := scaleway.NewLoadbalancerIp(ctx, "v4", nil)
//			if err != nil {
//				return err
//			}
//			v6, err := scaleway.NewLoadbalancerIp(ctx, "v6", &scaleway.LoadbalancerIpArgs{
//				IsIpv6: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewLoadbalancer(ctx, "main", &scaleway.LoadbalancerArgs{
//				IpIds: pulumi.StringArray{
//					v4.ID(),
//					v6.ID(),
//				},
//				Type: pulumi.String("LB-S"),
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
// ### Multiple configurations
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
//			// ## IP for Public Gateway
//			mainVpcPublicGatewayIp, err := scaleway.NewVpcPublicGatewayIp(ctx, "mainVpcPublicGatewayIp", nil)
//			if err != nil {
//				return err
//			}
//			// ## Scaleway Private Network
//			mainVpcPrivateNetwork, err := scaleway.NewVpcPrivateNetwork(ctx, "mainVpcPrivateNetwork", nil)
//			if err != nil {
//				return err
//			}
//			// ## VPC Public Gateway Network
//			mainVpcPublicGateway, err := scaleway.NewVpcPublicGateway(ctx, "mainVpcPublicGateway", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: mainVpcPublicGatewayIp.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			// ## VPC Public Gateway Network DHCP config
//			mainVpcPublicGatewayDhcp, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "mainVpcPublicGatewayDhcp", &scaleway.VpcPublicGatewayDhcpArgs{
//				Subnet: pulumi.String("10.0.0.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			// ## VPC Gateway Network
//			_, err = scaleway.NewVpcGatewayNetwork(ctx, "mainVpcGatewayNetwork", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        mainVpcPublicGateway.ID(),
//				PrivateNetworkId: mainVpcPrivateNetwork.ID(),
//				DhcpId:           mainVpcPublicGatewayDhcp.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			// ## Scaleway Instance
//			_, err = scaleway.NewInstanceServer(ctx, "mainInstanceServer", &scaleway.InstanceServerArgs{
//				Type:       pulumi.String("DEV1-S"),
//				Image:      pulumi.String("debian_bullseye"),
//				EnableIpv6: pulumi.Bool(false),
//				PrivateNetworks: scaleway.InstanceServerPrivateNetworkArray{
//					&scaleway.InstanceServerPrivateNetworkArgs{
//						PnId: mainVpcPrivateNetwork.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// ## IP for LB IP
//			mainLoadbalancerIp, err := scaleway.NewLoadbalancerIp(ctx, "mainLoadbalancerIp", nil)
//			if err != nil {
//				return err
//			}
//			// ## Scaleway Private Network
//			_, err = scaleway.NewVpcPrivateNetwork(ctx, "mainIndex/vpcPrivateNetworkVpcPrivateNetwork", nil)
//			if err != nil {
//				return err
//			}
//			// ## Scaleway Load Balancer
//			_, err = scaleway.NewLoadbalancer(ctx, "mainLoadbalancer", &scaleway.LoadbalancerArgs{
//				IpId: mainLoadbalancerIp.ID(),
//				Type: pulumi.String("LB-S"),
//				PrivateNetworks: scaleway.LoadbalancerPrivateNetworkArray{
//					&scaleway.LoadbalancerPrivateNetworkArgs{
//						PrivateNetworkId: mainVpcPrivateNetwork.ID(),
//						DhcpConfig:       pulumi.Bool(true),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				mainVpcPublicGateway,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Migration
//
// In order to migrate to other Load Balancer types, you can check upwards or downwards migration via our CLI `scw lb lb-types list`.
// This change will not recreate your Load Balancer.
//
// Please check our [documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-migrate-a-load-balancer) for further details
//
// ## IP ID
//
// Since v1.15.0, `ipId` is a required field. This means that now a separate `LoadbalancerIp` is required.
// When importing, the IP needs to be imported as well as the Load Balancer.
// When upgrading to v1.15.0, you will need to create a new `LoadbalancerIp` resource and import it.
//
// For instance, if you had the following:
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
//			_, err := scaleway.NewLoadbalancer(ctx, "main", &scaleway.LoadbalancerArgs{
//				Type: pulumi.String("LB-S"),
//				Zone: pulumi.String("fr-par-1"),
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
// You will need to update it to:
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
//			mainLoadbalancerIp, err := scaleway.NewLoadbalancerIp(ctx, "mainLoadbalancerIp", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewLoadbalancer(ctx, "mainLoadbalancer", &scaleway.LoadbalancerArgs{
//				IpId:      mainLoadbalancerIp.ID(),
//				Zone:      pulumi.String("fr-par-1"),
//				Type:      pulumi.String("LB-S"),
//				ReleaseIp: pulumi.Bool(false),
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
// Load Balancers can be imported using `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/loadbalancer:Loadbalancer main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
//
// Be aware that you will also need to import the `scaleway_lb_ip` resource.
type Loadbalancer struct {
	pulumi.CustomResourceState

	// Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
	AssignFlexibleIp pulumi.BoolPtrOutput `pulumi:"assignFlexibleIp"`
	// Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
	AssignFlexibleIpv6 pulumi.BoolPtrOutput `pulumi:"assignFlexibleIpv6"`
	// The description of the Load Balancer.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Load Balancer public IPv4 address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The ID of the associated Load Balancer IP. See below.
	//
	// > **Important:** Updates to `ipId` will recreate the Load Balancer.
	//
	// Deprecated: Please use ip_ids
	IpId pulumi.StringOutput `pulumi:"ipId"`
	// The List of IP IDs to attach to the Load Balancer.
	IpIds pulumi.StringArrayOutput `pulumi:"ipIds"`
	// The Load Balancer public IPv6 address.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// The name of the Load Balancer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Organization ID the Load Balancer is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayOutput `pulumi:"privateNetworks"`
	// `projectId`) The ID of the Project the Load Balancer is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region of the resource
	Region pulumi.StringOutput `pulumi:"region"`
	// The `releaseIp` allow the release of the IP address associated with the Load Balancer.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp pulumi.BoolPtrOutput `pulumi:"releaseIp"`
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
	SslCompatibilityLevel pulumi.StringPtrOutput `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the Load Balancer.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of the Load Balancer. Please check the migration section to upgrade the type.
	Type pulumi.StringOutput `pulumi:"type"`
	// `zone`) The zone of the Load Balancer.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewLoadbalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadbalancer(ctx *pulumi.Context,
	name string, args *LoadbalancerArgs, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Loadbalancer
	err := ctx.RegisterResource("scaleway:index/loadbalancer:Loadbalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadbalancer gets an existing Loadbalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadbalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadbalancerState, opts ...pulumi.ResourceOption) (*Loadbalancer, error) {
	var resource Loadbalancer
	err := ctx.ReadResource("scaleway:index/loadbalancer:Loadbalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Loadbalancer resources.
type loadbalancerState struct {
	// Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
	AssignFlexibleIp *bool `pulumi:"assignFlexibleIp"`
	// Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
	AssignFlexibleIpv6 *bool `pulumi:"assignFlexibleIpv6"`
	// The description of the Load Balancer.
	Description *string `pulumi:"description"`
	// The Load Balancer public IPv4 address.
	IpAddress *string `pulumi:"ipAddress"`
	// The ID of the associated Load Balancer IP. See below.
	//
	// > **Important:** Updates to `ipId` will recreate the Load Balancer.
	//
	// Deprecated: Please use ip_ids
	IpId *string `pulumi:"ipId"`
	// The List of IP IDs to attach to the Load Balancer.
	IpIds []string `pulumi:"ipIds"`
	// The Load Balancer public IPv6 address.
	Ipv6Address *string `pulumi:"ipv6Address"`
	// The name of the Load Balancer.
	Name *string `pulumi:"name"`
	// The ID of the Organization ID the Load Balancer is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// List of private network to connect with your load balancer
	PrivateNetworks []LoadbalancerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the Project the Load Balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The region of the resource
	Region *string `pulumi:"region"`
	// The `releaseIp` allow the release of the IP address associated with the Load Balancer.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp *bool `pulumi:"releaseIp"`
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
	SslCompatibilityLevel *string `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the Load Balancer.
	Tags []string `pulumi:"tags"`
	// The type of the Load Balancer. Please check the migration section to upgrade the type.
	Type *string `pulumi:"type"`
	// `zone`) The zone of the Load Balancer.
	Zone *string `pulumi:"zone"`
}

type LoadbalancerState struct {
	// Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
	AssignFlexibleIp pulumi.BoolPtrInput
	// Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
	AssignFlexibleIpv6 pulumi.BoolPtrInput
	// The description of the Load Balancer.
	Description pulumi.StringPtrInput
	// The Load Balancer public IPv4 address.
	IpAddress pulumi.StringPtrInput
	// The ID of the associated Load Balancer IP. See below.
	//
	// > **Important:** Updates to `ipId` will recreate the Load Balancer.
	//
	// Deprecated: Please use ip_ids
	IpId pulumi.StringPtrInput
	// The List of IP IDs to attach to the Load Balancer.
	IpIds pulumi.StringArrayInput
	// The Load Balancer public IPv6 address.
	Ipv6Address pulumi.StringPtrInput
	// The name of the Load Balancer.
	Name pulumi.StringPtrInput
	// The ID of the Organization ID the Load Balancer is associated with.
	OrganizationId pulumi.StringPtrInput
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayInput
	// `projectId`) The ID of the Project the Load Balancer is associated with.
	ProjectId pulumi.StringPtrInput
	// The region of the resource
	Region pulumi.StringPtrInput
	// The `releaseIp` allow the release of the IP address associated with the Load Balancer.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp pulumi.BoolPtrInput
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
	SslCompatibilityLevel pulumi.StringPtrInput
	// The tags associated with the Load Balancer.
	Tags pulumi.StringArrayInput
	// The type of the Load Balancer. Please check the migration section to upgrade the type.
	Type pulumi.StringPtrInput
	// `zone`) The zone of the Load Balancer.
	Zone pulumi.StringPtrInput
}

func (LoadbalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerState)(nil)).Elem()
}

type loadbalancerArgs struct {
	// Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
	AssignFlexibleIp *bool `pulumi:"assignFlexibleIp"`
	// Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
	AssignFlexibleIpv6 *bool `pulumi:"assignFlexibleIpv6"`
	// The description of the Load Balancer.
	Description *string `pulumi:"description"`
	// The ID of the associated Load Balancer IP. See below.
	//
	// > **Important:** Updates to `ipId` will recreate the Load Balancer.
	//
	// Deprecated: Please use ip_ids
	IpId *string `pulumi:"ipId"`
	// The List of IP IDs to attach to the Load Balancer.
	IpIds []string `pulumi:"ipIds"`
	// The name of the Load Balancer.
	Name *string `pulumi:"name"`
	// List of private network to connect with your load balancer
	PrivateNetworks []LoadbalancerPrivateNetwork `pulumi:"privateNetworks"`
	// `projectId`) The ID of the Project the Load Balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The `releaseIp` allow the release of the IP address associated with the Load Balancer.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp *bool `pulumi:"releaseIp"`
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
	SslCompatibilityLevel *string `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the Load Balancer.
	Tags []string `pulumi:"tags"`
	// The type of the Load Balancer. Please check the migration section to upgrade the type.
	Type string `pulumi:"type"`
	// `zone`) The zone of the Load Balancer.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Loadbalancer resource.
type LoadbalancerArgs struct {
	// Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
	AssignFlexibleIp pulumi.BoolPtrInput
	// Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
	AssignFlexibleIpv6 pulumi.BoolPtrInput
	// The description of the Load Balancer.
	Description pulumi.StringPtrInput
	// The ID of the associated Load Balancer IP. See below.
	//
	// > **Important:** Updates to `ipId` will recreate the Load Balancer.
	//
	// Deprecated: Please use ip_ids
	IpId pulumi.StringPtrInput
	// The List of IP IDs to attach to the Load Balancer.
	IpIds pulumi.StringArrayInput
	// The name of the Load Balancer.
	Name pulumi.StringPtrInput
	// List of private network to connect with your load balancer
	PrivateNetworks LoadbalancerPrivateNetworkArrayInput
	// `projectId`) The ID of the Project the Load Balancer is associated with.
	ProjectId pulumi.StringPtrInput
	// The `releaseIp` allow the release of the IP address associated with the Load Balancer.
	//
	// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
	ReleaseIp pulumi.BoolPtrInput
	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
	SslCompatibilityLevel pulumi.StringPtrInput
	// The tags associated with the Load Balancer.
	Tags pulumi.StringArrayInput
	// The type of the Load Balancer. Please check the migration section to upgrade the type.
	Type pulumi.StringInput
	// `zone`) The zone of the Load Balancer.
	Zone pulumi.StringPtrInput
}

func (LoadbalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerArgs)(nil)).Elem()
}

type LoadbalancerInput interface {
	pulumi.Input

	ToLoadbalancerOutput() LoadbalancerOutput
	ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput
}

func (*Loadbalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (i *Loadbalancer) ToLoadbalancerOutput() LoadbalancerOutput {
	return i.ToLoadbalancerOutputWithContext(context.Background())
}

func (i *Loadbalancer) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerOutput)
}

// LoadbalancerArrayInput is an input type that accepts LoadbalancerArray and LoadbalancerArrayOutput values.
// You can construct a concrete instance of `LoadbalancerArrayInput` via:
//
//	LoadbalancerArray{ LoadbalancerArgs{...} }
type LoadbalancerArrayInput interface {
	pulumi.Input

	ToLoadbalancerArrayOutput() LoadbalancerArrayOutput
	ToLoadbalancerArrayOutputWithContext(context.Context) LoadbalancerArrayOutput
}

type LoadbalancerArray []LoadbalancerInput

func (LoadbalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Loadbalancer)(nil)).Elem()
}

func (i LoadbalancerArray) ToLoadbalancerArrayOutput() LoadbalancerArrayOutput {
	return i.ToLoadbalancerArrayOutputWithContext(context.Background())
}

func (i LoadbalancerArray) ToLoadbalancerArrayOutputWithContext(ctx context.Context) LoadbalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerArrayOutput)
}

// LoadbalancerMapInput is an input type that accepts LoadbalancerMap and LoadbalancerMapOutput values.
// You can construct a concrete instance of `LoadbalancerMapInput` via:
//
//	LoadbalancerMap{ "key": LoadbalancerArgs{...} }
type LoadbalancerMapInput interface {
	pulumi.Input

	ToLoadbalancerMapOutput() LoadbalancerMapOutput
	ToLoadbalancerMapOutputWithContext(context.Context) LoadbalancerMapOutput
}

type LoadbalancerMap map[string]LoadbalancerInput

func (LoadbalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Loadbalancer)(nil)).Elem()
}

func (i LoadbalancerMap) ToLoadbalancerMapOutput() LoadbalancerMapOutput {
	return i.ToLoadbalancerMapOutputWithContext(context.Background())
}

func (i LoadbalancerMap) ToLoadbalancerMapOutputWithContext(ctx context.Context) LoadbalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerMapOutput)
}

type LoadbalancerOutput struct{ *pulumi.OutputState }

func (LoadbalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerOutput) ToLoadbalancerOutput() LoadbalancerOutput {
	return o
}

func (o LoadbalancerOutput) ToLoadbalancerOutputWithContext(ctx context.Context) LoadbalancerOutput {
	return o
}

// Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
func (o LoadbalancerOutput) AssignFlexibleIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.BoolPtrOutput { return v.AssignFlexibleIp }).(pulumi.BoolPtrOutput)
}

// Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
func (o LoadbalancerOutput) AssignFlexibleIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.BoolPtrOutput { return v.AssignFlexibleIpv6 }).(pulumi.BoolPtrOutput)
}

// The description of the Load Balancer.
func (o LoadbalancerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Load Balancer public IPv4 address.
func (o LoadbalancerOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The ID of the associated Load Balancer IP. See below.
//
// > **Important:** Updates to `ipId` will recreate the Load Balancer.
//
// Deprecated: Please use ip_ids
func (o LoadbalancerOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.IpId }).(pulumi.StringOutput)
}

// The List of IP IDs to attach to the Load Balancer.
func (o LoadbalancerOutput) IpIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringArrayOutput { return v.IpIds }).(pulumi.StringArrayOutput)
}

// The Load Balancer public IPv6 address.
func (o LoadbalancerOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

// The name of the Load Balancer.
func (o LoadbalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Organization ID the Load Balancer is associated with.
func (o LoadbalancerOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// List of private network to connect with your load balancer
func (o LoadbalancerOutput) PrivateNetworks() LoadbalancerPrivateNetworkArrayOutput {
	return o.ApplyT(func(v *Loadbalancer) LoadbalancerPrivateNetworkArrayOutput { return v.PrivateNetworks }).(LoadbalancerPrivateNetworkArrayOutput)
}

// `projectId`) The ID of the Project the Load Balancer is associated with.
func (o LoadbalancerOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region of the resource
func (o LoadbalancerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The `releaseIp` allow the release of the IP address associated with the Load Balancer.
//
// Deprecated: The resource ip will be destroyed by it's own resource. Please set this to `false`
func (o LoadbalancerOutput) ReleaseIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.BoolPtrOutput { return v.ReleaseIp }).(pulumi.BoolPtrOutput)
}

// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
func (o LoadbalancerOutput) SslCompatibilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringPtrOutput { return v.SslCompatibilityLevel }).(pulumi.StringPtrOutput)
}

// The tags associated with the Load Balancer.
func (o LoadbalancerOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the Load Balancer. Please check the migration section to upgrade the type.
func (o LoadbalancerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// `zone`) The zone of the Load Balancer.
func (o LoadbalancerOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Loadbalancer) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type LoadbalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadbalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerArrayOutput) ToLoadbalancerArrayOutput() LoadbalancerArrayOutput {
	return o
}

func (o LoadbalancerArrayOutput) ToLoadbalancerArrayOutputWithContext(ctx context.Context) LoadbalancerArrayOutput {
	return o
}

func (o LoadbalancerArrayOutput) Index(i pulumi.IntInput) LoadbalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Loadbalancer {
		return vs[0].([]*Loadbalancer)[vs[1].(int)]
	}).(LoadbalancerOutput)
}

type LoadbalancerMapOutput struct{ *pulumi.OutputState }

func (LoadbalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Loadbalancer)(nil)).Elem()
}

func (o LoadbalancerMapOutput) ToLoadbalancerMapOutput() LoadbalancerMapOutput {
	return o
}

func (o LoadbalancerMapOutput) ToLoadbalancerMapOutputWithContext(ctx context.Context) LoadbalancerMapOutput {
	return o
}

func (o LoadbalancerMapOutput) MapIndex(k pulumi.StringInput) LoadbalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Loadbalancer {
		return vs[0].(map[string]*Loadbalancer)[vs[1].(string)]
	}).(LoadbalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerInput)(nil)).Elem(), &Loadbalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerArrayInput)(nil)).Elem(), LoadbalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerMapInput)(nil)).Elem(), LoadbalancerMap{})
	pulumi.RegisterOutputType(LoadbalancerOutput{})
	pulumi.RegisterOutputType(LoadbalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadbalancerMapOutput{})
}
