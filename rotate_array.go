package compute

func Rotate(arr []string, index int) []string {
	index = index % len(arr)
	new_arr := []string(nil)
	new_arr = append(new_arr, arr[index:]...)
	new_arr = append(new_arr, arr[:index]...)
	return new_arr
}
