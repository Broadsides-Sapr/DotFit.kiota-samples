package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

// ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\attachments
type ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderGetQueryParameters retrieve a list of attachment objects.
type ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderGetQueryParameters
}
// ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAttachmentId gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.childFolders.item.messages.item.attachments.item collection
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) ByAttachmentId(attachmentId string)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attachmentId != "" {
        urlTplParams["attachment%2Did"] = attachmentId
    }
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderInternal instantiates a new AttachmentsRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/attachments{?%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder instantiates a new AttachmentsRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) Count()(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of attachment objects.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/eventmessage-list-attachments?view=graph-rest-1.0
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.AttachmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateAttachmentCollectionResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.AttachmentCollectionResponseable), nil
}
// Post use this API to add an attachment to a message.  An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource.  You can add an attachment to an existing message by posting to its attachments collection, or you can add an attachment to a message that is being created and sent on the fly. This operation limits the size of the attachment you can add to under 3 MB.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-post-attachments?view=graph-rest-1.0
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) Post(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachmentable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateAttachmentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachmentable), nil
}
// ToGetRequestInformation retrieve a list of attachment objects.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation use this API to add an attachment to a message.  An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource.  You can add an attachment to an existing message by posting to its attachments collection, or you can add an attachment to a message that is being created and sent on the fly. This operation limits the size of the attachment you can add to under 3 MB.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachmentable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
