
함수의 다중 return value 반환
```go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLenth, upperName := lenAndUpper("nico")
	fmt.Println(totalLenth, upperName)
}

// result : 4 NICO
```

