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
//			_, err := scaleway.NewTemDomainValidation(ctx, "example", &scaleway.TemDomainValidationArgs{
//				DomainId: pulumi.String("your-domain-id"),
//				Region:   pulumi.String("fr-par"),
//				Timeout:  pulumi.Int(300),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TemDomainValidation struct {
	pulumi.CustomResourceState

	// The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	Region pulumi.StringOutput `pulumi:"region"`
	// The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
	Validated pulumi.BoolOutput `pulumi:"validated"`
}

// NewTemDomainValidation registers a new resource with the given unique name, arguments, and options.
func NewTemDomainValidation(ctx *pulumi.Context,
	name string, args *TemDomainValidationArgs, opts ...pulumi.ResourceOption) (*TemDomainValidation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainId == nil {
		return nil, errors.New("invalid value for required argument 'DomainId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TemDomainValidation
	err := ctx.RegisterResource("scaleway:index/temDomainValidation:TemDomainValidation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemDomainValidation gets an existing TemDomainValidation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemDomainValidation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemDomainValidationState, opts ...pulumi.ResourceOption) (*TemDomainValidation, error) {
	var resource TemDomainValidation
	err := ctx.ReadResource("scaleway:index/temDomainValidation:TemDomainValidation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemDomainValidation resources.
type temDomainValidationState struct {
	// The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
	DomainId *string `pulumi:"domainId"`
	// `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	Region *string `pulumi:"region"`
	// The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
	Timeout *int `pulumi:"timeout"`
	// Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
	Validated *bool `pulumi:"validated"`
}

type TemDomainValidationState struct {
	// The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
	DomainId pulumi.StringPtrInput
	// `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	Region pulumi.StringPtrInput
	// The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
	Timeout pulumi.IntPtrInput
	// Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
	Validated pulumi.BoolPtrInput
}

func (TemDomainValidationState) ElementType() reflect.Type {
	return reflect.TypeOf((*temDomainValidationState)(nil)).Elem()
}

type temDomainValidationArgs struct {
	// The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
	DomainId string `pulumi:"domainId"`
	// `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	Region *string `pulumi:"region"`
	// The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a TemDomainValidation resource.
type TemDomainValidationArgs struct {
	// The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
	DomainId pulumi.StringInput
	// `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	Region pulumi.StringPtrInput
	// The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
	Timeout pulumi.IntPtrInput
}

func (TemDomainValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*temDomainValidationArgs)(nil)).Elem()
}

type TemDomainValidationInput interface {
	pulumi.Input

	ToTemDomainValidationOutput() TemDomainValidationOutput
	ToTemDomainValidationOutputWithContext(ctx context.Context) TemDomainValidationOutput
}

func (*TemDomainValidation) ElementType() reflect.Type {
	return reflect.TypeOf((**TemDomainValidation)(nil)).Elem()
}

func (i *TemDomainValidation) ToTemDomainValidationOutput() TemDomainValidationOutput {
	return i.ToTemDomainValidationOutputWithContext(context.Background())
}

func (i *TemDomainValidation) ToTemDomainValidationOutputWithContext(ctx context.Context) TemDomainValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemDomainValidationOutput)
}

// TemDomainValidationArrayInput is an input type that accepts TemDomainValidationArray and TemDomainValidationArrayOutput values.
// You can construct a concrete instance of `TemDomainValidationArrayInput` via:
//
//	TemDomainValidationArray{ TemDomainValidationArgs{...} }
type TemDomainValidationArrayInput interface {
	pulumi.Input

	ToTemDomainValidationArrayOutput() TemDomainValidationArrayOutput
	ToTemDomainValidationArrayOutputWithContext(context.Context) TemDomainValidationArrayOutput
}

type TemDomainValidationArray []TemDomainValidationInput

func (TemDomainValidationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemDomainValidation)(nil)).Elem()
}

func (i TemDomainValidationArray) ToTemDomainValidationArrayOutput() TemDomainValidationArrayOutput {
	return i.ToTemDomainValidationArrayOutputWithContext(context.Background())
}

func (i TemDomainValidationArray) ToTemDomainValidationArrayOutputWithContext(ctx context.Context) TemDomainValidationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemDomainValidationArrayOutput)
}

// TemDomainValidationMapInput is an input type that accepts TemDomainValidationMap and TemDomainValidationMapOutput values.
// You can construct a concrete instance of `TemDomainValidationMapInput` via:
//
//	TemDomainValidationMap{ "key": TemDomainValidationArgs{...} }
type TemDomainValidationMapInput interface {
	pulumi.Input

	ToTemDomainValidationMapOutput() TemDomainValidationMapOutput
	ToTemDomainValidationMapOutputWithContext(context.Context) TemDomainValidationMapOutput
}

type TemDomainValidationMap map[string]TemDomainValidationInput

func (TemDomainValidationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemDomainValidation)(nil)).Elem()
}

func (i TemDomainValidationMap) ToTemDomainValidationMapOutput() TemDomainValidationMapOutput {
	return i.ToTemDomainValidationMapOutputWithContext(context.Background())
}

func (i TemDomainValidationMap) ToTemDomainValidationMapOutputWithContext(ctx context.Context) TemDomainValidationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemDomainValidationMapOutput)
}

type TemDomainValidationOutput struct{ *pulumi.OutputState }

func (TemDomainValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemDomainValidation)(nil)).Elem()
}

func (o TemDomainValidationOutput) ToTemDomainValidationOutput() TemDomainValidationOutput {
	return o
}

func (o TemDomainValidationOutput) ToTemDomainValidationOutputWithContext(ctx context.Context) TemDomainValidationOutput {
	return o
}

// The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
func (o TemDomainValidationOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomainValidation) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
func (o TemDomainValidationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomainValidation) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
func (o TemDomainValidationOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TemDomainValidation) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
func (o TemDomainValidationOutput) Validated() pulumi.BoolOutput {
	return o.ApplyT(func(v *TemDomainValidation) pulumi.BoolOutput { return v.Validated }).(pulumi.BoolOutput)
}

type TemDomainValidationArrayOutput struct{ *pulumi.OutputState }

func (TemDomainValidationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemDomainValidation)(nil)).Elem()
}

func (o TemDomainValidationArrayOutput) ToTemDomainValidationArrayOutput() TemDomainValidationArrayOutput {
	return o
}

func (o TemDomainValidationArrayOutput) ToTemDomainValidationArrayOutputWithContext(ctx context.Context) TemDomainValidationArrayOutput {
	return o
}

func (o TemDomainValidationArrayOutput) Index(i pulumi.IntInput) TemDomainValidationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TemDomainValidation {
		return vs[0].([]*TemDomainValidation)[vs[1].(int)]
	}).(TemDomainValidationOutput)
}

type TemDomainValidationMapOutput struct{ *pulumi.OutputState }

func (TemDomainValidationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemDomainValidation)(nil)).Elem()
}

func (o TemDomainValidationMapOutput) ToTemDomainValidationMapOutput() TemDomainValidationMapOutput {
	return o
}

func (o TemDomainValidationMapOutput) ToTemDomainValidationMapOutputWithContext(ctx context.Context) TemDomainValidationMapOutput {
	return o
}

func (o TemDomainValidationMapOutput) MapIndex(k pulumi.StringInput) TemDomainValidationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TemDomainValidation {
		return vs[0].(map[string]*TemDomainValidation)[vs[1].(string)]
	}).(TemDomainValidationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemDomainValidationInput)(nil)).Elem(), &TemDomainValidation{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemDomainValidationArrayInput)(nil)).Elem(), TemDomainValidationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemDomainValidationMapInput)(nil)).Elem(), TemDomainValidationMap{})
	pulumi.RegisterOutputType(TemDomainValidationOutput{})
	pulumi.RegisterOutputType(TemDomainValidationArrayOutput{})
	pulumi.RegisterOutputType(TemDomainValidationMapOutput{})
}
