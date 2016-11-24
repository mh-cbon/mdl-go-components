package components_common

import (
	"html/template"
	"strconv"
	"time"
)

type Translator interface {
	T(id string, args ...interface{}) string
}
type NodeTranslator interface {
	Translate(t Translator)
}

type Namer interface {
	GetName() string
}
type ValueSetter interface {
	SetValue(s string)
	SetIntValue(value int)
	SetInt64Value(value int64)
	SetUint64Value(value uint64)
}
type ValueDateSetter interface {
	SetDate(s time.Time)
	GetDate() (time.Time, error)
	MustGetDate() time.Time
}
type ValueSliceSetter interface {
	SetValues(s []string)
	SetIntValues(s []int)
	SetInt64Values(s []int64)
	SetUint64Values(s []uint64)
}
type Optionner interface {
	SetOptions(s []*NodeOption)
	GetOptions() []*NodeOption
	AddOption(value string, label string)
	AddIntOption(value int, label string)
	AddInt64Option(value int64, label string)
	AddUint64Option(value uint64, label string)
	GetOptionForValue(s string) *NodeOption
	CheckOptions(values []string)
	CheckIntOptions(values []int)
	CheckInt64Options(values []int64)
	CheckUint64Options(values []uint64)
}
type ErrorProvider interface {
	// GetFailure() interface{}
	GetError(name string) interface{}
}
type NodeErrorsSetter interface {
	SetErrors(provider ErrorProvider)
}

type Node struct {
	Classes ClassList
	Attr    AttrList
}

func NewNode() *Node {
	return &Node{
		Classes: ClassList{},
		Attr:    AttrList{},
	}
}

func (i *Node) SetName(s string) {
	i.Attr.Set("name", s)
}
func (i *Node) GetName() string {
	return i.Attr.GetValue("name")
}
func (i *Node) SetId(s string) {
	i.Attr.Set("id", s)
}
func (i *Node) GetId() string {
	return i.Attr.GetValue("id")
}

type NodeType struct {
	Type string
}

func (i *NodeType) SetType(s string) {
	i.Type = s
}
func (i *NodeType) GetType() string {
	return i.Type
}

type NodeSingleValue struct {
	Value string
}

func (i *NodeSingleValue) GetValue() string {
	return i.Value
}
func (i *NodeSingleValue) SetValue(s string) {
	i.Value = s
}
func (i *NodeSingleValue) SetIntValue(value int) {
	i.SetValue(strconv.Itoa(value))
}
func (i *NodeSingleValue) SetInt64Value(value int64) {
	i.SetValue(strconv.FormatInt(value, 10))
}
func (i *NodeSingleValue) SetUint64Value(value uint64) {
	i.SetValue(strconv.FormatUint(value, 10))
}
func (i *NodeSingleValue) SetDate(b time.Time) {
	i.SetValue(b.Format(time.RFC3339))
}
func (i *NodeSingleValue) GetDate() (time.Time, error) {
	return time.Parse(time.RFC3339, i.GetValue())
}
func (i *NodeSingleValue) MustGetDate() time.Time {
	t, err := time.Parse(time.RFC3339, i.GetValue())
	if err != nil {
		panic(err)
	}
	return t
}

type NodeMutlipleValues struct {
	Values []string
}

func (i *NodeMutlipleValues) SetValue(s string) {
	i.Values = []string{s}
}
func (i *NodeMutlipleValues) GetValues() []string {
	return i.Values
}
func (i *NodeMutlipleValues) HasValue(s string) bool {
	for _, v := range i.Values {
		if v == s {
			return true
		}
	}
	return false
}
func (i *NodeMutlipleValues) SetValues(s []string) {
	i.Values = s
}
func (i *NodeMutlipleValues) SetIntValues(values []int) {
	sValues := make([]string, 0)
	for _, v := range values {
		sValues = append(sValues, strconv.Itoa(v))
	}
	i.SetValues(sValues)
}
func (i *NodeMutlipleValues) SetInt64Values(values []int64) {
	sValues := make([]string, 0)
	for _, v := range values {
		sValues = append(sValues, strconv.FormatInt(v, 10))
	}
	i.SetValues(sValues)
}
func (i *NodeMutlipleValues) SetUint64Values(values []uint64) {
	sValues := make([]string, 0)
	for _, v := range values {
		sValues = append(sValues, strconv.FormatUint(v, 10))
	}
	i.SetValues(sValues)
}

type NodeWithOption struct {
	Option NodeOption
}

func (i *NodeWithOption) SetOption(option NodeOption) {
	i.Option = option
}
func (i *NodeWithOption) GetOption() *NodeOption {
	return &i.Option
}
func (i *NodeWithOption) Translate(t Translator) {
	i.Option.Translate(t)
}

type NodeWithOptions struct {
	Options []*NodeOption
}

func (i *NodeWithOptions) SetOptions(s []*NodeOption) {
	i.Options = s
}
func (i *NodeWithOptions) GetOptions() []*NodeOption {
	return i.Options
}
func (i *NodeWithOptions) AddOption(value string, label string) {
	i.Options = append(i.Options, &NodeOption{
		NodeLabel:       NodeLabel{Label: label},
		NodeSingleValue: NodeSingleValue{Value: value},
	})
}
func (i *NodeWithOptions) AddIntOption(value int, label string) {
	i.AddOption(strconv.Itoa(value), label)
}
func (i *NodeWithOptions) AddInt64Option(value int64, label string) {
	i.AddOption(strconv.FormatInt(value, 10), label)
}
func (i *NodeWithOptions) AddUint64Option(value uint64, label string) {
	i.AddOption(strconv.FormatUint(value, 10), label)
}
func (i *NodeWithOptions) GetOptionForValue(s string) *NodeOption {
	for _, v := range i.Options {
		if v.Value == s {
			return v
		}
	}
	return nil
}
func (i *NodeWithOptions) CheckOptions(values []string) {
	for _, v := range values {
		option := i.GetOptionForValue(v)
		if option != nil {
			option.SetChecked(true)
		}
	}
}
func (i *NodeWithOptions) CheckIntOptions(values []int) {
	sValues := make([]string, 0)
	for _, v := range values {
		sValues = append(sValues, strconv.Itoa(v))
	}
	i.CheckOptions(sValues)
}
func (i *NodeWithOptions) CheckInt64Options(values []int64) {
	sValues := make([]string, 0)
	for _, v := range values {
		sValues = append(sValues, strconv.FormatInt(v, 10))
	}
	i.CheckOptions(sValues)
}
func (i *NodeWithOptions) CheckUint64Options(values []uint64) {
	sValues := make([]string, 0)
	for _, v := range values {
		sValues = append(sValues, strconv.FormatUint(v, 10))
	}
	i.CheckOptions(sValues)
}
func (i *NodeWithOptions) SetValue(s string) {
	i.CheckOptions([]string{s})
}
func (i *NodeWithOptions) SetIntValue(value int) {
	i.SetValue(strconv.Itoa(value))
}
func (i *NodeWithOptions) SetInt64Value(value int64) {
	i.SetValue(strconv.FormatInt(value, 10))
}
func (i *NodeWithOptions) SetUint64Value(value uint64) {
	i.SetValue(strconv.FormatUint(value, 10))
}
func (i *NodeWithOptions) Translate(t Translator) {
	for _, v := range i.Options {
		v.Translate(t)
	}
}

type NodeOption struct {
	NodeLabel
	NodeSingleValue
	NodeCheckable
}

func (i *NodeOption) Eq(s string) bool {
	return i.Value == s
}

type NodeSingleError struct {
	Error interface{}
}

func (i *NodeSingleError) SetError(s interface{}) {
	i.Error = s
}
func (i *NodeSingleError) GetError() interface{} {
	return i.Error
}

type NodeLabel struct {
	Label interface{}
}

func (i *NodeLabel) SetSafeLabel(s string) {
	i.Label = template.HTML(s)
}
func (i *NodeLabel) SetLabel(s interface{}) {
	i.Label = s
}
func (i *NodeLabel) GetLabel() interface{} {
	return i.Label
}
func (i *NodeLabel) Translate(t Translator) {
	if x, ok := i.Label.(template.HTML); ok {
		i.SetSafeLabel(t.T(string(x)))
	} else if x, ok := i.Label.(string); ok {
		i.SetLabel(t.T(x))
	}
}

type NodePlaceholder struct {
	PlaceHolderOnly bool
}

func (i *NodePlaceholder) SetPlaceHolderOnly(b bool) {
	i.PlaceHolderOnly = b
}
func (i *NodePlaceholder) IsPlaceHolderOnly() bool {
	return i.PlaceHolderOnly
}

type NodeCheckable struct {
	Checked bool
}

func (i *NodeCheckable) SetChecked(b bool) {
	i.Checked = b
}
func (i *NodeCheckable) IsChecked() bool {
	return i.Checked
}
