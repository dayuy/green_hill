package unit_test

import (
	"fmt"
	"testing"
)

func TestErrorInCode(t *testing.T) {
	fmt.Printf("Start")
	t.Error("Error") // 继续执行
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Fatal") // 中止测试
	fmt.Println("End")
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func flatten(root *TreeNode, result *TreeNode) {
	result.val = root.val
	result.right = &TreeNode{}
	if root.left != nil {
		flatten(root.left, result.right)
	}
	if root.right != nil {
		flatten(root.right, result.right)
	}
}

func flatten_2(root *TreeNode) []int {
	ret := []int{}
	var order func(*TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.val)
		order(node.left)
		order(node.right)
	}
	order(root)
	return ret
}

func TestFlatten(t *testing.T) {
	root := &TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,
			left: &TreeNode{
				val: 3,
			},
			right: &TreeNode{
				val: 4,
			},
		},
		right: &TreeNode{
			val: 5,
			right: &TreeNode{
				val: 6,
			},
		},
	}
	//result := &TreeNode{}
	//flatten(root, result)
	//for {
	//	fmt.Println(result.val)
	//	if result.right == nil {
	//		break
	//	}
	//	result = result.right
	//}
	flatten_1(root)
}

func flatten_1(root *TreeNode) {
	vals := []int{}
	s := []*TreeNode{}
	node := root
	for node != nil || len(s) > 0 {
		for node != nil {
			vals = append(vals, node.val)
			s = append(s, node)
			node = node.left
		}
		node = s[len(s)-1].right
		s = s[:len(s)-1]
	}
	fmt.Println(vals)
}

func TestMap(t *testing.T) {
	var labels map[string]string
	if labels["test"] != "test" {
		//if labels == nil {
		//	labels = make(map[string]string)
		//}
		//labels["test"] = "test" // panic: assignment to entry in nil map
	}
	fmt.Println(labels)
	fmt.Println(labels == nil)
	fmt.Println(labels["test"])
	fmt.Println(labels["test"] == "")
	fmt.Printf("%T\n", labels["test"])
}

func TestSlice(t *testing.T) {
	var slice []map[string]string
	var slice1 = make([]map[string]string, 3, 5)
	fmt.Println(slice == nil)
	fmt.Println(len(slice) == 0)
	slice1 = slice1[0:0]
	fmt.Println(slice1 == nil)
	fmt.Println(len(slice1) == 0)
}

type VersionedFilterCond struct {
	// Accurately match each item in the versions
	Versions []string `json:"versions,omitempty"`
	// Filter version by regexp
	VersionRegexp string `json:"regexp,omitempty"`
	// VersionConstraint Support for user-defined version ranges, etc.
	// Refer to the documentation for more details
	// https://github.com/Masterminds/semver#semver
	VersionConstraint string `json:"versionConstraint,omitempty"`
}
type FilterCond struct {
	// VersionedFilterCond filters which version in component are pulled/ignored from the repository
	VersionedFilterCond *VersionedFilterCond `json:"versionedFilterCond,omitempty"`
}

func TestStructNil(t *testing.T) {
	var st = map[string]FilterCond{"test": {VersionedFilterCond: &VersionedFilterCond{}}}
	var st1 = map[string]FilterCond{"test": {}}
	var st2 = map[string]FilterCond{}
	fmt.Println(st["test"].VersionedFilterCond == nil)
	fmt.Println(st["test"].VersionedFilterCond)
	fmt.Println(st1["test"].VersionedFilterCond == nil)
	fmt.Println(st1["test"].VersionedFilterCond)
	fmt.Println(st2["test"].VersionedFilterCond == nil)
	fmt.Println(st2["test"].VersionedFilterCond)
}

func TestStructK(t *testing.T) {
	var obj = make(map[string]string)
	obj["test"] =
		println(obj["test"])
	if v, ok := obj["test"]; !ok || v == "" {
		fmt.Println("----")
		fmt.Println(v)
	}
}
