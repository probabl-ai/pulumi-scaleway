// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway object storage objects.
 * For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
 *
 * ## Import
 *
 * Objects can be imported using the `{region}/{bucketName}/{objectKey}` identifier, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/objectItem:ObjectItem some_object fr-par/some-bucket/some-file
 * ```
 */
export class ObjectItem extends pulumi.CustomResource {
    /**
     * Get an existing ObjectItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectItemState, opts?: pulumi.CustomResourceOptions): ObjectItem {
        return new ObjectItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/objectItem:ObjectItem';

    /**
     * Returns true if the given object is an instance of ObjectItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectItem.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The name of the file to upload, defaults to an empty file
     */
    public readonly file!: pulumi.Output<string | undefined>;
    /**
     * Hash of the file, used to trigger upload on file change
     */
    public readonly hash!: pulumi.Output<string | undefined>;
    /**
     * The path of the object.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Map of metadata used for the object, keys must be lowercase
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Scaleway region this bucket resides in.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     */
    public readonly storageClass!: pulumi.Output<string | undefined>;
    /**
     * Map of tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Visibility of the object, `public-read` or `private`
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a ObjectItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectItemArgs | ObjectItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectItemState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["file"] = state ? state.file : undefined;
            resourceInputs["hash"] = state ? state.hash : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["storageClass"] = state ? state.storageClass : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ObjectItemArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["file"] = args ? args.file : undefined;
            resourceInputs["hash"] = args ? args.hash : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["storageClass"] = args ? args.storageClass : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectItem resources.
 */
export interface ObjectItemState {
    /**
     * The name of the bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The name of the file to upload, defaults to an empty file
     */
    file?: pulumi.Input<string>;
    /**
     * Hash of the file, used to trigger upload on file change
     */
    hash?: pulumi.Input<string>;
    /**
     * The path of the object.
     */
    key?: pulumi.Input<string>;
    /**
     * Map of metadata used for the object, keys must be lowercase
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Scaleway region this bucket resides in.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     */
    storageClass?: pulumi.Input<string>;
    /**
     * Map of tags
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Visibility of the object, `public-read` or `private`
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectItem resource.
 */
export interface ObjectItemArgs {
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * The name of the file to upload, defaults to an empty file
     */
    file?: pulumi.Input<string>;
    /**
     * Hash of the file, used to trigger upload on file change
     */
    hash?: pulumi.Input<string>;
    /**
     * The path of the object.
     */
    key: pulumi.Input<string>;
    /**
     * Map of metadata used for the object, keys must be lowercase
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Scaleway region this bucket resides in.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     */
    storageClass?: pulumi.Input<string>;
    /**
     * Map of tags
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Visibility of the object, `public-read` or `private`
     */
    visibility?: pulumi.Input<string>;
}
