package payload

type PayloadType string

const (
	TextMsg          PayloadType = "text"
	BadPayload                   = "bad_payload"
	RecipientOffline             = "recipient_offline"
)

const (
	SenderService string = "__service__"
)

const (
	BadPayloadErr       string = "bad payload"
	RecipientOfflineErr        = "recipient not connected"
)

type Payload interface {
	From() string
	To() string
	GetData() string
	SetError(t PayloadType, e string)
	IsError() bool
	Error() error
}
