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

// Manage Scaleway Messaging and queuing SNS topics.
// For further information, see
// our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/how-to/create-manage-topics/).
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
//			mainMnqSns, err := scaleway.NewMnqSns(ctx, "mainMnqSns", nil)
//			if err != nil {
//				return err
//			}
//			mainMnqSnsCredentials, err := scaleway.NewMnqSnsCredentials(ctx, "mainMnqSnsCredentials", &scaleway.MnqSnsCredentialsArgs{
//				ProjectId: mainMnqSns.ProjectId,
//				Permissions: &scaleway.MnqSnsCredentialsPermissionsArgs{
//					CanManage: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewMnqSnsTopic(ctx, "topic", &scaleway.MnqSnsTopicArgs{
//				ProjectId: mainMnqSns.ProjectId,
//				AccessKey: mainMnqSnsCredentials.AccessKey,
//				SecretKey: mainMnqSnsCredentials.SecretKey,
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
// SNS topics can be imported using `{region}/{project-id}/{topic-name}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/mnqSnsTopic:MnqSnsTopic main fr-par/11111111111111111111111111111111/my-topic
// ```
type MnqSnsTopic struct {
	pulumi.CustomResourceState

	// The access key of the SNS credentials.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// The ARN of the topic
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication pulumi.BoolOutput `pulumi:"contentBasedDeduplication"`
	// Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
	FifoTopic pulumi.BoolOutput `pulumi:"fifoTopic"`
	// The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Owner of the SNS topic, should have format 'project-${project_id}'
	Owner pulumi.StringOutput `pulumi:"owner"`
	// `projectId`) The ID of the Project in which SNS is enabled.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region
	// in which SNS is enabled.
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrOutput `pulumi:"snsEndpoint"`
}

// NewMnqSnsTopic registers a new resource with the given unique name, arguments, and options.
func NewMnqSnsTopic(ctx *pulumi.Context,
	name string, args *MnqSnsTopicArgs, opts ...pulumi.ResourceOption) (*MnqSnsTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	if args.AccessKey != nil {
		args.AccessKey = pulumi.ToSecret(args.AccessKey).(pulumi.StringInput)
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKey",
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqSnsTopic
	err := ctx.RegisterResource("scaleway:index/mnqSnsTopic:MnqSnsTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqSnsTopic gets an existing MnqSnsTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqSnsTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqSnsTopicState, opts ...pulumi.ResourceOption) (*MnqSnsTopic, error) {
	var resource MnqSnsTopic
	err := ctx.ReadResource("scaleway:index/mnqSnsTopic:MnqSnsTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqSnsTopic resources.
type mnqSnsTopicState struct {
	// The access key of the SNS credentials.
	AccessKey *string `pulumi:"accessKey"`
	// The ARN of the topic
	Arn *string `pulumi:"arn"`
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
	FifoTopic *bool `pulumi:"fifoTopic"`
	// The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Owner of the SNS topic, should have format 'project-${project_id}'
	Owner *string `pulumi:"owner"`
	// `projectId`) The ID of the Project in which SNS is enabled.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which SNS is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey *string `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint *string `pulumi:"snsEndpoint"`
}

type MnqSnsTopicState struct {
	// The access key of the SNS credentials.
	AccessKey pulumi.StringPtrInput
	// The ARN of the topic
	Arn pulumi.StringPtrInput
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
	FifoTopic pulumi.BoolPtrInput
	// The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Owner of the SNS topic, should have format 'project-${project_id}'
	Owner pulumi.StringPtrInput
	// `projectId`) The ID of the Project in which SNS is enabled.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which SNS is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringPtrInput
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrInput
}

func (MnqSnsTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSnsTopicState)(nil)).Elem()
}

type mnqSnsTopicArgs struct {
	// The access key of the SNS credentials.
	AccessKey string `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
	FifoTopic *bool `pulumi:"fifoTopic"`
	// The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// `projectId`) The ID of the Project in which SNS is enabled.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which SNS is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey string `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint *string `pulumi:"snsEndpoint"`
}

// The set of arguments for constructing a MnqSnsTopic resource.
type MnqSnsTopicArgs struct {
	// The access key of the SNS credentials.
	AccessKey pulumi.StringInput
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
	FifoTopic pulumi.BoolPtrInput
	// The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// `projectId`) The ID of the Project in which SNS is enabled.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which SNS is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringInput
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrInput
}

func (MnqSnsTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqSnsTopicArgs)(nil)).Elem()
}

type MnqSnsTopicInput interface {
	pulumi.Input

	ToMnqSnsTopicOutput() MnqSnsTopicOutput
	ToMnqSnsTopicOutputWithContext(ctx context.Context) MnqSnsTopicOutput
}

func (*MnqSnsTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSnsTopic)(nil)).Elem()
}

func (i *MnqSnsTopic) ToMnqSnsTopicOutput() MnqSnsTopicOutput {
	return i.ToMnqSnsTopicOutputWithContext(context.Background())
}

func (i *MnqSnsTopic) ToMnqSnsTopicOutputWithContext(ctx context.Context) MnqSnsTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSnsTopicOutput)
}

// MnqSnsTopicArrayInput is an input type that accepts MnqSnsTopicArray and MnqSnsTopicArrayOutput values.
// You can construct a concrete instance of `MnqSnsTopicArrayInput` via:
//
//	MnqSnsTopicArray{ MnqSnsTopicArgs{...} }
type MnqSnsTopicArrayInput interface {
	pulumi.Input

	ToMnqSnsTopicArrayOutput() MnqSnsTopicArrayOutput
	ToMnqSnsTopicArrayOutputWithContext(context.Context) MnqSnsTopicArrayOutput
}

type MnqSnsTopicArray []MnqSnsTopicInput

func (MnqSnsTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSnsTopic)(nil)).Elem()
}

func (i MnqSnsTopicArray) ToMnqSnsTopicArrayOutput() MnqSnsTopicArrayOutput {
	return i.ToMnqSnsTopicArrayOutputWithContext(context.Background())
}

func (i MnqSnsTopicArray) ToMnqSnsTopicArrayOutputWithContext(ctx context.Context) MnqSnsTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSnsTopicArrayOutput)
}

// MnqSnsTopicMapInput is an input type that accepts MnqSnsTopicMap and MnqSnsTopicMapOutput values.
// You can construct a concrete instance of `MnqSnsTopicMapInput` via:
//
//	MnqSnsTopicMap{ "key": MnqSnsTopicArgs{...} }
type MnqSnsTopicMapInput interface {
	pulumi.Input

	ToMnqSnsTopicMapOutput() MnqSnsTopicMapOutput
	ToMnqSnsTopicMapOutputWithContext(context.Context) MnqSnsTopicMapOutput
}

type MnqSnsTopicMap map[string]MnqSnsTopicInput

func (MnqSnsTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSnsTopic)(nil)).Elem()
}

func (i MnqSnsTopicMap) ToMnqSnsTopicMapOutput() MnqSnsTopicMapOutput {
	return i.ToMnqSnsTopicMapOutputWithContext(context.Background())
}

func (i MnqSnsTopicMap) ToMnqSnsTopicMapOutputWithContext(ctx context.Context) MnqSnsTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqSnsTopicMapOutput)
}

type MnqSnsTopicOutput struct{ *pulumi.OutputState }

func (MnqSnsTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqSnsTopic)(nil)).Elem()
}

func (o MnqSnsTopicOutput) ToMnqSnsTopicOutput() MnqSnsTopicOutput {
	return o
}

func (o MnqSnsTopicOutput) ToMnqSnsTopicOutputWithContext(ctx context.Context) MnqSnsTopicOutput {
	return o
}

// The access key of the SNS credentials.
func (o MnqSnsTopicOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// The ARN of the topic
func (o MnqSnsTopicOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies whether to enable content-based deduplication.
func (o MnqSnsTopicOutput) ContentBasedDeduplication() pulumi.BoolOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.BoolOutput { return v.ContentBasedDeduplication }).(pulumi.BoolOutput)
}

// Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
func (o MnqSnsTopicOutput) FifoTopic() pulumi.BoolOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.BoolOutput { return v.FifoTopic }).(pulumi.BoolOutput)
}

// The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
func (o MnqSnsTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o MnqSnsTopicOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// Owner of the SNS topic, should have format 'project-${project_id}'
func (o MnqSnsTopicOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// `projectId`) The ID of the Project in which SNS is enabled.
func (o MnqSnsTopicOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region
// in which SNS is enabled.
func (o MnqSnsTopicOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The secret key of the SNS credentials.
func (o MnqSnsTopicOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
func (o MnqSnsTopicOutput) SnsEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MnqSnsTopic) pulumi.StringPtrOutput { return v.SnsEndpoint }).(pulumi.StringPtrOutput)
}

type MnqSnsTopicArrayOutput struct{ *pulumi.OutputState }

func (MnqSnsTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqSnsTopic)(nil)).Elem()
}

func (o MnqSnsTopicArrayOutput) ToMnqSnsTopicArrayOutput() MnqSnsTopicArrayOutput {
	return o
}

func (o MnqSnsTopicArrayOutput) ToMnqSnsTopicArrayOutputWithContext(ctx context.Context) MnqSnsTopicArrayOutput {
	return o
}

func (o MnqSnsTopicArrayOutput) Index(i pulumi.IntInput) MnqSnsTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqSnsTopic {
		return vs[0].([]*MnqSnsTopic)[vs[1].(int)]
	}).(MnqSnsTopicOutput)
}

type MnqSnsTopicMapOutput struct{ *pulumi.OutputState }

func (MnqSnsTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqSnsTopic)(nil)).Elem()
}

func (o MnqSnsTopicMapOutput) ToMnqSnsTopicMapOutput() MnqSnsTopicMapOutput {
	return o
}

func (o MnqSnsTopicMapOutput) ToMnqSnsTopicMapOutputWithContext(ctx context.Context) MnqSnsTopicMapOutput {
	return o
}

func (o MnqSnsTopicMapOutput) MapIndex(k pulumi.StringInput) MnqSnsTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqSnsTopic {
		return vs[0].(map[string]*MnqSnsTopic)[vs[1].(string)]
	}).(MnqSnsTopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSnsTopicInput)(nil)).Elem(), &MnqSnsTopic{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSnsTopicArrayInput)(nil)).Elem(), MnqSnsTopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqSnsTopicMapInput)(nil)).Elem(), MnqSnsTopicMap{})
	pulumi.RegisterOutputType(MnqSnsTopicOutput{})
	pulumi.RegisterOutputType(MnqSnsTopicArrayOutput{})
	pulumi.RegisterOutputType(MnqSnsTopicMapOutput{})
}
