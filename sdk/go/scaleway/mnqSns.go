// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Activate Scaleway Messaging and queuing SNS for a project.
// For further information please check
// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
//
// ## Example Usage
//
// ### Basic
//
// # Activate SNS for default project
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
//			_, err := scaleway.NewMnqSns(ctx, "main", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Activate SNS for a specific project
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
//			project, err := scaleway.LookupAccountProject(ctx, &scaleway.LookupAccountProjectArgs{
//				Name: pulumi.StringRef("default"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// For specific project in default region
//			_, err = scaleway.NewMnqSns(ctx, "forProject", &scaleway.MnqSnsArgs{
//				ProjectId: pulumi.String(project.Id),
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
// SNS status can be imported using the `{region}/{project_id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/mnqSns:MnqSns main fr-par/11111111111111111111111111111111
// ```
type MnqSns struct {
	pulumi.CustomResourceState

	// The endpoint of the SNS service for this project.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// `projectId`) The ID of the project the sns will be enabled for.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region
	// in which sns will be enabled.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewMnqSns registers a new resource with the given unique name, arguments, and options.
func NewMnqSns(ctx *pulumi.Context,
	name string, args *MnqSnsArgs, opts ...pulumi.ResourceOption) (*MnqSns, error) {
	if args == nil {
		args = &MnqSnsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqSns
	err := ctx.RegisterResource("scaleway:index/mnqSns:MnqSns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqSns gets an existing MnqSns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqSns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqSnsState, opts ...pulumi.ResourceOption) (*MnqSns, error) {
	var resource MnqSns
	err := ctx.ReadResource("scaleway:index/mnqSns:MnqSns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqSns resources.
type mnqSnsState struct {
	// The endpoint of the SNS service for this project.
	Endpoint *string `pulumi:"endpoint"`
	// `projectId`) The ID of the project the sns will be enabled for.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which sns will be enabled.
	Region *string `pulumi:"region"`
}

type MnqSnsState struct {
	// The endpoint of the SNS service for this project.
	Endpoint pulumi.StringPtrInput
	// `projectId`) The ID of the project the sns will be enabled for.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which sns will be enabled.
	Region pulumi.StringPtrInput
}

func (MnqSnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSnsState)(nil)).Elem()
}

type mnqSnsArgs struct {
	// `projectId`) The ID of the project the sns will be enabled for.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which sns will be enabled.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a MnqSns resource.
type MnqSnsArgs struct {
	// `projectId`) The ID of the project the sns will be enabled for.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which sns will be enabled.
	Region pulumi.StringPtrInput
}

func (MnqSnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSnsArgs)(nil)).Elem()
}

type MnqSnsInput interface {
	pulumi.Input

	ToMnqSnsOutput() MnqSnsOutput
	ToMnqSnsOutputWithContext(ctx context.Context) MnqSnsOutput
}

func (*MnqSns) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSns)(nil)).Elem()
}

func (i *MnqSns) ToMnqSnsOutput() MnqSnsOutput {
	return i.ToMnqSnsOutputWithContext(context.Background())
}

func (i *MnqSns) ToMnqSnsOutputWithContext(ctx context.Context) MnqSnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSnsOutput)
}

// MnqSnsArrayInput is an input type that accepts MnqSnsArray and MnqSnsArrayOutput values.
// You can construct a concrete instance of `MnqSnsArrayInput` via:
//
//	MnqSnsArray{ MnqSnsArgs{...} }
type MnqSnsArrayInput interface {
	pulumi.Input

	ToMnqSnsArrayOutput() MnqSnsArrayOutput
	ToMnqSnsArrayOutputWithContext(context.Context) MnqSnsArrayOutput
}

type MnqSnsArray []MnqSnsInput

func (MnqSnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSns)(nil)).Elem()
}

func (i MnqSnsArray) ToMnqSnsArrayOutput() MnqSnsArrayOutput {
	return i.ToMnqSnsArrayOutputWithContext(context.Background())
}

func (i MnqSnsArray) ToMnqSnsArrayOutputWithContext(ctx context.Context) MnqSnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSnsArrayOutput)
}

// MnqSnsMapInput is an input type that accepts MnqSnsMap and MnqSnsMapOutput values.
// You can construct a concrete instance of `MnqSnsMapInput` via:
//
//	MnqSnsMap{ "key": MnqSnsArgs{...} }
type MnqSnsMapInput interface {
	pulumi.Input

	ToMnqSnsMapOutput() MnqSnsMapOutput
	ToMnqSnsMapOutputWithContext(context.Context) MnqSnsMapOutput
}

type MnqSnsMap map[string]MnqSnsInput

func (MnqSnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSns)(nil)).Elem()
}

func (i MnqSnsMap) ToMnqSnsMapOutput() MnqSnsMapOutput {
	return i.ToMnqSnsMapOutputWithContext(context.Background())
}

func (i MnqSnsMap) ToMnqSnsMapOutputWithContext(ctx context.Context) MnqSnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSnsMapOutput)
}

type MnqSnsOutput struct{ *pulumi.OutputState }

func (MnqSnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSns)(nil)).Elem()
}

func (o MnqSnsOutput) ToMnqSnsOutput() MnqSnsOutput {
	return o
}

func (o MnqSnsOutput) ToMnqSnsOutputWithContext(ctx context.Context) MnqSnsOutput {
	return o
}

// The endpoint of the SNS service for this project.
func (o MnqSnsOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSns) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the sns will be enabled for.
func (o MnqSnsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSns) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region
// in which sns will be enabled.
func (o MnqSnsOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSns) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type MnqSnsArrayOutput struct{ *pulumi.OutputState }

func (MnqSnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSns)(nil)).Elem()
}

func (o MnqSnsArrayOutput) ToMnqSnsArrayOutput() MnqSnsArrayOutput {
	return o
}

func (o MnqSnsArrayOutput) ToMnqSnsArrayOutputWithContext(ctx context.Context) MnqSnsArrayOutput {
	return o
}

func (o MnqSnsArrayOutput) Index(i pulumi.IntInput) MnqSnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqSns {
		return vs[0].([]*MnqSns)[vs[1].(int)]
	}).(MnqSnsOutput)
}

type MnqSnsMapOutput struct{ *pulumi.OutputState }

func (MnqSnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSns)(nil)).Elem()
}

func (o MnqSnsMapOutput) ToMnqSnsMapOutput() MnqSnsMapOutput {
	return o
}

func (o MnqSnsMapOutput) ToMnqSnsMapOutputWithContext(ctx context.Context) MnqSnsMapOutput {
	return o
}

func (o MnqSnsMapOutput) MapIndex(k pulumi.StringInput) MnqSnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqSns {
		return vs[0].(map[string]*MnqSns)[vs[1].(string)]
	}).(MnqSnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSnsInput)(nil)).Elem(), &MnqSns{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSnsArrayInput)(nil)).Elem(), MnqSnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSnsMapInput)(nil)).Elem(), MnqSnsMap{})
	pulumi.RegisterOutputType(MnqSnsOutput{})
	pulumi.RegisterOutputType(MnqSnsArrayOutput{})
	pulumi.RegisterOutputType(MnqSnsMapOutput{})
}
