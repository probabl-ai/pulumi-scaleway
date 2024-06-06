// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Compute Instance IPs. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/instance/#ips-268151).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const serverIp = new scaleway.InstanceIp("serverIp", {});
 * ```
 *
 * ## Import
 *
 * IPs can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/instanceIp:InstanceIp server_ip fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class InstanceIp extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIpState, opts?: pulumi.CustomResourceOptions): InstanceIp {
        return new InstanceIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instanceIp:InstanceIp';

    /**
     * Returns true if the given object is an instance of InstanceIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIp.__pulumiType;
    }

    /**
     * The IP address.
     */
    public /*out*/ readonly address!: pulumi.Output<string>;
    /**
     * The organization ID the IP is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The IP Prefix.
     */
    public /*out*/ readonly prefix!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The reverse dns attached to this IP
     */
    public /*out*/ readonly reverse!: pulumi.Output<string>;
    /**
     * The server associated with this IP
     */
    public /*out*/ readonly serverId!: pulumi.Output<string>;
    /**
     * The tags associated with the IP.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of the IP (`nat`, `routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
     *
     * > **Important:** An IP can migrate from `nat` to `routedIpv4` but cannot be converted back
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIpArgs | InstanceIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceIpState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["prefix"] = state ? state.prefix : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["reverse"] = state ? state.reverse : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceIpArgs | undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["address"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["prefix"] = undefined /*out*/;
            resourceInputs["reverse"] = undefined /*out*/;
            resourceInputs["serverId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIp resources.
 */
export interface InstanceIpState {
    /**
     * The IP address.
     */
    address?: pulumi.Input<string>;
    /**
     * The organization ID the IP is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The IP Prefix.
     */
    prefix?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The reverse dns attached to this IP
     */
    reverse?: pulumi.Input<string>;
    /**
     * The server associated with this IP
     */
    serverId?: pulumi.Input<string>;
    /**
     * The tags associated with the IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the IP (`nat`, `routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
     *
     * > **Important:** An IP can migrate from `nat` to `routedIpv4` but cannot be converted back
     */
    type?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIp resource.
 */
export interface InstanceIpArgs {
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The tags associated with the IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the IP (`nat`, `routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
     *
     * > **Important:** An IP can migrate from `nat` to `routedIpv4` but cannot be converted back
     */
    type?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    zone?: pulumi.Input<string>;
}
