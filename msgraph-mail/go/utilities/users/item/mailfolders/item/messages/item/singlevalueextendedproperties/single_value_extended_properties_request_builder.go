package singlevalueextendedproperties

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

type SingleValueExtendedPropertiesRequestBuilder struct {
    currentPath string;
    isRawUrl bool;
    pathSegment string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
}
type SingleValueExtendedPropertiesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewSingleValueExtendedPropertiesRequestBuilder(currentPath string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, isRawUrl bool)(*SingleValueExtendedPropertiesRequestBuilder) {
    m := &SingleValueExtendedPropertiesRequestBuilder{
    }
    m.pathSegment = "/singleValueExtendedProperties";
    m.requestAdapter = requestAdapter;
    m.currentPath = currentPath;
    m.isRawUrl = isRawUrl;
    return m
}
func (m *SingleValueExtendedPropertiesRequestBuilder) CreateGetRequestInformation(q func (value *SingleValueExtendedPropertiesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if err != nil {
        return nil, err
    }
    if q != nil {
        qParams := new(SingleValueExtendedPropertiesRequestBuilderGetQueryParameters)
        err = q(qParams)
        if err != nil {
            return nil, err
        }
        err := qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err = h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err = requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *SingleValueExtendedPropertiesRequestBuilder) CreatePostRequestInformation(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.SingleValueLegacyExtendedProperty, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    if err != nil {
        return nil, err
    }
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err = h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err = requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *SingleValueExtendedPropertiesRequestBuilder) Get(q func (value *SingleValueExtendedPropertiesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*SingleValueExtendedPropertiesResponse, error)) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return func() (*SingleValueExtendedPropertiesResponse, error) { return nil, err }
    }
    return func() (*SingleValueExtendedPropertiesResponse, error) {
        res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(SingleValueExtendedPropertiesResponse) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*SingleValueExtendedPropertiesResponse), nil
    }
}
func (m *SingleValueExtendedPropertiesRequestBuilder) Post(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.SingleValueLegacyExtendedProperty, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.SingleValueLegacyExtendedProperty, error)) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.SingleValueLegacyExtendedProperty, error) { return nil, err }
    }
    return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.SingleValueLegacyExtendedProperty, error) {
        res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.SingleValueLegacyExtendedProperty) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.SingleValueLegacyExtendedProperty), nil
    }
}
