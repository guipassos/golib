package kafkamod

type TopicPartition struct {
	Topic     string
	Partition int32
	Offset    int64
}
