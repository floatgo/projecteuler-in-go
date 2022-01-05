package main
import "fmt"
import "math"

func get_factors(num int) []int {
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
  return factors
}

func number_of_factors(num int) int {
  factors := get_factors(num)

  dict := make(map[int]int)
  for i := 0; i < len(factors); i++ {
    if _, ok := dict[factors[i]]; ok {
      dict[factors[i]] += 1
    } else {
      dict[factors[i]] = 1
    }
  }

  prod := 1
  for _, v := range dict {
    prod *= (v+1)
  }

  return prod
}

func main() {
  num := 1
  req := 500

  for {
    sum := int(num*(num+1)/2)
    if number_of_factors(sum) > req {
      fmt.Println(sum)
      break
    }
    num += 1
  }

}
