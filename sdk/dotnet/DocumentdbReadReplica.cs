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
    /// Creates and manages Scaleway DocumentDB Database read replicas.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
    /// 
    /// ## Example Usage
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
    ///     var replica = new Scaleway.DocumentdbReadReplica("replica", new()
    ///     {
    ///         DirectAccess = null,
    ///         InstanceId = "11111111-1111-1111-1111-111111111111",
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
    ///     var pn = new Scaleway.VpcPrivateNetwork("pn");
    /// 
    ///     var replica = new Scaleway.DocumentdbReadReplica("replica", new()
    ///     {
    ///         InstanceId = scaleway_rdb_instance.Instance.Id,
    ///         PrivateNetwork = new Scaleway.Inputs.DocumentdbReadReplicaPrivateNetworkArgs
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
    /// $ pulumi import scaleway:index/documentdbReadReplica:DocumentdbReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/documentdbReadReplica:DocumentdbReadReplica")]
    public partial class DocumentdbReadReplica : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creates a direct access endpoint to documentdb replica.
        /// </summary>
        [Output("directAccess")]
        public Output<Outputs.DocumentdbReadReplicaDirectAccess?> DirectAccess { get; private set; } = null!;

        /// <summary>
        /// UUID of the documentdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.DocumentdbReadReplicaPrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentdbReadReplica resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentdbReadReplica(string name, DocumentdbReadReplicaArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbReadReplica:DocumentdbReadReplica", name, args ?? new DocumentdbReadReplicaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentdbReadReplica(string name, Input<string> id, DocumentdbReadReplicaState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbReadReplica:DocumentdbReadReplica", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DocumentdbReadReplica resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentdbReadReplica Get(string name, Input<string> id, DocumentdbReadReplicaState? state = null, CustomResourceOptions? options = null)
        {
            return new DocumentdbReadReplica(name, id, state, options);
        }
    }

    public sealed class DocumentdbReadReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to documentdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DocumentdbReadReplicaDirectAccessArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the documentdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentdbReadReplicaPrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentdbReadReplicaArgs()
        {
        }
        public static new DocumentdbReadReplicaArgs Empty => new DocumentdbReadReplicaArgs();
    }

    public sealed class DocumentdbReadReplicaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to documentdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DocumentdbReadReplicaDirectAccessGetArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the documentdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentdbReadReplicaPrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentdbReadReplicaState()
        {
        }
        public static new DocumentdbReadReplicaState Empty => new DocumentdbReadReplicaState();
    }
}
