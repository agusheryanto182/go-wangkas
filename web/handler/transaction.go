package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-wangkas/transaction"
	"github.com/gin-gonic/gin"
)

type TransactionHandlerWeb interface {
	New(c *gin.Context)
	Index(c *gin.Context)
	SearchByWeek(c *gin.Context)
	Create(c *gin.Context)
	Edit(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type WebTransactionHandler struct {
	transactionsService transaction.Service
}

func NewTransactionsHandler2(transactionsService transaction.Service) *WebTransactionHandler {
	return &WebTransactionHandler{transactionsService}
}

func (h *WebTransactionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "transaction_new.html", nil)
}

func (h *WebTransactionHandler) Index(c *gin.Context) {
	transactions, err := h.transactionsService.GetAllDataTransactions()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transactions})
}

func (h *WebTransactionHandler) SearchByWeek(c *gin.Context) {
	MingguKe := c.Query("MingguKe")
	ID, _ := strconv.Atoi(MingguKe)

	transactions, err := h.transactionsService.GetByWeekID(ID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transactions})
}

func (h *WebTransactionHandler) Create(c *gin.Context) {
	var input transaction.RegisterData

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "transaction_new.html", input)
		return
	}

	registerInput := transaction.Transaction{}
	registerInput.Nama = input.Nama
	registerInput.TanggalTransaksi = input.TanggalTransaksi
	registerInput.Keterangan = input.Keterangan
	registerInput.MingguKe = input.MingguKe
	registerInput.JumlahMasuk = input.JumlahMasuk

	_, err = h.transactionsService.CreateData(registerInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/transactions")
}

func (h *WebTransactionHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.transactionsService.GetByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := transaction.FormUpdateDataInput{}
	input.ID = registeredUser.ID
	input.Nama = registeredUser.Nama
	input.TanggalTransaksi = registeredUser.TanggalTransaksi
	input.Keterangan = registeredUser.Keterangan
	input.MingguKe = registeredUser.MingguKe
	input.JumlahMasuk = registeredUser.JumlahMasuk

	c.HTML(http.StatusOK, "transaction_edit.html", input)
}

func (h *WebTransactionHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input transaction.FormUpdateDataInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "transaction_edit.html", input)
		return
	}
	input.ID = id

	_, err = h.transactionsService.UpdateData(id, input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/transactions")
}
func (h *WebTransactionHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	fmt.Println(id)

	err := h.transactionsService.DeleteData(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/transactions")
}
