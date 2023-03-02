package structure

import (
	"github.com/heromicro/omgind/pkg/helper/deepcopier"
	"github.com/jinzhu/copier"
)

// Copy 结构体映射
func Copy(s, ts any) error {
	return copier.Copy(ts, s)
}

// Copy 结构体映射
func DeepCopy(s, ts any) error {
	return deepcopier.Copy(s).To(ts)
}
