package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    i64ecde285b206ca053eecbd6e7f197e884287a3d0eaa42aae2406e7b3007eba1 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/odataerrors"
)

// ItemMailFoldersItemMessageRulesRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messageRules
type ItemMailFoldersItemMessageRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemMessageRulesRequestBuilderGetQueryParameters get all the messageRule objects defined for the user's inbox.
type ItemMailFoldersItemMessageRulesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemMessageRulesRequestBuilderGetQueryParameters
}
// ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMessageRuleId gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.messageRules.item collection
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) ByMessageRuleId(messageRuleId string)(*ItemMailFoldersItemMessageRulesMessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if messageRuleId != "" {
        urlTplParams["messageRule%2Did"] = messageRuleId
    }
    return NewItemMailFoldersItemMessageRulesMessageRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailFoldersItemMessageRulesRequestBuilderInternal instantiates a new MessageRulesRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessageRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessageRulesRequestBuilder) {
    m := &ItemMailFoldersItemMessageRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messageRules{?%24count,%24filter,%24orderby,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemMessageRulesRequestBuilder instantiates a new MessageRulesRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessageRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessageRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemMessageRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) Count()(*ItemMailFoldersItemMessageRulesCountRequestBuilder) {
    return NewItemMailFoldersItemMessageRulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get all the messageRule objects defined for the user's inbox.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailfolder-list-messagerules?view=graph-rest-1.0
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i64ecde285b206ca053eecbd6e7f197e884287a3d0eaa42aae2406e7b3007eba1.CreateODataErrorFromDiscriminatorValue,
        "5XX": i64ecde285b206ca053eecbd6e7f197e884287a3d0eaa42aae2406e7b3007eba1.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMessageRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleCollectionResponseable), nil
}
// Post create a messageRule object by specifying a set of conditions and actions. Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailfolder-post-messagerules?view=graph-rest-1.0
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) Post(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i64ecde285b206ca053eecbd6e7f197e884287a3d0eaa42aae2406e7b3007eba1.CreateODataErrorFromDiscriminatorValue,
        "5XX": i64ecde285b206ca053eecbd6e7f197e884287a3d0eaa42aae2406e7b3007eba1.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMessageRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable), nil
}
// ToGetRequestInformation get all the messageRule objects defined for the user's inbox.
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a messageRule object by specifying a set of conditions and actions. Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable, requestConfiguration *ItemMailFoldersItemMessageRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMailFoldersItemMessageRulesRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemMessageRulesRequestBuilder) {
    return NewItemMailFoldersItemMessageRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
