// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Load-Balancer ACLs. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls).
 *
 * ## Examples Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const acl01 = new scaleway.LoadbalancerAcl("acl01", {
 *     frontendId: scaleway_lb_frontend.frt01.id,
 *     description: "Exclude well-known IPs",
 *     index: 0,
 *     action: {
 *         type: "allow",
 *     },
 *     match: {
 *         ipSubnets: [
 *             "192.168.0.1",
 *             "192.168.0.2",
 *             "192.168.10.0/24",
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Load-Balancer ACL can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/loadbalancerAcl:LoadbalancerAcl acl01 fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class LoadbalancerAcl extends pulumi.CustomResource {
    /**
     * Get an existing LoadbalancerAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadbalancerAclState, opts?: pulumi.CustomResourceOptions): LoadbalancerAcl {
        return new LoadbalancerAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/loadbalancerAcl:LoadbalancerAcl';

    /**
     * Returns true if the given object is an instance of LoadbalancerAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadbalancerAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadbalancerAcl.__pulumiType;
    }

    /**
     * Action to undertake when an ACL filter matches.
     */
    public readonly action!: pulumi.Output<outputs.LoadbalancerAclAction>;
    /**
     * Date and time of ACL's creation (RFC 3339 format)
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The ACL description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The load-balancer Frontend ID to attach the ACL to.
     */
    public readonly frontendId!: pulumi.Output<string>;
    /**
     * The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
     */
    public readonly index!: pulumi.Output<number>;
    /**
     * The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
     */
    public readonly match!: pulumi.Output<outputs.LoadbalancerAclMatch | undefined>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Date and time of ACL's update (RFC 3339 format)
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a LoadbalancerAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadbalancerAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadbalancerAclArgs | LoadbalancerAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadbalancerAclState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["frontendId"] = state ? state.frontendId : undefined;
            resourceInputs["index"] = state ? state.index : undefined;
            resourceInputs["match"] = state ? state.match : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as LoadbalancerAclArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.frontendId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'frontendId'");
            }
            if ((!args || args.index === undefined) && !opts.urn) {
                throw new Error("Missing required property 'index'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["frontendId"] = args ? args.frontendId : undefined;
            resourceInputs["index"] = args ? args.index : undefined;
            resourceInputs["match"] = args ? args.match : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadbalancerAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadbalancerAcl resources.
 */
export interface LoadbalancerAclState {
    /**
     * Action to undertake when an ACL filter matches.
     */
    action?: pulumi.Input<inputs.LoadbalancerAclAction>;
    /**
     * Date and time of ACL's creation (RFC 3339 format)
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The ACL description.
     */
    description?: pulumi.Input<string>;
    /**
     * The load-balancer Frontend ID to attach the ACL to.
     */
    frontendId?: pulumi.Input<string>;
    /**
     * The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
     */
    index?: pulumi.Input<number>;
    /**
     * The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
     */
    match?: pulumi.Input<inputs.LoadbalancerAclMatch>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * Date and time of ACL's update (RFC 3339 format)
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadbalancerAcl resource.
 */
export interface LoadbalancerAclArgs {
    /**
     * Action to undertake when an ACL filter matches.
     */
    action: pulumi.Input<inputs.LoadbalancerAclAction>;
    /**
     * The ACL description.
     */
    description?: pulumi.Input<string>;
    /**
     * The load-balancer Frontend ID to attach the ACL to.
     */
    frontendId: pulumi.Input<string>;
    /**
     * The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
     */
    index: pulumi.Input<number>;
    /**
     * The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
     */
    match?: pulumi.Input<inputs.LoadbalancerAclMatch>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
}
