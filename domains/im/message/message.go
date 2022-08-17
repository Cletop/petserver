package message

type (
	MessageType    string // 消息类型
	MessageContent struct {
		Content interface{ any } // 消息内容
	} // 消息内容类型
)

type Message struct {
	Id        int64          `json:"id"`
	From      string         `json:"from"`
	To        string         `json:"to"`
	Type      MessageType    `json:"type"`
	Content   MessageContent `json:"content"`
	CreatedAt int64          `json:"created_at"`
}
