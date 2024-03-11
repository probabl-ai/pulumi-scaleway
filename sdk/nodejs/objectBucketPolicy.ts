// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway object storage bucket policy.
 * For more information, see [the documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/using-bucket-policies/).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const bucket = new scaleway.ObjectBucket("bucket", {});
 * const main = new scaleway.IamApplication("main", {description: "a description"});
 * const policy = new scaleway.ObjectBucketPolicy("policy", {
 *     bucket: bucket.name,
 *     policy: pulumi.jsonStringify({
 *         Version: "2023-04-17",
 *         Id: "MyBucketPolicy",
 *         Statement: [{
 *             Sid: "Delegate access",
 *             Effect: "Allow",
 *             Principal: {
 *                 SCW: pulumi.interpolate`application_id:${main.id}`,
 *             },
 *             Action: "s3:ListBucket",
 *             Resource: [
 *                 bucket.name,
 *                 pulumi.interpolate`${bucket.name}/*`,
 *             ],
 *         }],
 *     }),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Buckets can be imported using the `{region}/{bucketName}` identifier, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/objectBucketPolicy:ObjectBucketPolicy some_bucket fr-par/some-bucket
 * ```
 */
export class ObjectBucketPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ObjectBucketPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectBucketPolicyState, opts?: pulumi.CustomResourceOptions): ObjectBucketPolicy {
        return new ObjectBucketPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/objectBucketPolicy:ObjectBucketPolicy';

    /**
     * Returns true if the given object is an instance of ObjectBucketPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectBucketPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectBucketPolicy.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The text of the policy.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     *
     * > **Important:** The awsIamPolicyDocument data source may be used, so long as it specifies a principal.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The Scaleway region this bucket resides in.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a ObjectBucketPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectBucketPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectBucketPolicyArgs | ObjectBucketPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectBucketPolicyState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ObjectBucketPolicyArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectBucketPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectBucketPolicy resources.
 */
export interface ObjectBucketPolicyState {
    /**
     * The name of the bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The text of the policy.
     */
    policy?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     *
     * > **Important:** The awsIamPolicyDocument data source may be used, so long as it specifies a principal.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Scaleway region this bucket resides in.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectBucketPolicy resource.
 */
export interface ObjectBucketPolicyArgs {
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * The text of the policy.
     */
    policy: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     *
     * > **Important:** The awsIamPolicyDocument data source may be used, so long as it specifies a principal.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The Scaleway region this bucket resides in.
     */
    region?: pulumi.Input<string>;
}
