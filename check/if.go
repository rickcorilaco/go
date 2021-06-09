package check

// If simulates a ternary operator
/* Example:
var canFly = true

result := If(canFly, "so fly", "will fall").(string)

fmt.Println(result) // print: so fly
*/
func If(condition bool, trueValue, falseValue interface{}) (value interface{}) {
	if condition {
		return trueValue
	}

	return falseValue
}
