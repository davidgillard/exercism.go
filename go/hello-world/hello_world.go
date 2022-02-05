// fichier : hello_world.go
// auteur : David GILLARD
// date : 05/02/2022

package greeting

import "fmt"

// HelloWorld greets the world.
func HelloWorld() string {
	message := "Hello, World!"
	fmt.Println(message)
	return message
}
