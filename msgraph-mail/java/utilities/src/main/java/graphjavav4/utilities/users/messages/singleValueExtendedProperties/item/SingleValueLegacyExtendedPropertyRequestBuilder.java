package graphjavav4.utilities.users.messages.singleValueExtendedProperties.item;

import java.util.Objects;
import graphjavav4.utilities.users.SingleValueLegacyExtendedProperty;
import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.QueryParametersBase;
import java.util.Map;
import java.net.URI;
import java.net.URISyntaxException;
public class SingleValueLegacyExtendedPropertyRequestBuilder {
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<SingleValueLegacyExtendedProperty> get(@javax.annotation.Nonnull final java.util.function.Consumer<GetQueryParameters> q, @javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(q);
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            final GetQueryParameters qParams = new GetQueryParameters();
            q.accept(qParams);
            qParams.AddQueryParameters(requestInfo.queryParameters);
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> patch(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.PATCH;
            }};
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> delete(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.DELETE;
            }};
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nonnull
    private final String pathSegment = "";
    @javax.annotation.Nullable
    public String currentPath;
    @javax.annotation.Nullable
    public HttpCore httpCore;
    public class GetQueryParameters extends QueryParametersBase {
        @javax.annotation.Nullable
        public String[] select;
        @javax.annotation.Nullable
        public String[] expand;
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<SingleValueLegacyExtendedProperty> get(@javax.annotation.Nonnull final java.util.function.Consumer<Map<String, String>> h, @javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(h);
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            h.accept(requestInfo.headers);
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<SingleValueLegacyExtendedProperty> get(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> patch(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.PATCH;
            }};
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> delete(@javax.annotation.Nonnull final ResponseHandler responseHandler) {
        Objects.requireNonNull(responseHandler);
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.DELETE;
            }};
            return this.httpCore.sendAsync(requestInfo, responseHandler);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<SingleValueLegacyExtendedProperty> get() {
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.GET;
            }};
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> patch() {
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.PATCH;
            }};
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    @javax.annotation.Nullable
    public java.util.concurrent.CompletableFuture<Object> delete() {
        try {
            final RequestInfo requestInfo = new RequestInfo() {{
                uri = new URI(currentPath);
                httpMethod = HttpMethod.DELETE;
            }};
            return this.httpCore.sendAsync(requestInfo, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
}
