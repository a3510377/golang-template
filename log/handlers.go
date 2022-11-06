package log

type Handler struct {
	BaseRecord
	name string
}

func (h *Handler) GetName() string {
	return h.name
}

func (h *Handler) SetName() string {
	return h.name
}
