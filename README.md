# Groupie Tracker

Groupie Tracker est une application web permettant de rechercher des artistes, leurs informations, et d'afficher leurs concerts sur une carte interactive. Elle récupère les données via des API externes et offre une interface simple pour filtrer les résultats selon différents critères.

## Fonctionnalités

- **Recherche d'artistes** : Permet de rechercher des artistes en fonction de différents critères comme le nom, l'année de création, le premier album, ou le nombre de membres.
- **Carte interactive** : Accédez à une carte interactive pour visualiser les dates et lieux des concerts des artistes a partir de 5 villes.
- **Filtrage avancé** : Filtrez les résultats de recherche selon des critères tels que l'année de création de l'artiste, son premier album, ou le nombre de membres dans le groupe.

## Technologies utilisées

- **Go (Golang)** : Langage de programmation principal pour la logique backend.
- **HTML/CSS** : Utilisé pour structurer et styliser les pages.
- **JavaScript (AJAX)** : Permet de charger dynamiquement les résultats de recherche sans recharger la page.
- **API externes** : Utilisation d'APIs publiques pour récupérer les informations sur les artistes, leurs concerts, et les lieux.

## Installation

1. Clonez ce dépôt sur votre machine locale :
    
    git clone https://github.com/anaschk4/groupie-tracker.git
    
2. Allez dans le répertoire du projet :
    
    cd groupie-tracker
    
3. Compilez et exécutez l'application :
   
    go run main.go
    
4. Ouvrez votre navigateur et accédez à `http://localhost:8080` pour voir l'application en action.

## Auteurs

Anas Chakir, Kenzi Bendjelloul.
