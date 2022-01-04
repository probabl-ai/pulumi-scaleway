// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Scaleway
{
    public static class GetInstanceSecurityGroup
    {
        /// <summary>
        /// Gets information about a Security Group.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myKey = Output.Create(Scaleway.GetInstanceSecurityGroup.InvokeAsync(new Scaleway.GetInstanceSecurityGroupArgs
        ///         {
        ///             SecurityGroupId = "11111111-1111-1111-1111-111111111111",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceSecurityGroupResult> InvokeAsync(GetInstanceSecurityGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceSecurityGroupResult>("scaleway:index/getInstanceSecurityGroup:getInstanceSecurityGroup", args ?? new GetInstanceSecurityGroupArgs(), options.WithVersion());

        /// <summary>
        /// Gets information about a Security Group.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myKey = Output.Create(Scaleway.GetInstanceSecurityGroup.InvokeAsync(new Scaleway.GetInstanceSecurityGroupArgs
        ///         {
        ///             SecurityGroupId = "11111111-1111-1111-1111-111111111111",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceSecurityGroupResult> Invoke(GetInstanceSecurityGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceSecurityGroupResult>("scaleway:index/getInstanceSecurityGroup:getInstanceSecurityGroup", args ?? new GetInstanceSecurityGroupInvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstanceSecurityGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The security group name. Only one of `name` and `security_group_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The security group id. Only one of `name` and `security_group_id` should be specified.
        /// </summary>
        [Input("securityGroupId")]
        public string? SecurityGroupId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the security group exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstanceSecurityGroupArgs()
        {
        }
    }

    public sealed class GetInstanceSecurityGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The security group name. Only one of `name` and `security_group_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The security group id. Only one of `name` and `security_group_id` should be specified.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the security group exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetInstanceSecurityGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceSecurityGroupResult
    {
        public readonly string Description;
        public readonly bool EnableDefaultSecurity;
        public readonly bool ExternalRules;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        public readonly string InboundDefaultPolicy;
        /// <summary>
        /// A list of inbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceSecurityGroupInboundRuleResult> InboundRules;
        public readonly string? Name;
        /// <summary>
        /// The ID of the organization the security group is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        public readonly string OutboundDefaultPolicy;
        /// <summary>
        /// A list of outbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceSecurityGroupOutboundRuleResult> OutboundRules;
        /// <summary>
        /// The ID of the project the security group is associated with.
        /// </summary>
        public readonly string ProjectId;
        public readonly string? SecurityGroupId;
        public readonly bool Stateful;
        public readonly string? Zone;

        [OutputConstructor]
        private GetInstanceSecurityGroupResult(
            string description,

            bool enableDefaultSecurity,

            bool externalRules,

            string id,

            string inboundDefaultPolicy,

            ImmutableArray<Outputs.GetInstanceSecurityGroupInboundRuleResult> inboundRules,

            string? name,

            string organizationId,

            string outboundDefaultPolicy,

            ImmutableArray<Outputs.GetInstanceSecurityGroupOutboundRuleResult> outboundRules,

            string projectId,

            string? securityGroupId,

            bool stateful,

            string? zone)
        {
            Description = description;
            EnableDefaultSecurity = enableDefaultSecurity;
            ExternalRules = externalRules;
            Id = id;
            InboundDefaultPolicy = inboundDefaultPolicy;
            InboundRules = inboundRules;
            Name = name;
            OrganizationId = organizationId;
            OutboundDefaultPolicy = outboundDefaultPolicy;
            OutboundRules = outboundRules;
            ProjectId = projectId;
            SecurityGroupId = securityGroupId;
            Stateful = stateful;
            Zone = zone;
        }
    }
}
