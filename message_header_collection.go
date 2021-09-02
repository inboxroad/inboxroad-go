package inboxroad

type MessageHeaderCollectionInterface interface {
	Add(header MessageHeaderInterface) MessageHeaderCollectionInterface
	GetItems() []MessageHeaderInterface
	ToSliceMap() []map[string]string
	ToInboxroadMap() map[string]string
}

type MessageHeaderCollection struct {
	collection []MessageHeaderInterface
}

func (m *MessageHeaderCollection) Add(header MessageHeaderInterface) MessageHeaderCollectionInterface {
	m.collection = append(m.collection, header)

	return m
}

func (m *MessageHeaderCollection) GetItems() []MessageHeaderInterface {
	return m.collection
}

func (m MessageHeaderCollection) ToSliceMap() []map[string]string {
	data := make([]map[string]string, 0, len(m.GetItems()))
	for _, header := range m.GetItems() {
		data = append(data, header.ToMap())
	}

	return data
}

func (m MessageHeaderCollection) ToInboxroadMap() map[string]string {
	data := map[string]string{}
	for _, header := range m.GetItems() {
		data[header.GetKey()] = header.GetValue()
	}

	return data
}

func NewMessageHeaderCollection() MessageHeaderCollectionInterface {
	return &MessageHeaderCollection{
		collection: []MessageHeaderInterface{},
	}
}

func NewMessageHeaderCollectionFromSliceMap(rawHeaders []map[string]string) MessageHeaderCollectionInterface {
	headers := NewMessageHeaderCollection()

	for _, h := range rawHeaders {
		headerKey, headerValue := "", ""
		if key, ok := h["key"]; ok {
			headerKey = key
		}

		if value, ok := h["value"]; ok {
			headerValue = value
		}

		if headerKey != "" && headerValue != "" {
			headers.Add(NewMessageHeader(headerKey, headerValue))
		}
	}

	return headers
}
