// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class IamPolicyRule
    {
        /// <summary>
        /// ID of organization scoped to the rule.
        /// </summary>
        public readonly string? OrganizationId;
        /// <summary>
        /// Names of permission sets bound to the rule.
        /// 
        /// **_TIP:_**  You can use the Scaleway CLI to list the permissions details. e.g:
        /// 
        /// ```shell
        /// $ scw iam permission-set list
        /// ```
        /// </summary>
        public readonly ImmutableArray<string> PermissionSetNames;
        /// <summary>
        /// List of project IDs scoped to the rule.
        /// 
        /// &gt; **Important** One of `organization_id` or `project_ids`  must be set per rule.
        /// </summary>
        public readonly ImmutableArray<string> ProjectIds;

        [OutputConstructor]
        private IamPolicyRule(
            string? organizationId,

            ImmutableArray<string> permissionSetNames,

            ImmutableArray<string> projectIds)
        {
            OrganizationId = organizationId;
            PermissionSetNames = permissionSetNames;
            ProjectIds = projectIds;
        }
    }
}
