package repo

import (
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
)

// ParseOrder 解析排序字段
func ParseOrder(items []*schema.OrderField) []ent.OrderFunc {
	orders := make([]ent.OrderFunc, len(items))

	for i, item := range items {
		key := item.Key
		if item.Direction == schema.OrderByDESC {
			orders[i] = ent.Desc(key)
		} else {
			orders[i] = ent.Asc(key)
		}
	}

	return orders
}

func MakeUpOrderField(field string, order string) *schema.OrderField {
	switch order {
	case "asc":
		fallthrough
	case "ascend":
		return schema.NewOrderField(field, schema.OrderByASC)
	case "desc":
		fallthrough
	case "descend":
		return schema.NewOrderField(field, schema.OrderByDESC)
	}
	return nil
}
