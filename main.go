package main

import (
	"encoding/csv"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Question string
	Answer   string
}

var questions []Question

func loadQuestions(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		questions = append(questions, Question{
			Question: record[0],
			Answer:   record[1],
		})
	}
	return nil
}

func main() {
	err := loadQuestions("questions.csv")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/quiz", quizHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func quizHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		score := 0
		for i, question := range questions {
			userAnswer := r.FormValue("answer" + strconv.Itoa(i))
			if strings.ToLower(userAnswer) == strings.ToLower(question.Answer) {
				score++
			}
		}
		tmpl := template.Must(template.ParseFiles("templates/quiz.html"))
		tmpl.Execute(w, score)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/quiz.html"))
	tmpl.Execute(w, questions)
}
