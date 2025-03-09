package main

import (
	"fmt"
	"slices"
)

var tasks []string

func addTask(task ...string) {
	tasks = append(tasks, task...)
}

func showTasks() {
	fmt.Printf("Tasks: %v\n", tasks)
}

func completeTask(task string) {
	var index int = -1
	for i, v := range tasks {
		if v == task {
			index = i
			break
		}
	}

	if index >= 0 {
		// tasks[:index] -> take everything before index
		// takes[index+1:] -> take everyting after index
		// `:` indicates range
		// :index -> from 0 to index
		// index: -> from index to the end
		// tasks = append(tasks[:index], tasks[index+1:]...)
		// newer approach
		tasks = slices.Delete(tasks, index, index+1)
	} else {
		fmt.Println("There is no task with name: ", task)
	}

}

func main() {
	addTask("Kupić chlib", "Kupić mleko")
	addTask("Odkurzyć mieszkanie")
	showTasks()

	completeTask("Umyć auto")
	completeTask("Kupić chlib")
	showTasks()
}
