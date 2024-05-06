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
    public static class GetMnqSqs
    {
        /// <summary>
        /// Gets information about SQS for a project
        /// 
        /// ## Examples
        /// 
        /// ### Basic
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Scaleway.GetMnqSqs.Invoke();
        /// 
        ///     var forProject = Scaleway.GetMnqSqs.Invoke(new()
        ///     {
        ///         ProjectId = scaleway_account_project.Main.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetMnqSqsResult> InvokeAsync(GetMnqSqsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMnqSqsResult>("scaleway:index/getMnqSqs:getMnqSqs", args ?? new GetMnqSqsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about SQS for a project
        /// 
        /// ## Examples
        /// 
        /// ### Basic
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Scaleway.GetMnqSqs.Invoke();
        /// 
        ///     var forProject = Scaleway.GetMnqSqs.Invoke(new()
        ///     {
        ///         ProjectId = scaleway_account_project.Main.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetMnqSqsResult> Invoke(GetMnqSqsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMnqSqsResult>("scaleway:index/getMnqSqs:getMnqSqs", args ?? new GetMnqSqsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMnqSqsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project for which sqs is enabled.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which sqs is enabled.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetMnqSqsArgs()
        {
        }
        public static new GetMnqSqsArgs Empty => new GetMnqSqsArgs();
    }

    public sealed class GetMnqSqsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project for which sqs is enabled.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which sqs is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetMnqSqsInvokeArgs()
        {
        }
        public static new GetMnqSqsInvokeArgs Empty => new GetMnqSqsInvokeArgs();
    }


    [OutputType]
    public sealed class GetMnqSqsResult
    {
        /// <summary>
        /// The endpoint of the SQS service for this project.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ProjectId;
        public readonly string? Region;

        [OutputConstructor]
        private GetMnqSqsResult(
            string endpoint,

            string id,

            string? projectId,

            string? region)
        {
            Endpoint = endpoint;
            Id = id;
            ProjectId = projectId;
            Region = region;
        }
    }
}
