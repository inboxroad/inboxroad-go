package inboxroad

type MessageAttachmentCollectionInterface interface {
	Add(header MessageAttachmentInterface) MessageAttachmentCollectionInterface
	GetItems() []MessageAttachmentInterface
	ToSliceMap() SliceStringMap
	ToInboxroadSliceMap() SliceStringMap
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

func (m MessageAttachmentCollection) ToSliceMap() SliceStringMap {
	data := make(SliceStringMap, 0, len(m.GetItems()))
	for _, header := range m.GetItems() {
		data = append(data, header.ToMap())
	}

	return data
}

func (m MessageAttachmentCollection) ToInboxroadSliceMap() SliceStringMap {
	data := make(SliceStringMap, 0, len(m.GetItems()))
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

func NewMessageAttachmentCollectionFromSliceMap(rawAttachments SliceStringMap) MessageAttachmentCollectionInterface {
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
