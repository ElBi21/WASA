package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	// WASAText APIs paths
	rt.router.POST("/session", rt.doLogin)
	rt.router.POST("/user/:user_id/id", rt.setMyUserName)
	rt.router.POST("/user/:user_id/name", rt.setMyDisplayName)
	rt.router.POST("/user/:user_id/bio", rt.setMyBio)
	rt.router.GET("/user/:user_id/chats", rt.getMyConversations)
	rt.router.POST("/user/:user_id/photo", rt.setMyPhoto)

	rt.router.PUT("/chat", rt.createConversation)
	rt.router.GET("/chat/:chat_id", rt.getConversation)

	rt.router.PUT("/message", rt.sendMessage)
	rt.router.POST("/message/:message_id", rt.forwardMessage)
	rt.router.DELETE("/message/:message_id", rt.deleteMessage)

	rt.router.PUT("/comment", rt.commentMessage)
	rt.router.DELETE("/comment/:comment_id", rt.uncommentMessage)

	/*	rt.router.DELETE("/chat/:chat_id", rt.leaveGroup)
		rt.router.POST("/chat/:chat_id/users", rt.addToGroup)
		rt.router.POST("/chat/:chat_id/name", rt.setGroupName)
		rt.router.POST("/chat/:chat_id/photo", rt.setGroupPhoto)
		rt.router.POST("/chat/:chat_id/description", rt.setGroupDescription)



	*/

	return rt.router
}
