package something

import (
	"fmt"
)

func sayBye() {
	fmt.Println("Bye")
}

//SayHello gonna exported.
func SayHello() {
	fmt.Println("Hello, i'm exported")
}
