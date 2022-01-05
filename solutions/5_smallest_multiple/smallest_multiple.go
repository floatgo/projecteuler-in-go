package main
import "fmt"

func gcd(a int, b int) int {
  if b == 0 {
    return a
  }
  return gcd(b, a%b)
}

func lcm(a int, b int) int {
  return a*b/gcd(a,b)
}

func main() {
  num := 20
  result := 1
  for i := 2; i <= num; i++ {
    result = lcm(result, i)
  }
  fmt.Println(result)
}
