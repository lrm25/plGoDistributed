package service

const (
	COMMAND_ID_LOG   = 1
	COMMAND_ID_PRINT = 2
)

const (
	REPLY_OK = 1
)

type Command struct {
	id      int
	message *string
}

type Reply struct {
	reply int
	data  *[]string
}
