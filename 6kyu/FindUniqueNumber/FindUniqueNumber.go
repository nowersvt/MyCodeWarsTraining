package kata

func FindUniq(arr []float32) float32 {
	if arr[0] == arr[1] {
		for _, num := range arr {
			if num != arr[0] {
				return  num
			}
		}
	}
	if arr[0] == arr[2] {
		return arr[1]
	}
  return arr[0]
}