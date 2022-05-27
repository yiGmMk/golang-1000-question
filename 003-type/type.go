package typetest

import (
	"fmt"

	"github.com/yiGmMk/leetcode/golang/util"
)

// see now we have two customize type,and what is the difference?
type NodeType1 util.TreeNode

type NodeType2 = util.TreeNode

func Trans1(val interface{}) (*NodeType1, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%+v", r)
		}
	}()
	if val == nil {
		return nil, nil
	}
	out, _ := val.(*NodeType1)
	return out, err
}

func Trans2(val interface{}) (*NodeType2, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%+v", r)
		}
	}()
	if val == nil {
		return nil, nil
	}
	out, _ := val.(*NodeType2)
	return out, err
}
