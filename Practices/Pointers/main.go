package main

import "fmt"

func main() {
	var p *int
	num := 5
	p = &num
	fmt.Printf("p=%v num=%v\n", p, num) //=> p=0xc420016078 num=5
	fmt.Printf("*p = %v\n", *p)

	*p = 10
	fmt.Printf("p=%v num=%v\n", p, num) //=> p=0xc420016078 num=10

	// 配列へのポインタ
	p2 := &[3]int{1, 2, 3}

	fmt.Println((*p2)[0]) //=> 1
	fmt.Println((p2)[0])  //=> 1

	for i, v := range p2 {
		fmt.Println(i, v)
		/* 出力結果
		0 1
		1 2
		2 3
		*/
	}
}
