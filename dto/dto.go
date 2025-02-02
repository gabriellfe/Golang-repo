package dto

type ReturnDto struct {
	Name string `json:"nome"`
}

type ErrorDto struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type ErrorSchemaDto struct {
	Errors []string `json:"Errors"`
}

type UserSearchDto struct {
	Code   string  `json:"code" validate:"string,min=5,max=20"`
	Status float64 `json:"status" validate:"number,min=5,max=20"`
	Inn    uint8   `json:"inn" validate:"number,min=5,max=20"`
	Mail   string  `validate:"email"`
}
