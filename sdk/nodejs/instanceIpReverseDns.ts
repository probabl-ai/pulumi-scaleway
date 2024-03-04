// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages Scaleway Compute Instance IPs Reverse DNS.
 *
 * Please check our [guide](https://www.scaleway.com/en/docs/compute/instances/how-to/configure-reverse-dns/) for more details
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const serverIp = new scaleway.InstanceIp("serverIp", {});
 * const tfA = new scaleway.DomainRecord("tfA", {
 *     dnsZone: "scaleway.com",
 *     type: "A",
 *     data: serverIp.address,
 *     ttl: 3600,
 *     priority: 1,
 * });
 * const reverse = new scaleway.InstanceIpReverseDns("reverse", {
 *     ipId: serverIp.id,
 *     reverse: "www.scaleway.com",
 * });
 * ```
 *
 * ## Import
 *
 * IPs reverse DNS can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/instanceIpReverseDns:InstanceIpReverseDns reverse fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class InstanceIpReverseDns extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIpReverseDns resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIpReverseDnsState, opts?: pulumi.CustomResourceOptions): InstanceIpReverseDns {
        return new InstanceIpReverseDns(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instanceIpReverseDns:InstanceIpReverseDns';

    /**
     * Returns true if the given object is an instance of InstanceIpReverseDns.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIpReverseDns {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIpReverseDns.__pulumiType;
    }

    /**
     * The IP ID
     */
    public readonly ipId!: pulumi.Output<string>;
    /**
     * The reverse DNS for this IP.
     */
    public readonly reverse!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceIpReverseDns resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIpReverseDnsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIpReverseDnsArgs | InstanceIpReverseDnsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceIpReverseDnsState | undefined;
            resourceInputs["ipId"] = state ? state.ipId : undefined;
            resourceInputs["reverse"] = state ? state.reverse : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceIpReverseDnsArgs | undefined;
            if ((!args || args.ipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipId'");
            }
            if ((!args || args.reverse === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reverse'");
            }
            resourceInputs["ipId"] = args ? args.ipId : undefined;
            resourceInputs["reverse"] = args ? args.reverse : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceIpReverseDns.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIpReverseDns resources.
 */
export interface InstanceIpReverseDnsState {
    /**
     * The IP ID
     */
    ipId?: pulumi.Input<string>;
    /**
     * The reverse DNS for this IP.
     */
    reverse?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIpReverseDns resource.
 */
export interface InstanceIpReverseDnsArgs {
    /**
     * The IP ID
     */
    ipId: pulumi.Input<string>;
    /**
     * The reverse DNS for this IP.
     */
    reverse: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    zone?: pulumi.Input<string>;
}
