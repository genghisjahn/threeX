package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

const port = ":3030"

// HTML template for input form and result display
var tmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>3x + 1 Calculator</title>
</head>
<body>
	<h1>3x + 1 Calculator</h1>
	<form method="POST" action="/">
		<label for="number">Enter an integer:</label>
		<input type="number" name="number" id="number" required>
		<button type="submit">Submit</button>
	</form>
	{{if .}}
		<h2>Results:</h2>
		<p>{{.}}</p>
	{{end}}
</body>
</html>
`))

// Collatz function to calculate the 3x+1 sequence
func collatz(n int) []int {
	steps := []int{n}
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps = append(steps, n)
	}
	return steps
}

// Handler to render the form and results
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		input := r.FormValue("number")

		// Convert input to an integer
		n, err := strconv.Atoi(input)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Get the collatz result
		result := collatz(n)

		// Convert result to string for display
		resultString := fmt.Sprintf("%v", result)

		// Render the template with the result
		tmpl.Execute(w, resultString)
	} else {
		tmpl.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
