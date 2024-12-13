package main

import (
	"fmt"
	"net/http"
)

func main() {
	tasks := []string{"Watch Go Crash Course", "Watch Golang Full Course", "Build a TodoList App", "Reward myself with a cup of tea"}
	
	// Command Line Version
	// 
	// fmt.Println("### Welcome to our TodoList App! ###")
	// 
	// fmt.Println()
	// printTasks("Today's Tasks", tasks)
	// fmt.Println()
	// tasks = addTask("Learn Go", tasks)
	// printTasks("Updated Tasks", tasks)
	// fmt.Println()
	// printTasks("Today's Tasks", tasks)
	
	// HTTP Version
	http.HandleFunc("/", helloUser)
	
	http.HandleFunc("/show-tasks", showTasks("Tasks for today",tasks))
	
	http.ListenAndServe(":8080", nil)
	
}

// Http Handlers
func helloUser(writer http.ResponseWriter, request *http.Request) {
	greeting := "Hello, Welcome to our TodoList App!"
	
	fmt.Fprintln(writer, greeting)
}

func showTasks(title string,tasks []string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, title)
		for i, task := range tasks {
			fmt.Fprintf(writer, "%d. %s\n", i+1, task)
		}
	}
}

// Command Line Functions
// 
// func printTasks(title string,tasks []string) {
// 	fmt.Println(title)
// 	for i, task := range tasks {
// 		fmt.Printf("%d. %s\n", i+1, task)
// 	}
// }

// func addTask(newTask string, tasks []string) []string {
// 	updatedTasks := append(tasks, newTask)
// 	return updatedTasks
// }