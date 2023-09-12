package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubha-intelops/grpcx/bill/pkg/rest/server/models"
	"github.com/shubha-intelops/grpcx/bill/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type BillController struct {
	billService *services.BillService
}

func NewBillController() (*BillController, error) {
	billService, err := services.NewBillService()
	if err != nil {
		return nil, err
	}
	return &BillController{
		billService: billService,
	}, nil
}

func (billController *BillController) CreateBill(context *gin.Context) {
	// validate input
	var input models.Bill
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger bill creation
	if _, err := billController.billService.CreateBill(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Bill created successfully"})
}

func (billController *BillController) UpdateBill(context *gin.Context) {
	// validate input
	var input models.Bill
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger bill update
	if _, err := billController.billService.UpdateBill(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Bill updated successfully"})
}

func (billController *BillController) FetchBill(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger bill fetching
	bill, err := billController.billService.GetBill(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, bill)
}

func (billController *BillController) DeleteBill(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger bill deletion
	if err := billController.billService.DeleteBill(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Bill deleted successfully",
	})
}

func (billController *BillController) ListBills(context *gin.Context) {
	// trigger all bills fetching
	bills, err := billController.billService.ListBills()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, bills)
}

func (*BillController) PatchBill(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*BillController) OptionsBill(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*BillController) HeadBill(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
