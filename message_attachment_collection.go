package inboxroad

type MessageAttachmentCollectionInterface interface {
	Add(header MessageAttachmentInterface) MessageAttachmentCollectionInterface
	GetItems() []MessageAttachmentInterface
	ToArray() []map[string]string
	ToInboxroadArray() []map[string]string
}

type MessageAttachmentCollection struct {
	collection []MessageAttachmentInterface
}

func (m *MessageAttachmentCollection) Add(header MessageAttachmentInterface) MessageAttachmentCollectionInterface {
	m.collection = append(m.collection, header)

	return m
}

func (m *MessageAttachmentCollection) GetItems() []MessageAttachmentInterface {
	return m.collection
}

func (m MessageAttachmentCollection) ToArray() []map[string]string {
	data := make([]map[string]string, 0, len(m.GetItems()))
	for _, header := range m.GetItems() {
		data = append(data, header.ToArray())
	}

	return data
}

func (m MessageAttachmentCollection) ToInboxroadArray() []map[string]string {
	data := make([]map[string]string, 0, len(m.GetItems()))
	for _, header := range m.GetItems() {
		data = append(data, header.ToInboxroadArray())
	}

	return data
}

func NewMessageAttachmentCollection() MessageAttachmentCollectionInterface {
	return &MessageAttachmentCollection{
		collection: []MessageAttachmentInterface{},
	}
}

func NewMessageAttachmentCollectionFromArray(rawAttachments []map[string]string) MessageAttachmentCollectionInterface {
	attachments := NewMessageAttachmentCollection()

	for _, a := range rawAttachments {
		attachmentName, attachmentContent := "", ""
		if name, ok := a["name"]; ok {
			attachmentName = name
		}

		if content, ok := a["content"]; ok {
			attachmentContent = content
		}

		attachmentMimeType := "application/octet-stream"
		if mimeType, ok := a["mimeType"]; ok {
			attachmentMimeType = mimeType
		}

		if attachmentName != "" && attachmentContent != "" {
			attachments.Add(NewMessageAttachment(attachmentName, attachmentContent, attachmentMimeType))
		}
	}

	return attachments
}
