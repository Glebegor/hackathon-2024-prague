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

func NewUsersubRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup) {
	tr := repository.NewUsersubRepository(db, domain.CollectionSub)
	tc := &controller.UsersubController{
		UsersubUsecase: usecase.NewUsersubUsecase(tr, timeout),
	}
	group.GET("/usersub/ByUserId/:user_id", tc.GetAll)
	// group.GET("/usersub/ByChannelId/:user_id", tc.GetAll)
	group.GET("/usersub/:user_id/:id_channel", tc.GetById)
	group.POST("/usersub", tc.Create)
	group.DELETE("/usersub/:user_id/:id_channel", tc.Delete)
	group.PUT("/usersub/:user_id/:id_channel", tc.Update)
}
