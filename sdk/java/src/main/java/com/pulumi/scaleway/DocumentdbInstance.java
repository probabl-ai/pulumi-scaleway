// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.DocumentdbInstanceArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.DocumentdbInstanceState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Database Instances.
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
 * 
 * ## Examples
 * 
 * ### Example Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.DocumentdbInstance;
 * import com.pulumi.scaleway.DocumentdbInstanceArgs;
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
 *         var main = new DocumentdbInstance(&#34;main&#34;, DocumentdbInstanceArgs.builder()        
 *             .engine(&#34;FerretDB-1&#34;)
 *             .nodeType(&#34;docdb-play2-pico&#34;)
 *             .password(&#34;thiZ_is_v&amp;ry_s3cret&#34;)
 *             .tags(            
 *                 &#34;terraform-test&#34;,
 *                 &#34;scaleway_documentdb_instance&#34;,
 *                 &#34;minimal&#34;)
 *             .userName(&#34;my_initial_user&#34;)
 *             .volumeSizeInGb(20)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Database Instance can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/documentdbInstance:DocumentdbInstance db fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/documentdbInstance:DocumentdbInstance")
public class DocumentdbInstance extends com.pulumi.resources.CustomResource {
    /**
     * Database Instance&#39;s engine version (e.g. `FerretDB-1`).
     * 
     * &gt; **Important:** Updates to `engine` will recreate the Database Instance.
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output<String> engine;

    /**
     * @return Database Instance&#39;s engine version (e.g. `FerretDB-1`).
     * 
     * &gt; **Important:** Updates to `engine` will recreate the Database Instance.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Enable or disable high availability for the database instance.
     * 
     */
    @Export(name="isHaCluster", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isHaCluster;

    /**
     * @return Enable or disable high availability for the database instance.
     * 
     */
    public Output<Optional<Boolean>> isHaCluster() {
        return Codegen.optional(this.isHaCluster);
    }
    /**
     * The name of the Database Instance.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Database Instance.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The type of database instance you want to create (e.g. `docdb-play2-pico`).
     * 
     * &gt; **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
     * interruption. Keep in mind that you cannot downgrade a Database Instance.
     * 
     */
    @Export(name="nodeType", refs={String.class}, tree="[0]")
    private Output<String> nodeType;

    /**
     * @return The type of database instance you want to create (e.g. `docdb-play2-pico`).
     * 
     * &gt; **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
     * interruption. Keep in mind that you cannot downgrade a Database Instance.
     * 
     */
    public Output<String> nodeType() {
        return this.nodeType;
    }
    /**
     * Password for the first user of the database instance.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return Password for the first user of the database instance.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * `project_id`) The ID of the project the Database
     * Instance is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the Database
     * Instance is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * `region`) The region
     * in which the Database Instance should be created.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`) The region
     * in which the Database Instance should be created.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The tags associated with the Database Instance.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The tags associated with the Database Instance.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
     * 
     * &gt; **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
     * 
     */
    @Export(name="telemetryEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> telemetryEnabled;

    /**
     * @return Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
     * 
     * &gt; **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
     * 
     */
    public Output<Optional<Boolean>> telemetryEnabled() {
        return Codegen.optional(this.telemetryEnabled);
    }
    /**
     * Identifier for the first user of the database instance.
     * 
     * &gt; **Important:** Updates to `user_name` will recreate the Database Instance.
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userName;

    /**
     * @return Identifier for the first user of the database instance.
     * 
     * &gt; **Important:** Updates to `user_name` will recreate the Database Instance.
     * 
     */
    public Output<Optional<String>> userName() {
        return Codegen.optional(this.userName);
    }
    /**
     * Volume size (in GB) when `volume_type` is set to `bssd`.
     * 
     */
    @Export(name="volumeSizeInGb", refs={Integer.class}, tree="[0]")
    private Output<Integer> volumeSizeInGb;

    /**
     * @return Volume size (in GB) when `volume_type` is set to `bssd`.
     * 
     */
    public Output<Integer> volumeSizeInGb() {
        return this.volumeSizeInGb;
    }
    /**
     * Type of volume where data are stored (`bssd` or `lssd`).
     * 
     */
    @Export(name="volumeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> volumeType;

    /**
     * @return Type of volume where data are stored (`bssd` or `lssd`).
     * 
     */
    public Output<Optional<String>> volumeType() {
        return Codegen.optional(this.volumeType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DocumentdbInstance(String name) {
        this(name, DocumentdbInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DocumentdbInstance(String name, DocumentdbInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DocumentdbInstance(String name, DocumentdbInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/documentdbInstance:DocumentdbInstance", name, args == null ? DocumentdbInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DocumentdbInstance(String name, Output<String> id, @Nullable DocumentdbInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/documentdbInstance:DocumentdbInstance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static DocumentdbInstance get(String name, Output<String> id, @Nullable DocumentdbInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DocumentdbInstance(name, id, state, options);
    }
}
