package main

import "fmt"

// import "net/http"
// import "fmt"

// func soma(x int, y int) (int , bool){
// 	if(x > 10){
// 		return x + y , true
// 	}else{
// 		return x + y , false
// 	}
// }

//That's a end-point 
// func home(resp http.ResponseWriter, req *http.Request){
// 	resp.Write([]byte("Hello, World"))
// }

type Course struct{
	name string
	description string
	price float32
}

func (course Course) getCourseInfo() string{
	return fmt.Sprintf("Name: %s, Description: %s, Price: %f" , course.name , course.description, course.price)
}

func main() {
	// println("Hello, World!")

	// var a string
	// a= "hello world"
	// print(a)

	// phase := "Hello, World"
	// println(phase)

	// result , status := soma(11, 9)
	// println(result)
	// println(status)
	
	// a := 10
	// fmt.Println("Hello, World" , a)

	// http.HandleFunc("/" , home)
	// http.ListenAndServe(":9999" , nil) //Build a web server 

	course := Course{
		name: "Golang",
		description: "A good course for learning golang",
		price: 299.99,
	}

	// fmt.Println("Courses: " , course.name)
	fmt.Println(course.getCourseInfo())
}
