package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func StartUpdates(group *gin.RouterGroup) {
	group.GET("", UpdateBase)
}

func UpdateBase(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err);
		return
	}

	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break;
		}

		// if string(message) == "hello" {
		// 	message = []byte("Whats up")
		// 	fmt.Println("whats up")
		// }
		fmt.Println("Received " + string(message))

		err = ws.WriteMessage(mt, []byte("received"))
		if err != nil {
			fmt.Println(err)
			break
		}

	}

}
