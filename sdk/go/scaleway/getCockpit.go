// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about the Scaleway Cockpit.
//
// For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/).
//
// ## Example Usage
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
//			_, err := scaleway.LookupCockpit(ctx, nil, nil)
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
//			_, err := scaleway.LookupCockpit(ctx, &scaleway.LookupCockpitArgs{
//				ProjectId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupCockpit(ctx *pulumi.Context, args *LookupCockpitArgs, opts ...pulumi.InvokeOption) (*LookupCockpitResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCockpitResult
	err := ctx.Invoke("scaleway:index/getCockpit:getCockpit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCockpit.
type LookupCockpitArgs struct {
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getCockpit.
type LookupCockpitResult struct {
	// Endpoints
	Endpoints []GetCockpitEndpoint `pulumi:"endpoints"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the current plan
	PlanId    string              `pulumi:"planId"`
	ProjectId *string             `pulumi:"projectId"`
	PushUrls  []GetCockpitPushUrl `pulumi:"pushUrls"`
}

func LookupCockpitOutput(ctx *pulumi.Context, args LookupCockpitOutputArgs, opts ...pulumi.InvokeOption) LookupCockpitResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCockpitResult, error) {
			args := v.(LookupCockpitArgs)
			r, err := LookupCockpit(ctx, &args, opts...)
			var s LookupCockpitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCockpitResultOutput)
}

// A collection of arguments for invoking getCockpit.
type LookupCockpitOutputArgs struct {
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupCockpitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCockpitArgs)(nil)).Elem()
}

// A collection of values returned by getCockpit.
type LookupCockpitResultOutput struct{ *pulumi.OutputState }

func (LookupCockpitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCockpitResult)(nil)).Elem()
}

func (o LookupCockpitResultOutput) ToLookupCockpitResultOutput() LookupCockpitResultOutput {
	return o
}

func (o LookupCockpitResultOutput) ToLookupCockpitResultOutputWithContext(ctx context.Context) LookupCockpitResultOutput {
	return o
}

// Endpoints
func (o LookupCockpitResultOutput) Endpoints() GetCockpitEndpointArrayOutput {
	return o.ApplyT(func(v LookupCockpitResult) []GetCockpitEndpoint { return v.Endpoints }).(GetCockpitEndpointArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCockpitResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCockpitResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the current plan
func (o LookupCockpitResultOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCockpitResult) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o LookupCockpitResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCockpitResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupCockpitResultOutput) PushUrls() GetCockpitPushUrlArrayOutput {
	return o.ApplyT(func(v LookupCockpitResult) []GetCockpitPushUrl { return v.PushUrls }).(GetCockpitPushUrlArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCockpitResultOutput{})
}
