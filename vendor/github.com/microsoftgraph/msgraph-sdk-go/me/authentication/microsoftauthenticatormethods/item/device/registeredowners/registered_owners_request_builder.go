package registeredowners

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i150b9f1af48ddb235685b77fb1b00d358e372810bd9f27029550ad39abd1d33d "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/approleassignment"
    i1eaeb4bcdb0e9babbddcf8b9bfbcf8adb9c0ca55a607c6f4051ac9c03334cf06 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/user"
    i2ba099f283d6d4ae5be3f4a85570c156a4a0a6390f739d778c9f3d71725a1bdd "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/ref"
    i6780186fccc1c7882c6ce38c73f3d809888fc02dd75fe76ed23ff63667abdb87 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/count"
    ibb6ca7a85e75ff7f533805c8818710c5bd0a345fb3048f2d43fe5f776cc08e89 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/endpoint"
    idba5dc65de5625db79712d9c51616d6b6471f92cf19e54550e4dfced5cb3e9bb "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/serviceprincipal"
)

// RegisteredOwnersRequestBuilder provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
type RegisteredOwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RegisteredOwnersRequestBuilderGetQueryParameters the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
type RegisteredOwnersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// RegisteredOwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RegisteredOwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RegisteredOwnersRequestBuilderGetQueryParameters
}
// AppRoleAssignment the appRoleAssignment property
func (m *RegisteredOwnersRequestBuilder) AppRoleAssignment()(*i150b9f1af48ddb235685b77fb1b00d358e372810bd9f27029550ad39abd1d33d.AppRoleAssignmentRequestBuilder) {
    return i150b9f1af48ddb235685b77fb1b00d358e372810bd9f27029550ad39abd1d33d.NewAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRegisteredOwnersRequestBuilderInternal instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewRegisteredOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredOwnersRequestBuilder) {
    m := &RegisteredOwnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRegisteredOwnersRequestBuilder instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewRegisteredOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredOwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegisteredOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RegisteredOwnersRequestBuilder) Count()(*i6780186fccc1c7882c6ce38c73f3d809888fc02dd75fe76ed23ff63667abdb87.CountRequestBuilder) {
    return i6780186fccc1c7882c6ce38c73f3d809888fc02dd75fe76ed23ff63667abdb87.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *RegisteredOwnersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RegisteredOwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Endpoint the endpoint property
func (m *RegisteredOwnersRequestBuilder) Endpoint()(*ibb6ca7a85e75ff7f533805c8818710c5bd0a345fb3048f2d43fe5f776cc08e89.EndpointRequestBuilder) {
    return ibb6ca7a85e75ff7f533805c8818710c5bd0a345fb3048f2d43fe5f776cc08e89.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *RegisteredOwnersRequestBuilder) Get(ctx context.Context, requestConfiguration *RegisteredOwnersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// Ref the Ref property
func (m *RegisteredOwnersRequestBuilder) Ref()(*i2ba099f283d6d4ae5be3f4a85570c156a4a0a6390f739d778c9f3d71725a1bdd.RefRequestBuilder) {
    return i2ba099f283d6d4ae5be3f4a85570c156a4a0a6390f739d778c9f3d71725a1bdd.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *RegisteredOwnersRequestBuilder) ServicePrincipal()(*idba5dc65de5625db79712d9c51616d6b6471f92cf19e54550e4dfced5cb3e9bb.ServicePrincipalRequestBuilder) {
    return idba5dc65de5625db79712d9c51616d6b6471f92cf19e54550e4dfced5cb3e9bb.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *RegisteredOwnersRequestBuilder) User()(*i1eaeb4bcdb0e9babbddcf8b9bfbcf8adb9c0ca55a607c6f4051ac9c03334cf06.UserRequestBuilder) {
    return i1eaeb4bcdb0e9babbddcf8b9bfbcf8adb9c0ca55a607c6f4051ac9c03334cf06.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}