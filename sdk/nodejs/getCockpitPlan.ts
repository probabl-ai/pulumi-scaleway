// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a Scaleway Cockpit plan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const premium = scaleway.getCockpitPlan({
 *     name: "premium",
 * });
 * const main = new scaleway.Cockpit("main", {plan: premium.then(premium => premium.id)});
 * ```
 */
export function getCockpitPlan(args: GetCockpitPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetCockpitPlanResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getCockpitPlan:getCockpitPlan", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockpitPlan.
 */
export interface GetCockpitPlanArgs {
    /**
     * The name of the plan.
     */
    name: string;
}

/**
 * A collection of values returned by getCockpitPlan.
 */
export interface GetCockpitPlanResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
/**
 * Gets information about a Scaleway Cockpit plan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const premium = scaleway.getCockpitPlan({
 *     name: "premium",
 * });
 * const main = new scaleway.Cockpit("main", {plan: premium.then(premium => premium.id)});
 * ```
 */
export function getCockpitPlanOutput(args: GetCockpitPlanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCockpitPlanResult> {
    return pulumi.output(args).apply((a: any) => getCockpitPlan(a, opts))
}

/**
 * A collection of arguments for invoking getCockpitPlan.
 */
export interface GetCockpitPlanOutputArgs {
    /**
     * The name of the plan.
     */
    name: pulumi.Input<string>;
}
