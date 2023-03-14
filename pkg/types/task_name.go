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
	EventProcessor         TaskName = "EventProcessor"
	DeadLetterProcessor    TaskName = "DeadLetterProcessor"
	CreateEventProcessor   TaskName = "CreateEventProcessor"
	NotificationProcessor  TaskName = "NotificationProcessor"
	IndexDocument          TaskName = "index document"
	DailyAnalytics         TaskName = "daily analytics"
	MonitorTwitterSources  TaskName = "monitor twitter sources"
	RetentionPolicies      TaskName = "retention_policies"
	EmailProcessor         TaskName = "EmailProcessor"
	ExpireSecretsProcessor TaskName = "ExpireSecretsProcessor"
)
