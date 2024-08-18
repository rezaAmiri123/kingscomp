package webapp

func(w *Webapp)urls(){
	lobby:=w.e.Group("/lobby")
	lobby.GET("/:lobby_id", w.lobbyIndex)

	auth := w.e.Group("/auth")
	auth.POST("/validate", w.validatInitData)
}
