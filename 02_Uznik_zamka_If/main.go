package main

import "fmt"

func Pass(A, B, C, D, E int) string {
	if (A <= D && B <= E) || (A <= E && B <= D) ||
		(A <= D && C <= E) || (A <= E && C <= D) ||
		(B <= D && C <= E) || (B <= E && C <= D) ||
		(B <= D && A <= E) || (B <= E && A <= D) ||
		(C <= D && A <= E) || (C <= E && A <= D) ||
		(C <= D && B <= E) || (C <= E && B <= D) {
		return "YES"
	}
	return "NO"
}

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Println(Pass(a, b, c, d, e))
}
