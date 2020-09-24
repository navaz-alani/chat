package payload

import "fmt"

type WsPayload struct {
	Type      PayloadType `json:"type"`
	Err       bool        `json:"err"`
	ErrMsg    string      `json:"errMsg"`
	Sender    string      `json:"sender"`
	Recipient string      `json:"recipient"`
	Data      string      `json:"data"`
}

func NewWsPayload(t PayloadType, recipient, sender, data string,
	err bool, errMsg string) *WsPayload {
	return &WsPayload{
		Err:       err,
		ErrMsg:    errMsg,
		Type:      t,
		Sender:    sender,
		Recipient: recipient,
		Data:      data,
	}
}

func NewTextMsg(recipient, sender, msg string) *WsPayload {
	return NewWsPayload(TextMsg, recipient, sender, msg, false, "")
}

func NewBadPayload(recipient string) *WsPayload {
	return NewWsPayload(BadPayload, recipient, SenderService, "",
		true, BadPayloadErr)
}

// todo: add messageId field in data - recipient should know which message this
// is in relation to
func NewRecipientNotConnected(recipient string) *WsPayload {
	return NewWsPayload(RecipientOffline, recipient, SenderService, "",
		true, RecipientOfflineErr)
}

func (p *WsPayload) To() string { return p.Recipient }

func (p *WsPayload) From() string { return p.Sender }

func (p *WsPayload) IsError() bool { return p.Err }

func (p *WsPayload) GetData() string { return p.Data }

func (p *WsPayload) SetError(t PayloadType, e string) {
	p.Type = t
	p.Err = true
	p.ErrMsg = e
}

func (p *WsPayload) Error() error {
	if !p.Err {
		return nil
	}
	return fmt.Errorf(p.ErrMsg)
}
