// payment パッケージは、支払金額の計算機能を提供します。
package payment

import "errors"

// Calclator は１人当たりの支払額を計算します。
type Calclator struct {
	// membersとpriceは小文字で始まるので同じpackageでしか参照できない(C#でいう名前空間レベル)
	// packageが最小のアクセスレベルでprivateはない

	members []string
	price   int
}

// NewCalculator は新しいCalculatorのインスタンスを生成して返します。
func NewCalculator(members []string, price int) (*Calclator, error) {
	if len(members) == 0 {
		// errorが発生しうる場合、多値を返すようにして最後をerrorにする
		return nil, errors.New("no one joined in this party")
	}

	c := Calclator{
		members: members,
		price:   price,
	}
	return &c, nil
}

// Calculate は計算を実行してResponsibilityのスライスを返します。
func (c *Calclator) Calculate() []Responsibility {
	num := len(c.members)
	value := c.price / num
	mod := c.price % num

	// make()を使うと最初から指定の要素数(この場合はnum)を持ったスライス(可変長配列のようなもの)を作れる
	rr := make([]Responsibility, num)

	// スライス(c.members)のループ
	// iがインデックスでmが要素
	for i, m := range c.members {
		v := value
		if i == 0 {
			v += mod
		}

		rr[i] = Responsibility{
			Name:  m,
			Value: v,
		}
	}

	return rr
}
