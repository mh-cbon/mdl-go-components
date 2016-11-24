package components_common

import "strings"

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
func (c *ClassList) Render() string {
	ret := ""
	for _, v := range *c {
		ret += v + " "
	}
	if len(ret) > 0 {
		return ret[0 : len(ret)-1]
	}
	return ret
}
