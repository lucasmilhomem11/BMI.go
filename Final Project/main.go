/* This is a program to calculate your IMC (BMI)
Author: Lucas Milhomem
Date Started: 04/04/21
*/

package main

import (
	"strconv"
	"log"
	"math"
	"net/http"
	"html/template"
)

const x = 703
var tpl *template.Template


func init() {
	tpl = template.Must(template.ParseGlob("view/*.gohtml"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "FirstView.gohtml", nil)
}
func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc("/process", processor)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	fheight := r.FormValue("height")
	lweight := r.FormValue("weight")
	height, _ := strconv.ParseFloat(fheight, 64)
	weight, _ := strconv.ParseFloat(lweight, 64)
	bmi := (weight * x) / math.Pow(height, 2)

	d := struct{
		Height float64
		Weight float64
		BMI float64
	}{
		Height: height,
		Weight: weight,
		BMI: bmi,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml",d )
}
