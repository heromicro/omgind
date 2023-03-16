package types

import (
	"strings"
)

type TaskName string

func (t TaskName) SetPrefix(prefix string) TaskName {

	var name strings.Builder
	delim := "-"

	name.WriteString(prefix)
	name.WriteString(delim)
	name.WriteString(string(t))

	return TaskName(name.String())
}

const (
	TaskName_REPAIR_DISTRICT_TREE_PATH TaskName = "TN_REPAIR_DISTRICT_TREE_PATH"
)
