package main

import (
	"fmt"
	"strings"
)

// 表示组织机构的接口
type Organization interface {
	display() // 显示组织架构
	duty()    // 显示职责
}

// 组合组织对象
type CompositeOrganization struct {
	orgName string
	depth   int
	list    []Organization
}

// 创建并返回一个新的组合组织对象
func NewCompositeOrganization(name string, depth int) *CompositeOrganization {
	return &CompositeOrganization{name, depth, []Organization{}}
}

// 在组合组织中添加一个组织对象
func (c *CompositeOrganization) add(org Organization) {
	if c == nil {
		return
	}
	c.list = append(c.list, org)
}

// 从组合组织中移除指定的组织对象
func (c *CompositeOrganization) remove(org Organization) {
	if c == nil {
		return
	}
	for i, val := range c.list {
		if val == org {
			c.list = append(c.list[:i], c.list[i+1:]...)
			return
		}
	}
}

// 递归地显示组织架构
func (c *CompositeOrganization) display() {
	if c == nil {
		return
	}
	fmt.Println(strings.Repeat("-", c.depth*2), " ", c.orgName)
	for _, val := range c.list {
		val.display()
	}
}

// 递归地显示组织的职责
func (c *CompositeOrganization) duty() {
	if c == nil {
		return
	}
	for _, val := range c.list {
		val.duty()
	}
}

// 人力资源部门对象
type HRDOrg struct {
	orgName string
	depth   int
}

// 显示人力资源部门的组织架构
func (o *HRDOrg) display() {
	if o == nil {
		return
	}
	fmt.Println(strings.Repeat("-", o.depth*2), " ", o.orgName)
}

// 显示人力资源部门的职责
func (o *HRDOrg) duty() {
	if o == nil {
		return
	}
	fmt.Println(o.orgName, "员工招聘培训管理")
}

// 财务部门对象
type FinanceOrg struct {
	orgName string
	depth   int
}

// 显示财务部门的组织架构
func (f *FinanceOrg) display() {
	if f == nil {
		return
	}
	fmt.Println(strings.Repeat("-", f.depth*2), " ", f.orgName)
}

// 显示财务部门的职责
func (f *FinanceOrg) duty() {
	if f == nil {
		return
	}
	fmt.Println(f.orgName, "财务管理")
}
