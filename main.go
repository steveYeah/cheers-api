package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Toast struct {
	Quote  string `json:"quote"`
	Cheers string `json:"cheers"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	rand.Seed(time.Now().Unix())

	http.HandleFunc("/api", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	toast := createToast()
	toastJSON, err := json.Marshal(toast)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(toastJSON)
}

func createToast() Toast {
	toasts := [10]string{
		"Here's to those who've seen us at our best and seen us at our worst and can't tell the difference.",
		"To our wives and girlfriends ... may they never meet",
		"Here's to the floor, who will hold you when no one else will.",
		"May we get what we want, but never what we deserve.",
		"A toast to those who wish me well, and all the rest can go to hell.",
		"To Hell. May the stay there be as enjoyable as the way there.",
		"I drank to your health in company. I drank to your health alone. I drank to your health so many times...I nearly ruined my own.",
		"May the best of your past be the worse of your future.",
		"Here's to staying positive and testing negative.",
		"May you be in heaven half an hour before the devil knows you're dead.",
	}

	cheers := [10]string{
		"Na zdravi",
		"Proost",
		"Santé",
		"乾杯 (Kanpai)",
		"Noroc",
		"Skål",
		"Cheers",
		"干杯 (Gan bay)",
		"ΥΓΕΙΑ",
		"į sveikatą",
	}

	return Toast{
		toasts[rand.Intn(len(toasts))],
		cheers[rand.Intn(len(cheers))],
	}
}
