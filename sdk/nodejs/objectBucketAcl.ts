// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const someBucket = new scaleway.ObjectBucket("someBucket", {});
 * const main = new scaleway.ObjectBucketAcl("main", {
 *     bucket: scaleway_object_bucket.main.id,
 *     acl: "private",
 * });
 * ```
 *
 * ### With Grants
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const mainObjectBucket = new scaleway.ObjectBucket("mainObjectBucket", {});
 * const mainObjectBucketAcl = new scaleway.ObjectBucketAcl("mainObjectBucketAcl", {
 *     bucket: mainObjectBucket.id,
 *     accessControlPolicy: {
 *         grants: [
 *             {
 *                 grantee: {
 *                     id: "<project-id>:<project-id>",
 *                     type: "CanonicalUser",
 *                 },
 *                 permission: "FULL_CONTROL",
 *             },
 *             {
 *                 grantee: {
 *                     id: "<project-id>",
 *                     type: "CanonicalUser",
 *                 },
 *                 permission: "WRITE",
 *             },
 *         ],
 *         owner: {
 *             id: "<project-id>",
 *         },
 *     },
 * });
 * ```
 *
 * ## The ACL
 *
 * Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl)
 *
 * ## The Access Control policy
 *
 * The `accessControlPolicy` configuration block supports the following arguments:
 *
 * * `grant` - (Required) Set of grant configuration blocks documented below.
 * * `owner` - (Required) Configuration block of the bucket owner's display name and ID documented below.
 *
 * ## The Grant
 *
 * The `grant` configuration block supports the following arguments:
 *
 * * `grantee` - (Required) Configuration block for the project being granted permissions documented below.
 * * `permission` - (Required) Logging permissions assigned to the grantee for the bucket.
 *
 * ## The permission
 *
 * The following list shows each access policy permissions supported.
 *
 * `READ`, `WRITE`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`
 *
 * For more information about ACL permissions in the S3 bucket, see [ACL permissions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html).
 *
 * ## The owner
 *
 * The `owner` configuration block supports the following arguments:
 *
 * * `id` - (Required) The ID of the project owner.
 * * `displayName` - (Optional) The display name of the owner.
 *
 * ## the grantee
 *
 * The `grantee` configuration block supports the following arguments:
 *
 * * `id` - (Required) The canonical user ID of the grantee.
 * * `type` - (Required) Type of grantee. Valid values: CanonicalUser.
 *
 * ## Import
 *
 * Bucket ACLs can be imported using the `{region}/{bucketName}/{acl}` identifier, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/objectBucketAcl:ObjectBucketAcl some_bucket fr-par/some-bucket/private
 * ```
 *
 * ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
 *
 * If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/objectBucketAcl:ObjectBucketAcl some_bucket fr-par/some-bucket/private@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
 * ```
 */
export class ObjectBucketAcl extends pulumi.CustomResource {
    /**
     * Get an existing ObjectBucketAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectBucketAclState, opts?: pulumi.CustomResourceOptions): ObjectBucketAcl {
        return new ObjectBucketAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/objectBucketAcl:ObjectBucketAcl';

    /**
     * Returns true if the given object is an instance of ObjectBucketAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectBucketAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectBucketAcl.__pulumiType;
    }

    /**
     * A configuration block that sets the ACL permissions for an object per grantee documented below.
     */
    public readonly accessControlPolicy!: pulumi.Output<outputs.ObjectBucketAclAccessControlPolicy>;
    /**
     * The canned ACL you want to apply to the bucket.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * The bucket's name or regional ID.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The project ID of the expected bucket owner.
     */
    public readonly expectedBucketOwner!: pulumi.Output<string | undefined>;
    /**
     * The project_id you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a ObjectBucketAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectBucketAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectBucketAclArgs | ObjectBucketAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectBucketAclState | undefined;
            resourceInputs["accessControlPolicy"] = state ? state.accessControlPolicy : undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["expectedBucketOwner"] = state ? state.expectedBucketOwner : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ObjectBucketAclArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["accessControlPolicy"] = args ? args.accessControlPolicy : undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["expectedBucketOwner"] = args ? args.expectedBucketOwner : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectBucketAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectBucketAcl resources.
 */
export interface ObjectBucketAclState {
    /**
     * A configuration block that sets the ACL permissions for an object per grantee documented below.
     */
    accessControlPolicy?: pulumi.Input<inputs.ObjectBucketAclAccessControlPolicy>;
    /**
     * The canned ACL you want to apply to the bucket.
     */
    acl?: pulumi.Input<string>;
    /**
     * The bucket's name or regional ID.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The project ID of the expected bucket owner.
     */
    expectedBucketOwner?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectBucketAcl resource.
 */
export interface ObjectBucketAclArgs {
    /**
     * A configuration block that sets the ACL permissions for an object per grantee documented below.
     */
    accessControlPolicy?: pulumi.Input<inputs.ObjectBucketAclAccessControlPolicy>;
    /**
     * The canned ACL you want to apply to the bucket.
     */
    acl?: pulumi.Input<string>;
    /**
     * The bucket's name or regional ID.
     */
    bucket: pulumi.Input<string>;
    /**
     * The project ID of the expected bucket owner.
     */
    expectedBucketOwner?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    region?: pulumi.Input<string>;
}
