package messageSender

import (
	_ "github.com/rabbits0209/komari/utils/messageSender/bark"
	_ "github.com/rabbits0209/komari/utils/messageSender/email"
	_ "github.com/rabbits0209/komari/utils/messageSender/empty"
	_ "github.com/rabbits0209/komari/utils/messageSender/serverchan3"
	_ "github.com/rabbits0209/komari/utils/messageSender/serverchanturbo"
	_ "github.com/rabbits0209/komari/utils/messageSender/javascript"
	_ "github.com/rabbits0209/komari/utils/messageSender/telegram"
	_ "github.com/rabbits0209/komari/utils/messageSender/webhook"
)

func All() {
}
