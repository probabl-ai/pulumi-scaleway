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
    /// Creates and manages Scaleway Messaging and Queuing NATS credentials.
    /// For further information, see
    /// our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/).
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainMnqNatsAccount = new Scaleway.MnqNatsAccount("mainMnqNatsAccount");
    /// 
    ///     var mainMnqNatsCredentials = new Scaleway.MnqNatsCredentials("mainMnqNatsCredentials", new()
    ///     {
    ///         AccountId = mainMnqNatsAccount.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Namespaces can be imported using `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/mnqNatsCredentials:MnqNatsCredentials main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mnqNatsCredentials:MnqNatsCredentials")]
    public partial class MnqNatsCredentials : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the NATS account the credentials are generated from
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The content of the credentials file.
        /// </summary>
        [Output("file")]
        public Output<string> File { get; private set; } = null!;

        /// <summary>
        /// The unique name of the NATS credentials.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which the account exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a MnqNatsCredentials resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqNatsCredentials(string name, MnqNatsCredentialsArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqNatsCredentials:MnqNatsCredentials", name, args ?? new MnqNatsCredentialsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqNatsCredentials(string name, Input<string> id, MnqNatsCredentialsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqNatsCredentials:MnqNatsCredentials", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MnqNatsCredentials resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqNatsCredentials Get(string name, Input<string> id, MnqNatsCredentialsState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqNatsCredentials(name, id, state, options);
        }
    }

    public sealed class MnqNatsCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the NATS account the credentials are generated from
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The unique name of the NATS credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which the account exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqNatsCredentialsArgs()
        {
        }
        public static new MnqNatsCredentialsArgs Empty => new MnqNatsCredentialsArgs();
    }

    public sealed class MnqNatsCredentialsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the NATS account the credentials are generated from
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The content of the credentials file.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// The unique name of the NATS credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which the account exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqNatsCredentialsState()
        {
        }
        public static new MnqNatsCredentialsState Empty => new MnqNatsCredentialsState();
    }
}
