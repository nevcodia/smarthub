package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nevcodia/smarthub/service"
)

type SmartController interface {
	StorageTypes(ctx *gin.Context)
	StoreNames(ctx *gin.Context)
	Objects(ctx *gin.Context)
	ObjectsWithMetadata(ctx *gin.Context)
	GetObject(ctx *gin.Context)
	Upload(ctx *gin.Context)
	PresignUploadLink(ctx *gin.Context)
	Download(ctx *gin.Context)
	PresignDownloadLink(ctx *gin.Context)
	PresignDownloadLinkWithDuration(ctx *gin.Context)
	DeleteAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Copy(ctx *gin.Context)
	CopyAll(ctx *gin.Context)
	Move(ctx *gin.Context)
}

type smartController struct {
	service service.SmartService
}

func NewSmartController(service service.SmartService) SmartController {
	return &smartController{
		service: service,
	}
}

func (s *smartController) StorageTypes(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) StoreNames(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) Objects(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) ObjectsWithMetadata(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) GetObject(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) Upload(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) PresignUploadLink(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) Download(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) PresignDownloadLink(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) PresignDownloadLinkWithDuration(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) DeleteAll(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) Copy(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) CopyAll(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *smartController) Move(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
