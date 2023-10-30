// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway
{
    public static class GetVpcs
    {
        /// <summary>
        /// Gets information about multiple Virtual Private Clouds.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetVpcs.Invoke(new()
        ///     {
        ///         Name = "tf-vpc-datasource",
        ///         Region = "nl-ams",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcsResult> InvokeAsync(GetVpcsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcsResult>("scaleway:index/getVpcs:getVpcs", args ?? new GetVpcsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Virtual Private Clouds.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetVpcs.Invoke(new()
        ///     {
        ///         Name = "tf-vpc-datasource",
        ///         Region = "nl-ams",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpcsResult> Invoke(GetVpcsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcsResult>("scaleway:index/getVpcs:getVpcs", args ?? new GetVpcsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The VPC name used as filter. VPCs with a name like it are listed.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which vpcs exist.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// List of tags used as filter. VPCs with these exact tags are listed.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public GetVpcsArgs()
        {
        }
        public static new GetVpcsArgs Empty => new GetVpcsArgs();
    }

    public sealed class GetVpcsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The VPC name used as filter. VPCs with a name like it are listed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which vpcs exist.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags used as filter. VPCs with these exact tags are listed.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public GetVpcsInvokeArgs()
        {
        }
        public static new GetVpcsInvokeArgs Empty => new GetVpcsInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The organization ID the VPC is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The ID of the project the VPC is associated with.
        /// </summary>
        public readonly string ProjectId;
        public readonly string Region;
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// List of found vpcs
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcsVpcResult> Vpcs;

        [OutputConstructor]
        private GetVpcsResult(
            string id,

            string? name,

            string organizationId,

            string projectId,

            string region,

            ImmutableArray<string> tags,

            ImmutableArray<Outputs.GetVpcsVpcResult> vpcs)
        {
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Tags = tags;
            Vpcs = vpcs;
        }
    }
}
