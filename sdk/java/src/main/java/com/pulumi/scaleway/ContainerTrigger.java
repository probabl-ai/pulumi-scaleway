// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.ContainerTriggerArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.ContainerTriggerState;
import com.pulumi.scaleway.outputs.ContainerTriggerNats;
import com.pulumi.scaleway.outputs.ContainerTriggerSqs;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Container Triggers.
 * For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-triggers).
 * 
 * ## Examples
 * 
 * ### SQS
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ContainerTrigger;
 * import com.pulumi.scaleway.ContainerTriggerArgs;
 * import com.pulumi.scaleway.inputs.ContainerTriggerSqsArgs;
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
 *         var main = new ContainerTrigger(&#34;main&#34;, ContainerTriggerArgs.builder()        
 *             .containerId(scaleway_container.main().id())
 *             .sqs(ContainerTriggerSqsArgs.builder()
 *                 .projectId(scaleway_mnq_sqs.main().project_id())
 *                 .queue(&#34;MyQueue&#34;)
 *                 .region(scaleway_mnq_sqs.main().region())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### Nats
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ContainerTrigger;
 * import com.pulumi.scaleway.ContainerTriggerArgs;
 * import com.pulumi.scaleway.inputs.ContainerTriggerNatsArgs;
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
 *         var main = new ContainerTrigger(&#34;main&#34;, ContainerTriggerArgs.builder()        
 *             .containerId(scaleway_container.main().id())
 *             .nats(ContainerTriggerNatsArgs.builder()
 *                 .accountId(scaleway_mnq_nats_account.main().id())
 *                 .subject(&#34;MySubject&#34;)
 *                 .region(scaleway_mnq_nats_account.main().region())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Container Triggers can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/containerTrigger:ContainerTrigger main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/containerTrigger:ContainerTrigger")
public class ContainerTrigger extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the container to create a trigger for
     * 
     */
    @Export(name="containerId", refs={String.class}, tree="[0]")
    private Output<String> containerId;

    /**
     * @return The ID of the container to create a trigger for
     * 
     */
    public Output<String> containerId() {
        return this.containerId;
    }
    /**
     * The description of the trigger.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the trigger.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The unique name of the trigger. Default to a generated name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique name of the trigger. Default to a generated name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The configuration for the Scaleway&#39;s Nats used by the trigger
     * 
     */
    @Export(name="nats", refs={ContainerTriggerNats.class}, tree="[0]")
    private Output</* @Nullable */ ContainerTriggerNats> nats;

    /**
     * @return The configuration for the Scaleway&#39;s Nats used by the trigger
     * 
     */
    public Output<Optional<ContainerTriggerNats>> nats() {
        return Codegen.optional(this.nats);
    }
    /**
     * `region`). The region in which the namespace should be created.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`). The region in which the namespace should be created.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The configuration of the Scaleway&#39;s SQS used by the trigger
     * 
     */
    @Export(name="sqs", refs={ContainerTriggerSqs.class}, tree="[0]")
    private Output</* @Nullable */ ContainerTriggerSqs> sqs;

    /**
     * @return The configuration of the Scaleway&#39;s SQS used by the trigger
     * 
     */
    public Output<Optional<ContainerTriggerSqs>> sqs() {
        return Codegen.optional(this.sqs);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerTrigger(String name) {
        this(name, ContainerTriggerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerTrigger(String name, ContainerTriggerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerTrigger(String name, ContainerTriggerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/containerTrigger:ContainerTrigger", name, args == null ? ContainerTriggerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerTrigger(String name, Output<String> id, @Nullable ContainerTriggerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/containerTrigger:ContainerTrigger", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static ContainerTrigger get(String name, Output<String> id, @Nullable ContainerTriggerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerTrigger(name, id, state, options);
    }
}
