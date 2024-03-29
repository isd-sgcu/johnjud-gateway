package image

import (
	"github.com/isd-sgcu/johnjud-gateway/src/app/dto"
)

type Service interface {
	FindByPetId(string) ([]*dto.ImageResponse, *dto.ResponseErr)
	Upload(*dto.UploadImageRequest) (*dto.ImageResponse, *dto.ResponseErr)
	Delete(string) (*dto.DeleteImageResponse, *dto.ResponseErr)
	AssignPet(*dto.AssignPetRequest) (*dto.AssignPetResponse, *dto.ResponseErr)
}
