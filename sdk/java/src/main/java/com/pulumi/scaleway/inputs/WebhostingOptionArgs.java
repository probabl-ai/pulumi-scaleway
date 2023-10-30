// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WebhostingOptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final WebhostingOptionArgs Empty = new WebhostingOptionArgs();

    /**
     * The option ID.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The option ID.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The option name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The option name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private WebhostingOptionArgs() {}

    private WebhostingOptionArgs(WebhostingOptionArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebhostingOptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebhostingOptionArgs $;

        public Builder() {
            $ = new WebhostingOptionArgs();
        }

        public Builder(WebhostingOptionArgs defaults) {
            $ = new WebhostingOptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The option ID.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The option ID.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param name The option name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The option name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public WebhostingOptionArgs build() {
            return $;
        }
    }

}
