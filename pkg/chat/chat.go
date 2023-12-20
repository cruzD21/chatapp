package chat

type UserInfo struct {
	ID       int
	UserName string
}

type Message struct {
	From    UserInfo
	Content []byte
}
