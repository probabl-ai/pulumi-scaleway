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
    public static class GetIamSshKey
    {
        /// <summary>
        /// Use this data source to get SSH key information based on its ID or name.
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
        ///     var myKey = Scaleway.GetIamSshKey.Invoke(new()
        ///     {
        ///         SshKeyId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIamSshKeyResult> InvokeAsync(GetIamSshKeyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIamSshKeyResult>("scaleway:index/getIamSshKey:getIamSshKey", args ?? new GetIamSshKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get SSH key information based on its ID or name.
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
        ///     var myKey = Scaleway.GetIamSshKey.Invoke(new()
        ///     {
        ///         SshKeyId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIamSshKeyResult> Invoke(GetIamSshKeyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIamSshKeyResult>("scaleway:index/getIamSshKey:getIamSshKey", args ?? new GetIamSshKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIamSshKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The SSH key name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the SSH
        /// key is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The SSH key id.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `ssh_key_id`.
        /// </summary>
        [Input("sshKeyId")]
        public string? SshKeyId { get; set; }

        public GetIamSshKeyArgs()
        {
        }
        public static new GetIamSshKeyArgs Empty => new GetIamSshKeyArgs();
    }

    public sealed class GetIamSshKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The SSH key name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the SSH
        /// key is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The SSH key id.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `ssh_key_id`.
        /// </summary>
        [Input("sshKeyId")]
        public Input<string>? SshKeyId { get; set; }

        public GetIamSshKeyInvokeArgs()
        {
        }
        public static new GetIamSshKeyInvokeArgs Empty => new GetIamSshKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetIamSshKeyResult
    {
        /// <summary>
        /// The date and time of the creation of the SSH key.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The SSH key status.
        /// </summary>
        public readonly bool Disabled;
        public readonly string Fingerprint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The ID of the organization the SSH key is associated with.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string? ProjectId;
        /// <summary>
        /// The SSH public key string
        /// </summary>
        public readonly string PublicKey;
        public readonly string? SshKeyId;
        /// <summary>
        /// The date and time of the last update of the SSH key.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetIamSshKeyResult(
            string createdAt,

            bool disabled,

            string fingerprint,

            string id,

            string? name,

            string organizationId,

            string? projectId,

            string publicKey,

            string? sshKeyId,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Disabled = disabled;
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            PublicKey = publicKey;
            SshKeyId = sshKeyId;
            UpdatedAt = updatedAt;
        }
    }
}
