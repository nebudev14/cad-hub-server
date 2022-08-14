package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
	"github.com/NebuDev14/cad-hub-server/helpers"
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

	// Grab size of file
	_, buffSize, err := ws.ReadMessage()

	if err != nil {
		fmt.Println(err);
	}

	size, err := strconv.Atoi(string(buffSize))

	fmt.Println("size: ")
	fmt.Println(size)

	_, tag, err := ws.ReadMessage()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tag)

	data := make([]byte, 0)

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break;
		}
		
		for i := 0; i < len(message); i++ {
			data = append(data, message[i])
		}

		// End
		if len(data) == size { 
			result := helpers.Decrypt(data, tag)
			fmt.Println(result)
		}

		if err != nil {
			fmt.Println(err)
			break
		}
	}

}
