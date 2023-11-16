package msg

type ReqAuth struct {
	AppID string `json:"app_id"`
	Token string `json:"token"`
}
