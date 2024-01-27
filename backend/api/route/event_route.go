package route

import (
	"time"
	
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"ton-event-bot/domain"
	"ton-event-bot/usecase"
	"ton-event-bot/api/controller"
	"ton-event-bot/repository"
	bootstrap "ton-event-bot/bootstrap"

)

func NewEventsRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup) {
	tr := repository.NewEventRepository(db, domain.CollectionEvent)
	tc := &controller.EventController{
		TaskUsecase: usecase.NewEventUsecase(tr, timeout),
	}
	group.GET("/event", tc.GetAll)
	group.GET("/event/:id", tc.GetById)
	group.POST("/event", tc.Create)
	group.DELETE("/event/:id", tc.Delete)
	group.PUT("/event/:id", tc.Update)
}

