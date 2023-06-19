package repo

import (
	"entgo.io/ent/dialect/sql"
)

type OrderDire string

const (
	// OrderByASC 升序排序
	OrderByASC OrderDire = "asc"
	// OrderByDESC 降序排序
	OrderByDESC OrderDire = "desc"
)

func (o OrderDire) String() string {
	switch o {
	case OrderByASC:
		return "asc"
	case OrderByDESC:
		return "desc"
	default:
		return ""
	}
}

// OrderField 排序字段
type OrderField struct {
	Key       string    // 字段名(字段名约束为小写蛇形)
	Direction OrderDire // 排序方向
}

func OrderDirection(v string) sql.OrderTermOption {
	switch v {
	case "a":
		fallthrough
	case "asc":
		fallthrough
	case "ascend":
		return sql.OrderAsc()
	case "d":
		fallthrough
	case "desc":
		fallthrough
	case "descend":
		return sql.OrderDesc()
	default:
		return sql.OrderDesc()
	}
}

func OrderDirections() []sql.OrderTermOption {
	ods := []sql.OrderTermOption{}

	return ods
}
