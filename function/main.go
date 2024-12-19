package main

import (
    "encoding/json"
    "html/template"
    "net/http"
    "strconv"
)

type Artist struct {
    ID             int                 `json:"id"`
    Name           string              `json:"name"`
    Image          string              `json:"image"`
    Biography      string              `json:"biography"`  // Ajout de la biographie
    DatesLocations map[string][]string `json:"datesLocations"`
}

var artists = []Artist{
    {
        ID:    1,
        Name:  "WeRenoi",
        Image: "/static/images/werenoi.jpg",
        Biography: "Werenoi est un jeune talent qui a rapidement su se faire une place dans la scène musicale. Son univers est marqué par une fusion unique de sonorités trap et mélodieuses, qui lui permet de toucher un large public.",
        DatesLocations: map[string][]string{
            "Défense ARENA":      {"2026-01-30"},
            "Thonon les Bains":         {"2025-06-04"},
            "Sion SUISSE":      {"2025-06-19"},
        },
    },
    {
        ID:    2,
        Name:  "Ninho",
        Image: "/static/images/ninho.jpg",
        Biography: "Ninho est un artiste incontournable dans le paysage du rap français. Véritable machine à hits, il s'est fait connaître avec ses projets puissants et ses collaborations avec les plus grands noms du milieu.",
        DatesLocations: map[string][]string{
            "Le dome MARSEILLE": {"2025-01-9"},
            "Halle Tony Garnier LYON":  {"2025-01-10"},
        },
    },
    {
        ID:    3,
        Name:  "Gazo",
        Image: "/static/images/gazo (copie).jpg",
        Biography: "Gazo est le pionnier du drill en France. Son style brut et percutant fait de lui une figure montante dans le rap français. Avec des morceaux puissants et des textes sans concession, il a su se forger une place à part dans la scène musicale urbaine.",
        DatesLocations: map[string][]string{
            "Le dome MARSEILLE": {"2025-01-9"},
            "Halle Tony Garnier LYON":  {"2025-01-10"},
        },
    },
    {
        ID:    4,
        Name:  "Hamza",
        Image: "/static/images/hamza.jpg",
        Biography: "",
        DatesLocations: map[string][]string{
            "Le dome MARSEILLE": {"2025-01-9"},
            "Halle Tony Garnier LYON":  {"2025-01-10"},
        },
    },
    {
        ID:    4,
        Name:  "maes",
        Image: "/static/images/hamza.jpg",
        Biography: "",
        DatesLocations: map[string][]string{
            "Le dome MARSEILLE": {"2025-01-9"},
            "Halle Tony Garnier LYON":  {"2025-01-10"},
        },
    },
    // Ajoutez les autres artistes ici
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/api/artists", artistsHandler)
    http.HandleFunc("/api/artist/", artistHandler)

    // Servir les fichiers statiques (images, CSS, JS)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Lancer le serveur
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, artists)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(artists)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
    // Extraire l'ID de l'URL
    id := r.URL.Path[len("/api/artist/"):]

    // Chercher l'artiste correspondant à l'ID dans la liste
    for _, artist := range artists {
        if strconv.Itoa(artist.ID) == id {  // Convertir l'ID de l'artiste en string pour la comparaison
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(artist)
            return
        }
    }
    http.NotFound(w, r) // Si l'artiste n'est pas trouvé, renvoyer une erreur 404
}
