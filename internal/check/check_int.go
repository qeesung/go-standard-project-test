package check

func CheckInt(obj interface{}) bool {
  value, err := obj.(int)
  return true
} 
