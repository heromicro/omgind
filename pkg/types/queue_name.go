package types

type QueueName string

// queues
const (
	EventQueue       QueueName = "EventQueue"
	CreateEventQueue QueueName = "CreateEventQueue"
	PriorityQueue    QueueName = "PriorityQueue"
	ScheduleQueue    QueueName = "ScheduleQueue"
	DefaultQueue     QueueName = "DefaultQueue"
)
