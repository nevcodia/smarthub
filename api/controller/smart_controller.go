package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nevcodia/smarthub/domain"
	"github.com/nevcodia/smarthub/service"
	"net/http"
	"strconv"
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
	types := []string{
		domain.S3.String(),
		//string(domain.FTP),
		//string(domain.SHAREPOINT),
	}
	ctx.JSON(http.StatusOK, types)
}

func (s *smartController) StoreNames(ctx *gin.Context) {
	storageType := s.ExtractStorageType(ctx)
	storeNames, err := s.service.StoreNames(storageType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, storeNames)
	}
}

func (s *smartController) Objects(ctx *gin.Context) {
	storageType := s.ExtractStorageType(ctx)
	storeName := ctx.Query("storeName")
	maxObjectPerPage := ctx.DefaultQuery("maxObjectPerPage", "1000")
	maxKeys, err := strconv.ParseInt(maxObjectPerPage, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	page := ctx.DefaultQuery("page", "0")
	currentPage, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	prefix := ctx.Query("prefix")
	objects, err := s.service.Objects(storageType, storeName, int32(maxKeys), int32(currentPage), prefix)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, objects)
}

func (s *smartController) ObjectsWithMetadata(ctx *gin.Context) {
	storageType := s.ExtractStorageType(ctx)
	storeName := ctx.Query("storeName")
	maxObjectPerPage := ctx.DefaultQuery("maxObjectPerPage", "1000")
	maxKeys, err := strconv.ParseInt(maxObjectPerPage, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	page := ctx.DefaultQuery("page", "0")
	currentPage, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	prefix := ctx.Query("prefix")
	objects, err := s.service.ObjectsWithMetadata(storageType, storeName, int32(maxKeys), int32(currentPage), prefix)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, objects)
}

func (s *smartController) GetObject(ctx *gin.Context) {
	storageType := s.ExtractStorageType(ctx)
	storeName := ctx.Query("storeName")
	key := ctx.Query("key")
	params := &domain.ObjectParams{
		StoreName: storeName,
		Key:       key,
	}
	objects, err := s.service.GetObject(storageType, params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, objects)
}

func (s *smartController) Upload(ctx *gin.Context) {
	storageType := s.ExtractStorageType(ctx)
	file, err := ctx.FormFile("file")
	storeName := ctx.Request.PostFormValue("storeName")
	key := ctx.Request.PostFormValue("key")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	params := &domain.ObjectParams{
		StoreName: storeName,
		Key:       key,
	}
	response, err := s.service.UploadMultiPart(storageType, params, map[string]string{}, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
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

func (s *smartController) ExtractStorageType(ctx *gin.Context) domain.StorageType {
	sType := ctx.Param("type")
	return domain.StorageTypeFromValue(sType)
}
