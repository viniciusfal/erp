package routes

import (
	"github.com/gin-gonic/gin"
	factories "github.com/viniciusfal/erp/factories/accountability"
	"github.com/viniciusfal/erp/middleware"
)

func AccountabilityRoutes(router *gin.RouterGroup) {
	CreateACController := factories.MakeCreateACC()
	GetAccsByDateController := factories.MakeGetAccsByDate()
	GetAccsByUser := factories.MakeGetAccsByUser()
	SetAccCntroller := factories.MakeAccSet()
	ChangeAccRequestController := factories.MakeAccChangeRequest()
	GetPendingController := factories.MakeGetPendingRequests()
	ApprovedAcc := factories.MakeAccApproved()
	RejectAcc := factories.MakeAccReject()

	router.POST("/accountability", middleware.RBAC("accountability.create"), CreateACController.CreateAcc)
	router.GET("/accountability/:start_date/:end_date", middleware.RBAC("accountability.view"), GetAccsByDateController.GetByDate)
	router.GET("/accountability/self/:start_date/:end_date/:resp_id", middleware.RBAC("accountability.view_self"), GetAccsByUser.GetByUser)
	router.PUT("/accountability/:id", middleware.RBAC("accountability.update"), SetAccCntroller.SetAcc)
	router.POST("/accountability/change-request", middleware.RBAC("accountability.create"), ChangeAccRequestController.ChangeACC)
	router.GET("/accountability/requests", middleware.RBAC("accountability.view_requests"), GetPendingController.GetPendingRequests)
	router.PUT("/accountability/request/:requestId/:adminId", middleware.RBAC("accountability.approve"), ApprovedAcc.ApprovedACC)
	router.PUT("/accountability/request/reject/:requestId/:adminId", middleware.RBAC("accountability.reject"), RejectAcc.Reject)
}
