package main

import "fmt"

func main() {
    // var task1 = "Task 1 is to create a todo app" 
	// var emptyVar string
	//newTask := "Task 2 is to create a todo app"
	fmt.Println("##### welcome to our Todo App #####")
	// fmt.Println(1)
	// fmt.Println(task1)
	// fmt.Println(emptyVar)
	// fmt.Println(newTask)
	//fmt.Println(task2)

	arr := []string{"Task 1", "Task 2", "Task 3"}
	// fmt.Println("tasks",arr)
	// fmt.Println(arr[0])

	for index, task := range arr{
		points := index + 1
		//fmt.Println(points,task)
		fmt.Printf("%d. %s\n", points, task)
	}



}
