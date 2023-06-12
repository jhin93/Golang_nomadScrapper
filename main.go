package main

import "fmt"

func main() {
	//slice는 길이 제한 없는 array
	names := []string{"a", "b", "c"}
	// 요소 추가를 위해선 append 함수 사용
	names = append(names, "ee") // 새로운 slice 를 return하는 것
	fmt.Println(names)
	// [a b c ee]
}
