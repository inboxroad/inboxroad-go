package inboxroad

type MessageHeaderInterface interface {
	SetKey(key string) MessageHeaderInterface
	GetKey() string
	SetValue(value string) MessageHeaderInterface
	GetValue() string
	ToArray() map[string]string
	ToInboxroadArray() map[string]string
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

func (m MessageHeader) ToArray() map[string]string {
	return map[string]string{
		"key":   m.GetKey(),
		"value": m.GetValue(),
	}
}

func (m MessageHeader) ToInboxroadArray() map[string]string {
	return map[string]string{
		m.GetKey(): m.GetValue(),
	}
}

func NewMessageHeader(key, value string) MessageHeaderInterface {
	return (&MessageHeader{}).SetKey(key).SetValue(value) //nolint:exhaustivestruct
}
