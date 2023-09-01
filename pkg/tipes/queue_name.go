package tipes

type QueueName string

// queues
const (
	EventQueue       QueueName = "EventQueue"
	RepaireTreeQueue QueueName = "RepaireTreeQueue"
	PriorityQueue    QueueName = "PriorityQueue"
	ScheduleQueue    QueueName = "ScheduleQueue"
	DefaultQueue     QueueName = "DefaultQueue"
)
