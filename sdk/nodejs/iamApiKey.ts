// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway IAM API Keys. For more information, please
 * check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#api-keys-3665ae)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const ciCd = new scaleway.IamApplication("ciCd", {});
 * const main = new scaleway.IamApiKey("main", {
 *     applicationId: scaleway_iam_application.main.id,
 *     description: "a description",
 * });
 * ```
 *
 * ## Import
 *
 * Api keys can be imported using the `{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/iamApiKey:IamApiKey main 11111111111111111111
 * ```
 */
export class IamApiKey extends pulumi.CustomResource {
    /**
     * Get an existing IamApiKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamApiKeyState, opts?: pulumi.CustomResourceOptions): IamApiKey {
        return new IamApiKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/iamApiKey:IamApiKey';

    /**
     * Returns true if the given object is an instance of IamApiKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamApiKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamApiKey.__pulumiType;
    }

    /**
     * The access key of the iam api key.
     */
    public /*out*/ readonly accessKey!: pulumi.Output<string>;
    /**
     * ID of the application attached to the api key.
     * Only one of the `applicationId` and `userId` should be specified.
     */
    public readonly applicationId!: pulumi.Output<string | undefined>;
    /**
     * The date and time of the creation of the iam api key.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The IP Address of the device which created the API key.
     */
    public /*out*/ readonly creationIp!: pulumi.Output<string>;
    /**
     * The default project ID to use with object storage.
     */
    public readonly defaultProjectId!: pulumi.Output<string>;
    /**
     * The description of the iam api key.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the iam api key is editable.
     */
    public /*out*/ readonly editable!: pulumi.Output<boolean>;
    /**
     * The date and time of the expiration of the iam api key. Please note that in case of change,
     * the resource will be recreated.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The secret Key of the iam api key.
     */
    public /*out*/ readonly secretKey!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the iam api key.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * ID of the user attached to the api key.
     * Only one of the `applicationId` and `userId` should be specified.
     */
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a IamApiKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IamApiKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamApiKeyArgs | IamApiKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamApiKeyState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["creationIp"] = state ? state.creationIp : undefined;
            resourceInputs["defaultProjectId"] = state ? state.defaultProjectId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["editable"] = state ? state.editable : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as IamApiKeyArgs | undefined;
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["defaultProjectId"] = args ? args.defaultProjectId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["accessKey"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["creationIp"] = undefined /*out*/;
            resourceInputs["editable"] = undefined /*out*/;
            resourceInputs["secretKey"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IamApiKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamApiKey resources.
 */
export interface IamApiKeyState {
    /**
     * The access key of the iam api key.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * ID of the application attached to the api key.
     * Only one of the `applicationId` and `userId` should be specified.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The date and time of the creation of the iam api key.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The IP Address of the device which created the API key.
     */
    creationIp?: pulumi.Input<string>;
    /**
     * The default project ID to use with object storage.
     */
    defaultProjectId?: pulumi.Input<string>;
    /**
     * The description of the iam api key.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the iam api key is editable.
     */
    editable?: pulumi.Input<boolean>;
    /**
     * The date and time of the expiration of the iam api key. Please note that in case of change,
     * the resource will be recreated.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The secret Key of the iam api key.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the iam api key.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * ID of the user attached to the api key.
     * Only one of the `applicationId` and `userId` should be specified.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamApiKey resource.
 */
export interface IamApiKeyArgs {
    /**
     * ID of the application attached to the api key.
     * Only one of the `applicationId` and `userId` should be specified.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The default project ID to use with object storage.
     */
    defaultProjectId?: pulumi.Input<string>;
    /**
     * The description of the iam api key.
     */
    description?: pulumi.Input<string>;
    /**
     * The date and time of the expiration of the iam api key. Please note that in case of change,
     * the resource will be recreated.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * ID of the user attached to the api key.
     * Only one of the `applicationId` and `userId` should be specified.
     */
    userId?: pulumi.Input<string>;
}
