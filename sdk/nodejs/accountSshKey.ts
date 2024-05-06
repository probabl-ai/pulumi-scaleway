// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages user SSH keys to access servers provisioned on Scaleway.
 *
 * > **Important:**  The resource `scaleway.AccountSshKey` has been deprecated and will no longer be supported. Instead, use `scaleway.IamSshKey`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.AccountSshKey("main", {publicKey: "<YOUR-PUBLIC-SSH-KEY>"});
 * ```
 *
 * ## Import
 *
 * SSH keys can be imported using the `id`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/accountSshKey:AccountSshKey main 11111111-1111-1111-1111-111111111111
 * ```
 */
export class AccountSshKey extends pulumi.CustomResource {
    /**
     * Get an existing AccountSshKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountSshKeyState, opts?: pulumi.CustomResourceOptions): AccountSshKey {
        return new AccountSshKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/accountSshKey:AccountSshKey';

    /**
     * Returns true if the given object is an instance of AccountSshKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountSshKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountSshKey.__pulumiType;
    }

    /**
     * The date and time of the creation of the iam SSH Key
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The SSH key status
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * The fingerprint of the iam SSH key
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The name of the SSH key.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the SSH key is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the SSH key is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The public SSH key to be added.
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the iam SSH Key
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a AccountSshKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountSshKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountSshKeyArgs | AccountSshKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountSshKeyState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as AccountSshKeyArgs | undefined;
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountSshKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountSshKey resources.
 */
export interface AccountSshKeyState {
    /**
     * The date and time of the creation of the iam SSH Key
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The SSH key status
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The fingerprint of the iam SSH key
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The name of the SSH key.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the SSH key is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the SSH key is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The public SSH key to be added.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the iam SSH Key
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountSshKey resource.
 */
export interface AccountSshKeyArgs {
    /**
     * The SSH key status
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The name of the SSH key.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the SSH key is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The public SSH key to be added.
     */
    publicKey: pulumi.Input<string>;
}
