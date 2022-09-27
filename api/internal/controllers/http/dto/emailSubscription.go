package dto

type SubscribeEmail struct {
	Email string `form:"email" binding:"required"`
}
