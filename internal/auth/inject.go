package auth

import (
	handler "github.com/ant-02/myreel-plus/internal/auth/api/rpc"
	"github.com/ant-02/myreel-plus/internal/auth/application"
	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
	"github.com/ant-02/myreel-plus/internal/auth/domain/service"
	"github.com/ant-02/myreel-plus/internal/auth/infrastructure/persistence/mysql"
	"github.com/ant-02/myreel-plus/pkg/client"
)

func InjectAuthHandler() *handler.AuthHandler {
	mySQL, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}

	db := mysql.NewUserDB(mySQL)
	if err := db.Magrate(); err != nil {
		panic(err)
	}

	us := service.NewAuthenticationService(db)

	aas := application.NewAuthAppService(us, &value.IdentityFactory{})

	return handler.NewAuthController(aas)
}
