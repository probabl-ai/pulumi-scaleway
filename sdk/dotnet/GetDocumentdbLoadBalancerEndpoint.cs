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
    public static class GetDocumentdbLoadBalancerEndpoint
    {
        /// <summary>
        /// Gets information about an DocumentDB load balancer endpoint.
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
        ///     var myEndpoint = Scaleway.GetDocumentdbLoadBalancerEndpoint.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDocumentdbLoadBalancerEndpointResult> InvokeAsync(GetDocumentdbLoadBalancerEndpointArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDocumentdbLoadBalancerEndpointResult>("scaleway:index/getDocumentdbLoadBalancerEndpoint:getDocumentdbLoadBalancerEndpoint", args ?? new GetDocumentdbLoadBalancerEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an DocumentDB load balancer endpoint.
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
        ///     var myEndpoint = Scaleway.GetDocumentdbLoadBalancerEndpoint.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDocumentdbLoadBalancerEndpointResult> Invoke(GetDocumentdbLoadBalancerEndpointInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDocumentdbLoadBalancerEndpointResult>("scaleway:index/getDocumentdbLoadBalancerEndpoint:getDocumentdbLoadBalancerEndpoint", args ?? new GetDocumentdbLoadBalancerEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDocumentdbLoadBalancerEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DocumentDB Instance on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The DocumentDB Instance Name on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceName")]
        public string? InstanceName { get; set; }

        /// <summary>
        /// The ID of the project the DocumentDB endpoint is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the DocumentDB endpoint exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetDocumentdbLoadBalancerEndpointArgs()
        {
        }
        public static new GetDocumentdbLoadBalancerEndpointArgs Empty => new GetDocumentdbLoadBalancerEndpointArgs();
    }

    public sealed class GetDocumentdbLoadBalancerEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DocumentDB Instance on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The DocumentDB Instance Name on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The ID of the project the DocumentDB endpoint is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the DocumentDB endpoint exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetDocumentdbLoadBalancerEndpointInvokeArgs()
        {
        }
        public static new GetDocumentdbLoadBalancerEndpointInvokeArgs Empty => new GetDocumentdbLoadBalancerEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetDocumentdbLoadBalancerEndpointResult
    {
        /// <summary>
        /// The hostname of your endpoint.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string InstanceName;
        /// <summary>
        /// The IP of your load balancer service.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The name of your load balancer service.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port of your load balancer service.
        /// </summary>
        public readonly int Port;
        public readonly string ProjectId;
        public readonly string Region;

        [OutputConstructor]
        private GetDocumentdbLoadBalancerEndpointResult(
            string hostname,

            string id,

            string instanceId,

            string instanceName,

            string ip,

            string name,

            int port,

            string projectId,

            string region)
        {
            Hostname = hostname;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            Ip = ip;
            Name = name;
            Port = port;
            ProjectId = projectId;
            Region = region;
        }
    }
}
