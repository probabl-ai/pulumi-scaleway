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
    /// Activate Scaleway Messaging and queuing SNS for a project.
    /// For further information please check
    /// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// Activate SNS for default project
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.MnqSns("main");
    /// 
    /// });
    /// ```
    /// 
    /// Activate SNS for a specific project
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = Scaleway.GetAccountProject.Invoke(new()
    ///     {
    ///         Name = "default",
    ///     });
    /// 
    ///     // For specific project in default region
    ///     var forProject = new Scaleway.MnqSns("forProject", new()
    ///     {
    ///         ProjectId = project.Apply(getAccountProjectResult =&gt; getAccountProjectResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SNS status can be imported using the `{region}/{project_id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/mnqSns:MnqSns main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mnqSns:MnqSns")]
    public partial class MnqSns : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint of the SNS service for this project.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the sns will be enabled for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which sns will be enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a MnqSns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqSns(string name, MnqSnsArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSns:MnqSns", name, args ?? new MnqSnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqSns(string name, Input<string> id, MnqSnsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSns:MnqSns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MnqSns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqSns Get(string name, Input<string> id, MnqSnsState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqSns(name, id, state, options);
        }
    }

    public sealed class MnqSnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the sns will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sns will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqSnsArgs()
        {
        }
        public static new MnqSnsArgs Empty => new MnqSnsArgs();
    }

    public sealed class MnqSnsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint of the SNS service for this project.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sns will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sns will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqSnsState()
        {
        }
        public static new MnqSnsState Empty => new MnqSnsState();
    }
}
