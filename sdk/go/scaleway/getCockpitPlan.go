// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Scaleway Cockpit plan.
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
//			premium, err := scaleway.GetCockpitPlan(ctx, &scaleway.GetCockpitPlanArgs{
//				Name: "premium",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewCockpit(ctx, "main", &scaleway.CockpitArgs{
//				Plan: *pulumi.String(premium.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCockpitPlan(ctx *pulumi.Context, args *GetCockpitPlanArgs, opts ...pulumi.InvokeOption) (*GetCockpitPlanResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCockpitPlanResult
	err := ctx.Invoke("scaleway:index/getCockpitPlan:getCockpitPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCockpitPlan.
type GetCockpitPlanArgs struct {
	// The name of the plan.
	Name string `pulumi:"name"`
}

// A collection of values returned by getCockpitPlan.
type GetCockpitPlanResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func GetCockpitPlanOutput(ctx *pulumi.Context, args GetCockpitPlanOutputArgs, opts ...pulumi.InvokeOption) GetCockpitPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCockpitPlanResult, error) {
			args := v.(GetCockpitPlanArgs)
			r, err := GetCockpitPlan(ctx, &args, opts...)
			var s GetCockpitPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCockpitPlanResultOutput)
}

// A collection of arguments for invoking getCockpitPlan.
type GetCockpitPlanOutputArgs struct {
	// The name of the plan.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetCockpitPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCockpitPlanArgs)(nil)).Elem()
}

// A collection of values returned by getCockpitPlan.
type GetCockpitPlanResultOutput struct{ *pulumi.OutputState }

func (GetCockpitPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCockpitPlanResult)(nil)).Elem()
}

func (o GetCockpitPlanResultOutput) ToGetCockpitPlanResultOutput() GetCockpitPlanResultOutput {
	return o
}

func (o GetCockpitPlanResultOutput) ToGetCockpitPlanResultOutputWithContext(ctx context.Context) GetCockpitPlanResultOutput {
	return o
}

func (o GetCockpitPlanResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetCockpitPlanResult] {
	return pulumix.Output[GetCockpitPlanResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetCockpitPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockpitPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCockpitPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockpitPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCockpitPlanResultOutput{})
}
