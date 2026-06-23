package basics

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, Interfaces, Enums

	// snake_case
	// Eg. user_id, first_name, http_request

	// UPPERCASE -> Constants

	// mixedCase
	// Eg. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5
}
