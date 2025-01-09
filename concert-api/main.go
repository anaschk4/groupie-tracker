package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Structure pour représenter un concert
type Concert struct {
	Artist   string `json:"artist"`
	Date     string `json:"date"`
	Location string `json:"location"`
}

// Liste de lieux (fixe pour l'exemple)
var locations = []string{"Paris, France", "New York, USA", "Berlin, Germany", "Tokyo, Japan", "London, UK"}

// Fonction pour générer une date aléatoire
func randomDate() string {
	now := time.Now()
	future := now.AddDate(1, 0, 0)
	randomTime := now.Add(time.Duration(rand.Int63n(future.Unix()-now.Unix())) * time.Second)
	return randomTime.Format("2006-01-02")
}

// Fonction pour charger les artistes depuis un fichier JSON
func loadArtists() ([]string, error) {
	file, err := os.Open("artists.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var artists []string
	if err := json.NewDecoder(file).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

// Fonction pour générer des concerts pour un artiste
func generateConcerts(artist string) []Concert {
	numConcerts := rand.Intn(5) + 1
	var concerts []Concert
	for i := 0; i < numConcerts; i++ {
		concerts = append(concerts, Concert{
			Artist:   artist,
			Date:     randomDate(),
			Location: locations[rand.Intn(len(locations))],
		})
	}
	return concerts
}

// Handler pour l'API
func concertsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Charger les artistes dynamiquement
	artists, err := loadArtists()
	if err != nil {
		http.Error(w, "500: Could not load artists", http.StatusInternalServerError)
		return
	}

	// Générer des concerts pour chaque artiste
	var allConcerts []Concert
	for _, artist := range artists {
		allConcerts = append(allConcerts, generateConcerts(artist)...)
	}

	// Retourner les données JSON
	json.NewEncoder(w).Encode(allConcerts)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/api/concerts", concertsHandler)
	http.ListenAndServe(":8080", nil)
}
