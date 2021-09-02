package inboxroad

type MessageHeaderInterface interface {
	SetKey(key string) MessageHeaderInterface
	GetKey() string
	SetValue(value string) MessageHeaderInterface
	GetValue() string
	ToMap() StringMap
	ToInboxroadMap() StringMap
}

type MessageHeader struct {
	key, value string
}

func (m *MessageHeader) SetKey(key string) MessageHeaderInterface {
	m.key = key

	return m
}

func (m MessageHeader) GetKey() string {
	return m.key
}

func (m *MessageHeader) SetValue(value string) MessageHeaderInterface {
	m.value = value

	return m
}

func (m MessageHeader) GetValue() string {
	return m.value
}

func (m MessageHeader) ToMap() StringMap {
	return StringMap{
		"key":   m.GetKey(),
		"value": m.GetValue(),
	}
}

func (m MessageHeader) ToInboxroadMap() StringMap {
	return StringMap{
		m.GetKey(): m.GetValue(),
	}
}

func NewMessageHeader(key, value string) MessageHeaderInterface {
	return (&MessageHeader{}).SetKey(key).SetValue(value) //nolint:exhaustivestruct
}
