package invalid

type Invalid interface {
	Error() string
}

type Invalids struct {
	Field   string `json:"code"`
	Message string `json:"message"`
}

func New(field string, message string) Invalids {
	return Invalids{field, message}
}

func NewInternal(message string) Invalids {
	return Invalids{"internal", message}
}

func Parse(err error) *Invalids {
	switch v := err.(type) {
	case Invalids:
		return &v
	default:
		return nil
	}
}

func (e Invalids) Error() string {
	return e.Message
}
