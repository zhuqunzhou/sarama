package sarama

type serde struct {
	config       *Config
	serializer   *serializer
	deserializer *deserializer
}

type Serde interface {
	Serializer() Serializer
	Deserializer() Deserializer
}

func NewSerde(config *Config) (Serde, error) {
	serde := &serde{
		config: config,
	}
	serializer, err := NewSerializer(config)
	if err == nil {
		serde.serializer = serializer
	}
	deserializer, err := NewDeserializer(config)
	if err == nil {
		serde.deserializer = deserializer
	}
	return serde, err
}

func (serde *serde) Serializer() Serializer {
	return serde.serializer
}

func (serde *serde) Deserializer() Deserializer {
	return serde.deserializer
}
