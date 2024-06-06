// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Secrets.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/secret-manager/).
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
//			_, err := scaleway.NewSecret(ctx, "main", &scaleway.SecretArgs{
//				Description: pulumi.String("barr"),
//				Tags: pulumi.StringArray{
//					pulumi.String("foo"),
//					pulumi.String("terraform"),
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
// The Secret can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/secret:Secret main fr-par/11111111-1111-1111-1111-111111111111
// ```
type Secret struct {
	pulumi.CustomResourceState

	// Date and time of secret's creation (RFC 3339 format).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the secret (e.g. `my-new-description`).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the secret (e.g. `my-secret`).
	Name pulumi.StringOutput `pulumi:"name"`
	// Path of the secret, defaults to `/`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The project ID containing is the secret.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region
	// in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// The status of the Secret.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags of the secret (e.g. `["tag", "secret"]`).
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Date and time of secret's last update (RFC 3339 format).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The number of versions for this Secret.
	VersionCount pulumi.IntOutput `pulumi:"versionCount"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		args = &SecretArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Secret
	err := ctx.RegisterResource("scaleway:index/secret:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("scaleway:index/secret:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
	// Date and time of secret's creation (RFC 3339 format).
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the secret (e.g. `my-new-description`).
	Description *string `pulumi:"description"`
	// Name of the secret (e.g. `my-secret`).
	Name *string `pulumi:"name"`
	// Path of the secret, defaults to `/`.
	Path *string `pulumi:"path"`
	// The project ID containing is the secret.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region
	// in which the resource exists.
	Region *string `pulumi:"region"`
	// The status of the Secret.
	Status *string `pulumi:"status"`
	// Tags of the secret (e.g. `["tag", "secret"]`).
	Tags []string `pulumi:"tags"`
	// Date and time of secret's last update (RFC 3339 format).
	UpdatedAt *string `pulumi:"updatedAt"`
	// The number of versions for this Secret.
	VersionCount *int `pulumi:"versionCount"`
}

type SecretState struct {
	// Date and time of secret's creation (RFC 3339 format).
	CreatedAt pulumi.StringPtrInput
	// Description of the secret (e.g. `my-new-description`).
	Description pulumi.StringPtrInput
	// Name of the secret (e.g. `my-secret`).
	Name pulumi.StringPtrInput
	// Path of the secret, defaults to `/`.
	Path pulumi.StringPtrInput
	// The project ID containing is the secret.
	ProjectId pulumi.StringPtrInput
	// `region`) The region
	// in which the resource exists.
	Region pulumi.StringPtrInput
	// The status of the Secret.
	Status pulumi.StringPtrInput
	// Tags of the secret (e.g. `["tag", "secret"]`).
	Tags pulumi.StringArrayInput
	// Date and time of secret's last update (RFC 3339 format).
	UpdatedAt pulumi.StringPtrInput
	// The number of versions for this Secret.
	VersionCount pulumi.IntPtrInput
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// Description of the secret (e.g. `my-new-description`).
	Description *string `pulumi:"description"`
	// Name of the secret (e.g. `my-secret`).
	Name *string `pulumi:"name"`
	// Path of the secret, defaults to `/`.
	Path *string `pulumi:"path"`
	// The project ID containing is the secret.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region
	// in which the resource exists.
	Region *string `pulumi:"region"`
	// Tags of the secret (e.g. `["tag", "secret"]`).
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// Description of the secret (e.g. `my-new-description`).
	Description pulumi.StringPtrInput
	// Name of the secret (e.g. `my-secret`).
	Name pulumi.StringPtrInput
	// Path of the secret, defaults to `/`.
	Path pulumi.StringPtrInput
	// The project ID containing is the secret.
	ProjectId pulumi.StringPtrInput
	// `region`) The region
	// in which the resource exists.
	Region pulumi.StringPtrInput
	// Tags of the secret (e.g. `["tag", "secret"]`).
	Tags pulumi.StringArrayInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(ctx context.Context) SecretOutput
}

func (*Secret) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

// SecretArrayInput is an input type that accepts SecretArray and SecretArrayOutput values.
// You can construct a concrete instance of `SecretArrayInput` via:
//
//	SecretArray{ SecretArgs{...} }
type SecretArrayInput interface {
	pulumi.Input

	ToSecretArrayOutput() SecretArrayOutput
	ToSecretArrayOutputWithContext(context.Context) SecretArrayOutput
}

type SecretArray []SecretInput

func (SecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Secret)(nil)).Elem()
}

func (i SecretArray) ToSecretArrayOutput() SecretArrayOutput {
	return i.ToSecretArrayOutputWithContext(context.Background())
}

func (i SecretArray) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretArrayOutput)
}

// SecretMapInput is an input type that accepts SecretMap and SecretMapOutput values.
// You can construct a concrete instance of `SecretMapInput` via:
//
//	SecretMap{ "key": SecretArgs{...} }
type SecretMapInput interface {
	pulumi.Input

	ToSecretMapOutput() SecretMapOutput
	ToSecretMapOutputWithContext(context.Context) SecretMapOutput
}

type SecretMap map[string]SecretInput

func (SecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Secret)(nil)).Elem()
}

func (i SecretMap) ToSecretMapOutput() SecretMapOutput {
	return i.ToSecretMapOutputWithContext(context.Background())
}

func (i SecretMap) ToSecretMapOutputWithContext(ctx context.Context) SecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretMapOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

// Date and time of secret's creation (RFC 3339 format).
func (o SecretOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the secret (e.g. `my-new-description`).
func (o SecretOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the secret (e.g. `my-secret`).
func (o SecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path of the secret, defaults to `/`.
func (o SecretOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The project ID containing is the secret.
func (o SecretOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region
// in which the resource exists.
func (o SecretOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The status of the Secret.
func (o SecretOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags of the secret (e.g. `["tag", "secret"]`).
func (o SecretOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Date and time of secret's last update (RFC 3339 format).
func (o SecretOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The number of versions for this Secret.
func (o SecretOutput) VersionCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Secret) pulumi.IntOutput { return v.VersionCount }).(pulumi.IntOutput)
}

type SecretArrayOutput struct{ *pulumi.OutputState }

func (SecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Secret)(nil)).Elem()
}

func (o SecretArrayOutput) ToSecretArrayOutput() SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) Index(i pulumi.IntInput) SecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Secret {
		return vs[0].([]*Secret)[vs[1].(int)]
	}).(SecretOutput)
}

type SecretMapOutput struct{ *pulumi.OutputState }

func (SecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Secret)(nil)).Elem()
}

func (o SecretMapOutput) ToSecretMapOutput() SecretMapOutput {
	return o
}

func (o SecretMapOutput) ToSecretMapOutputWithContext(ctx context.Context) SecretMapOutput {
	return o
}

func (o SecretMapOutput) MapIndex(k pulumi.StringInput) SecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Secret {
		return vs[0].(map[string]*Secret)[vs[1].(string)]
	}).(SecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretInput)(nil)).Elem(), &Secret{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretArrayInput)(nil)).Elem(), SecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretMapInput)(nil)).Elem(), SecretMap{})
	pulumi.RegisterOutputType(SecretOutput{})
	pulumi.RegisterOutputType(SecretArrayOutput{})
	pulumi.RegisterOutputType(SecretMapOutput{})
}
