package model

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResWithOutData struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type UserData struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Id string `json:"id"`

}

type LoginResponse struct {
	User        UserData `json:"user"`
	AccessToken string      `json:"accessToken"`
}
