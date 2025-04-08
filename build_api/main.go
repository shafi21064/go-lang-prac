package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// / Course model
type Course struct {
	CourseId   string  `json:"course_id"`
	CourseName string  `json:"course_name"`
	Price      float32 `json:"price"`
	Password   string  `json:"-"`
	Author     *Author `json:"author"`
}

// Author model
type Author struct {
	FullName string `json:"full_name"`
	Avater   string `json:"avater"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Run server")

	router := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "Dart", Price: 532, Password: "123", Author: &Author{FullName: "shafi", Avater: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSPQi3bgz1IBGVltXBwceHqJseIgLVLxGgsW0_xeFb7zZ8vlai2J3eVPeqTmluGcYLpZ0g&usqp=CAU", Website: "linkdin.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "go-lang", Price: 232, Password: "252", Author: &Author{FullName: "shafi", Avater: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSPQi3bgz1IBGVltXBwceHqJseIgLVLxGgsW0_xeFb7zZ8vlai2J3eVPeqTmluGcYLpZ0g&usqp=CAU", Website: "linkdin.com"}})

	router.HandleFunc("/", ServeHome).Methods("GET")
	router.HandleFunc("/courses", GetAllCourse).Methods("GET")
	router.HandleFunc("/course/{id}", GetSingleCourse).Methods("GET")
	router.HandleFunc("/course", CreateNewCourse).Methods("POST")
	router.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", DeleteSingleCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

//controllers

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the course of go lang backend</h1>"))
}

// serve courses route
func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get single course data
func GetSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get single course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			break
		} else {
			json.NewEncoder(w).Encode("No course found with the id")
		}
	}
}

// create course
func CreateNewCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a new course")
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Input some data")
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
	}

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000000))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

// update course
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a new course")
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Input some data")
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
	}

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {

			// courses = append(courses[:index], courses[index+1:]... )
			courses = slices.Delete(courses, index, index+1)
			course.CourseId = params["id"]
			json.NewEncoder(w).Encode(course)
		} else {
			json.NewEncoder(w).Encode("No course found with the id")
		}
	}
}

func DeleteSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a new course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {

			// courses = append(courses[:index], courses[index+1:]... )
			courses = slices.Delete(courses, index, index+1)
			json.NewEncoder(w).Encode("Deleted the course")
			break
		} else {
			json.NewEncoder(w).Encode("No course found with the id")
		}
	}
}
