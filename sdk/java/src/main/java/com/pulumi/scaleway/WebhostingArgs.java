// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WebhostingArgs extends com.pulumi.resources.ResourceArgs {

    public static final WebhostingArgs Empty = new WebhostingArgs();

    /**
     * The domain name of the hosting.
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return The domain name of the hosting.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    /**
     * The contact email of the client for the hosting.
     * 
     */
    @Import(name="email", required=true)
    private Output<String> email;

    /**
     * @return The contact email of the client for the hosting.
     * 
     */
    public Output<String> email() {
        return this.email;
    }

    /**
     * The ID of the selected offer for the hosting.
     * 
     */
    @Import(name="offerId", required=true)
    private Output<String> offerId;

    /**
     * @return The ID of the selected offer for the hosting.
     * 
     */
    public Output<String> offerId() {
        return this.offerId;
    }

    /**
     * The IDs of the selected options for the hosting.
     * 
     */
    @Import(name="optionIds")
    private @Nullable Output<List<String>> optionIds;

    /**
     * @return The IDs of the selected options for the hosting.
     * 
     */
    public Optional<Output<List<String>>> optionIds() {
        return Optional.ofNullable(this.optionIds);
    }

    /**
     * `project_id`) The ID of the project the VPC is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the VPC is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`) The region of the Hosting.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region of the Hosting.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The tags associated with the hosting.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags associated with the hosting.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private WebhostingArgs() {}

    private WebhostingArgs(WebhostingArgs $) {
        this.domain = $.domain;
        this.email = $.email;
        this.offerId = $.offerId;
        this.optionIds = $.optionIds;
        this.projectId = $.projectId;
        this.region = $.region;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebhostingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebhostingArgs $;

        public Builder() {
            $ = new WebhostingArgs();
        }

        public Builder(WebhostingArgs defaults) {
            $ = new WebhostingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The domain name of the hosting.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The domain name of the hosting.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param email The contact email of the client for the hosting.
         * 
         * @return builder
         * 
         */
        public Builder email(Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The contact email of the client for the hosting.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param offerId The ID of the selected offer for the hosting.
         * 
         * @return builder
         * 
         */
        public Builder offerId(Output<String> offerId) {
            $.offerId = offerId;
            return this;
        }

        /**
         * @param offerId The ID of the selected offer for the hosting.
         * 
         * @return builder
         * 
         */
        public Builder offerId(String offerId) {
            return offerId(Output.of(offerId));
        }

        /**
         * @param optionIds The IDs of the selected options for the hosting.
         * 
         * @return builder
         * 
         */
        public Builder optionIds(@Nullable Output<List<String>> optionIds) {
            $.optionIds = optionIds;
            return this;
        }

        /**
         * @param optionIds The IDs of the selected options for the hosting.
         * 
         * @return builder
         * 
         */
        public Builder optionIds(List<String> optionIds) {
            return optionIds(Output.of(optionIds));
        }

        /**
         * @param optionIds The IDs of the selected options for the hosting.
         * 
         * @return builder
         * 
         */
        public Builder optionIds(String... optionIds) {
            return optionIds(List.of(optionIds));
        }

        /**
         * @param projectId `project_id`) The ID of the project the VPC is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the VPC is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region `region`) The region of the Hosting.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region of the Hosting.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tags The tags associated with the hosting.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags associated with the hosting.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags associated with the hosting.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public WebhostingArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.email = Objects.requireNonNull($.email, "expected parameter 'email' to be non-null");
            $.offerId = Objects.requireNonNull($.offerId, "expected parameter 'offerId' to be non-null");
            return $;
        }
    }

}
