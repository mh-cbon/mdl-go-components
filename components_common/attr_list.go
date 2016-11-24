package components_common

import "html/template"

type AttributeNode struct {
	Name  string
	Value string
}

func (c AttributeNode) SafeName() interface{} {
	return template.HTMLAttr(c.Name)
}

type AttrList []*AttributeNode

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
			return v
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
		*c = append(*c, &AttributeNode{
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
