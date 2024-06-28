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

// Creates and manages Scaleway Transactional Email Domains.
// For more information refer to [the API documentation](https://www.scaleway.com/en/developers/api/transactional-email).
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
//			_, err := scaleway.NewTemDomain(ctx, "main", &scaleway.TemDomainArgs{
//				AcceptTos: pulumi.Bool(true),
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
// ### Add the required records to your DNS zone
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			domainName := cfg.Require("domainName")
//			main, err := scaleway.NewTemDomain(ctx, "main", &scaleway.TemDomainArgs{
//				AcceptTos: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDomainRecord(ctx, "spf", &scaleway.DomainRecordArgs{
//				DnsZone: pulumi.String(domainName),
//				Type:    pulumi.String("TXT"),
//				Data: main.SpfConfig.ApplyT(func(spfConfig string) (string, error) {
//					return fmt.Sprintf("v=spf1 %v -all", spfConfig), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDomainRecord(ctx, "dkim", &scaleway.DomainRecordArgs{
//				DnsZone: pulumi.String(domainName),
//				Type:    pulumi.String("TXT"),
//				Data:    main.DkimConfig,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDomainRecord(ctx, "mx", &scaleway.DomainRecordArgs{
//				DnsZone: pulumi.String(domainName),
//				Type:    pulumi.String("MX"),
//				Data:    pulumi.String("."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDomainRecord(ctx, "dmarc", &scaleway.DomainRecordArgs{
//				DnsZone: pulumi.String(domainName),
//				Type:    pulumi.String("TXT"),
//				Data:    main.DmarcConfig,
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
// Domains can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/temDomain:TemDomain main fr-par/11111111-1111-1111-1111-111111111111
// ```
type TemDomain struct {
	pulumi.CustomResourceState

	// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
	// > **Important:** This attribute must be set to `true`.
	AcceptTos pulumi.BoolOutput `pulumi:"acceptTos"`
	// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The DKIM public key, as should be recorded in the DNS zone.
	DkimConfig pulumi.StringOutput `pulumi:"dkimConfig"`
	// DMARC record for the domain, as should be recorded in the DNS zone.
	DmarcConfig pulumi.StringOutput `pulumi:"dmarcConfig"`
	// DMARC name for the domain, as should be recorded in the DNS zone.
	DmarcName pulumi.StringOutput `pulumi:"dmarcName"`
	// The error message if the last check failed.
	LastError pulumi.StringOutput `pulumi:"lastError"`
	// The date and time the domain was last found to be valid (RFC 3339 format).
	LastValidAt pulumi.StringOutput `pulumi:"lastValidAt"`
	// The Scaleway's blackhole MX server to use if you do not have one.
	MxBlackhole pulumi.StringOutput `pulumi:"mxBlackhole"`
	// The domain name, must not be used in another Transactional Email Domain.
	// > **Important:** Updates to `name` will recreate the domain.
	Name pulumi.StringOutput `pulumi:"name"`
	// The date and time of the next scheduled check (RFC 3339 format).
	NextCheckAt pulumi.StringOutput `pulumi:"nextCheckAt"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region in which the domain should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The domain's reputation.
	Reputations TemDomainReputationArrayOutput `pulumi:"reputations"`
	// The date and time of the revocation of the domain (RFC 3339 format).
	RevokedAt pulumi.StringOutput `pulumi:"revokedAt"`
	// The SMTP host to use to send emails.
	SmtpHost pulumi.StringOutput `pulumi:"smtpHost"`
	// The SMTP port to use to send emails over TLS.
	SmtpPort pulumi.IntOutput `pulumi:"smtpPort"`
	// The SMTP port to use to send emails over TLS.
	SmtpPortAlternative pulumi.IntOutput `pulumi:"smtpPortAlternative"`
	// The SMTP port to use to send emails.
	SmtpPortUnsecure pulumi.IntOutput `pulumi:"smtpPortUnsecure"`
	// SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
	SmtpsAuthUser pulumi.StringOutput `pulumi:"smtpsAuthUser"`
	// The SMTPS port to use to send emails over TLS Wrapper.
	SmtpsPort pulumi.IntOutput `pulumi:"smtpsPort"`
	// The SMTPS port to use to send emails over TLS Wrapper.
	SmtpsPortAlternative pulumi.IntOutput `pulumi:"smtpsPortAlternative"`
	// The snippet of the SPF record that should be registered in the DNS zone.
	SpfConfig pulumi.StringOutput `pulumi:"spfConfig"`
	// The status of the domain's reputation.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewTemDomain registers a new resource with the given unique name, arguments, and options.
func NewTemDomain(ctx *pulumi.Context,
	name string, args *TemDomainArgs, opts ...pulumi.ResourceOption) (*TemDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceptTos == nil {
		return nil, errors.New("invalid value for required argument 'AcceptTos'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TemDomain
	err := ctx.RegisterResource("scaleway:index/temDomain:TemDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemDomain gets an existing TemDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemDomainState, opts ...pulumi.ResourceOption) (*TemDomain, error) {
	var resource TemDomain
	err := ctx.ReadResource("scaleway:index/temDomain:TemDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemDomain resources.
type temDomainState struct {
	// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
	// > **Important:** This attribute must be set to `true`.
	AcceptTos *bool `pulumi:"acceptTos"`
	// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
	CreatedAt *string `pulumi:"createdAt"`
	// The DKIM public key, as should be recorded in the DNS zone.
	DkimConfig *string `pulumi:"dkimConfig"`
	// DMARC record for the domain, as should be recorded in the DNS zone.
	DmarcConfig *string `pulumi:"dmarcConfig"`
	// DMARC name for the domain, as should be recorded in the DNS zone.
	DmarcName *string `pulumi:"dmarcName"`
	// The error message if the last check failed.
	LastError *string `pulumi:"lastError"`
	// The date and time the domain was last found to be valid (RFC 3339 format).
	LastValidAt *string `pulumi:"lastValidAt"`
	// The Scaleway's blackhole MX server to use if you do not have one.
	MxBlackhole *string `pulumi:"mxBlackhole"`
	// The domain name, must not be used in another Transactional Email Domain.
	// > **Important:** Updates to `name` will recreate the domain.
	Name *string `pulumi:"name"`
	// The date and time of the next scheduled check (RFC 3339 format).
	NextCheckAt *string `pulumi:"nextCheckAt"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the domain should be created.
	Region *string `pulumi:"region"`
	// The domain's reputation.
	Reputations []TemDomainReputation `pulumi:"reputations"`
	// The date and time of the revocation of the domain (RFC 3339 format).
	RevokedAt *string `pulumi:"revokedAt"`
	// The SMTP host to use to send emails.
	SmtpHost *string `pulumi:"smtpHost"`
	// The SMTP port to use to send emails over TLS.
	SmtpPort *int `pulumi:"smtpPort"`
	// The SMTP port to use to send emails over TLS.
	SmtpPortAlternative *int `pulumi:"smtpPortAlternative"`
	// The SMTP port to use to send emails.
	SmtpPortUnsecure *int `pulumi:"smtpPortUnsecure"`
	// SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
	SmtpsAuthUser *string `pulumi:"smtpsAuthUser"`
	// The SMTPS port to use to send emails over TLS Wrapper.
	SmtpsPort *int `pulumi:"smtpsPort"`
	// The SMTPS port to use to send emails over TLS Wrapper.
	SmtpsPortAlternative *int `pulumi:"smtpsPortAlternative"`
	// The snippet of the SPF record that should be registered in the DNS zone.
	SpfConfig *string `pulumi:"spfConfig"`
	// The status of the domain's reputation.
	Status *string `pulumi:"status"`
}

type TemDomainState struct {
	// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
	// > **Important:** This attribute must be set to `true`.
	AcceptTos pulumi.BoolPtrInput
	// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
	CreatedAt pulumi.StringPtrInput
	// The DKIM public key, as should be recorded in the DNS zone.
	DkimConfig pulumi.StringPtrInput
	// DMARC record for the domain, as should be recorded in the DNS zone.
	DmarcConfig pulumi.StringPtrInput
	// DMARC name for the domain, as should be recorded in the DNS zone.
	DmarcName pulumi.StringPtrInput
	// The error message if the last check failed.
	LastError pulumi.StringPtrInput
	// The date and time the domain was last found to be valid (RFC 3339 format).
	LastValidAt pulumi.StringPtrInput
	// The Scaleway's blackhole MX server to use if you do not have one.
	MxBlackhole pulumi.StringPtrInput
	// The domain name, must not be used in another Transactional Email Domain.
	// > **Important:** Updates to `name` will recreate the domain.
	Name pulumi.StringPtrInput
	// The date and time of the next scheduled check (RFC 3339 format).
	NextCheckAt pulumi.StringPtrInput
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the domain should be created.
	Region pulumi.StringPtrInput
	// The domain's reputation.
	Reputations TemDomainReputationArrayInput
	// The date and time of the revocation of the domain (RFC 3339 format).
	RevokedAt pulumi.StringPtrInput
	// The SMTP host to use to send emails.
	SmtpHost pulumi.StringPtrInput
	// The SMTP port to use to send emails over TLS.
	SmtpPort pulumi.IntPtrInput
	// The SMTP port to use to send emails over TLS.
	SmtpPortAlternative pulumi.IntPtrInput
	// The SMTP port to use to send emails.
	SmtpPortUnsecure pulumi.IntPtrInput
	// SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
	SmtpsAuthUser pulumi.StringPtrInput
	// The SMTPS port to use to send emails over TLS Wrapper.
	SmtpsPort pulumi.IntPtrInput
	// The SMTPS port to use to send emails over TLS Wrapper.
	SmtpsPortAlternative pulumi.IntPtrInput
	// The snippet of the SPF record that should be registered in the DNS zone.
	SpfConfig pulumi.StringPtrInput
	// The status of the domain's reputation.
	Status pulumi.StringPtrInput
}

func (TemDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*temDomainState)(nil)).Elem()
}

type temDomainArgs struct {
	// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
	// > **Important:** This attribute must be set to `true`.
	AcceptTos bool `pulumi:"acceptTos"`
	// The domain name, must not be used in another Transactional Email Domain.
	// > **Important:** Updates to `name` will recreate the domain.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the domain should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a TemDomain resource.
type TemDomainArgs struct {
	// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
	// > **Important:** This attribute must be set to `true`.
	AcceptTos pulumi.BoolInput
	// The domain name, must not be used in another Transactional Email Domain.
	// > **Important:** Updates to `name` will recreate the domain.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the domain should be created.
	Region pulumi.StringPtrInput
}

func (TemDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*temDomainArgs)(nil)).Elem()
}

type TemDomainInput interface {
	pulumi.Input

	ToTemDomainOutput() TemDomainOutput
	ToTemDomainOutputWithContext(ctx context.Context) TemDomainOutput
}

func (*TemDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**TemDomain)(nil)).Elem()
}

func (i *TemDomain) ToTemDomainOutput() TemDomainOutput {
	return i.ToTemDomainOutputWithContext(context.Background())
}

func (i *TemDomain) ToTemDomainOutputWithContext(ctx context.Context) TemDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemDomainOutput)
}

// TemDomainArrayInput is an input type that accepts TemDomainArray and TemDomainArrayOutput values.
// You can construct a concrete instance of `TemDomainArrayInput` via:
//
//	TemDomainArray{ TemDomainArgs{...} }
type TemDomainArrayInput interface {
	pulumi.Input

	ToTemDomainArrayOutput() TemDomainArrayOutput
	ToTemDomainArrayOutputWithContext(context.Context) TemDomainArrayOutput
}

type TemDomainArray []TemDomainInput

func (TemDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemDomain)(nil)).Elem()
}

func (i TemDomainArray) ToTemDomainArrayOutput() TemDomainArrayOutput {
	return i.ToTemDomainArrayOutputWithContext(context.Background())
}

func (i TemDomainArray) ToTemDomainArrayOutputWithContext(ctx context.Context) TemDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemDomainArrayOutput)
}

// TemDomainMapInput is an input type that accepts TemDomainMap and TemDomainMapOutput values.
// You can construct a concrete instance of `TemDomainMapInput` via:
//
//	TemDomainMap{ "key": TemDomainArgs{...} }
type TemDomainMapInput interface {
	pulumi.Input

	ToTemDomainMapOutput() TemDomainMapOutput
	ToTemDomainMapOutputWithContext(context.Context) TemDomainMapOutput
}

type TemDomainMap map[string]TemDomainInput

func (TemDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemDomain)(nil)).Elem()
}

func (i TemDomainMap) ToTemDomainMapOutput() TemDomainMapOutput {
	return i.ToTemDomainMapOutputWithContext(context.Background())
}

func (i TemDomainMap) ToTemDomainMapOutputWithContext(ctx context.Context) TemDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemDomainMapOutput)
}

type TemDomainOutput struct{ *pulumi.OutputState }

func (TemDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemDomain)(nil)).Elem()
}

func (o TemDomainOutput) ToTemDomainOutput() TemDomainOutput {
	return o
}

func (o TemDomainOutput) ToTemDomainOutputWithContext(ctx context.Context) TemDomainOutput {
	return o
}

// Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
// > **Important:** This attribute must be set to `true`.
func (o TemDomainOutput) AcceptTos() pulumi.BoolOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.BoolOutput { return v.AcceptTos }).(pulumi.BoolOutput)
}

// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
func (o TemDomainOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The DKIM public key, as should be recorded in the DNS zone.
func (o TemDomainOutput) DkimConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.DkimConfig }).(pulumi.StringOutput)
}

// DMARC record for the domain, as should be recorded in the DNS zone.
func (o TemDomainOutput) DmarcConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.DmarcConfig }).(pulumi.StringOutput)
}

// DMARC name for the domain, as should be recorded in the DNS zone.
func (o TemDomainOutput) DmarcName() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.DmarcName }).(pulumi.StringOutput)
}

// The error message if the last check failed.
func (o TemDomainOutput) LastError() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.LastError }).(pulumi.StringOutput)
}

// The date and time the domain was last found to be valid (RFC 3339 format).
func (o TemDomainOutput) LastValidAt() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.LastValidAt }).(pulumi.StringOutput)
}

// The Scaleway's blackhole MX server to use if you do not have one.
func (o TemDomainOutput) MxBlackhole() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.MxBlackhole }).(pulumi.StringOutput)
}

// The domain name, must not be used in another Transactional Email Domain.
// > **Important:** Updates to `name` will recreate the domain.
func (o TemDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The date and time of the next scheduled check (RFC 3339 format).
func (o TemDomainOutput) NextCheckAt() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.NextCheckAt }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the domain is associated with.
func (o TemDomainOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region in which the domain should be created.
func (o TemDomainOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The domain's reputation.
func (o TemDomainOutput) Reputations() TemDomainReputationArrayOutput {
	return o.ApplyT(func(v *TemDomain) TemDomainReputationArrayOutput { return v.Reputations }).(TemDomainReputationArrayOutput)
}

// The date and time of the revocation of the domain (RFC 3339 format).
func (o TemDomainOutput) RevokedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.RevokedAt }).(pulumi.StringOutput)
}

// The SMTP host to use to send emails.
func (o TemDomainOutput) SmtpHost() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.SmtpHost }).(pulumi.StringOutput)
}

// The SMTP port to use to send emails over TLS.
func (o TemDomainOutput) SmtpPort() pulumi.IntOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.IntOutput { return v.SmtpPort }).(pulumi.IntOutput)
}

// The SMTP port to use to send emails over TLS.
func (o TemDomainOutput) SmtpPortAlternative() pulumi.IntOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.IntOutput { return v.SmtpPortAlternative }).(pulumi.IntOutput)
}

// The SMTP port to use to send emails.
func (o TemDomainOutput) SmtpPortUnsecure() pulumi.IntOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.IntOutput { return v.SmtpPortUnsecure }).(pulumi.IntOutput)
}

// SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
func (o TemDomainOutput) SmtpsAuthUser() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.SmtpsAuthUser }).(pulumi.StringOutput)
}

// The SMTPS port to use to send emails over TLS Wrapper.
func (o TemDomainOutput) SmtpsPort() pulumi.IntOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.IntOutput { return v.SmtpsPort }).(pulumi.IntOutput)
}

// The SMTPS port to use to send emails over TLS Wrapper.
func (o TemDomainOutput) SmtpsPortAlternative() pulumi.IntOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.IntOutput { return v.SmtpsPortAlternative }).(pulumi.IntOutput)
}

// The snippet of the SPF record that should be registered in the DNS zone.
func (o TemDomainOutput) SpfConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.SpfConfig }).(pulumi.StringOutput)
}

// The status of the domain's reputation.
func (o TemDomainOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TemDomain) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type TemDomainArrayOutput struct{ *pulumi.OutputState }

func (TemDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemDomain)(nil)).Elem()
}

func (o TemDomainArrayOutput) ToTemDomainArrayOutput() TemDomainArrayOutput {
	return o
}

func (o TemDomainArrayOutput) ToTemDomainArrayOutputWithContext(ctx context.Context) TemDomainArrayOutput {
	return o
}

func (o TemDomainArrayOutput) Index(i pulumi.IntInput) TemDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TemDomain {
		return vs[0].([]*TemDomain)[vs[1].(int)]
	}).(TemDomainOutput)
}

type TemDomainMapOutput struct{ *pulumi.OutputState }

func (TemDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemDomain)(nil)).Elem()
}

func (o TemDomainMapOutput) ToTemDomainMapOutput() TemDomainMapOutput {
	return o
}

func (o TemDomainMapOutput) ToTemDomainMapOutputWithContext(ctx context.Context) TemDomainMapOutput {
	return o
}

func (o TemDomainMapOutput) MapIndex(k pulumi.StringInput) TemDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TemDomain {
		return vs[0].(map[string]*TemDomain)[vs[1].(string)]
	}).(TemDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemDomainInput)(nil)).Elem(), &TemDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemDomainArrayInput)(nil)).Elem(), TemDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemDomainMapInput)(nil)).Elem(), TemDomainMap{})
	pulumi.RegisterOutputType(TemDomainOutput{})
	pulumi.RegisterOutputType(TemDomainArrayOutput{})
	pulumi.RegisterOutputType(TemDomainMapOutput{})
}
