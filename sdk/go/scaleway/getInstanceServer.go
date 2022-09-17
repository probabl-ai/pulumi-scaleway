// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an instance server.
//
// ## Example Usage
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
//			_, err := scaleway.LookupInstanceServer(ctx, &GetInstanceServerArgs{
//				ServerId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInstanceServer(ctx *pulumi.Context, args *LookupInstanceServerArgs, opts ...pulumi.InvokeOption) (*LookupInstanceServerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupInstanceServerResult
	err := ctx.Invoke("scaleway:index/getInstanceServer:getInstanceServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceServer.
type LookupInstanceServerArgs struct {
	// The server name. Only one of `name` and `serverId` should be specified.
	Name *string `pulumi:"name"`
	// The server id. Only one of `name` and `serverId` should be specified.
	ServerId *string `pulumi:"serverId"`
	// `zone`) The zone in which the server exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceServer.
type LookupInstanceServerResult struct {
	// The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
	// attached to the server.
	AdditionalVolumeIds []string `pulumi:"additionalVolumeIds"`
	BootType            string   `pulumi:"bootType"`
	BootscriptId        string   `pulumi:"bootscriptId"`
	// The cloud init script associated with this server.
	CloudInit string `pulumi:"cloudInit"`
	// True is dynamic IP in enable on the server.
	EnableDynamicIp bool `pulumi:"enableDynamicIp"`
	// Determines if IPv6 is enabled for the server.
	EnableIpv6 bool `pulumi:"enableIpv6"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The UUID and the label of the base image used by the server.
	Image string `pulumi:"image"`
	IpId  string `pulumi:"ipId"`
	// The default ipv6 address routed to the server. ( Only set when enableIpv6 is set to true )
	Ipv6Address string `pulumi:"ipv6Address"`
	// The ipv6 gateway address. ( Only set when enableIpv6 is set to true )
	Ipv6Gateway string `pulumi:"ipv6Gateway"`
	// The prefix length of the ipv6 subnet routed to the server. ( Only set when enableIpv6 is set to true )
	Ipv6PrefixLength int     `pulumi:"ipv6PrefixLength"`
	Name             *string `pulumi:"name"`
	// The ID of the organization the server is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
	PlacementGroupId string `pulumi:"placementGroupId"`
	// True when the placement group policy is respected.
	PlacementGroupPolicyRespected bool `pulumi:"placementGroupPolicyRespected"`
	// The Scaleway internal IP address of the server.
	PrivateIp       string                            `pulumi:"privateIp"`
	PrivateNetworks []GetInstanceServerPrivateNetwork `pulumi:"privateNetworks"`
	// The ID of the project the server is associated with.
	ProjectId string `pulumi:"projectId"`
	// The public IPv4 address of the server.
	PublicIp string `pulumi:"publicIp"`
	// Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
	RootVolumes []GetInstanceServerRootVolume `pulumi:"rootVolumes"`
	// The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
	SecurityGroupId string  `pulumi:"securityGroupId"`
	ServerId        *string `pulumi:"serverId"`
	// The state of the server. Possible values are: `started`, `stopped` or `standby`.
	State string `pulumi:"state"`
	// The tags associated with the server.
	Tags []string `pulumi:"tags"`
	// The commercial type of the server.
	// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
	Type string `pulumi:"type"`
	// The user data associated with the server.
	UserData map[string]string `pulumi:"userData"`
	Zone     *string           `pulumi:"zone"`
}

func LookupInstanceServerOutput(ctx *pulumi.Context, args LookupInstanceServerOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceServerResult, error) {
			args := v.(LookupInstanceServerArgs)
			r, err := LookupInstanceServer(ctx, &args, opts...)
			var s LookupInstanceServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceServerResultOutput)
}

// A collection of arguments for invoking getInstanceServer.
type LookupInstanceServerOutputArgs struct {
	// The server name. Only one of `name` and `serverId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The server id. Only one of `name` and `serverId` should be specified.
	ServerId pulumi.StringPtrInput `pulumi:"serverId"`
	// `zone`) The zone in which the server exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupInstanceServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceServerArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceServer.
type LookupInstanceServerResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceServerResult)(nil)).Elem()
}

func (o LookupInstanceServerResultOutput) ToLookupInstanceServerResultOutput() LookupInstanceServerResultOutput {
	return o
}

func (o LookupInstanceServerResultOutput) ToLookupInstanceServerResultOutputWithContext(ctx context.Context) LookupInstanceServerResultOutput {
	return o
}

// The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
// attached to the server.
func (o LookupInstanceServerResultOutput) AdditionalVolumeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) []string { return v.AdditionalVolumeIds }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceServerResultOutput) BootType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.BootType }).(pulumi.StringOutput)
}

func (o LookupInstanceServerResultOutput) BootscriptId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.BootscriptId }).(pulumi.StringOutput)
}

// The cloud init script associated with this server.
func (o LookupInstanceServerResultOutput) CloudInit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.CloudInit }).(pulumi.StringOutput)
}

// True is dynamic IP in enable on the server.
func (o LookupInstanceServerResultOutput) EnableDynamicIp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) bool { return v.EnableDynamicIp }).(pulumi.BoolOutput)
}

// Determines if IPv6 is enabled for the server.
func (o LookupInstanceServerResultOutput) EnableIpv6() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) bool { return v.EnableIpv6 }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The UUID and the label of the base image used by the server.
func (o LookupInstanceServerResultOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.Image }).(pulumi.StringOutput)
}

func (o LookupInstanceServerResultOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.IpId }).(pulumi.StringOutput)
}

// The default ipv6 address routed to the server. ( Only set when enableIpv6 is set to true )
func (o LookupInstanceServerResultOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.Ipv6Address }).(pulumi.StringOutput)
}

// The ipv6 gateway address. ( Only set when enableIpv6 is set to true )
func (o LookupInstanceServerResultOutput) Ipv6Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.Ipv6Gateway }).(pulumi.StringOutput)
}

// The prefix length of the ipv6 subnet routed to the server. ( Only set when enableIpv6 is set to true )
func (o LookupInstanceServerResultOutput) Ipv6PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) int { return v.Ipv6PrefixLength }).(pulumi.IntOutput)
}

func (o LookupInstanceServerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the organization the server is associated with.
func (o LookupInstanceServerResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
func (o LookupInstanceServerResultOutput) PlacementGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.PlacementGroupId }).(pulumi.StringOutput)
}

// True when the placement group policy is respected.
func (o LookupInstanceServerResultOutput) PlacementGroupPolicyRespected() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) bool { return v.PlacementGroupPolicyRespected }).(pulumi.BoolOutput)
}

// The Scaleway internal IP address of the server.
func (o LookupInstanceServerResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o LookupInstanceServerResultOutput) PrivateNetworks() GetInstanceServerPrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) []GetInstanceServerPrivateNetwork { return v.PrivateNetworks }).(GetInstanceServerPrivateNetworkArrayOutput)
}

// The ID of the project the server is associated with.
func (o LookupInstanceServerResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The public IPv4 address of the server.
func (o LookupInstanceServerResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

// Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
func (o LookupInstanceServerResultOutput) RootVolumes() GetInstanceServerRootVolumeArrayOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) []GetInstanceServerRootVolume { return v.RootVolumes }).(GetInstanceServerRootVolumeArrayOutput)
}

// The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
func (o LookupInstanceServerResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

func (o LookupInstanceServerResultOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) *string { return v.ServerId }).(pulumi.StringPtrOutput)
}

// The state of the server. Possible values are: `started`, `stopped` or `standby`.
func (o LookupInstanceServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.State }).(pulumi.StringOutput)
}

// The tags associated with the server.
func (o LookupInstanceServerResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The commercial type of the server.
// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
func (o LookupInstanceServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) string { return v.Type }).(pulumi.StringOutput)
}

// The user data associated with the server.
func (o LookupInstanceServerResultOutput) UserData() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) map[string]string { return v.UserData }).(pulumi.StringMapOutput)
}

func (o LookupInstanceServerResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceServerResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceServerResultOutput{})
}
