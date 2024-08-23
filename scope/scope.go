package scope

// https://go.dev/ref/spec#Declarations_and_scope

const (
	Max int = 100
	min int = 1
)

func ReturnMin() int {
	return min
}
