// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// ## Import
//
// IoT Hubs can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/iotHub:IotHub hub01 fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type IotHub struct {
	pulumi.CustomResourceState

	// The current number of connected devices in the Hub.
	ConnectedDeviceCount pulumi.IntOutput `pulumi:"connectedDeviceCount"`
	// The date and time the Hub was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Wether to enable the device auto provisioning or not
	DeviceAutoProvisioning pulumi.BoolPtrOutput `pulumi:"deviceAutoProvisioning"`
	// The number of registered devices in the Hub.
	DeviceCount pulumi.IntOutput `pulumi:"deviceCount"`
	// Whether to enable the hub events or not
	DisableEvents pulumi.BoolPtrOutput `pulumi:"disableEvents"`
	// Wether the IoT Hub instance should be enabled or not.
	//
	// > **Important:** Updates to `enabled` will disconnect eventually connected devices.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The MQTT network endpoint to connect MQTT devices to.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Topic prefix for the hub events
	EventsTopicPrefix pulumi.StringPtrOutput `pulumi:"eventsTopicPrefix"`
	// Custom user provided certificate authority
	HubCa pulumi.StringPtrOutput `pulumi:"hubCa"`
	// Challenge certificate for the user provided hub CA
	HubCaChallenge pulumi.StringPtrOutput `pulumi:"hubCaChallenge"`
	// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
	//
	// > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
	ProductPlan pulumi.StringOutput `pulumi:"productPlan"`
	// `projectId`) The ID of the project the IoT Hub Instance is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The current status of the Hub.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date and time the Hub resource was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIotHub registers a new resource with the given unique name, arguments, and options.
func NewIotHub(ctx *pulumi.Context,
	name string, args *IotHubArgs, opts ...pulumi.ResourceOption) (*IotHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductPlan == nil {
		return nil, errors.New("invalid value for required argument 'ProductPlan'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IotHub
	err := ctx.RegisterResource("scaleway:index/iotHub:IotHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHub gets an existing IotHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubState, opts ...pulumi.ResourceOption) (*IotHub, error) {
	var resource IotHub
	err := ctx.ReadResource("scaleway:index/iotHub:IotHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHub resources.
type iotHubState struct {
	// The current number of connected devices in the Hub.
	ConnectedDeviceCount *int `pulumi:"connectedDeviceCount"`
	// The date and time the Hub was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Wether to enable the device auto provisioning or not
	DeviceAutoProvisioning *bool `pulumi:"deviceAutoProvisioning"`
	// The number of registered devices in the Hub.
	DeviceCount *int `pulumi:"deviceCount"`
	// Whether to enable the hub events or not
	DisableEvents *bool `pulumi:"disableEvents"`
	// Wether the IoT Hub instance should be enabled or not.
	//
	// > **Important:** Updates to `enabled` will disconnect eventually connected devices.
	Enabled *bool `pulumi:"enabled"`
	// The MQTT network endpoint to connect MQTT devices to.
	Endpoint *string `pulumi:"endpoint"`
	// Topic prefix for the hub events
	EventsTopicPrefix *string `pulumi:"eventsTopicPrefix"`
	// Custom user provided certificate authority
	HubCa *string `pulumi:"hubCa"`
	// Challenge certificate for the user provided hub CA
	HubCaChallenge *string `pulumi:"hubCaChallenge"`
	// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
	//
	// > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
	ProductPlan *string `pulumi:"productPlan"`
	// `projectId`) The ID of the project the IoT Hub Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// The current status of the Hub.
	Status *string `pulumi:"status"`
	// The date and time the Hub resource was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IotHubState struct {
	// The current number of connected devices in the Hub.
	ConnectedDeviceCount pulumi.IntPtrInput
	// The date and time the Hub was created.
	CreatedAt pulumi.StringPtrInput
	// Wether to enable the device auto provisioning or not
	DeviceAutoProvisioning pulumi.BoolPtrInput
	// The number of registered devices in the Hub.
	DeviceCount pulumi.IntPtrInput
	// Whether to enable the hub events or not
	DisableEvents pulumi.BoolPtrInput
	// Wether the IoT Hub instance should be enabled or not.
	//
	// > **Important:** Updates to `enabled` will disconnect eventually connected devices.
	Enabled pulumi.BoolPtrInput
	// The MQTT network endpoint to connect MQTT devices to.
	Endpoint pulumi.StringPtrInput
	// Topic prefix for the hub events
	EventsTopicPrefix pulumi.StringPtrInput
	// Custom user provided certificate authority
	HubCa pulumi.StringPtrInput
	// Challenge certificate for the user provided hub CA
	HubCaChallenge pulumi.StringPtrInput
	// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
	//
	// > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
	ProductPlan pulumi.StringPtrInput
	// `projectId`) The ID of the project the IoT Hub Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// The current status of the Hub.
	Status pulumi.StringPtrInput
	// The date and time the Hub resource was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (IotHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubState)(nil)).Elem()
}

type iotHubArgs struct {
	// Wether to enable the device auto provisioning or not
	DeviceAutoProvisioning *bool `pulumi:"deviceAutoProvisioning"`
	// Whether to enable the hub events or not
	DisableEvents *bool `pulumi:"disableEvents"`
	// Wether the IoT Hub instance should be enabled or not.
	//
	// > **Important:** Updates to `enabled` will disconnect eventually connected devices.
	Enabled *bool `pulumi:"enabled"`
	// Topic prefix for the hub events
	EventsTopicPrefix *string `pulumi:"eventsTopicPrefix"`
	// Custom user provided certificate authority
	HubCa *string `pulumi:"hubCa"`
	// Challenge certificate for the user provided hub CA
	HubCaChallenge *string `pulumi:"hubCaChallenge"`
	// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
	Name *string `pulumi:"name"`
	// Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
	//
	// > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
	ProductPlan string `pulumi:"productPlan"`
	// `projectId`) The ID of the project the IoT Hub Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the Database Instance should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a IotHub resource.
type IotHubArgs struct {
	// Wether to enable the device auto provisioning or not
	DeviceAutoProvisioning pulumi.BoolPtrInput
	// Whether to enable the hub events or not
	DisableEvents pulumi.BoolPtrInput
	// Wether the IoT Hub instance should be enabled or not.
	//
	// > **Important:** Updates to `enabled` will disconnect eventually connected devices.
	Enabled pulumi.BoolPtrInput
	// Topic prefix for the hub events
	EventsTopicPrefix pulumi.StringPtrInput
	// Custom user provided certificate authority
	HubCa pulumi.StringPtrInput
	// Challenge certificate for the user provided hub CA
	HubCaChallenge pulumi.StringPtrInput
	// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
	Name pulumi.StringPtrInput
	// Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
	//
	// > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
	ProductPlan pulumi.StringInput
	// `projectId`) The ID of the project the IoT Hub Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringPtrInput
}

func (IotHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubArgs)(nil)).Elem()
}

type IotHubInput interface {
	pulumi.Input

	ToIotHubOutput() IotHubOutput
	ToIotHubOutputWithContext(ctx context.Context) IotHubOutput
}

func (*IotHub) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHub)(nil)).Elem()
}

func (i *IotHub) ToIotHubOutput() IotHubOutput {
	return i.ToIotHubOutputWithContext(context.Background())
}

func (i *IotHub) ToIotHubOutputWithContext(ctx context.Context) IotHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubOutput)
}

func (i *IotHub) ToOutput(ctx context.Context) pulumix.Output[*IotHub] {
	return pulumix.Output[*IotHub]{
		OutputState: i.ToIotHubOutputWithContext(ctx).OutputState,
	}
}

// IotHubArrayInput is an input type that accepts IotHubArray and IotHubArrayOutput values.
// You can construct a concrete instance of `IotHubArrayInput` via:
//
//	IotHubArray{ IotHubArgs{...} }
type IotHubArrayInput interface {
	pulumi.Input

	ToIotHubArrayOutput() IotHubArrayOutput
	ToIotHubArrayOutputWithContext(context.Context) IotHubArrayOutput
}

type IotHubArray []IotHubInput

func (IotHubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotHub)(nil)).Elem()
}

func (i IotHubArray) ToIotHubArrayOutput() IotHubArrayOutput {
	return i.ToIotHubArrayOutputWithContext(context.Background())
}

func (i IotHubArray) ToIotHubArrayOutputWithContext(ctx context.Context) IotHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubArrayOutput)
}

func (i IotHubArray) ToOutput(ctx context.Context) pulumix.Output[[]*IotHub] {
	return pulumix.Output[[]*IotHub]{
		OutputState: i.ToIotHubArrayOutputWithContext(ctx).OutputState,
	}
}

// IotHubMapInput is an input type that accepts IotHubMap and IotHubMapOutput values.
// You can construct a concrete instance of `IotHubMapInput` via:
//
//	IotHubMap{ "key": IotHubArgs{...} }
type IotHubMapInput interface {
	pulumi.Input

	ToIotHubMapOutput() IotHubMapOutput
	ToIotHubMapOutputWithContext(context.Context) IotHubMapOutput
}

type IotHubMap map[string]IotHubInput

func (IotHubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotHub)(nil)).Elem()
}

func (i IotHubMap) ToIotHubMapOutput() IotHubMapOutput {
	return i.ToIotHubMapOutputWithContext(context.Background())
}

func (i IotHubMap) ToIotHubMapOutputWithContext(ctx context.Context) IotHubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubMapOutput)
}

func (i IotHubMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IotHub] {
	return pulumix.Output[map[string]*IotHub]{
		OutputState: i.ToIotHubMapOutputWithContext(ctx).OutputState,
	}
}

type IotHubOutput struct{ *pulumi.OutputState }

func (IotHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHub)(nil)).Elem()
}

func (o IotHubOutput) ToIotHubOutput() IotHubOutput {
	return o
}

func (o IotHubOutput) ToIotHubOutputWithContext(ctx context.Context) IotHubOutput {
	return o
}

func (o IotHubOutput) ToOutput(ctx context.Context) pulumix.Output[*IotHub] {
	return pulumix.Output[*IotHub]{
		OutputState: o.OutputState,
	}
}

// The current number of connected devices in the Hub.
func (o IotHubOutput) ConnectedDeviceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *IotHub) pulumi.IntOutput { return v.ConnectedDeviceCount }).(pulumi.IntOutput)
}

// The date and time the Hub was created.
func (o IotHubOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Wether to enable the device auto provisioning or not
func (o IotHubOutput) DeviceAutoProvisioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotHub) pulumi.BoolPtrOutput { return v.DeviceAutoProvisioning }).(pulumi.BoolPtrOutput)
}

// The number of registered devices in the Hub.
func (o IotHubOutput) DeviceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *IotHub) pulumi.IntOutput { return v.DeviceCount }).(pulumi.IntOutput)
}

// Whether to enable the hub events or not
func (o IotHubOutput) DisableEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotHub) pulumi.BoolPtrOutput { return v.DisableEvents }).(pulumi.BoolPtrOutput)
}

// Wether the IoT Hub instance should be enabled or not.
//
// > **Important:** Updates to `enabled` will disconnect eventually connected devices.
func (o IotHubOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotHub) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The MQTT network endpoint to connect MQTT devices to.
func (o IotHubOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Topic prefix for the hub events
func (o IotHubOutput) EventsTopicPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringPtrOutput { return v.EventsTopicPrefix }).(pulumi.StringPtrOutput)
}

// Custom user provided certificate authority
func (o IotHubOutput) HubCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringPtrOutput { return v.HubCa }).(pulumi.StringPtrOutput)
}

// Challenge certificate for the user provided hub CA
func (o IotHubOutput) HubCaChallenge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringPtrOutput { return v.HubCaChallenge }).(pulumi.StringPtrOutput)
}

// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
func (o IotHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization_id you want to attach the resource to
func (o IotHubOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
//
// > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
func (o IotHubOutput) ProductPlan() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.ProductPlan }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the IoT Hub Instance is associated with.
func (o IotHubOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region in which the Database Instance should be created.
func (o IotHubOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current status of the Hub.
func (o IotHubOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The date and time the Hub resource was updated.
func (o IotHubOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IotHub) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IotHubArrayOutput struct{ *pulumi.OutputState }

func (IotHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotHub)(nil)).Elem()
}

func (o IotHubArrayOutput) ToIotHubArrayOutput() IotHubArrayOutput {
	return o
}

func (o IotHubArrayOutput) ToIotHubArrayOutputWithContext(ctx context.Context) IotHubArrayOutput {
	return o
}

func (o IotHubArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IotHub] {
	return pulumix.Output[[]*IotHub]{
		OutputState: o.OutputState,
	}
}

func (o IotHubArrayOutput) Index(i pulumi.IntInput) IotHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IotHub {
		return vs[0].([]*IotHub)[vs[1].(int)]
	}).(IotHubOutput)
}

type IotHubMapOutput struct{ *pulumi.OutputState }

func (IotHubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotHub)(nil)).Elem()
}

func (o IotHubMapOutput) ToIotHubMapOutput() IotHubMapOutput {
	return o
}

func (o IotHubMapOutput) ToIotHubMapOutputWithContext(ctx context.Context) IotHubMapOutput {
	return o
}

func (o IotHubMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IotHub] {
	return pulumix.Output[map[string]*IotHub]{
		OutputState: o.OutputState,
	}
}

func (o IotHubMapOutput) MapIndex(k pulumi.StringInput) IotHubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IotHub {
		return vs[0].(map[string]*IotHub)[vs[1].(string)]
	}).(IotHubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubInput)(nil)).Elem(), &IotHub{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubArrayInput)(nil)).Elem(), IotHubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubMapInput)(nil)).Elem(), IotHubMap{})
	pulumi.RegisterOutputType(IotHubOutput{})
	pulumi.RegisterOutputType(IotHubArrayOutput{})
	pulumi.RegisterOutputType(IotHubMapOutput{})
}
