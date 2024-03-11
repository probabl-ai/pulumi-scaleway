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

// Creates and manages Scaleway Secret Versions.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/#secret-versions-079501).
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
//			main, err := scaleway.NewSecret(ctx, "main", &scaleway.SecretArgs{
//				Description: pulumi.String("barr"),
//				Tags: pulumi.StringArray{
//					pulumi.String("foo"),
//					pulumi.String("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewSecretVersion(ctx, "v1", &scaleway.SecretVersionArgs{
//				Description: pulumi.String("version1"),
//				SecretId:    main.ID(),
//				Data:        pulumi.String("my_new_secret"),
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
// ## Data
//
// Note: The `data` should be a base64 encoded string when sent from the API. **It is already handled by the provider so you don't need to encode it yourself.**
//
// Updating `data` will force creating a new the secret version.
//
// Be aware that this is a sensitive attribute. For more information, see Sensitive Data in State.
//
// > **Important:**  This property is sensitive and will not be displayed in the plan.
//
// ## Import
//
// The Secret Version can be imported using the `{region}/{id}/{revision}`, e.g.
//
// ~> **Important:** Be aware if you import with revision `latest` you will overwrite the version you used before.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/secretVersion:SecretVersion main fr-par/11111111-1111-1111-1111-111111111111/2
// ```
type SecretVersion struct {
	pulumi.CustomResourceState

	// Date and time of secret version's creation (RFC 3339 format).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
	Data pulumi.StringOutput `pulumi:"data"`
	// Description of the secret version (e.g. `my-new-description`).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// `region`) The region
	// in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// The revision for this Secret Version.
	Revision pulumi.StringOutput `pulumi:"revision"`
	// The Secret ID associated wit the secret version.
	SecretId pulumi.StringOutput `pulumi:"secretId"`
	// The status of the Secret Version.
	Status pulumi.StringOutput `pulumi:"status"`
	// Date and time of secret version's last update (RFC 3339 format).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewSecretVersion registers a new resource with the given unique name, arguments, and options.
func NewSecretVersion(ctx *pulumi.Context,
	name string, args *SecretVersionArgs, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.SecretId == nil {
		return nil, errors.New("invalid value for required argument 'SecretId'")
	}
	if args.Data != nil {
		args.Data = pulumi.ToSecret(args.Data).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"data",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretVersion
	err := ctx.RegisterResource("scaleway:index/secretVersion:SecretVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretVersion gets an existing SecretVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretVersionState, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	var resource SecretVersion
	err := ctx.ReadResource("scaleway:index/secretVersion:SecretVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretVersion resources.
type secretVersionState struct {
	// Date and time of secret version's creation (RFC 3339 format).
	CreatedAt *string `pulumi:"createdAt"`
	// The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
	Data *string `pulumi:"data"`
	// Description of the secret version (e.g. `my-new-description`).
	Description *string `pulumi:"description"`
	// `region`) The region
	// in which the resource exists.
	Region *string `pulumi:"region"`
	// The revision for this Secret Version.
	Revision *string `pulumi:"revision"`
	// The Secret ID associated wit the secret version.
	SecretId *string `pulumi:"secretId"`
	// The status of the Secret Version.
	Status *string `pulumi:"status"`
	// Date and time of secret version's last update (RFC 3339 format).
	UpdatedAt *string `pulumi:"updatedAt"`
}

type SecretVersionState struct {
	// Date and time of secret version's creation (RFC 3339 format).
	CreatedAt pulumi.StringPtrInput
	// The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
	Data pulumi.StringPtrInput
	// Description of the secret version (e.g. `my-new-description`).
	Description pulumi.StringPtrInput
	// `region`) The region
	// in which the resource exists.
	Region pulumi.StringPtrInput
	// The revision for this Secret Version.
	Revision pulumi.StringPtrInput
	// The Secret ID associated wit the secret version.
	SecretId pulumi.StringPtrInput
	// The status of the Secret Version.
	Status pulumi.StringPtrInput
	// Date and time of secret version's last update (RFC 3339 format).
	UpdatedAt pulumi.StringPtrInput
}

func (SecretVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretVersionState)(nil)).Elem()
}

type secretVersionArgs struct {
	// The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
	Data string `pulumi:"data"`
	// Description of the secret version (e.g. `my-new-description`).
	Description *string `pulumi:"description"`
	// `region`) The region
	// in which the resource exists.
	Region *string `pulumi:"region"`
	// The Secret ID associated wit the secret version.
	SecretId string `pulumi:"secretId"`
}

// The set of arguments for constructing a SecretVersion resource.
type SecretVersionArgs struct {
	// The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
	Data pulumi.StringInput
	// Description of the secret version (e.g. `my-new-description`).
	Description pulumi.StringPtrInput
	// `region`) The region
	// in which the resource exists.
	Region pulumi.StringPtrInput
	// The Secret ID associated wit the secret version.
	SecretId pulumi.StringInput
}

func (SecretVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretVersionArgs)(nil)).Elem()
}

type SecretVersionInput interface {
	pulumi.Input

	ToSecretVersionOutput() SecretVersionOutput
	ToSecretVersionOutputWithContext(ctx context.Context) SecretVersionOutput
}

func (*SecretVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretVersion)(nil)).Elem()
}

func (i *SecretVersion) ToSecretVersionOutput() SecretVersionOutput {
	return i.ToSecretVersionOutputWithContext(context.Background())
}

func (i *SecretVersion) ToSecretVersionOutputWithContext(ctx context.Context) SecretVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionOutput)
}

// SecretVersionArrayInput is an input type that accepts SecretVersionArray and SecretVersionArrayOutput values.
// You can construct a concrete instance of `SecretVersionArrayInput` via:
//
//	SecretVersionArray{ SecretVersionArgs{...} }
type SecretVersionArrayInput interface {
	pulumi.Input

	ToSecretVersionArrayOutput() SecretVersionArrayOutput
	ToSecretVersionArrayOutputWithContext(context.Context) SecretVersionArrayOutput
}

type SecretVersionArray []SecretVersionInput

func (SecretVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretVersion)(nil)).Elem()
}

func (i SecretVersionArray) ToSecretVersionArrayOutput() SecretVersionArrayOutput {
	return i.ToSecretVersionArrayOutputWithContext(context.Background())
}

func (i SecretVersionArray) ToSecretVersionArrayOutputWithContext(ctx context.Context) SecretVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionArrayOutput)
}

// SecretVersionMapInput is an input type that accepts SecretVersionMap and SecretVersionMapOutput values.
// You can construct a concrete instance of `SecretVersionMapInput` via:
//
//	SecretVersionMap{ "key": SecretVersionArgs{...} }
type SecretVersionMapInput interface {
	pulumi.Input

	ToSecretVersionMapOutput() SecretVersionMapOutput
	ToSecretVersionMapOutputWithContext(context.Context) SecretVersionMapOutput
}

type SecretVersionMap map[string]SecretVersionInput

func (SecretVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretVersion)(nil)).Elem()
}

func (i SecretVersionMap) ToSecretVersionMapOutput() SecretVersionMapOutput {
	return i.ToSecretVersionMapOutputWithContext(context.Background())
}

func (i SecretVersionMap) ToSecretVersionMapOutputWithContext(ctx context.Context) SecretVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionMapOutput)
}

type SecretVersionOutput struct{ *pulumi.OutputState }

func (SecretVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretVersion)(nil)).Elem()
}

func (o SecretVersionOutput) ToSecretVersionOutput() SecretVersionOutput {
	return o
}

func (o SecretVersionOutput) ToSecretVersionOutputWithContext(ctx context.Context) SecretVersionOutput {
	return o
}

// Date and time of secret version's creation (RFC 3339 format).
func (o SecretVersionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
func (o SecretVersionOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

// Description of the secret version (e.g. `my-new-description`).
func (o SecretVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// `region`) The region
// in which the resource exists.
func (o SecretVersionOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The revision for this Secret Version.
func (o SecretVersionOutput) Revision() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.Revision }).(pulumi.StringOutput)
}

// The Secret ID associated wit the secret version.
func (o SecretVersionOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.SecretId }).(pulumi.StringOutput)
}

// The status of the Secret Version.
func (o SecretVersionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Date and time of secret version's last update (RFC 3339 format).
func (o SecretVersionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type SecretVersionArrayOutput struct{ *pulumi.OutputState }

func (SecretVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretVersion)(nil)).Elem()
}

func (o SecretVersionArrayOutput) ToSecretVersionArrayOutput() SecretVersionArrayOutput {
	return o
}

func (o SecretVersionArrayOutput) ToSecretVersionArrayOutputWithContext(ctx context.Context) SecretVersionArrayOutput {
	return o
}

func (o SecretVersionArrayOutput) Index(i pulumi.IntInput) SecretVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretVersion {
		return vs[0].([]*SecretVersion)[vs[1].(int)]
	}).(SecretVersionOutput)
}

type SecretVersionMapOutput struct{ *pulumi.OutputState }

func (SecretVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretVersion)(nil)).Elem()
}

func (o SecretVersionMapOutput) ToSecretVersionMapOutput() SecretVersionMapOutput {
	return o
}

func (o SecretVersionMapOutput) ToSecretVersionMapOutputWithContext(ctx context.Context) SecretVersionMapOutput {
	return o
}

func (o SecretVersionMapOutput) MapIndex(k pulumi.StringInput) SecretVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretVersion {
		return vs[0].(map[string]*SecretVersion)[vs[1].(string)]
	}).(SecretVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionInput)(nil)).Elem(), &SecretVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionArrayInput)(nil)).Elem(), SecretVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionMapInput)(nil)).Elem(), SecretVersionMap{})
	pulumi.RegisterOutputType(SecretVersionOutput{})
	pulumi.RegisterOutputType(SecretVersionArrayOutput{})
	pulumi.RegisterOutputType(SecretVersionMapOutput{})
}
