package inboxroad

type MessageAttachmentCollectionInterface interface {
	Add(header MessageAttachmentInterface) MessageAttachmentCollectionInterface
	GetItems() []MessageAttachmentInterface
	ToSliceMap() []map[string]string
	ToInboxroadSliceMap() []map[string]string
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

func (m MessageAttachmentCollection) ToSliceMap() []map[string]string {
	data := make([]map[string]string, 0, len(m.GetItems()))
	for _, header := range m.GetItems() {
		data = append(data, header.ToMap())
	}

	return data
}

func (m MessageAttachmentCollection) ToInboxroadSliceMap() []map[string]string {
	data := make([]map[string]string, 0, len(m.GetItems()))
	for _, header := range m.GetItems() {
		data = append(data, header.ToInboxroadMap())
	}

	return data
}

func NewMessageAttachmentCollection() MessageAttachmentCollectionInterface {
	return &MessageAttachmentCollection{
		collection: []MessageAttachmentInterface{},
	}
}

func NewMessageAttachmentCollectionFromSliceMap(
	rawAttachments []map[string]string,
) MessageAttachmentCollectionInterface {
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
