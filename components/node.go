package components

import (
	"html/template"
	"strconv"
	"strings"
)

type Translator interface {
	T(id string, args ...interface{}) string
}

type ViewTranslator interface {
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
type ValueSliceSetter interface {
	SetValues(s []string)
}

type ClassList []string

func (c *ClassList) Contains(some string) bool {
	if some == "" {
		return false
	}
	for _, v := range *c {
		if v == some {
			return true
		}
	}
	return false
}
func (c *ClassList) Add(some string) {
	for _, s := range strings.Split(some, " ") {
		if c.Contains(s) == false {
			*c = append(*c, s)
		}
	}
}
func (c *ClassList) Toggle(some string) {
	for _, s := range strings.Split(some, " ") {
		if c.Contains(s) {
			c.Remove(s)
		} else {
			c.Add(s)
		}
	}
}
func (c *ClassList) Remove(some string) {
	for _, s := range strings.Split(some, " ") {
		for i, v := range *c {
			if v == s {
				k := *c
				*c = append(k[:i], k[i+1:]...)
				break
			}
		}
	}
}
func (c *ClassList) RemoveStartingWith(some string) {
	for _, s := range strings.Split(some, " ") {
		for i, v := range *c {
			if len(v) >= len(s) && v[0:len(s)] == s {
				k := *c
				*c = append(k[:i], k[i+1:]...)
				break
			}
		}
	}
}
func (c *ClassList) CopyFrom(some ClassList) {
	k := *c
	*c = k[0:0]
	for _, v := range some {
		c.Add(v)
	}
}
func (c *ClassList) MergeFrom(some ClassList) {
	for _, v := range some {
		c.Add(v)
	}
}
func (c *ClassList) Render() interface{} {
	ret := ""
	for _, v := range *c {
		ret += v + " "
	}
	if len(ret) > 0 {
		return ret[0 : len(ret)-1]
	}
	return ret
}

type AttributeNode struct {
	Name  string
	Value string
}

func (c AttributeNode) Render() interface{} {
	ret := c.Name + "=\"" + c.Value + "\" "
	return template.HTMLAttr(ret)
}

type AttrList []AttributeNode

func (c *AttrList) Has(name string) bool {
	for _, v := range *c {
		if v.Name == name {
			return true
		}
	}
	return false
}
func (c *AttrList) Get(name string) *AttributeNode {
	for _, v := range *c {
		if v.Name == name {
			return &v
		}
	}
	return nil
}
func (c *AttrList) GetValue(name string) string {
	if attr := c.Get(name); attr != nil {
		return attr.Value
	}
	return ""
}
func (c *AttrList) Set(name string, value string) {
	if attr := c.Get(name); attr == nil {
		*c = append(*c, AttributeNode{
			Name:  name,
			Value: value,
		})
	} else {
		attr.Value = value
	}
}
func (c *AttrList) Remove(name string) {
	for i, v := range *c {
		if v.Name == name {
			k := *c
			*c = append(k[:i], k[i+1:]...)
			break
		}
	}
}
func (c *AttrList) CopyFrom(some AttrList) {
	for _, v := range some {
		c.Set(v.Name, v.Value)
	}
}
func (c *AttrList) MergeFrom(some AttrList) {
	for _, v := range some {
		if c.Has(v.Name) == false {
			c.Set(v.Name, v.Value)
		}
	}
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

type NodeMutlipleValues struct {
	Values []string
}

func (i *NodeMutlipleValues) SetValue(s string) {
	i.Values = []string{s}
}
func (i *NodeMutlipleValues) SetValues(s []string) {
	i.Values = s
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

type NodeWithOption struct {
	Option NodeOption
}

func (i *NodeWithOption) SetOption(option NodeOption) {
	i.Option = option
}
func (i *NodeWithOption) GetOption() *NodeOption {
	return &i.Option
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
func (i *NodeWithOptions) AddInt64Option(value int64, label string) {
	i.AddOption(strconv.FormatInt(value, 10), label)
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
	Error string
}

func (i *NodeSingleError) SetError(s string) {
	i.Error = s
}
func (i *NodeSingleError) GetError() string {
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
