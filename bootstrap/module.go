package bootstrap

import (
	"github.com/puspa222/high_concurrency_chat/api/controllers"
	"github.com/puspa222/high_concurrency_chat/services"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(services.NewHub),
	fx.Provide(controllers.NewChatController),
	fx.Invoke(func(hub *services.Hub) {
		go hub.Run() // Start the hub in the background
	}),
)
