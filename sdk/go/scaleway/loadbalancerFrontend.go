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

// Creates and manages Scaleway Load-Balancer Frontends. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
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
//			_, err := scaleway.NewLoadbalancerFrontend(ctx, "frontend01", &scaleway.LoadbalancerFrontendArgs{
//				LbId:        pulumi.Any(scaleway_lb.Lb01.Id),
//				BackendId:   pulumi.Any(scaleway_lb_backend.Backend01.Id),
//				InboundPort: pulumi.Int(80),
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
// ## With ACLs
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
//			_, err := scaleway.NewLoadbalancerFrontend(ctx, "frontend01", &scaleway.LoadbalancerFrontendArgs{
//				LbId:        pulumi.Any(scaleway_lb.Lb01.Id),
//				BackendId:   pulumi.Any(scaleway_lb_backend.Backend01.Id),
//				InboundPort: pulumi.Int(80),
//				Acls: scaleway.LoadbalancerFrontendAclArray{
//					&scaleway.LoadbalancerFrontendAclArgs{
//						Name: pulumi.String("blacklist wellknwon IPs"),
//						Action: &scaleway.LoadbalancerFrontendAclActionArgs{
//							Type: pulumi.String("allow"),
//						},
//						Match: &scaleway.LoadbalancerFrontendAclMatchArgs{
//							IpSubnets: pulumi.StringArray{
//								pulumi.String("192.168.0.1"),
//								pulumi.String("192.168.0.2"),
//								pulumi.String("192.168.10.0/24"),
//							},
//						},
//					},
//					&scaleway.LoadbalancerFrontendAclArgs{
//						Action: &scaleway.LoadbalancerFrontendAclActionArgs{
//							Type: pulumi.String("deny"),
//						},
//						Match: &scaleway.LoadbalancerFrontendAclMatchArgs{
//							IpSubnets: pulumi.StringArray{
//								pulumi.String("51.51.51.51"),
//							},
//							HttpFilter: pulumi.String("regex"),
//							HttpFilterValues: pulumi.StringArray{
//								pulumi.String("^foo*bar$"),
//							},
//						},
//					},
//					&scaleway.LoadbalancerFrontendAclArgs{
//						Action: &scaleway.LoadbalancerFrontendAclActionArgs{
//							Type: pulumi.String("allow"),
//						},
//						Match: &scaleway.LoadbalancerFrontendAclMatchArgs{
//							HttpFilter: pulumi.String("path_begin"),
//							HttpFilterValues: pulumi.StringArray{
//								pulumi.String("foo"),
//								pulumi.String("bar"),
//							},
//						},
//					},
//					&scaleway.LoadbalancerFrontendAclArgs{
//						Action: &scaleway.LoadbalancerFrontendAclActionArgs{
//							Type: pulumi.String("allow"),
//						},
//						Match: &scaleway.LoadbalancerFrontendAclMatchArgs{
//							HttpFilter: pulumi.String("path_begin"),
//							HttpFilterValues: pulumi.StringArray{
//								pulumi.String("hi"),
//							},
//							Invert: pulumi.Bool(true),
//						},
//					},
//					&scaleway.LoadbalancerFrontendAclArgs{
//						Action: &scaleway.LoadbalancerFrontendAclActionArgs{
//							Type: pulumi.String("allow"),
//						},
//						Match: &scaleway.LoadbalancerFrontendAclMatchArgs{
//							HttpFilter:       pulumi.String("http_header_match"),
//							HttpFilterValues: pulumi.StringArray("foo"),
//							HttpFilterOption: pulumi.String("bar"),
//						},
//					},
//					&scaleway.LoadbalancerFrontendAclArgs{
//						Action: &scaleway.LoadbalancerFrontendAclActionArgs{
//							Type: pulumi.String("redirect"),
//							Redirects: scaleway.LoadbalancerFrontendAclActionRedirectArray{
//								&scaleway.LoadbalancerFrontendAclActionRedirectArgs{
//									Type:   pulumi.String("location"),
//									Target: pulumi.String("https://example.com"),
//									Code:   pulumi.Int(307),
//								},
//							},
//						},
//						Match: &scaleway.LoadbalancerFrontendAclMatchArgs{
//							IpSubnets: pulumi.StringArray{
//								pulumi.String("10.0.0.10"),
//							},
//							HttpFilter: pulumi.String("path_begin"),
//							HttpFilterValues: pulumi.StringArray{
//								pulumi.String("foo"),
//								pulumi.String("bar"),
//							},
//						},
//					},
//				},
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
// Load-Balancer frontend can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/loadbalancerFrontend:LoadbalancerFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type LoadbalancerFrontend struct {
	pulumi.CustomResourceState

	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls LoadbalancerFrontendAclArrayOutput `pulumi:"acls"`
	// The load-balancer backend ID this frontend is attached to.
	//
	// > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
	BackendId pulumi.StringOutput `pulumi:"backendId"`
	// (Deprecated) first certificate ID used by the frontend.
	//
	// Deprecated: Please use certificate_ids
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// List of Certificate IDs that should be used by the frontend.
	//
	// > **Important:** Certificates are not allowed on port 80.
	CertificateIds pulumi.StringArrayOutput `pulumi:"certificateIds"`
	// Activates HTTP/3 protocol.
	EnableHttp3 pulumi.BoolPtrOutput `pulumi:"enableHttp3"`
	// A boolean to specify whether to use lb_acl.
	// If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
	ExternalAcls pulumi.BoolPtrOutput `pulumi:"externalAcls"`
	// TCP port to listen on the front side.
	InboundPort pulumi.IntOutput `pulumi:"inboundPort"`
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringOutput `pulumi:"lbId"`
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient pulumi.StringPtrOutput `pulumi:"timeoutClient"`
}

// NewLoadbalancerFrontend registers a new resource with the given unique name, arguments, and options.
func NewLoadbalancerFrontend(ctx *pulumi.Context,
	name string, args *LoadbalancerFrontendArgs, opts ...pulumi.ResourceOption) (*LoadbalancerFrontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendId == nil {
		return nil, errors.New("invalid value for required argument 'BackendId'")
	}
	if args.InboundPort == nil {
		return nil, errors.New("invalid value for required argument 'InboundPort'")
	}
	if args.LbId == nil {
		return nil, errors.New("invalid value for required argument 'LbId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoadbalancerFrontend
	err := ctx.RegisterResource("scaleway:index/loadbalancerFrontend:LoadbalancerFrontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadbalancerFrontend gets an existing LoadbalancerFrontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadbalancerFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadbalancerFrontendState, opts ...pulumi.ResourceOption) (*LoadbalancerFrontend, error) {
	var resource LoadbalancerFrontend
	err := ctx.ReadResource("scaleway:index/loadbalancerFrontend:LoadbalancerFrontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadbalancerFrontend resources.
type loadbalancerFrontendState struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls []LoadbalancerFrontendAcl `pulumi:"acls"`
	// The load-balancer backend ID this frontend is attached to.
	//
	// > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
	BackendId *string `pulumi:"backendId"`
	// (Deprecated) first certificate ID used by the frontend.
	//
	// Deprecated: Please use certificate_ids
	CertificateId *string `pulumi:"certificateId"`
	// List of Certificate IDs that should be used by the frontend.
	//
	// > **Important:** Certificates are not allowed on port 80.
	CertificateIds []string `pulumi:"certificateIds"`
	// Activates HTTP/3 protocol.
	EnableHttp3 *bool `pulumi:"enableHttp3"`
	// A boolean to specify whether to use lb_acl.
	// If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
	ExternalAcls *bool `pulumi:"externalAcls"`
	// TCP port to listen on the front side.
	InboundPort *int `pulumi:"inboundPort"`
	// The load-balancer ID this frontend is attached to.
	LbId *string `pulumi:"lbId"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient *string `pulumi:"timeoutClient"`
}

type LoadbalancerFrontendState struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls LoadbalancerFrontendAclArrayInput
	// The load-balancer backend ID this frontend is attached to.
	//
	// > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
	BackendId pulumi.StringPtrInput
	// (Deprecated) first certificate ID used by the frontend.
	//
	// Deprecated: Please use certificate_ids
	CertificateId pulumi.StringPtrInput
	// List of Certificate IDs that should be used by the frontend.
	//
	// > **Important:** Certificates are not allowed on port 80.
	CertificateIds pulumi.StringArrayInput
	// Activates HTTP/3 protocol.
	EnableHttp3 pulumi.BoolPtrInput
	// A boolean to specify whether to use lb_acl.
	// If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
	ExternalAcls pulumi.BoolPtrInput
	// TCP port to listen on the front side.
	InboundPort pulumi.IntPtrInput
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringPtrInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient pulumi.StringPtrInput
}

func (LoadbalancerFrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerFrontendState)(nil)).Elem()
}

type loadbalancerFrontendArgs struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls []LoadbalancerFrontendAcl `pulumi:"acls"`
	// The load-balancer backend ID this frontend is attached to.
	//
	// > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
	BackendId string `pulumi:"backendId"`
	// List of Certificate IDs that should be used by the frontend.
	//
	// > **Important:** Certificates are not allowed on port 80.
	CertificateIds []string `pulumi:"certificateIds"`
	// Activates HTTP/3 protocol.
	EnableHttp3 *bool `pulumi:"enableHttp3"`
	// A boolean to specify whether to use lb_acl.
	// If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
	ExternalAcls *bool `pulumi:"externalAcls"`
	// TCP port to listen on the front side.
	InboundPort int `pulumi:"inboundPort"`
	// The load-balancer ID this frontend is attached to.
	LbId string `pulumi:"lbId"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient *string `pulumi:"timeoutClient"`
}

// The set of arguments for constructing a LoadbalancerFrontend resource.
type LoadbalancerFrontendArgs struct {
	// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
	Acls LoadbalancerFrontendAclArrayInput
	// The load-balancer backend ID this frontend is attached to.
	//
	// > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
	BackendId pulumi.StringInput
	// List of Certificate IDs that should be used by the frontend.
	//
	// > **Important:** Certificates are not allowed on port 80.
	CertificateIds pulumi.StringArrayInput
	// Activates HTTP/3 protocol.
	EnableHttp3 pulumi.BoolPtrInput
	// A boolean to specify whether to use lb_acl.
	// If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
	ExternalAcls pulumi.BoolPtrInput
	// TCP port to listen on the front side.
	InboundPort pulumi.IntInput
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// Maximum inactivity time on the client side. (e.g.: `1s`)
	TimeoutClient pulumi.StringPtrInput
}

func (LoadbalancerFrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadbalancerFrontendArgs)(nil)).Elem()
}

type LoadbalancerFrontendInput interface {
	pulumi.Input

	ToLoadbalancerFrontendOutput() LoadbalancerFrontendOutput
	ToLoadbalancerFrontendOutputWithContext(ctx context.Context) LoadbalancerFrontendOutput
}

func (*LoadbalancerFrontend) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadbalancerFrontend)(nil)).Elem()
}

func (i *LoadbalancerFrontend) ToLoadbalancerFrontendOutput() LoadbalancerFrontendOutput {
	return i.ToLoadbalancerFrontendOutputWithContext(context.Background())
}

func (i *LoadbalancerFrontend) ToLoadbalancerFrontendOutputWithContext(ctx context.Context) LoadbalancerFrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerFrontendOutput)
}

// LoadbalancerFrontendArrayInput is an input type that accepts LoadbalancerFrontendArray and LoadbalancerFrontendArrayOutput values.
// You can construct a concrete instance of `LoadbalancerFrontendArrayInput` via:
//
//	LoadbalancerFrontendArray{ LoadbalancerFrontendArgs{...} }
type LoadbalancerFrontendArrayInput interface {
	pulumi.Input

	ToLoadbalancerFrontendArrayOutput() LoadbalancerFrontendArrayOutput
	ToLoadbalancerFrontendArrayOutputWithContext(context.Context) LoadbalancerFrontendArrayOutput
}

type LoadbalancerFrontendArray []LoadbalancerFrontendInput

func (LoadbalancerFrontendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadbalancerFrontend)(nil)).Elem()
}

func (i LoadbalancerFrontendArray) ToLoadbalancerFrontendArrayOutput() LoadbalancerFrontendArrayOutput {
	return i.ToLoadbalancerFrontendArrayOutputWithContext(context.Background())
}

func (i LoadbalancerFrontendArray) ToLoadbalancerFrontendArrayOutputWithContext(ctx context.Context) LoadbalancerFrontendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerFrontendArrayOutput)
}

// LoadbalancerFrontendMapInput is an input type that accepts LoadbalancerFrontendMap and LoadbalancerFrontendMapOutput values.
// You can construct a concrete instance of `LoadbalancerFrontendMapInput` via:
//
//	LoadbalancerFrontendMap{ "key": LoadbalancerFrontendArgs{...} }
type LoadbalancerFrontendMapInput interface {
	pulumi.Input

	ToLoadbalancerFrontendMapOutput() LoadbalancerFrontendMapOutput
	ToLoadbalancerFrontendMapOutputWithContext(context.Context) LoadbalancerFrontendMapOutput
}

type LoadbalancerFrontendMap map[string]LoadbalancerFrontendInput

func (LoadbalancerFrontendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadbalancerFrontend)(nil)).Elem()
}

func (i LoadbalancerFrontendMap) ToLoadbalancerFrontendMapOutput() LoadbalancerFrontendMapOutput {
	return i.ToLoadbalancerFrontendMapOutputWithContext(context.Background())
}

func (i LoadbalancerFrontendMap) ToLoadbalancerFrontendMapOutputWithContext(ctx context.Context) LoadbalancerFrontendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadbalancerFrontendMapOutput)
}

type LoadbalancerFrontendOutput struct{ *pulumi.OutputState }

func (LoadbalancerFrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadbalancerFrontend)(nil)).Elem()
}

func (o LoadbalancerFrontendOutput) ToLoadbalancerFrontendOutput() LoadbalancerFrontendOutput {
	return o
}

func (o LoadbalancerFrontendOutput) ToLoadbalancerFrontendOutputWithContext(ctx context.Context) LoadbalancerFrontendOutput {
	return o
}

// A list of ACL rules to apply to the load-balancer frontend.  Defined below.
func (o LoadbalancerFrontendOutput) Acls() LoadbalancerFrontendAclArrayOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) LoadbalancerFrontendAclArrayOutput { return v.Acls }).(LoadbalancerFrontendAclArrayOutput)
}

// The load-balancer backend ID this frontend is attached to.
//
// > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
func (o LoadbalancerFrontendOutput) BackendId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.StringOutput { return v.BackendId }).(pulumi.StringOutput)
}

// (Deprecated) first certificate ID used by the frontend.
//
// Deprecated: Please use certificate_ids
func (o LoadbalancerFrontendOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// List of Certificate IDs that should be used by the frontend.
//
// > **Important:** Certificates are not allowed on port 80.
func (o LoadbalancerFrontendOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.StringArrayOutput { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

// Activates HTTP/3 protocol.
func (o LoadbalancerFrontendOutput) EnableHttp3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.BoolPtrOutput { return v.EnableHttp3 }).(pulumi.BoolPtrOutput)
}

// A boolean to specify whether to use lb_acl.
// If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
func (o LoadbalancerFrontendOutput) ExternalAcls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.BoolPtrOutput { return v.ExternalAcls }).(pulumi.BoolPtrOutput)
}

// TCP port to listen on the front side.
func (o LoadbalancerFrontendOutput) InboundPort() pulumi.IntOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.IntOutput { return v.InboundPort }).(pulumi.IntOutput)
}

// The load-balancer ID this frontend is attached to.
func (o LoadbalancerFrontendOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.StringOutput { return v.LbId }).(pulumi.StringOutput)
}

// The ACL name. If not provided it will be randomly generated.
func (o LoadbalancerFrontendOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Maximum inactivity time on the client side. (e.g.: `1s`)
func (o LoadbalancerFrontendOutput) TimeoutClient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadbalancerFrontend) pulumi.StringPtrOutput { return v.TimeoutClient }).(pulumi.StringPtrOutput)
}

type LoadbalancerFrontendArrayOutput struct{ *pulumi.OutputState }

func (LoadbalancerFrontendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadbalancerFrontend)(nil)).Elem()
}

func (o LoadbalancerFrontendArrayOutput) ToLoadbalancerFrontendArrayOutput() LoadbalancerFrontendArrayOutput {
	return o
}

func (o LoadbalancerFrontendArrayOutput) ToLoadbalancerFrontendArrayOutputWithContext(ctx context.Context) LoadbalancerFrontendArrayOutput {
	return o
}

func (o LoadbalancerFrontendArrayOutput) Index(i pulumi.IntInput) LoadbalancerFrontendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadbalancerFrontend {
		return vs[0].([]*LoadbalancerFrontend)[vs[1].(int)]
	}).(LoadbalancerFrontendOutput)
}

type LoadbalancerFrontendMapOutput struct{ *pulumi.OutputState }

func (LoadbalancerFrontendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadbalancerFrontend)(nil)).Elem()
}

func (o LoadbalancerFrontendMapOutput) ToLoadbalancerFrontendMapOutput() LoadbalancerFrontendMapOutput {
	return o
}

func (o LoadbalancerFrontendMapOutput) ToLoadbalancerFrontendMapOutputWithContext(ctx context.Context) LoadbalancerFrontendMapOutput {
	return o
}

func (o LoadbalancerFrontendMapOutput) MapIndex(k pulumi.StringInput) LoadbalancerFrontendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadbalancerFrontend {
		return vs[0].(map[string]*LoadbalancerFrontend)[vs[1].(string)]
	}).(LoadbalancerFrontendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerFrontendInput)(nil)).Elem(), &LoadbalancerFrontend{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerFrontendArrayInput)(nil)).Elem(), LoadbalancerFrontendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadbalancerFrontendMapInput)(nil)).Elem(), LoadbalancerFrontendMap{})
	pulumi.RegisterOutputType(LoadbalancerFrontendOutput{})
	pulumi.RegisterOutputType(LoadbalancerFrontendArrayOutput{})
	pulumi.RegisterOutputType(LoadbalancerFrontendMapOutput{})
}
