// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Serverless SQL Databases. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/serverless-databases/).
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
//			_, err := scaleway.NewSdbDatabase(ctx, "database", &scaleway.SdbDatabaseArgs{
//				MaxCpu: pulumi.Int(8),
//				MinCpu: pulumi.Int(0),
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
// Serverless SQL Database can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/sdbDatabase:SdbDatabase database fr-par/11111111-1111-1111-1111-111111111111
// ```
type SdbDatabase struct {
	pulumi.CustomResourceState

	// Endpoint of the database
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The maximum number of CPU units for your database. Defaults to 15.
	MaxCpu pulumi.IntPtrOutput `pulumi:"maxCpu"`
	// The minimum number of CPU units for your database. Defaults to 0.
	MinCpu pulumi.IntPtrOutput `pulumi:"minCpu"`
	// Name of the database (e.g. `my-new-database`).
	//
	// > **Important:** Updates to `name` will recreate the database.
	Name pulumi.StringOutput `pulumi:"name"`
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewSdbDatabase registers a new resource with the given unique name, arguments, and options.
func NewSdbDatabase(ctx *pulumi.Context,
	name string, args *SdbDatabaseArgs, opts ...pulumi.ResourceOption) (*SdbDatabase, error) {
	if args == nil {
		args = &SdbDatabaseArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SdbDatabase
	err := ctx.RegisterResource("scaleway:index/sdbDatabase:SdbDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSdbDatabase gets an existing SdbDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSdbDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SdbDatabaseState, opts ...pulumi.ResourceOption) (*SdbDatabase, error) {
	var resource SdbDatabase
	err := ctx.ReadResource("scaleway:index/sdbDatabase:SdbDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SdbDatabase resources.
type sdbDatabaseState struct {
	// Endpoint of the database
	Endpoint *string `pulumi:"endpoint"`
	// The maximum number of CPU units for your database. Defaults to 15.
	MaxCpu *int `pulumi:"maxCpu"`
	// The minimum number of CPU units for your database. Defaults to 0.
	MinCpu *int `pulumi:"minCpu"`
	// Name of the database (e.g. `my-new-database`).
	//
	// > **Important:** Updates to `name` will recreate the database.
	Name *string `pulumi:"name"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
}

type SdbDatabaseState struct {
	// Endpoint of the database
	Endpoint pulumi.StringPtrInput
	// The maximum number of CPU units for your database. Defaults to 15.
	MaxCpu pulumi.IntPtrInput
	// The minimum number of CPU units for your database. Defaults to 0.
	MinCpu pulumi.IntPtrInput
	// Name of the database (e.g. `my-new-database`).
	//
	// > **Important:** Updates to `name` will recreate the database.
	Name pulumi.StringPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
}

func (SdbDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sdbDatabaseState)(nil)).Elem()
}

type sdbDatabaseArgs struct {
	// The maximum number of CPU units for your database. Defaults to 15.
	MaxCpu *int `pulumi:"maxCpu"`
	// The minimum number of CPU units for your database. Defaults to 0.
	MinCpu *int `pulumi:"minCpu"`
	// Name of the database (e.g. `my-new-database`).
	//
	// > **Important:** Updates to `name` will recreate the database.
	Name *string `pulumi:"name"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a SdbDatabase resource.
type SdbDatabaseArgs struct {
	// The maximum number of CPU units for your database. Defaults to 15.
	MaxCpu pulumi.IntPtrInput
	// The minimum number of CPU units for your database. Defaults to 0.
	MinCpu pulumi.IntPtrInput
	// Name of the database (e.g. `my-new-database`).
	//
	// > **Important:** Updates to `name` will recreate the database.
	Name pulumi.StringPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
}

func (SdbDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sdbDatabaseArgs)(nil)).Elem()
}

type SdbDatabaseInput interface {
	pulumi.Input

	ToSdbDatabaseOutput() SdbDatabaseOutput
	ToSdbDatabaseOutputWithContext(ctx context.Context) SdbDatabaseOutput
}

func (*SdbDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**SdbDatabase)(nil)).Elem()
}

func (i *SdbDatabase) ToSdbDatabaseOutput() SdbDatabaseOutput {
	return i.ToSdbDatabaseOutputWithContext(context.Background())
}

func (i *SdbDatabase) ToSdbDatabaseOutputWithContext(ctx context.Context) SdbDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdbDatabaseOutput)
}

// SdbDatabaseArrayInput is an input type that accepts SdbDatabaseArray and SdbDatabaseArrayOutput values.
// You can construct a concrete instance of `SdbDatabaseArrayInput` via:
//
//	SdbDatabaseArray{ SdbDatabaseArgs{...} }
type SdbDatabaseArrayInput interface {
	pulumi.Input

	ToSdbDatabaseArrayOutput() SdbDatabaseArrayOutput
	ToSdbDatabaseArrayOutputWithContext(context.Context) SdbDatabaseArrayOutput
}

type SdbDatabaseArray []SdbDatabaseInput

func (SdbDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SdbDatabase)(nil)).Elem()
}

func (i SdbDatabaseArray) ToSdbDatabaseArrayOutput() SdbDatabaseArrayOutput {
	return i.ToSdbDatabaseArrayOutputWithContext(context.Background())
}

func (i SdbDatabaseArray) ToSdbDatabaseArrayOutputWithContext(ctx context.Context) SdbDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdbDatabaseArrayOutput)
}

// SdbDatabaseMapInput is an input type that accepts SdbDatabaseMap and SdbDatabaseMapOutput values.
// You can construct a concrete instance of `SdbDatabaseMapInput` via:
//
//	SdbDatabaseMap{ "key": SdbDatabaseArgs{...} }
type SdbDatabaseMapInput interface {
	pulumi.Input

	ToSdbDatabaseMapOutput() SdbDatabaseMapOutput
	ToSdbDatabaseMapOutputWithContext(context.Context) SdbDatabaseMapOutput
}

type SdbDatabaseMap map[string]SdbDatabaseInput

func (SdbDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SdbDatabase)(nil)).Elem()
}

func (i SdbDatabaseMap) ToSdbDatabaseMapOutput() SdbDatabaseMapOutput {
	return i.ToSdbDatabaseMapOutputWithContext(context.Background())
}

func (i SdbDatabaseMap) ToSdbDatabaseMapOutputWithContext(ctx context.Context) SdbDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SdbDatabaseMapOutput)
}

type SdbDatabaseOutput struct{ *pulumi.OutputState }

func (SdbDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SdbDatabase)(nil)).Elem()
}

func (o SdbDatabaseOutput) ToSdbDatabaseOutput() SdbDatabaseOutput {
	return o
}

func (o SdbDatabaseOutput) ToSdbDatabaseOutputWithContext(ctx context.Context) SdbDatabaseOutput {
	return o
}

// Endpoint of the database
func (o SdbDatabaseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *SdbDatabase) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The maximum number of CPU units for your database. Defaults to 15.
func (o SdbDatabaseOutput) MaxCpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SdbDatabase) pulumi.IntPtrOutput { return v.MaxCpu }).(pulumi.IntPtrOutput)
}

// The minimum number of CPU units for your database. Defaults to 0.
func (o SdbDatabaseOutput) MinCpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SdbDatabase) pulumi.IntPtrOutput { return v.MinCpu }).(pulumi.IntPtrOutput)
}

// Name of the database (e.g. `my-new-database`).
//
// > **Important:** Updates to `name` will recreate the database.
func (o SdbDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SdbDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The projectId you want to attach the resource to
func (o SdbDatabaseOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SdbDatabase) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o SdbDatabaseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SdbDatabase) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type SdbDatabaseArrayOutput struct{ *pulumi.OutputState }

func (SdbDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SdbDatabase)(nil)).Elem()
}

func (o SdbDatabaseArrayOutput) ToSdbDatabaseArrayOutput() SdbDatabaseArrayOutput {
	return o
}

func (o SdbDatabaseArrayOutput) ToSdbDatabaseArrayOutputWithContext(ctx context.Context) SdbDatabaseArrayOutput {
	return o
}

func (o SdbDatabaseArrayOutput) Index(i pulumi.IntInput) SdbDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SdbDatabase {
		return vs[0].([]*SdbDatabase)[vs[1].(int)]
	}).(SdbDatabaseOutput)
}

type SdbDatabaseMapOutput struct{ *pulumi.OutputState }

func (SdbDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SdbDatabase)(nil)).Elem()
}

func (o SdbDatabaseMapOutput) ToSdbDatabaseMapOutput() SdbDatabaseMapOutput {
	return o
}

func (o SdbDatabaseMapOutput) ToSdbDatabaseMapOutputWithContext(ctx context.Context) SdbDatabaseMapOutput {
	return o
}

func (o SdbDatabaseMapOutput) MapIndex(k pulumi.StringInput) SdbDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SdbDatabase {
		return vs[0].(map[string]*SdbDatabase)[vs[1].(string)]
	}).(SdbDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SdbDatabaseInput)(nil)).Elem(), &SdbDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*SdbDatabaseArrayInput)(nil)).Elem(), SdbDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SdbDatabaseMapInput)(nil)).Elem(), SdbDatabaseMap{})
	pulumi.RegisterOutputType(SdbDatabaseOutput{})
	pulumi.RegisterOutputType(SdbDatabaseArrayOutput{})
	pulumi.RegisterOutputType(SdbDatabaseMapOutput{})
}
