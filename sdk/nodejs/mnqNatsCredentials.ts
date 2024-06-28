// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Messaging and Queuing NATS credentials.
 * For further information, see
 * our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const mainMnqNatsAccount = new scaleway.MnqNatsAccount("mainMnqNatsAccount", {});
 * const mainMnqNatsCredentials = new scaleway.MnqNatsCredentials("mainMnqNatsCredentials", {accountId: mainMnqNatsAccount.id});
 * ```
 *
 * ## Import
 *
 * Namespaces can be imported using `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/mnqNatsCredentials:MnqNatsCredentials main fr-par/11111111111111111111111111111111
 * ```
 */
export class MnqNatsCredentials extends pulumi.CustomResource {
    /**
     * Get an existing MnqNatsCredentials resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MnqNatsCredentialsState, opts?: pulumi.CustomResourceOptions): MnqNatsCredentials {
        return new MnqNatsCredentials(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/mnqNatsCredentials:MnqNatsCredentials';

    /**
     * Returns true if the given object is an instance of MnqNatsCredentials.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MnqNatsCredentials {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MnqNatsCredentials.__pulumiType;
    }

    /**
     * The ID of the NATS account the credentials are generated from
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The content of the credentials file.
     */
    public /*out*/ readonly file!: pulumi.Output<string>;
    /**
     * The unique name of the NATS credentials.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `region`). The region
     * in which the account exists.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a MnqNatsCredentials resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MnqNatsCredentialsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MnqNatsCredentialsArgs | MnqNatsCredentialsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MnqNatsCredentialsState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["file"] = state ? state.file : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as MnqNatsCredentialsArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["file"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MnqNatsCredentials.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MnqNatsCredentials resources.
 */
export interface MnqNatsCredentialsState {
    /**
     * The ID of the NATS account the credentials are generated from
     */
    accountId?: pulumi.Input<string>;
    /**
     * The content of the credentials file.
     */
    file?: pulumi.Input<string>;
    /**
     * The unique name of the NATS credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which the account exists.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MnqNatsCredentials resource.
 */
export interface MnqNatsCredentialsArgs {
    /**
     * The ID of the NATS account the credentials are generated from
     */
    accountId: pulumi.Input<string>;
    /**
     * The unique name of the NATS credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which the account exists.
     */
    region?: pulumi.Input<string>;
}
