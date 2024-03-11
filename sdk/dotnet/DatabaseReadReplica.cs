// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    /// <summary>
    /// Creates and manages Scaleway Database read replicas.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Scaleway.DatabaseInstance("instance", new()
    ///     {
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         Tags = new[]
    ///         {
    ///             "terraform-test",
    ///             "scaleway_rdb_read_replica",
    ///             "minimal",
    ///         },
    ///     });
    /// 
    ///     var replica = new Scaleway.DatabaseReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         DirectAccess = null,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Private network
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Scaleway.DatabaseInstance("instance", new()
    ///     {
    ///         NodeType = "db-dev-s",
    ///         Engine = "PostgreSQL-14",
    ///         IsHaCluster = false,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///     });
    /// 
    ///     var pn = new Scaleway.VpcPrivateNetwork("pn");
    /// 
    ///     var replica = new Scaleway.DatabaseReadReplica("replica", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         PrivateNetwork = new Scaleway.Inputs.DatabaseReadReplicaPrivateNetworkArgs
    ///         {
    ///             PrivateNetworkId = pn.Id,
    ///             ServiceIp = "192.168.1.254/24",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Database Read replica can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/databaseReadReplica:DatabaseReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/databaseReadReplica:DatabaseReadReplica")]
    public partial class DatabaseReadReplica : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Output("directAccess")]
        public Output<Outputs.DatabaseReadReplicaDirectAccess?> DirectAccess { get; private set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.DatabaseReadReplicaPrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Output("sameZone")]
        public Output<bool?> SameZone { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseReadReplica resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseReadReplica(string name, DatabaseReadReplicaArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/databaseReadReplica:DatabaseReadReplica", name, args ?? new DatabaseReadReplicaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseReadReplica(string name, Input<string> id, DatabaseReadReplicaState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/databaseReadReplica:DatabaseReadReplica", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseReadReplica resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseReadReplica Get(string name, Input<string> id, DatabaseReadReplicaState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseReadReplica(name, id, state, options);
        }
    }

    public sealed class DatabaseReadReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DatabaseReadReplicaDirectAccessArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DatabaseReadReplicaPrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Input("sameZone")]
        public Input<bool>? SameZone { get; set; }

        public DatabaseReadReplicaArgs()
        {
        }
        public static new DatabaseReadReplicaArgs Empty => new DatabaseReadReplicaArgs();
    }

    public sealed class DatabaseReadReplicaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to rdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DatabaseReadReplicaDirectAccessGetArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DatabaseReadReplicaPrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        /// </summary>
        [Input("sameZone")]
        public Input<bool>? SameZone { get; set; }

        public DatabaseReadReplicaState()
        {
        }
        public static new DatabaseReadReplicaState Empty => new DatabaseReadReplicaState();
    }
}
