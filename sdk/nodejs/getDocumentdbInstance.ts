// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an DocumentDB instance. For further information see our [developers website](https://www.scaleway.com/en/developers/api/document_db/)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const db = scaleway.getDocumentdbInstance({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDocumentdbInstance(args?: GetDocumentdbInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetDocumentdbInstanceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getDocumentdbInstance:getDocumentdbInstance", {
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDocumentdbInstance.
 */
export interface GetDocumentdbInstanceArgs {
    /**
     * The DocumentDB instance ID.
     * Only one of `name` and `instanceId` should be specified.
     */
    instanceId?: string;
    /**
     * The name of the DocumentDB instance.
     * Only one of `name` and `instanceId` should be specified.
     */
    name?: string;
    /**
     * `region`) The region in which the DocumentDB instance exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getDocumentdbInstance.
 */
export interface GetDocumentdbInstanceResult {
    readonly engine: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly isHaCluster: boolean;
    readonly name?: string;
    readonly nodeType: string;
    readonly password: string;
    readonly projectId: string;
    readonly region?: string;
    readonly tags: string[];
    readonly telemetryEnabled: boolean;
    readonly userName: string;
    readonly volumeSizeInGb: number;
    readonly volumeType: string;
}
/**
 * Gets information about an DocumentDB instance. For further information see our [developers website](https://www.scaleway.com/en/developers/api/document_db/)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const db = scaleway.getDocumentdbInstance({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDocumentdbInstanceOutput(args?: GetDocumentdbInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDocumentdbInstanceResult> {
    return pulumi.output(args).apply((a: any) => getDocumentdbInstance(a, opts))
}

/**
 * A collection of arguments for invoking getDocumentdbInstance.
 */
export interface GetDocumentdbInstanceOutputArgs {
    /**
     * The DocumentDB instance ID.
     * Only one of `name` and `instanceId` should be specified.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the DocumentDB instance.
     * Only one of `name` and `instanceId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * `region`) The region in which the DocumentDB instance exists.
     */
    region?: pulumi.Input<string>;
}
