package main

import (
	"fmt"
	"testing"
)

func TestCombination(t *testing.T) {
	// 创建根组织
	root := NewCompositeOrganization("北京总公司", 1)

	// 添加人力资源部门和财务部门到根组织
	root.add(&HRDOrg{orgName: "总公司人力资源部", depth: 2})
	root.add(&FinanceOrg{orgName: "总公司财务部", depth: 2})

	// 创建上海分公司，并添加人力资源部门和财务部门到上海分公司
	compSh := NewCompositeOrganization("上海分公司", 2)
	compSh.add(&FinanceOrg{orgName: "上海分公司财务部", depth: 3})

	hrdSh := &HRDOrg{orgName: "上海分公司人力资源部", depth: 3}
	compSh.add(hrdSh)

	compSh.remove(hrdSh)

	root.add(compSh)

	// 创建广东分公司，并添加人力资源部门和财务部门到广东分公司
	compGd := NewCompositeOrganization("广东分公司", 2)
	compGd.add(&HRDOrg{orgName: "广东分公司人力资源部", depth: 3})
	compGd.add(&FinanceOrg{orgName: "广东分公司财务部", depth: 3})
	root.add(compGd)

	fmt.Println("公司组织架构：")
	root.display() // 显示公司组织架构

	fmt.Println("各组织的职责：")
	root.duty() // 显示各组织的职责

}
