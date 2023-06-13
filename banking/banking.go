// 컴파일하지 않을 것이기 때문에 main function이 존재하지 않음.
package banking

// Account struct
type Account struct { // 외부에서 접근할 수 있게 대문자로 설정
	Owner   string // 대문자로 외부접근 허용
	Balance int    // 대문자로 외부접근 허용
}
