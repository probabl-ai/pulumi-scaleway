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
    public static class GetAccountProject
    {
        /// <summary>
        /// Gets information about an existing Project.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetAccountProject.Invoke(new()
        ///     {
        ///         Name = "default",
        ///     });
        /// 
        ///     var byId = Scaleway.GetAccountProject.Invoke(new()
        ///     {
        ///         ProjectId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAccountProjectResult> InvokeAsync(GetAccountProjectArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountProjectResult>("scaleway:index/getAccountProject:getAccountProject", args ?? new GetAccountProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an existing Project.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetAccountProject.Invoke(new()
        ///     {
        ///         Name = "default",
        ///     });
        /// 
        ///     var byId = Scaleway.GetAccountProject.Invoke(new()
        ///     {
        ///         ProjectId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAccountProjectResult> Invoke(GetAccountProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountProjectResult>("scaleway:index/getAccountProject:getAccountProject", args ?? new GetAccountProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountProjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Project.
        /// Only one of the `name` and `project_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The organization ID the Project is associated with.
        /// If no default organization_id is set, one must be set explicitly in this datasource
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        /// <summary>
        /// The ID of the Project.
        /// Only one of the `name` and `project_id` should be specified.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetAccountProjectArgs()
        {
        }
        public static new GetAccountProjectArgs Empty => new GetAccountProjectArgs();
    }

    public sealed class GetAccountProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Project.
        /// Only one of the `name` and `project_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the Project is associated with.
        /// If no default organization_id is set, one must be set explicitly in this datasource
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The ID of the Project.
        /// Only one of the `name` and `project_id` should be specified.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetAccountProjectInvokeArgs()
        {
        }
        public static new GetAccountProjectInvokeArgs Empty => new GetAccountProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountProjectResult
    {
        public readonly string CreatedAt;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? OrganizationId;
        public readonly string ProjectId;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetAccountProjectResult(
            string createdAt,

            string description,

            string id,

            string? name,

            string? organizationId,

            string projectId,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            UpdatedAt = updatedAt;
        }
    }
}
