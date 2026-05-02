package oauth

import (
	_ "github.com/rabbits0209/komari/utils/oauth/cloudflare"
	_ "github.com/rabbits0209/komari/utils/oauth/factory"
	_ "github.com/rabbits0209/komari/utils/oauth/generic"
	_ "github.com/rabbits0209/komari/utils/oauth/github"
	_ "github.com/rabbits0209/komari/utils/oauth/qq"
)

func All() {
	//empty function to ensure all OIDC providers are registered
}