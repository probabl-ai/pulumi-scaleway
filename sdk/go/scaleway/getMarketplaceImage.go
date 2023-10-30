// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets local image ID of an image from its label name.
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
//			_, err := scaleway.GetMarketplaceImage(ctx, &scaleway.GetMarketplaceImageArgs{
//				Label: "ubuntu_jammy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMarketplaceImage(ctx *pulumi.Context, args *GetMarketplaceImageArgs, opts ...pulumi.InvokeOption) (*GetMarketplaceImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMarketplaceImageResult
	err := ctx.Invoke("scaleway:index/getMarketplaceImage:getMarketplaceImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMarketplaceImage.
type GetMarketplaceImageArgs struct {
	// The instance type the image is compatible with.
	// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
	InstanceType *string `pulumi:"instanceType"`
	// Exact label of the desired image. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&per_page=100)
	// to find the right `label`.
	Label string `pulumi:"label"`
	// `zone`) The zone in which the image exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getMarketplaceImage.
type GetMarketplaceImageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	InstanceType *string `pulumi:"instanceType"`
	Label        string  `pulumi:"label"`
	Zone         string  `pulumi:"zone"`
}

func GetMarketplaceImageOutput(ctx *pulumi.Context, args GetMarketplaceImageOutputArgs, opts ...pulumi.InvokeOption) GetMarketplaceImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMarketplaceImageResult, error) {
			args := v.(GetMarketplaceImageArgs)
			r, err := GetMarketplaceImage(ctx, &args, opts...)
			var s GetMarketplaceImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMarketplaceImageResultOutput)
}

// A collection of arguments for invoking getMarketplaceImage.
type GetMarketplaceImageOutputArgs struct {
	// The instance type the image is compatible with.
	// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// Exact label of the desired image. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&per_page=100)
	// to find the right `label`.
	Label pulumi.StringInput `pulumi:"label"`
	// `zone`) The zone in which the image exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetMarketplaceImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMarketplaceImageArgs)(nil)).Elem()
}

// A collection of values returned by getMarketplaceImage.
type GetMarketplaceImageResultOutput struct{ *pulumi.OutputState }

func (GetMarketplaceImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMarketplaceImageResult)(nil)).Elem()
}

func (o GetMarketplaceImageResultOutput) ToGetMarketplaceImageResultOutput() GetMarketplaceImageResultOutput {
	return o
}

func (o GetMarketplaceImageResultOutput) ToGetMarketplaceImageResultOutputWithContext(ctx context.Context) GetMarketplaceImageResultOutput {
	return o
}

func (o GetMarketplaceImageResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetMarketplaceImageResult] {
	return pulumix.Output[GetMarketplaceImageResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetMarketplaceImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMarketplaceImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMarketplaceImageResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMarketplaceImageResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o GetMarketplaceImageResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v GetMarketplaceImageResult) string { return v.Label }).(pulumi.StringOutput)
}

func (o GetMarketplaceImageResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetMarketplaceImageResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMarketplaceImageResultOutput{})
}
