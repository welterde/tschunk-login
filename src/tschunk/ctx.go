package ctx

import bans "tschunk/bans"


func NewContext() (ctx *Context) {
	ctx = &Context{
		Bans: bans.NewBanManager(),
	}
	return
}


type Context struct {
	Bans *bans.BanManager
}
