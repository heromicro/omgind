package structure

import (
	"github.com/jinzhu/copier"
)

// Copy 结构体映射
func Copy(s, ts any) error {
	return copier.Copy(ts, s)
}
