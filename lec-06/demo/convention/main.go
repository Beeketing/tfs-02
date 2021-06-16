package main

func main() {
	var path, resp string
	_, _ = path, resp

	// 1. Path should using only plural noun, without verb
	// Should
	path = "/students"
	path = "/students/1"

	// Should not
	path = "/getAllStudents"
	path = "/createStudent"
	path = "/student"

	// 2. Related resources pattern
	// Should
	path = "/schools/10/classes"
	path = "/schools/10/classes/20"

	// Should not
	path = "/schools/10/classes/20/students"
	path = "/schools/10/classes/20/students/30"

	// 3. Versioning
	// Should
	path = "/v1/students"
	path = "/v1/students/10"

	// Should not
	path = "/students/v1"
	path = "/v1.5/students/10"

	// 4. Response attribute convention
	// Should
	resp = `
	{
		"created_at": "2021-06-16T00:00:00",
		"updated_at": "2021-06-16T00:00:00"
	}
	`

	// Should not
	resp = `
	{
		"created_at": "2021-06-16T00:00:00",
		"updatedAt": "2021-06-16T00:00:00"
	}
	`

	// 5. Defined request, response content type

	// 6. Always using HTTP code for response
}
