package main

import "fmt"

// Arrays and Slices
// Go에서는 Array를 만들 때 길이와 타입을 명시해주어야 한다.

func main() {
	names := [5]string{"nico", "lynn", "dal"}
	fmt.Println(names)
	// 결과 [nico lynn dal  ]
	// 이 시점(fmt.Println(names)) 이후로 names에 추가로 값을 대입했다면
	// 새롭게 fmt.Println(names)를 실행해 줘야 추가로 대입한 값을 확인할 수 있다.
	fmt.Println(names[3])
	// 결과 (공백)
	// 오류가 나진 않고 인덱스에 밸류가 빈 상태로 출력됨.
	names[3] = "lalala"
	fmt.Println(names)
	// 결과 [nico lynn dal lalala ]
	// 처음 실행한 fmt.Println(names)에서는 'lalala'를 확인할 수 없다. 이후에 추가로 대입한 값이기 때문.
}
