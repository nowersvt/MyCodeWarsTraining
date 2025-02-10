package Sum_of_Digits_Digital_Root

func DigitalRoot(n int) int {
  if n == 0 {
    return 0
  }
  return (n-1)%9 + 1
}