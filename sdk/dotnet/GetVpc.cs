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
    public static class GetVpc
    {
        /// <summary>
        /// Gets information about a Scaleway Virtual Private Cloud.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetVpc.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.GetVpc.Invoke(new()
        ///     {
        ///         VpcId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var @default = Scaleway.GetVpc.Invoke(new()
        ///     {
        ///         IsDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVpcResult> InvokeAsync(GetVpcArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcResult>("scaleway:index/getVpc:getVpc", args ?? new GetVpcArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Scaleway Virtual Private Cloud.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetVpc.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.GetVpc.Invoke(new()
        ///     {
        ///         VpcId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var @default = Scaleway.GetVpc.Invoke(new()
        ///     {
        ///         IsDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVpcResult> Invoke(GetVpcInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcResult>("scaleway:index/getVpc:getVpc", args ?? new GetVpcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether the targeted VPC is the default VPC.
        /// </summary>
        [Input("isDefault")]
        public bool? IsDefault { get; set; }

        /// <summary>
        /// Name of the VPC. A maximum of one of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the Organization the VPC is associated with.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// ID of the VPC. A maximum of one of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVpcArgs()
        {
        }
        public static new GetVpcArgs Empty => new GetVpcArgs();
    }

    public sealed class GetVpcInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether the targeted VPC is the default VPC.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Name of the VPC. A maximum of one of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Organization the VPC is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// ID of the VPC. A maximum of one of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetVpcInvokeArgs()
        {
        }
        public static new GetVpcInvokeArgs Empty => new GetVpcInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcResult
    {
        public readonly string CreatedAt;
        public readonly bool EnableRouting;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IsDefault;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string? ProjectId;
        public readonly string? Region;
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdatedAt;
        public readonly string? VpcId;

        [OutputConstructor]
        private GetVpcResult(
            string createdAt,

            bool enableRouting,

            string id,

            bool? isDefault,

            string? name,

            string organizationId,

            string? projectId,

            string? region,

            ImmutableArray<string> tags,

            string updatedAt,

            string? vpcId)
        {
            CreatedAt = createdAt;
            EnableRouting = enableRouting;
            Id = id;
            IsDefault = isDefault;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Tags = tags;
            UpdatedAt = updatedAt;
            VpcId = vpcId;
        }
    }
}
