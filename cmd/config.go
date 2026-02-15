package cmd

type Config struct {
	HttpPort               string
	DbHost                 string
	DbPort                 string
	DbUser                 string
	DbPassword             string
	DbName                 string
	DbSslMode              string
	GeoServiceGrpcHost     string
	KafkaHost              string
	KafkaConsumerGroup     string
	KafkaBasketEventsTopic string
	KafkaOrderEventsTopic  string
}
