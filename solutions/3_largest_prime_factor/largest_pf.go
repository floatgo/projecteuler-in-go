package main
import "fmt"
import "math"


func main() {
  num := 600851475143
  
  factors := make([]int, 0)
  for num % 2 == 0 {
    factors = append(factors, 2)
    num = num/2
  }

  for i := 3; i < int(math.Sqrt(float64(num)))+1; i += 2 {
    for num % i == 0 {
      num = num / i
      factors = append(factors, i)
    }
  }

  if num > 2 {
    factors = append(factors, num)
  }

  max := factors[0]
  for i := 0; i < len(factors); i++ {
    if factors[i] > max {
      max = factors[i]
    }
  }

  fmt.Println(factors)
  fmt.Println(max)
}
