// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about the Scaleway Cockpit.
 *
 * For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getCockpit({});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getCockpit({
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCockpit(args?: GetCockpitArgs, opts?: pulumi.InvokeOptions): Promise<GetCockpitResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getCockpit:getCockpit", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockpit.
 */
export interface GetCockpitArgs {
    /**
     * `projectId`) The ID of the project the cockpit is associated with.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getCockpit.
 */
export interface GetCockpitResult {
    /**
     * Endpoints
     */
    readonly endpoints: outputs.GetCockpitEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID of the current plan
     */
    readonly planId: string;
    readonly projectId?: string;
    readonly pushUrls: outputs.GetCockpitPushUrl[];
}
/**
 * Gets information about the Scaleway Cockpit.
 *
 * For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getCockpit({});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getCockpit({
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCockpitOutput(args?: GetCockpitOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCockpitResult> {
    return pulumi.output(args).apply((a: any) => getCockpit(a, opts))
}

/**
 * A collection of arguments for invoking getCockpit.
 */
export interface GetCockpitOutputArgs {
    /**
     * `projectId`) The ID of the project the cockpit is associated with.
     */
    projectId?: pulumi.Input<string>;
}
