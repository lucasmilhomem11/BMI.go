/* This is a program to calculate your IMC (BMI)
Author: Lucas Milhomem
Date Started: 04/04/21
*/

package main

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

const x, y = 703, 12

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

	feet := r.FormValue("feet")
	fheight := r.FormValue("inches")
	lweight := r.FormValue("weight")
	height, _ := strconv.ParseFloat(fheight, 64)
	feetHeight, _ := strconv.ParseFloat(feet, 64)
	weight, _ := strconv.ParseFloat(lweight, 64)
	totalHeight := (feetHeight * y) + height
	bmi := (weight * x) / math.Pow(totalHeight, 2)

	d := struct {
		Height      float64
		Inches      float64
		Weight      float64
		totalHeight float64
		BMI         float64
	}{
		Height: height,
		Inches: feetHeight,
		Weight: weight,
		BMI:    bmi,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
