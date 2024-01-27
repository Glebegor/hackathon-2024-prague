package route

import (
	"time"

	"ton-event-bot/api/controller"
	bootstrap "ton-event-bot/bootstrap"
	"ton-event-bot/domain"
	"ton-event-bot/repository"
	"ton-event-bot/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewSubRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup) {
	tr := repository.NewSubRepository(db, domain.CollectionSub)
	tc := &controller.SubController{
		SubUsecase: usecase.NewSubUsecase(tr, timeout),
	}
	group.GET("/sub", tc.GetAll)
	group.GET("/sub/:id", tc.GetById)
	group.POST("/sub", tc.Create)
	group.DELETE("/sub/:id", tc.Delete)
	group.PUT("/sub/:id", tc.Update)
}
