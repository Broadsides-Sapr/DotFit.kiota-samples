package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\extensions\$count
type ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
}
// ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderGetQueryParameters
}
// NewItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder) {
    m := &ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/extensions/$count{?%24filter}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder) {
    return NewItemMailFoldersItemMessagesItemExtensionsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}