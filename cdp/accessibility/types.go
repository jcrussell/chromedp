package accessibility

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AXNodeID unique accessibility node identifier.
type AXNodeID string

// String returns the AXNodeID as string value.
func (t AXNodeID) String() string {
	return string(t)
}

// AXValueType enum of possible property types.
type AXValueType string

// String returns the AXValueType as string value.
func (t AXValueType) String() string {
	return string(t)
}

// AXValueType values.
const (
	AXValueTypeBoolean            AXValueType = "boolean"
	AXValueTypeTristate           AXValueType = "tristate"
	AXValueTypeBooleanOrUndefined AXValueType = "booleanOrUndefined"
	AXValueTypeIdref              AXValueType = "idref"
	AXValueTypeIdrefList          AXValueType = "idrefList"
	AXValueTypeInteger            AXValueType = "integer"
	AXValueTypeNode               AXValueType = "node"
	AXValueTypeNodeList           AXValueType = "nodeList"
	AXValueTypeNumber             AXValueType = "number"
	AXValueTypeString             AXValueType = "string"
	AXValueTypeComputedString     AXValueType = "computedString"
	AXValueTypeToken              AXValueType = "token"
	AXValueTypeTokenList          AXValueType = "tokenList"
	AXValueTypeDomRelation        AXValueType = "domRelation"
	AXValueTypeRole               AXValueType = "role"
	AXValueTypeInternalRole       AXValueType = "internalRole"
	AXValueTypeValueUndefined     AXValueType = "valueUndefined"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXValueType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXValueType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXValueType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXValueType(in.String()) {
	case AXValueTypeBoolean:
		*t = AXValueTypeBoolean
	case AXValueTypeTristate:
		*t = AXValueTypeTristate
	case AXValueTypeBooleanOrUndefined:
		*t = AXValueTypeBooleanOrUndefined
	case AXValueTypeIdref:
		*t = AXValueTypeIdref
	case AXValueTypeIdrefList:
		*t = AXValueTypeIdrefList
	case AXValueTypeInteger:
		*t = AXValueTypeInteger
	case AXValueTypeNode:
		*t = AXValueTypeNode
	case AXValueTypeNodeList:
		*t = AXValueTypeNodeList
	case AXValueTypeNumber:
		*t = AXValueTypeNumber
	case AXValueTypeString:
		*t = AXValueTypeString
	case AXValueTypeComputedString:
		*t = AXValueTypeComputedString
	case AXValueTypeToken:
		*t = AXValueTypeToken
	case AXValueTypeTokenList:
		*t = AXValueTypeTokenList
	case AXValueTypeDomRelation:
		*t = AXValueTypeDomRelation
	case AXValueTypeRole:
		*t = AXValueTypeRole
	case AXValueTypeInternalRole:
		*t = AXValueTypeInternalRole
	case AXValueTypeValueUndefined:
		*t = AXValueTypeValueUndefined

	default:
		in.AddError(errors.New("unknown AXValueType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXValueType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXValueSourceType enum of possible property sources.
type AXValueSourceType string

// String returns the AXValueSourceType as string value.
func (t AXValueSourceType) String() string {
	return string(t)
}

// AXValueSourceType values.
const (
	AXValueSourceTypeAttribute      AXValueSourceType = "attribute"
	AXValueSourceTypeImplicit       AXValueSourceType = "implicit"
	AXValueSourceTypeStyle          AXValueSourceType = "style"
	AXValueSourceTypeContents       AXValueSourceType = "contents"
	AXValueSourceTypePlaceholder    AXValueSourceType = "placeholder"
	AXValueSourceTypeRelatedElement AXValueSourceType = "relatedElement"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXValueSourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXValueSourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXValueSourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXValueSourceType(in.String()) {
	case AXValueSourceTypeAttribute:
		*t = AXValueSourceTypeAttribute
	case AXValueSourceTypeImplicit:
		*t = AXValueSourceTypeImplicit
	case AXValueSourceTypeStyle:
		*t = AXValueSourceTypeStyle
	case AXValueSourceTypeContents:
		*t = AXValueSourceTypeContents
	case AXValueSourceTypePlaceholder:
		*t = AXValueSourceTypePlaceholder
	case AXValueSourceTypeRelatedElement:
		*t = AXValueSourceTypeRelatedElement

	default:
		in.AddError(errors.New("unknown AXValueSourceType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXValueSourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXValueNativeSourceType enum of possible native property sources (as a
// subtype of a particular AXValueSourceType).
type AXValueNativeSourceType string

// String returns the AXValueNativeSourceType as string value.
func (t AXValueNativeSourceType) String() string {
	return string(t)
}

// AXValueNativeSourceType values.
const (
	AXValueNativeSourceTypeFigcaption   AXValueNativeSourceType = "figcaption"
	AXValueNativeSourceTypeLabel        AXValueNativeSourceType = "label"
	AXValueNativeSourceTypeLabelfor     AXValueNativeSourceType = "labelfor"
	AXValueNativeSourceTypeLabelwrapped AXValueNativeSourceType = "labelwrapped"
	AXValueNativeSourceTypeLegend       AXValueNativeSourceType = "legend"
	AXValueNativeSourceTypeTablecaption AXValueNativeSourceType = "tablecaption"
	AXValueNativeSourceTypeTitle        AXValueNativeSourceType = "title"
	AXValueNativeSourceTypeOther        AXValueNativeSourceType = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXValueNativeSourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXValueNativeSourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXValueNativeSourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXValueNativeSourceType(in.String()) {
	case AXValueNativeSourceTypeFigcaption:
		*t = AXValueNativeSourceTypeFigcaption
	case AXValueNativeSourceTypeLabel:
		*t = AXValueNativeSourceTypeLabel
	case AXValueNativeSourceTypeLabelfor:
		*t = AXValueNativeSourceTypeLabelfor
	case AXValueNativeSourceTypeLabelwrapped:
		*t = AXValueNativeSourceTypeLabelwrapped
	case AXValueNativeSourceTypeLegend:
		*t = AXValueNativeSourceTypeLegend
	case AXValueNativeSourceTypeTablecaption:
		*t = AXValueNativeSourceTypeTablecaption
	case AXValueNativeSourceTypeTitle:
		*t = AXValueNativeSourceTypeTitle
	case AXValueNativeSourceTypeOther:
		*t = AXValueNativeSourceTypeOther

	default:
		in.AddError(errors.New("unknown AXValueNativeSourceType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXValueNativeSourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXValueSource a single source for a computed AX property.
type AXValueSource struct {
	Type              AXValueSourceType       `json:"type"`                        // What type of source this is.
	Value             *AXValue                `json:"value,omitempty"`             // The value of this property source.
	Attribute         string                  `json:"attribute,omitempty"`         // The name of the relevant attribute, if any.
	AttributeValue    *AXValue                `json:"attributeValue,omitempty"`    // The value of the relevant attribute, if any.
	Superseded        bool                    `json:"superseded,omitempty"`        // Whether this source is superseded by a higher priority source.
	NativeSource      AXValueNativeSourceType `json:"nativeSource,omitempty"`      // The native markup source for this value, e.g. a <label> element.
	NativeSourceValue *AXValue                `json:"nativeSourceValue,omitempty"` // The value, such as a node or node list, of the native source.
	Invalid           bool                    `json:"invalid,omitempty"`           // Whether the value for this property is invalid.
	InvalidReason     string                  `json:"invalidReason,omitempty"`     // Reason for the value being invalid, if it is.
}

// AXRelatedNode [no description].
type AXRelatedNode struct {
	BackendDOMNodeID cdp.BackendNodeID `json:"backendDOMNodeId"` // The BackendNodeId of the related DOM node.
	Idref            string            `json:"idref,omitempty"`  // The IDRef value provided, if any.
	Text             string            `json:"text,omitempty"`   // The text alternative of this node in the current context.
}

// AXProperty [no description].
type AXProperty struct {
	Name  AXPropertyName `json:"name"`  // The name of this property.
	Value *AXValue       `json:"value"` // The value of this property.
}

// AXValue a single computed AX property.
type AXValue struct {
	Type         AXValueType         `json:"type"`                   // The type of this value.
	Value        easyjson.RawMessage `json:"value,omitempty"`        // The computed value of this property.
	RelatedNodes []*AXRelatedNode    `json:"relatedNodes,omitempty"` // One or more related nodes, if applicable.
	Sources      []*AXValueSource    `json:"sources,omitempty"`      // The sources which contributed to the computation of this property.
}

// AXPropertyName values of AXProperty name: from 'busy' to 'roledescription'
// - states which apply to every AX node, from 'live' to 'root' - attributes
// which apply to nodes in live regions, from 'autocomplete' to 'valuetext' -
// attributes which apply to widgets, from 'checked' to 'selected' - states
// which apply to widgets, from 'activedescendant' to 'owns' - relationships
// between elements other than parent/child/sibling.
type AXPropertyName string

// String returns the AXPropertyName as string value.
func (t AXPropertyName) String() string {
	return string(t)
}

// AXPropertyName values.
const (
	AXPropertyNameBusy             AXPropertyName = "busy"
	AXPropertyNameDisabled         AXPropertyName = "disabled"
	AXPropertyNameHidden           AXPropertyName = "hidden"
	AXPropertyNameHiddenRoot       AXPropertyName = "hiddenRoot"
	AXPropertyNameInvalid          AXPropertyName = "invalid"
	AXPropertyNameKeyshortcuts     AXPropertyName = "keyshortcuts"
	AXPropertyNameRoledescription  AXPropertyName = "roledescription"
	AXPropertyNameLive             AXPropertyName = "live"
	AXPropertyNameAtomic           AXPropertyName = "atomic"
	AXPropertyNameRelevant         AXPropertyName = "relevant"
	AXPropertyNameRoot             AXPropertyName = "root"
	AXPropertyNameAutocomplete     AXPropertyName = "autocomplete"
	AXPropertyNameHaspopup         AXPropertyName = "haspopup"
	AXPropertyNameLevel            AXPropertyName = "level"
	AXPropertyNameMultiselectable  AXPropertyName = "multiselectable"
	AXPropertyNameOrientation      AXPropertyName = "orientation"
	AXPropertyNameMultiline        AXPropertyName = "multiline"
	AXPropertyNameReadonly         AXPropertyName = "readonly"
	AXPropertyNameRequired         AXPropertyName = "required"
	AXPropertyNameValuemin         AXPropertyName = "valuemin"
	AXPropertyNameValuemax         AXPropertyName = "valuemax"
	AXPropertyNameValuetext        AXPropertyName = "valuetext"
	AXPropertyNameChecked          AXPropertyName = "checked"
	AXPropertyNameExpanded         AXPropertyName = "expanded"
	AXPropertyNameModal            AXPropertyName = "modal"
	AXPropertyNamePressed          AXPropertyName = "pressed"
	AXPropertyNameSelected         AXPropertyName = "selected"
	AXPropertyNameActivedescendant AXPropertyName = "activedescendant"
	AXPropertyNameControls         AXPropertyName = "controls"
	AXPropertyNameDescribedby      AXPropertyName = "describedby"
	AXPropertyNameDetails          AXPropertyName = "details"
	AXPropertyNameErrormessage     AXPropertyName = "errormessage"
	AXPropertyNameFlowto           AXPropertyName = "flowto"
	AXPropertyNameLabelledby       AXPropertyName = "labelledby"
	AXPropertyNameOwns             AXPropertyName = "owns"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXPropertyName) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXPropertyName) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXPropertyName) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXPropertyName(in.String()) {
	case AXPropertyNameBusy:
		*t = AXPropertyNameBusy
	case AXPropertyNameDisabled:
		*t = AXPropertyNameDisabled
	case AXPropertyNameHidden:
		*t = AXPropertyNameHidden
	case AXPropertyNameHiddenRoot:
		*t = AXPropertyNameHiddenRoot
	case AXPropertyNameInvalid:
		*t = AXPropertyNameInvalid
	case AXPropertyNameKeyshortcuts:
		*t = AXPropertyNameKeyshortcuts
	case AXPropertyNameRoledescription:
		*t = AXPropertyNameRoledescription
	case AXPropertyNameLive:
		*t = AXPropertyNameLive
	case AXPropertyNameAtomic:
		*t = AXPropertyNameAtomic
	case AXPropertyNameRelevant:
		*t = AXPropertyNameRelevant
	case AXPropertyNameRoot:
		*t = AXPropertyNameRoot
	case AXPropertyNameAutocomplete:
		*t = AXPropertyNameAutocomplete
	case AXPropertyNameHaspopup:
		*t = AXPropertyNameHaspopup
	case AXPropertyNameLevel:
		*t = AXPropertyNameLevel
	case AXPropertyNameMultiselectable:
		*t = AXPropertyNameMultiselectable
	case AXPropertyNameOrientation:
		*t = AXPropertyNameOrientation
	case AXPropertyNameMultiline:
		*t = AXPropertyNameMultiline
	case AXPropertyNameReadonly:
		*t = AXPropertyNameReadonly
	case AXPropertyNameRequired:
		*t = AXPropertyNameRequired
	case AXPropertyNameValuemin:
		*t = AXPropertyNameValuemin
	case AXPropertyNameValuemax:
		*t = AXPropertyNameValuemax
	case AXPropertyNameValuetext:
		*t = AXPropertyNameValuetext
	case AXPropertyNameChecked:
		*t = AXPropertyNameChecked
	case AXPropertyNameExpanded:
		*t = AXPropertyNameExpanded
	case AXPropertyNameModal:
		*t = AXPropertyNameModal
	case AXPropertyNamePressed:
		*t = AXPropertyNamePressed
	case AXPropertyNameSelected:
		*t = AXPropertyNameSelected
	case AXPropertyNameActivedescendant:
		*t = AXPropertyNameActivedescendant
	case AXPropertyNameControls:
		*t = AXPropertyNameControls
	case AXPropertyNameDescribedby:
		*t = AXPropertyNameDescribedby
	case AXPropertyNameDetails:
		*t = AXPropertyNameDetails
	case AXPropertyNameErrormessage:
		*t = AXPropertyNameErrormessage
	case AXPropertyNameFlowto:
		*t = AXPropertyNameFlowto
	case AXPropertyNameLabelledby:
		*t = AXPropertyNameLabelledby
	case AXPropertyNameOwns:
		*t = AXPropertyNameOwns

	default:
		in.AddError(errors.New("unknown AXPropertyName value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXPropertyName) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXNode a node in the accessibility tree.
type AXNode struct {
	NodeID           AXNodeID          `json:"nodeId"`                     // Unique identifier for this node.
	Ignored          bool              `json:"ignored"`                    // Whether this node is ignored for accessibility
	IgnoredReasons   []*AXProperty     `json:"ignoredReasons,omitempty"`   // Collection of reasons why this node is hidden.
	Role             *AXValue          `json:"role,omitempty"`             // This Node's role, whether explicit or implicit.
	Name             *AXValue          `json:"name,omitempty"`             // The accessible name for this Node.
	Description      *AXValue          `json:"description,omitempty"`      // The accessible description for this Node.
	Value            *AXValue          `json:"value,omitempty"`            // The value for this Node.
	Properties       []*AXProperty     `json:"properties,omitempty"`       // All other properties
	ChildIds         []AXNodeID        `json:"childIds,omitempty"`         // IDs for each of this node's child nodes.
	BackendDOMNodeID cdp.BackendNodeID `json:"backendDOMNodeId,omitempty"` // The backend ID for the associated DOM node, if any.
}
