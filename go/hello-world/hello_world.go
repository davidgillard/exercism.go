package greeting
import "fmt"

// HelloWorld greets the world.
func HelloWorld() string {
    message := "Hello, World!"
    fmt.Println(message)
    return message
}
