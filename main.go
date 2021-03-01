package main

import (
	"fmt"
)

// Slices. 대부분의 경우 slice를 사용하게 된다.
// Go에서 slice는 기본적으로 Array이다. 하지만 length가 없다. length를 미리 입력하지 않는 것.
// slice에서 item을 추가하기 위해 사용하는 것은 append()라는 function이다.

func main() {
	names := []string{"nico", "lynn", "dal"}
	// append()는 2개의 arguments를 요구한다. 첫번째는 요소를 넣을 slice, 두번째는 넣을 요소.
	// 그러나 append()는 slice를 수정하지는 않는다. append()는 새로운 값이 추가된 slice를 return한다.
	// 수정된 slice를 기존 slice에 대입한다.
	names = append(names, "flynn")
	fmt.Println(names)
}
