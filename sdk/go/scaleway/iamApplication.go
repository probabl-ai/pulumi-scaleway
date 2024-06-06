// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway IAM Applications. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#applications-83ce5e).
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
//			_, err := scaleway.NewIamApplication(ctx, "main", &scaleway.IamApplicationArgs{
//				Description: pulumi.String("a description"),
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
// Applications can be imported using the `{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/iamApplication:IamApplication main 11111111-1111-1111-1111-111111111111
// ```
type IamApplication struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the application.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the iam application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the application is editable.
	Editable pulumi.BoolOutput `pulumi:"editable"`
	// The name of the iam application.
	Name pulumi.StringOutput `pulumi:"name"`
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The tags associated with the application.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the application.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIamApplication registers a new resource with the given unique name, arguments, and options.
func NewIamApplication(ctx *pulumi.Context,
	name string, args *IamApplicationArgs, opts ...pulumi.ResourceOption) (*IamApplication, error) {
	if args == nil {
		args = &IamApplicationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamApplication
	err := ctx.RegisterResource("scaleway:index/iamApplication:IamApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamApplication gets an existing IamApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamApplicationState, opts ...pulumi.ResourceOption) (*IamApplication, error) {
	var resource IamApplication
	err := ctx.ReadResource("scaleway:index/iamApplication:IamApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamApplication resources.
type iamApplicationState struct {
	// The date and time of the creation of the application.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the iam application.
	Description *string `pulumi:"description"`
	// Whether the application is editable.
	Editable *bool `pulumi:"editable"`
	// The name of the iam application.
	Name *string `pulumi:"name"`
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The tags associated with the application.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the application.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IamApplicationState struct {
	// The date and time of the creation of the application.
	CreatedAt pulumi.StringPtrInput
	// The description of the iam application.
	Description pulumi.StringPtrInput
	// Whether the application is editable.
	Editable pulumi.BoolPtrInput
	// The name of the iam application.
	Name pulumi.StringPtrInput
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId pulumi.StringPtrInput
	// The tags associated with the application.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the application.
	UpdatedAt pulumi.StringPtrInput
}

func (IamApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamApplicationState)(nil)).Elem()
}

type iamApplicationArgs struct {
	// The description of the iam application.
	Description *string `pulumi:"description"`
	// The name of the iam application.
	Name *string `pulumi:"name"`
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The tags associated with the application.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a IamApplication resource.
type IamApplicationArgs struct {
	// The description of the iam application.
	Description pulumi.StringPtrInput
	// The name of the iam application.
	Name pulumi.StringPtrInput
	// `organizationId`) The ID of the organization the application is associated with.
	OrganizationId pulumi.StringPtrInput
	// The tags associated with the application.
	Tags pulumi.StringArrayInput
}

func (IamApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamApplicationArgs)(nil)).Elem()
}

type IamApplicationInput interface {
	pulumi.Input

	ToIamApplicationOutput() IamApplicationOutput
	ToIamApplicationOutputWithContext(ctx context.Context) IamApplicationOutput
}

func (*IamApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**IamApplication)(nil)).Elem()
}

func (i *IamApplication) ToIamApplicationOutput() IamApplicationOutput {
	return i.ToIamApplicationOutputWithContext(context.Background())
}

func (i *IamApplication) ToIamApplicationOutputWithContext(ctx context.Context) IamApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamApplicationOutput)
}

// IamApplicationArrayInput is an input type that accepts IamApplicationArray and IamApplicationArrayOutput values.
// You can construct a concrete instance of `IamApplicationArrayInput` via:
//
//	IamApplicationArray{ IamApplicationArgs{...} }
type IamApplicationArrayInput interface {
	pulumi.Input

	ToIamApplicationArrayOutput() IamApplicationArrayOutput
	ToIamApplicationArrayOutputWithContext(context.Context) IamApplicationArrayOutput
}

type IamApplicationArray []IamApplicationInput

func (IamApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamApplication)(nil)).Elem()
}

func (i IamApplicationArray) ToIamApplicationArrayOutput() IamApplicationArrayOutput {
	return i.ToIamApplicationArrayOutputWithContext(context.Background())
}

func (i IamApplicationArray) ToIamApplicationArrayOutputWithContext(ctx context.Context) IamApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamApplicationArrayOutput)
}

// IamApplicationMapInput is an input type that accepts IamApplicationMap and IamApplicationMapOutput values.
// You can construct a concrete instance of `IamApplicationMapInput` via:
//
//	IamApplicationMap{ "key": IamApplicationArgs{...} }
type IamApplicationMapInput interface {
	pulumi.Input

	ToIamApplicationMapOutput() IamApplicationMapOutput
	ToIamApplicationMapOutputWithContext(context.Context) IamApplicationMapOutput
}

type IamApplicationMap map[string]IamApplicationInput

func (IamApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamApplication)(nil)).Elem()
}

func (i IamApplicationMap) ToIamApplicationMapOutput() IamApplicationMapOutput {
	return i.ToIamApplicationMapOutputWithContext(context.Background())
}

func (i IamApplicationMap) ToIamApplicationMapOutputWithContext(ctx context.Context) IamApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamApplicationMapOutput)
}

type IamApplicationOutput struct{ *pulumi.OutputState }

func (IamApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamApplication)(nil)).Elem()
}

func (o IamApplicationOutput) ToIamApplicationOutput() IamApplicationOutput {
	return o
}

func (o IamApplicationOutput) ToIamApplicationOutputWithContext(ctx context.Context) IamApplicationOutput {
	return o
}

// The date and time of the creation of the application.
func (o IamApplicationOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApplication) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the iam application.
func (o IamApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamApplication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the application is editable.
func (o IamApplicationOutput) Editable() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamApplication) pulumi.BoolOutput { return v.Editable }).(pulumi.BoolOutput)
}

// The name of the iam application.
func (o IamApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `organizationId`) The ID of the organization the application is associated with.
func (o IamApplicationOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApplication) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The tags associated with the application.
func (o IamApplicationOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IamApplication) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the application.
func (o IamApplicationOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamApplication) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IamApplicationArrayOutput struct{ *pulumi.OutputState }

func (IamApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamApplication)(nil)).Elem()
}

func (o IamApplicationArrayOutput) ToIamApplicationArrayOutput() IamApplicationArrayOutput {
	return o
}

func (o IamApplicationArrayOutput) ToIamApplicationArrayOutputWithContext(ctx context.Context) IamApplicationArrayOutput {
	return o
}

func (o IamApplicationArrayOutput) Index(i pulumi.IntInput) IamApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamApplication {
		return vs[0].([]*IamApplication)[vs[1].(int)]
	}).(IamApplicationOutput)
}

type IamApplicationMapOutput struct{ *pulumi.OutputState }

func (IamApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamApplication)(nil)).Elem()
}

func (o IamApplicationMapOutput) ToIamApplicationMapOutput() IamApplicationMapOutput {
	return o
}

func (o IamApplicationMapOutput) ToIamApplicationMapOutputWithContext(ctx context.Context) IamApplicationMapOutput {
	return o
}

func (o IamApplicationMapOutput) MapIndex(k pulumi.StringInput) IamApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamApplication {
		return vs[0].(map[string]*IamApplication)[vs[1].(string)]
	}).(IamApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamApplicationInput)(nil)).Elem(), &IamApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamApplicationArrayInput)(nil)).Elem(), IamApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamApplicationMapInput)(nil)).Elem(), IamApplicationMap{})
	pulumi.RegisterOutputType(IamApplicationOutput{})
	pulumi.RegisterOutputType(IamApplicationArrayOutput{})
	pulumi.RegisterOutputType(IamApplicationMapOutput{})
}
