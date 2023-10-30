// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about SQS for a project
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getMnqSqs({});
 * const forProject = scaleway.getMnqSqs({
 *     projectId: scaleway_account_project.main.id,
 * });
 * ```
 */
export function getMnqSqs(args?: GetMnqSqsArgs, opts?: pulumi.InvokeOptions): Promise<GetMnqSqsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getMnqSqs:getMnqSqs", {
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getMnqSqs.
 */
export interface GetMnqSqsArgs {
    /**
     * `projectId`) The ID of the project for which sqs is enabled.
     */
    projectId?: string;
    /**
     * `region`). The region in which sqs is enabled.
     */
    region?: string;
}

/**
 * A collection of values returned by getMnqSqs.
 */
export interface GetMnqSqsResult {
    /**
     * The endpoint of the SQS service for this project.
     */
    readonly endpoint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId?: string;
    readonly region?: string;
}
/**
 * Gets information about SQS for a project
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getMnqSqs({});
 * const forProject = scaleway.getMnqSqs({
 *     projectId: scaleway_account_project.main.id,
 * });
 * ```
 */
export function getMnqSqsOutput(args?: GetMnqSqsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMnqSqsResult> {
    return pulumi.output(args).apply((a: any) => getMnqSqs(a, opts))
}

/**
 * A collection of arguments for invoking getMnqSqs.
 */
export interface GetMnqSqsOutputArgs {
    /**
     * `projectId`) The ID of the project for which sqs is enabled.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which sqs is enabled.
     */
    region?: pulumi.Input<string>;
}
