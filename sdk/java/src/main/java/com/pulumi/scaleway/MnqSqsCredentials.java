// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.MnqSqsCredentialsArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.MnqSqsCredentialsState;
import com.pulumi.scaleway.outputs.MnqSqsCredentialsPermissions;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Messaging and queuing SQS Credentials.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
 * 
 * ## Examples
 * 
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.MnqSqs;
 * import com.pulumi.scaleway.MnqSqsCredentials;
 * import com.pulumi.scaleway.MnqSqsCredentialsArgs;
 * import com.pulumi.scaleway.inputs.MnqSqsCredentialsPermissionsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var mainMnqSqs = new MnqSqs(&#34;mainMnqSqs&#34;);
 * 
 *         var mainMnqSqsCredentials = new MnqSqsCredentials(&#34;mainMnqSqsCredentials&#34;, MnqSqsCredentialsArgs.builder()        
 *             .projectId(mainMnqSqs.projectId())
 *             .permissions(MnqSqsCredentialsPermissionsArgs.builder()
 *                 .canManage(false)
 *                 .canReceive(true)
 *                 .canPublish(false)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SQS credentials can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/mnqSqsCredentials:MnqSqsCredentials main fr-par/11111111111111111111111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/mnqSqsCredentials:MnqSqsCredentials")
public class MnqSqsCredentials extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the key.
     * 
     */
    @Export(name="accessKey", refs={String.class}, tree="[0]")
    private Output<String> accessKey;

    /**
     * @return The ID of the key.
     * 
     */
    public Output<String> accessKey() {
        return this.accessKey;
    }
    /**
     * The unique name of the sqs credentials.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique name of the sqs credentials.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * . List of permissions associated to these credentials. Only one of permissions may be set.
     * 
     */
    @Export(name="permissions", refs={MnqSqsCredentialsPermissions.class}, tree="[0]")
    private Output<MnqSqsCredentialsPermissions> permissions;

    /**
     * @return . List of permissions associated to these credentials. Only one of permissions may be set.
     * 
     */
    public Output<MnqSqsCredentialsPermissions> permissions() {
        return this.permissions;
    }
    /**
     * `project_id`) The ID of the project the sqs is enabled for.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the sqs is enabled for.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * `region`). The region in which sqs is enabled.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`). The region in which sqs is enabled.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The secret value of the key.
     * 
     */
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output<String> secretKey;

    /**
     * @return The secret value of the key.
     * 
     */
    public Output<String> secretKey() {
        return this.secretKey;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MnqSqsCredentials(String name) {
        this(name, MnqSqsCredentialsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MnqSqsCredentials(String name, @Nullable MnqSqsCredentialsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MnqSqsCredentials(String name, @Nullable MnqSqsCredentialsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/mnqSqsCredentials:MnqSqsCredentials", name, args == null ? MnqSqsCredentialsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MnqSqsCredentials(String name, Output<String> id, @Nullable MnqSqsCredentialsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/mnqSqsCredentials:MnqSqsCredentials", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessKey",
                "secretKey"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static MnqSqsCredentials get(String name, Output<String> id, @Nullable MnqSqsCredentialsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MnqSqsCredentials(name, id, state, options);
    }
}
