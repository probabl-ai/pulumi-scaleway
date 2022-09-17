// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the Bucket.
// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewObjectBucket(ctx, "main", &scaleway.ObjectBucketArgs{
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupObjectBucket(ctx, &GetObjectBucketArgs{
//				Name: pulumi.StringRef("bucket.test.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupObjectBucket(ctx *pulumi.Context, args *LookupObjectBucketArgs, opts ...pulumi.InvokeOption) (*LookupObjectBucketResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupObjectBucketResult
	err := ctx.Invoke("scaleway:index/getObjectBucket:getObjectBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjectBucket.
type LookupObjectBucketArgs struct {
	// The bucket name.
	Name *string `pulumi:"name"`
	// `region`) The region in which the Object Storage exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getObjectBucket.
type LookupObjectBucketResult struct {
	Acl       string                    `pulumi:"acl"`
	CorsRules []GetObjectBucketCorsRule `pulumi:"corsRules"`
	// The endpoint URL of the bucket
	Endpoint     string `pulumi:"endpoint"`
	ForceDestroy bool   `pulumi:"forceDestroy"`
	// The provider-assigned unique ID for this managed resource.
	Id             string                         `pulumi:"id"`
	LifecycleRules []GetObjectBucketLifecycleRule `pulumi:"lifecycleRules"`
	Name           *string                        `pulumi:"name"`
	Region         *string                        `pulumi:"region"`
	Tags           map[string]string              `pulumi:"tags"`
	Versionings    []GetObjectBucketVersioning    `pulumi:"versionings"`
}

func LookupObjectBucketOutput(ctx *pulumi.Context, args LookupObjectBucketOutputArgs, opts ...pulumi.InvokeOption) LookupObjectBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectBucketResult, error) {
			args := v.(LookupObjectBucketArgs)
			r, err := LookupObjectBucket(ctx, &args, opts...)
			var s LookupObjectBucketResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObjectBucketResultOutput)
}

// A collection of arguments for invoking getObjectBucket.
type LookupObjectBucketOutputArgs struct {
	// The bucket name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `region`) The region in which the Object Storage exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupObjectBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectBucketArgs)(nil)).Elem()
}

// A collection of values returned by getObjectBucket.
type LookupObjectBucketResultOutput struct{ *pulumi.OutputState }

func (LookupObjectBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectBucketResult)(nil)).Elem()
}

func (o LookupObjectBucketResultOutput) ToLookupObjectBucketResultOutput() LookupObjectBucketResultOutput {
	return o
}

func (o LookupObjectBucketResultOutput) ToLookupObjectBucketResultOutputWithContext(ctx context.Context) LookupObjectBucketResultOutput {
	return o
}

func (o LookupObjectBucketResultOutput) Acl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) string { return v.Acl }).(pulumi.StringOutput)
}

func (o LookupObjectBucketResultOutput) CorsRules() GetObjectBucketCorsRuleArrayOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) []GetObjectBucketCorsRule { return v.CorsRules }).(GetObjectBucketCorsRuleArrayOutput)
}

// The endpoint URL of the bucket
func (o LookupObjectBucketResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupObjectBucketResultOutput) ForceDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) bool { return v.ForceDestroy }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupObjectBucketResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupObjectBucketResultOutput) LifecycleRules() GetObjectBucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) []GetObjectBucketLifecycleRule { return v.LifecycleRules }).(GetObjectBucketLifecycleRuleArrayOutput)
}

func (o LookupObjectBucketResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupObjectBucketResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupObjectBucketResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupObjectBucketResultOutput) Versionings() GetObjectBucketVersioningArrayOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) []GetObjectBucketVersioning { return v.Versionings }).(GetObjectBucketVersioningArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectBucketResultOutput{})
}
