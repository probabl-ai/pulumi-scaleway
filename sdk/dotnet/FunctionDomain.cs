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
    /// Creates and manages Scaleway Function Domain bindings.
    /// For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-functions).
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
    ///     var mainFunctionNamespace = new Scaleway.FunctionNamespace("mainFunctionNamespace");
    /// 
    ///     var mainFunction = new Scaleway.Function("mainFunction", new()
    ///     {
    ///         NamespaceId = mainFunctionNamespace.Id,
    ///         Runtime = "go118",
    ///         Privacy = "private",
    ///         Handler = "Handle",
    ///         ZipFile = "testfixture/gofunction.zip",
    ///         Deploy = true,
    ///     });
    /// 
    ///     var mainFunctionDomain = new Scaleway.FunctionDomain("mainFunctionDomain", new()
    ///     {
    ///         FunctionId = mainFunction.Id,
    ///         Hostname = "example.com",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             mainFunction,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Domain can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/functionDomain:FunctionDomain main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/functionDomain:FunctionDomain")]
    public partial class FunctionDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the function you want to create a domain with.
        /// </summary>
        [Output("functionId")]
        public Output<string> FunctionId { get; private set; } = null!;

        /// <summary>
        /// The hostname that should resolve to your function id native domain.
        /// You should use a CNAME domain record that point to your native function `domain_name` for it.
        /// 
        /// &gt; **Important** Updates to `function_id` or `hostname` will recreate the domain.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// (Defaults to provider `region`) The region in where the domain was created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URL that triggers the function
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionDomain(string name, FunctionDomainArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/functionDomain:FunctionDomain", name, args ?? new FunctionDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionDomain(string name, Input<string> id, FunctionDomainState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/functionDomain:FunctionDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionDomain Get(string name, Input<string> id, FunctionDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionDomain(name, id, state, options);
        }
    }

    public sealed class FunctionDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the function you want to create a domain with.
        /// </summary>
        [Input("functionId", required: true)]
        public Input<string> FunctionId { get; set; } = null!;

        /// <summary>
        /// The hostname that should resolve to your function id native domain.
        /// You should use a CNAME domain record that point to your native function `domain_name` for it.
        /// 
        /// &gt; **Important** Updates to `function_id` or `hostname` will recreate the domain.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// (Defaults to provider `region`) The region in where the domain was created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public FunctionDomainArgs()
        {
        }
        public static new FunctionDomainArgs Empty => new FunctionDomainArgs();
    }

    public sealed class FunctionDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the function you want to create a domain with.
        /// </summary>
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        /// <summary>
        /// The hostname that should resolve to your function id native domain.
        /// You should use a CNAME domain record that point to your native function `domain_name` for it.
        /// 
        /// &gt; **Important** Updates to `function_id` or `hostname` will recreate the domain.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// (Defaults to provider `region`) The region in where the domain was created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URL that triggers the function
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public FunctionDomainState()
        {
        }
        public static new FunctionDomainState Empty => new FunctionDomainState();
    }
}
