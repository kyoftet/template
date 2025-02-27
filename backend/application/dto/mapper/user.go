package mapper

import (
	"backend/application/dto"
	"backend/domain/entity"
)

func ToUserEntity(dtoUser *dto.User) *entity.User {
	return entity.NewUser(
		dtoUser.ID,
		dtoUser.Name,
		dtoUser.MailAddress,
		dtoUser.Password,
		dtoUser.Introduction,
	)
}
