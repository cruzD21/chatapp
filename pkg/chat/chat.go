package chat

type UserInfo struct {
	ID       string
	UserName string
}

type Message struct {
	From    UserInfo
	Content []byte
}
