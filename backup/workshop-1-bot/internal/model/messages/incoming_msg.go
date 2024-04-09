package messages

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		return s.tgClient.SendMessage("hello", msg.UserID)
	}
	return s.tgClient.SendMessage("не знаю эту команду", msg.UserID)
}
