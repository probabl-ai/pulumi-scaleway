// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Messaging and Queuing NATS accounts.
 * For further information, see
 * our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/)
 * To use the Scaleway provider with the official NATS JetStream provider, check out the corresponding guide.
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.MnqNatsAccount("main", {});
 * ```
 *
 * ## Import
 *
 * Namespaces can be imported using `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/mnqNatsAccount:MnqNatsAccount main fr-par/11111111111111111111111111111111
 * ```
 */
export class MnqNatsAccount extends pulumi.CustomResource {
    /**
     * Get an existing MnqNatsAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MnqNatsAccountState, opts?: pulumi.CustomResourceOptions): MnqNatsAccount {
        return new MnqNatsAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/mnqNatsAccount:MnqNatsAccount';

    /**
     * Returns true if the given object is an instance of MnqNatsAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MnqNatsAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MnqNatsAccount.__pulumiType;
    }

    /**
     * The endpoint of the NATS service for this account.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The unique name of the NATS account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the Project the
     * account is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region
     * in which the account should be created.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a MnqNatsAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MnqNatsAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MnqNatsAccountArgs | MnqNatsAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MnqNatsAccountState | undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as MnqNatsAccountArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["endpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MnqNatsAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MnqNatsAccount resources.
 */
export interface MnqNatsAccountState {
    /**
     * The endpoint of the NATS service for this account.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The unique name of the NATS account.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the Project the
     * account is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which the account should be created.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MnqNatsAccount resource.
 */
export interface MnqNatsAccountArgs {
    /**
     * The unique name of the NATS account.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the Project the
     * account is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which the account should be created.
     */
    region?: pulumi.Input<string>;
}
