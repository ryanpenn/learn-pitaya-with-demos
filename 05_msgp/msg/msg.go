package msg

//go:generate msgp

type (
	NewUser struct {
		Name string `msg:"name"`
	}

	Message struct {
		Content string `msg:"content"`
	}

	Response struct {
		Code int    `msg:"code"`
		Msg  string `msg:"msg"`
	}
)
