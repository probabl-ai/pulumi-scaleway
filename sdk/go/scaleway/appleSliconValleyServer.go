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

// Creates and manages Scaleway Apple silicon M1. For more information,
// see [the documentation](https://www.scaleway.com/en/docs/compute/apple-silicon/concepts).
//
// ## Examples
//
// ### Basic
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := scaleway.NewAppleSliconValleyServer(ctx, "server", &scaleway.AppleSliconValleyServerArgs{
//				Type: pulumi.String("M1-M"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Instance servers can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/appleSliconValleyServer:AppleSliconValleyServer server fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type AppleSliconValleyServer struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the Apple Silicon server.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt pulumi.StringOutput `pulumi:"deletableAt"`
	// IPv4 address of the server (IPv4 address).
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The name of the server.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the server is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The state of the server. Check the possible values on
	// our [sdk](https://github.com/scaleway/scaleway-sdk-go/blob/master/api/applesilicon/v1alpha1/applesilicon_sdk.go#L103).
	State pulumi.StringOutput `pulumi:"state"`
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the Apple Silicon server.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// URL of the VNC.
	VncUrl pulumi.StringOutput `pulumi:"vncUrl"`
	// `zone`) The zone in which
	// the server should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewAppleSliconValleyServer registers a new resource with the given unique name, arguments, and options.
func NewAppleSliconValleyServer(ctx *pulumi.Context,
	name string, args *AppleSliconValleyServerArgs, opts ...pulumi.ResourceOption) (*AppleSliconValleyServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppleSliconValleyServer
	err := ctx.RegisterResource("scaleway:index/appleSliconValleyServer:AppleSliconValleyServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppleSliconValleyServer gets an existing AppleSliconValleyServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppleSliconValleyServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppleSliconValleyServerState, opts ...pulumi.ResourceOption) (*AppleSliconValleyServer, error) {
	var resource AppleSliconValleyServer
	err := ctx.ReadResource("scaleway:index/appleSliconValleyServer:AppleSliconValleyServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppleSliconValleyServer resources.
type appleSliconValleyServerState struct {
	// The date and time of the creation of the Apple Silicon server.
	CreatedAt *string `pulumi:"createdAt"`
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt *string `pulumi:"deletableAt"`
	// IPv4 address of the server (IPv4 address).
	Ip *string `pulumi:"ip"`
	// The name of the server.
	Name *string `pulumi:"name"`
	// The organization ID the server is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// The state of the server. Check the possible values on
	// our [sdk](https://github.com/scaleway/scaleway-sdk-go/blob/master/api/applesilicon/v1alpha1/applesilicon_sdk.go#L103).
	State *string `pulumi:"state"`
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type *string `pulumi:"type"`
	// The date and time of the last update of the Apple Silicon server.
	UpdatedAt *string `pulumi:"updatedAt"`
	// URL of the VNC.
	VncUrl *string `pulumi:"vncUrl"`
	// `zone`) The zone in which
	// the server should be created.
	Zone *string `pulumi:"zone"`
}

type AppleSliconValleyServerState struct {
	// The date and time of the creation of the Apple Silicon server.
	CreatedAt pulumi.StringPtrInput
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt pulumi.StringPtrInput
	// IPv4 address of the server (IPv4 address).
	Ip pulumi.StringPtrInput
	// The name of the server.
	Name pulumi.StringPtrInput
	// The organization ID the server is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// The state of the server. Check the possible values on
	// our [sdk](https://github.com/scaleway/scaleway-sdk-go/blob/master/api/applesilicon/v1alpha1/applesilicon_sdk.go#L103).
	State pulumi.StringPtrInput
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type pulumi.StringPtrInput
	// The date and time of the last update of the Apple Silicon server.
	UpdatedAt pulumi.StringPtrInput
	// URL of the VNC.
	VncUrl pulumi.StringPtrInput
	// `zone`) The zone in which
	// the server should be created.
	Zone pulumi.StringPtrInput
}

func (AppleSliconValleyServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*appleSliconValleyServerState)(nil)).Elem()
}

type appleSliconValleyServerArgs struct {
	// The name of the server.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type string `pulumi:"type"`
	// `zone`) The zone in which
	// the server should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a AppleSliconValleyServer resource.
type AppleSliconValleyServerArgs struct {
	// The name of the server.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type pulumi.StringInput
	// `zone`) The zone in which
	// the server should be created.
	Zone pulumi.StringPtrInput
}

func (AppleSliconValleyServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appleSliconValleyServerArgs)(nil)).Elem()
}

type AppleSliconValleyServerInput interface {
	pulumi.Input

	ToAppleSliconValleyServerOutput() AppleSliconValleyServerOutput
	ToAppleSliconValleyServerOutputWithContext(ctx context.Context) AppleSliconValleyServerOutput
}

func (*AppleSliconValleyServer) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleSliconValleyServer)(nil)).Elem()
}

func (i *AppleSliconValleyServer) ToAppleSliconValleyServerOutput() AppleSliconValleyServerOutput {
	return i.ToAppleSliconValleyServerOutputWithContext(context.Background())
}

func (i *AppleSliconValleyServer) ToAppleSliconValleyServerOutputWithContext(ctx context.Context) AppleSliconValleyServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleSliconValleyServerOutput)
}

// AppleSliconValleyServerArrayInput is an input type that accepts AppleSliconValleyServerArray and AppleSliconValleyServerArrayOutput values.
// You can construct a concrete instance of `AppleSliconValleyServerArrayInput` via:
//
//	AppleSliconValleyServerArray{ AppleSliconValleyServerArgs{...} }
type AppleSliconValleyServerArrayInput interface {
	pulumi.Input

	ToAppleSliconValleyServerArrayOutput() AppleSliconValleyServerArrayOutput
	ToAppleSliconValleyServerArrayOutputWithContext(context.Context) AppleSliconValleyServerArrayOutput
}

type AppleSliconValleyServerArray []AppleSliconValleyServerInput

func (AppleSliconValleyServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppleSliconValleyServer)(nil)).Elem()
}

func (i AppleSliconValleyServerArray) ToAppleSliconValleyServerArrayOutput() AppleSliconValleyServerArrayOutput {
	return i.ToAppleSliconValleyServerArrayOutputWithContext(context.Background())
}

func (i AppleSliconValleyServerArray) ToAppleSliconValleyServerArrayOutputWithContext(ctx context.Context) AppleSliconValleyServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleSliconValleyServerArrayOutput)
}

// AppleSliconValleyServerMapInput is an input type that accepts AppleSliconValleyServerMap and AppleSliconValleyServerMapOutput values.
// You can construct a concrete instance of `AppleSliconValleyServerMapInput` via:
//
//	AppleSliconValleyServerMap{ "key": AppleSliconValleyServerArgs{...} }
type AppleSliconValleyServerMapInput interface {
	pulumi.Input

	ToAppleSliconValleyServerMapOutput() AppleSliconValleyServerMapOutput
	ToAppleSliconValleyServerMapOutputWithContext(context.Context) AppleSliconValleyServerMapOutput
}

type AppleSliconValleyServerMap map[string]AppleSliconValleyServerInput

func (AppleSliconValleyServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppleSliconValleyServer)(nil)).Elem()
}

func (i AppleSliconValleyServerMap) ToAppleSliconValleyServerMapOutput() AppleSliconValleyServerMapOutput {
	return i.ToAppleSliconValleyServerMapOutputWithContext(context.Background())
}

func (i AppleSliconValleyServerMap) ToAppleSliconValleyServerMapOutputWithContext(ctx context.Context) AppleSliconValleyServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleSliconValleyServerMapOutput)
}

type AppleSliconValleyServerOutput struct{ *pulumi.OutputState }

func (AppleSliconValleyServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleSliconValleyServer)(nil)).Elem()
}

func (o AppleSliconValleyServerOutput) ToAppleSliconValleyServerOutput() AppleSliconValleyServerOutput {
	return o
}

func (o AppleSliconValleyServerOutput) ToAppleSliconValleyServerOutputWithContext(ctx context.Context) AppleSliconValleyServerOutput {
	return o
}

// The date and time of the creation of the Apple Silicon server.
func (o AppleSliconValleyServerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The minimal date and time on which you can delete this server due to Apple licence
func (o AppleSliconValleyServerOutput) DeletableAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.DeletableAt }).(pulumi.StringOutput)
}

// IPv4 address of the server (IPv4 address).
func (o AppleSliconValleyServerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The name of the server.
func (o AppleSliconValleyServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the server is associated with.
func (o AppleSliconValleyServerOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the server is
// associated with.
func (o AppleSliconValleyServerOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The state of the server. Check the possible values on
// our [sdk](https://github.com/scaleway/scaleway-sdk-go/blob/master/api/applesilicon/v1alpha1/applesilicon_sdk.go#L103).
func (o AppleSliconValleyServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The commercial type of the server. You find all the available types on
// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
// resource.
func (o AppleSliconValleyServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date and time of the last update of the Apple Silicon server.
func (o AppleSliconValleyServerOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// URL of the VNC.
func (o AppleSliconValleyServerOutput) VncUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.VncUrl }).(pulumi.StringOutput)
}

// `zone`) The zone in which
// the server should be created.
func (o AppleSliconValleyServerOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *AppleSliconValleyServer) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type AppleSliconValleyServerArrayOutput struct{ *pulumi.OutputState }

func (AppleSliconValleyServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppleSliconValleyServer)(nil)).Elem()
}

func (o AppleSliconValleyServerArrayOutput) ToAppleSliconValleyServerArrayOutput() AppleSliconValleyServerArrayOutput {
	return o
}

func (o AppleSliconValleyServerArrayOutput) ToAppleSliconValleyServerArrayOutputWithContext(ctx context.Context) AppleSliconValleyServerArrayOutput {
	return o
}

func (o AppleSliconValleyServerArrayOutput) Index(i pulumi.IntInput) AppleSliconValleyServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppleSliconValleyServer {
		return vs[0].([]*AppleSliconValleyServer)[vs[1].(int)]
	}).(AppleSliconValleyServerOutput)
}

type AppleSliconValleyServerMapOutput struct{ *pulumi.OutputState }

func (AppleSliconValleyServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppleSliconValleyServer)(nil)).Elem()
}

func (o AppleSliconValleyServerMapOutput) ToAppleSliconValleyServerMapOutput() AppleSliconValleyServerMapOutput {
	return o
}

func (o AppleSliconValleyServerMapOutput) ToAppleSliconValleyServerMapOutputWithContext(ctx context.Context) AppleSliconValleyServerMapOutput {
	return o
}

func (o AppleSliconValleyServerMapOutput) MapIndex(k pulumi.StringInput) AppleSliconValleyServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppleSliconValleyServer {
		return vs[0].(map[string]*AppleSliconValleyServer)[vs[1].(string)]
	}).(AppleSliconValleyServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppleSliconValleyServerInput)(nil)).Elem(), &AppleSliconValleyServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppleSliconValleyServerArrayInput)(nil)).Elem(), AppleSliconValleyServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppleSliconValleyServerMapInput)(nil)).Elem(), AppleSliconValleyServerMap{})
	pulumi.RegisterOutputType(AppleSliconValleyServerOutput{})
	pulumi.RegisterOutputType(AppleSliconValleyServerArrayOutput{})
	pulumi.RegisterOutputType(AppleSliconValleyServerMapOutput{})
}
