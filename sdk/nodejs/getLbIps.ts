// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about multiple Load Balancer IP addresses.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getLbIps({
 *     ipCidrRange: "0.0.0.0/0",
 *     zone: "fr-par-2",
 * });
 * const ipsByTagsAndType = scaleway.getLbIps({
 *     ipType: "ipv4",
 *     tags: ["a tag"],
 * });
 * ```
 */
export function getLbIps(args?: GetLbIpsArgs, opts?: pulumi.InvokeOptions): Promise<GetLbIpsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getLbIps:getLbIps", {
        "ipCidrRange": args.ipCidrRange,
        "ipType": args.ipType,
        "projectId": args.projectId,
        "tags": args.tags,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbIps.
 */
export interface GetLbIpsArgs {
    /**
     * The IP CIDR range to filter for. IPs within a matching CIDR block are listed.
     */
    ipCidrRange?: string;
    /**
     * The IP type used as a filter.
     */
    ipType?: string;
    /**
     * The ID of the Project the Load Balancer is associated with.
     */
    projectId?: string;
    /**
     * List of tags used as filter. IPs with these exact tags are listed.
     */
    tags?: string[];
    /**
     * `zone`) The zone in which the IPs exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getLbIps.
 */
export interface GetLbIpsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipCidrRange?: string;
    readonly ipType?: string;
    /**
     * List of retrieved IPs
     */
    readonly ips: outputs.GetLbIpsIp[];
    /**
     * The ID of the Organization the Load Balancer is associated with.
     */
    readonly organizationId: string;
    /**
     * The ID of the Project the Load Balancer is associated with.
     */
    readonly projectId: string;
    readonly tags?: string[];
    /**
     * The zone of the Load Balancer.
     */
    readonly zone: string;
}
/**
 * Gets information about multiple Load Balancer IP addresses.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getLbIps({
 *     ipCidrRange: "0.0.0.0/0",
 *     zone: "fr-par-2",
 * });
 * const ipsByTagsAndType = scaleway.getLbIps({
 *     ipType: "ipv4",
 *     tags: ["a tag"],
 * });
 * ```
 */
export function getLbIpsOutput(args?: GetLbIpsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLbIpsResult> {
    return pulumi.output(args).apply((a: any) => getLbIps(a, opts))
}

/**
 * A collection of arguments for invoking getLbIps.
 */
export interface GetLbIpsOutputArgs {
    /**
     * The IP CIDR range to filter for. IPs within a matching CIDR block are listed.
     */
    ipCidrRange?: pulumi.Input<string>;
    /**
     * The IP type used as a filter.
     */
    ipType?: pulumi.Input<string>;
    /**
     * The ID of the Project the Load Balancer is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * List of tags used as filter. IPs with these exact tags are listed.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the IPs exist.
     */
    zone?: pulumi.Input<string>;
}
