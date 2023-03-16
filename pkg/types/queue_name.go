package types

type QueueName string

// queues
const (
	EventQueue    QueueName = "EventQueue"
	DistrictQueue QueueName = "DistrictQueue"
	PriorityQueue QueueName = "PriorityQueue"
	ScheduleQueue QueueName = "ScheduleQueue"
	DefaultQueue  QueueName = "DefaultQueue"
)
