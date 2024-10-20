package main

import( 
	"fmt"
	"net/http"
)



func main() {
    
	//fmt.Println()
	http.HandleFunc("/",hellUser)
	http.HandleFunc("/show-tasks",listTask)
	http.ListenAndServe(":8080", nil)


}

func hellUser(w http.ResponseWriter, r *http.Request){
	greeting := "##### welcome to our Todo App #####"
	fmt.Fprintf(w, greeting)
}

// func listTaskArr(w http.ResponseWriter, r *http.Request){
// 	arr := []string{"Task 1", "Task 2", "Task 3"}
// 	for index, task := range arr{
// 		points := index + 1
// 		//fmt.Println(points,task)
// 		fmt.Fprintf(w, "%d. %s\n", points, task)
// 	}
// }

func addTask(newTask string,arr []string) []string{
	updateArr := append(arr, newTask)
	return updateArr
}

func listTask(w http.ResponseWriter, r *http.Request){
	arr := []string{"Task 1", "Task 2", "Task 3"}
	for index, task := range arr{
		points := index + 1
		//fmt.Println(points,task)
		fmt.Fprintf(w, "%d. %s\n", points, task)
	}
}