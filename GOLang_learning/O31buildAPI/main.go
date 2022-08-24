package main

import (
	//"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `jsin:"courseid"`
	CourseName  string  `jsin:"coursename"`
	CoursePrice int     `json:"pricename"`
	Auther      *Auther `json:"author"`
}

type Auther struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB slice
var courses []Course

// middleware helper
func (C *Course) IsEmpty() bool {
	//C.CourseId == "" && C.CourseName == ""
	return C.CourseName == ""
}

func main() {
	fmt.Println("API")

	router := mux.NewRouter()

	// add data
	courses = append(courses, Course{CourseId: "2", CourseName: "GoLang", CoursePrice: 5000,
		Auther: &Auther{Fullname: "Rashmi deshmukh", Website: "demo.io"}})
	courses = append(courses, Course{CourseId: "6", CourseName: "data science", CoursePrice: 5000,
		Auther: &Auther{Fullname: "Rashmi deshmukh", Website: "DSdemo.io"}})

	//routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deletOneCourse).Methods("DELET")

	//listen to port
	log.Fatal(http.ListenAndServe(":8000", router))

}

//controllers -file
//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Countent-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get singal course")
	w.Header().Set("Countent-Type", "application/json")

	//grap id from request
	params := mux.Vars(r)
	fmt.Println("mux value :", params)

	//loop through courses, find matchingf id and return the response

	for _, course := range courses {

		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create singal course")
	w.Header().Set("Countent-Type", "application/json")

	//what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")

	}

	// what about data :{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in JSON")
		return
	}

	//generate unique id,string
	//append course inti coursee slice

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("updatesingal course")
	w.Header().Set("Countent-Type", "application/json")

	//first -grap id from request
	params := mux.Vars(r)

	//loop, id ,remove,add with my ID
	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//toda :send a responce when id is not found

}

func deletOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delet course")
	w.Header().Set("Countent-Type", "application/json")

	//grap course from request for delet
	params := mux.Vars(r)

	//loop , id, reove(index,index+1)
	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course is remove")
			break

		}
	}

}
