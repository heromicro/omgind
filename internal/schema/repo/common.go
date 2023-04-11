package repo

import (
	"entgo.io/ent/dialect/sql"
)

func OrderBy(v string) sql.OrderTermOption {
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
