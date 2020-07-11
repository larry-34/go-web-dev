package main

import "fmt"

func main() {
	name := "Larry Okongo"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello Larry's Back</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`
	fmt.Println(tpl)
}
