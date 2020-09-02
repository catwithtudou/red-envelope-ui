package red_envelope_ui

import (
	_ "github.com/catwithtudou/red-envelope-account/core/accounts"
	"github.com/catwithtudou/red-envelope-infra"
	"github.com/catwithtudou/red-envelope-infra/base"
	_ "red-envelope-ui/views"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
