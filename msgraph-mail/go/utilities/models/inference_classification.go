package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InferenceClassification struct {
    Entity
    overrides []InferenceClassificationOverride;
}
func NewInferenceClassification()(*InferenceClassification) {
    m := &InferenceClassification{
        Entity: *NewEntity(),
    }
    return m
}
func (m *InferenceClassification) GetOverrides()([]InferenceClassificationOverride) {
    return m.overrides
}
func (m *InferenceClassification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["overrides"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewInferenceClassificationOverride() })
        if err != nil {
            return err
        }
        res := make([]InferenceClassificationOverride, len(val))
        for i, v := range val {
            res[i] = *(v.(*InferenceClassificationOverride))
        }
        o.(*InferenceClassification).SetOverrides(res)
        return nil
    }
    return res
}
func (m *InferenceClassification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    err = writer.WriteCollectionOfObjectValues("overrides", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetOverrides()))
    if err != nil {
        return err
    }
    return nil
}
func (m *InferenceClassification) SetOverrides(value []InferenceClassificationOverride)() {
    m.overrides = value
}
