package payment

import "fmt"

// Responsibility は１人当たりの支払情報を保持します。
type Responsibility struct {
	// Name は支払者の名称です。
	Name string
	// Value は金額です。
	Value int
}

// String はResponsibilityの文字列表現を返します。
func (r Responsibility) String() string {
	// fmt.Stringerインタフェースを暗黙的に実装している
	// C#のimplementsのような明示的実装はしない
	return fmt.Sprintf("%s: %d", r.Name, r.Value)
}
