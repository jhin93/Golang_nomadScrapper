package main

// fmt는 go에서 표준출력 및 입력을 하기 위한 패키지이다.
// https://brownbears.tistory.com/175
import (
	"fmt"
	"strings"
)

// 컴파일러에게 a와 b가 무엇인지 말을 해주어야 한다. (a int, b int)
// multiply도 int인 값을 return 한다고 말해야 한다.
// --> func multiply (a int, b int) 'int' {}
// a,b 둘 다 숫자일 경우, (a int, b int) 가 (a, b int)로 요약될 수 있다.
// 항상 어떤 종류의 arguments를 받는지, 어떤 종류의 return값인지 작성해야 한다.

/*
func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))

}
*/

/*		Go의 특징 : function들이 여러 개의 return 값을 가질 수 있다.	*/

// func lenAndUpper. multiple value(4 NICO)를 반환하는 function
// 1. name이라는 string을 인자로 사용한다.
// 2. return 하는 것은 name이라는 string의 길이(int)와 name(string) 그 자체(대문자로).
// 3. (int, string) <-- 이런 식으로 return 받을 값의 type들을 정의해야 한다.
// 4. len은 문자열의 길이를 구하는 함수, strings는 name을 string으로 만들어주는 go package, ToUpper는 대문자로 만들어주는 함수이다.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	// totalLength, upperName는 변수
	// lenAndUpper 사용하기. 인자는 "nico".
	totalLength, upperName := lenAndUpper("nico")
	// Go는 뭔가를 만들었으면 사용해야만 한다. 변수를 위에서 39번째 줄에서 만들었으니 41번째에서 사용.
	fmt.Println(totalLength, upperName)
}

// 만약 totalLength만 출력하려 한다면 실패한다. 1개의 variable로 2개의 value를 담을 수 없기 때문.
// totalLength, _ := lenAndUpper("nico") <-- 이런식으로 _(underscore) 를 작성한다면 컴파일러가 value값을 무시하여 가능하다.
// 결과는 4만 출력된다.
