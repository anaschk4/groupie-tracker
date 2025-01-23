package main

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "strings"
)

type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
    Index []struct {
        ID        int      `json:"id"`
        Locations []string `json:"locations"`
    } `json:"index"`
}

type Date struct {
    Index []struct {
        ID    int      `json:"id"`
        Dates []string `json:"dates"`
    } `json:"index"`
}

type Relation struct {
    Index []struct {
        ID             int                 `json:"id"`
        DatesLocations map[string][]string `json:"datesLocations"`
    } `json:"index"`
}

type ArtistComplete struct {
    Artist
    Locations      []string
    Dates          []string
    DatesLocations map[string][]string
}

type PageData struct {
    Query   string
    Results []ArtistComplete
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/search", handleSearch)
    http.HandleFunc("/map", handleMap)
    http.HandleFunc("/profil", handleProfil)
    log.Println("Serveur démarré sur http://localhost:3030")
    log.Fatal(http.ListenAndServe(":3030", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    creationDate := r.URL.Query().Get("creationDate")
    firstAlbum := r.URL.Query().Get("firstAlbum")
    memberCount := r.URL.Query().Get("memberCount")

    log.Println("Requête reçue - Query:", query, "Creation Date:", creationDate, "First Album:", firstAlbum, "Member Count:", memberCount)

    if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
        artists, err := searchArtists(query, creationDate, firstAlbum, memberCount)
        if err != nil {
            log.Println("Erreur lors de la recherche des artistes :", err)
            http.Error(w, "Une erreur est survenue", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(artists)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/search.html"))
    data := PageData{
        Query: query,
    }
    if query != "" || creationDate != "" || firstAlbum != "" || memberCount != "" {
        artists, err := searchArtists(query, creationDate, firstAlbum, memberCount)
        if err == nil {
            data.Results = artists
        } else {
            log.Println("Erreur lors de la recherche (non-AJAX) :", err)
        }
    }
    tmpl.Execute(w, data)
}

func handleMap(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/map.html"))
    tmpl.Execute(w, nil)
}

// New handler for profile page
func handleProfil(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/profil.html"))
    tmpl.Execute(w, nil)
}

func fetchAPI[T any](url string) (T, error) {
    var result T
    resp, err := http.Get(url)
    if err != nil {
        return result, err
    }
    defer resp.Body.Close()

    log.Println("Appel API à l'URL :", url)
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        log.Println("Erreur lors du décodage JSON :", err)
    }
    return result, err
}

func searchArtists(query, creationDate, firstAlbum, memberCount string) ([]ArtistComplete, error) {
    artists, err := fetchAPI[[]Artist]("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        return nil, err
    }

    locations, err := fetchAPI[Location]("https://groupietrackers.herokuapp.com/api/locations")
    if err != nil {
        return nil, err
    }

    dates, err := fetchAPI[Date]("https://groupietrackers.herokuapp.com/api/dates")
    if err != nil {
        return nil, err
    }

    relations, err := fetchAPI[Relation]("https://groupietrackers.herokuapp.com/api/relation")
    if err != nil {
        return nil, err
    }

    var completeArtists []ArtistComplete
    query = strings.ToLower(query)

    for _, artist := range artists {
        creationDateStr := strconv.Itoa(artist.CreationDate)

        // Application des filtres
        if (query == "" || strings.Contains(strings.ToLower(artist.Name), query)) &&
            (creationDate == "" || creationDate == creationDateStr) &&
            (firstAlbum == "" || strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(firstAlbum))) {

            // Filtrage par nombre de membres
            if memberCount != "" {
                switch memberCount {
                case "1":
                    if len(artist.Members) != 1 {
                        continue
                    }
                case "2":
                    if len(artist.Members) != 2 {
                        continue
                    }
                case "3+":
                    if len(artist.Members) < 3 {
                        continue
                    }
                }
            }
            // Filtrage par nombre de membres
            if firstAlbum != "" {
                switch firstAlbum {
                case "1990":
                    if len(artist.FirstAlbum) != 1990 {
                        continue
                    }
                case "2000":
                    if len(artist.FirstAlbum) != 2000 {
                        continue
                    }
                case "2010":
                    if len(artist.FirstAlbum) < 2010 {
                        continue
                    }
                }
            }

            complete := ArtistComplete{
                Artist: artist,
            }

            for _, loc := range locations.Index {
                if loc.ID == artist.ID {
                    complete.Locations = loc.Locations
                    break
                }
            }

            for _, date := range dates.Index {
                if date.ID == artist.ID {
                    complete.Dates = date.Dates
                    break
                }
            }

            for _, relation := range relations.Index {
                if relation.ID == artist.ID {
                    complete.DatesLocations = relation.DatesLocations
                    break
                }
            }

            completeArtists = append(completeArtists, complete)
        }
    }

    log.Println("Résultats trouvés :", len(completeArtists))
    return completeArtists, nil
}