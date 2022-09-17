// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.scaleway.DatabaseReadReplicaArgs;
import com.pulumi.scaleway.Utilities;
import com.pulumi.scaleway.inputs.DatabaseReadReplicaState;
import com.pulumi.scaleway.outputs.DatabaseReadReplicaDirectAccess;
import com.pulumi.scaleway.outputs.DatabaseReadReplicaPrivateNetwork;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Database read replicas.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
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
 * import com.pulumi.scaleway.DatabaseInstance;
 * import com.pulumi.scaleway.DatabaseInstanceArgs;
 * import com.pulumi.scaleway.DatabaseReadReplica;
 * import com.pulumi.scaleway.DatabaseReadReplicaArgs;
 * import com.pulumi.scaleway.inputs.DatabaseReadReplicaDirectAccessArgs;
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
 *         var instance = new DatabaseInstance(&#34;instance&#34;, DatabaseInstanceArgs.builder()        
 *             .nodeType(&#34;db-dev-s&#34;)
 *             .engine(&#34;PostgreSQL-14&#34;)
 *             .isHaCluster(false)
 *             .disableBackup(true)
 *             .userName(&#34;my_initial_user&#34;)
 *             .password(&#34;thiZ_is_v&amp;ry_s3cret&#34;)
 *             .tags(            
 *                 &#34;terraform-test&#34;,
 *                 &#34;scaleway_rdb_read_replica&#34;,
 *                 &#34;minimal&#34;)
 *             .build());
 * 
 *         var replica = new DatabaseReadReplica(&#34;replica&#34;, DatabaseReadReplicaArgs.builder()        
 *             .instanceId(instance.id())
 *             .directAccess()
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### Private network
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.DatabaseInstance;
 * import com.pulumi.scaleway.DatabaseInstanceArgs;
 * import com.pulumi.scaleway.VpcPrivateNetwork;
 * import com.pulumi.scaleway.DatabaseReadReplica;
 * import com.pulumi.scaleway.DatabaseReadReplicaArgs;
 * import com.pulumi.scaleway.inputs.DatabaseReadReplicaPrivateNetworkArgs;
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
 *         var instance = new DatabaseInstance(&#34;instance&#34;, DatabaseInstanceArgs.builder()        
 *             .nodeType(&#34;db-dev-s&#34;)
 *             .engine(&#34;PostgreSQL-14&#34;)
 *             .isHaCluster(false)
 *             .disableBackup(true)
 *             .userName(&#34;my_initial_user&#34;)
 *             .password(&#34;thiZ_is_v&amp;ry_s3cret&#34;)
 *             .build());
 * 
 *         var pn = new VpcPrivateNetwork(&#34;pn&#34;);
 * 
 *         var replica = new DatabaseReadReplica(&#34;replica&#34;, DatabaseReadReplicaArgs.builder()        
 *             .instanceId(instance.id())
 *             .privateNetwork(DatabaseReadReplicaPrivateNetworkArgs.builder()
 *                 .privateNetworkId(pn.id())
 *                 .serviceIp(&#34;192.168.1.254/24&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Database Read replica can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/databaseReadReplica:DatabaseReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/databaseReadReplica:DatabaseReadReplica")
public class DatabaseReadReplica extends com.pulumi.resources.CustomResource {
    /**
     * Creates a direct access endpoint to rdb replica.
     * 
     */
    @Export(name="directAccess", type=DatabaseReadReplicaDirectAccess.class, parameters={})
    private Output</* @Nullable */ DatabaseReadReplicaDirectAccess> directAccess;

    /**
     * @return Creates a direct access endpoint to rdb replica.
     * 
     */
    public Output<Optional<DatabaseReadReplicaDirectAccess>> directAccess() {
        return Codegen.optional(this.directAccess);
    }
    /**
     * Id of the rdb instance to replicate.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return Id of the rdb instance to replicate.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Create an endpoint in a private network.
     * 
     */
    @Export(name="privateNetwork", type=DatabaseReadReplicaPrivateNetwork.class, parameters={})
    private Output</* @Nullable */ DatabaseReadReplicaPrivateNetwork> privateNetwork;

    /**
     * @return Create an endpoint in a private network.
     * 
     */
    public Output<Optional<DatabaseReadReplicaPrivateNetwork>> privateNetwork() {
        return Codegen.optional(this.privateNetwork);
    }
    /**
     * `region`) The region in which the Database read replica should be created.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return `region`) The region in which the Database read replica should be created.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatabaseReadReplica(String name) {
        this(name, DatabaseReadReplicaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatabaseReadReplica(String name, DatabaseReadReplicaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatabaseReadReplica(String name, DatabaseReadReplicaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/databaseReadReplica:DatabaseReadReplica", name, args == null ? DatabaseReadReplicaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DatabaseReadReplica(String name, Output<String> id, @Nullable DatabaseReadReplicaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/databaseReadReplica:DatabaseReadReplica", name, state, makeResourceOptions(options, id));
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
    public static DatabaseReadReplica get(String name, Output<String> id, @Nullable DatabaseReadReplicaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatabaseReadReplica(name, id, state, options);
    }
}
