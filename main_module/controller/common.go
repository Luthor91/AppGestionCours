package controller

import (
	"fmt"
	"main_module/model"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// Fonction utilitaire pour générer une chaîne aléatoire
func GenerateRandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func InsertTestDatas(db *gorm.DB) error {
	// Insérer des enregistrements pour la table Utilisateur
	utilisateurs := []model.Utilisateur{
		{Nom: GenerateRandomString(8), Email: "utilisateur1@example.com", MotDePasse: "password1"},
		{Nom: GenerateRandomString(8), Email: "utilisateur2@example.com", MotDePasse: "password2"},
	}
	for _, utilisateur := range utilisateurs {
		if err := db.Create(&utilisateur).Error; err != nil {
			fmt.Println("Erreur lors de l'insertion de l'utilisateur:", err)
			return err
		}
		fmt.Println("Utilisateur inséré avec succès:", utilisateur.Nom)
	}

	// Insérer des enregistrements pour la table Cours
	cours := []model.Cours{
		{Identifiant: GenerateRandomString(8), Titre: "Mathématique", Resume: "Resume 1", Niveau: "Débutant", PlaceDisponible: 10, TempsExpiration: 3600},
		{Identifiant: GenerateRandomString(8), Titre: "Titre 2", Resume: "Resume 2", Niveau: "Intermédiaire", PlaceDisponible: 20, TempsExpiration: 7200},
	}
	for _, cour := range cours {
		if err := db.Create(&cour).Error; err != nil {
			fmt.Println("Erreur lors de l'insertion du cours:", err)
			return err
		}
		fmt.Println("Cours inséré avec succès:", cour.Titre)
	}

	// Insérer des enregistrements pour la table Etudiant
	etudiants := []model.Etudiant{
		{Nom: "Etudiant 1", Identifiant: GenerateRandomString(8), UtilisateurID: 1, Cours: []model.Cours{cours[0]}}, // Associer un cours à l'étudiant
		{Nom: "Etudiant 2", Identifiant: GenerateRandomString(8), UtilisateurID: 1, Cours: []model.Cours{cours[1]}}, // Associer un autre cours à un autre étudiant
	}
	for _, etudiant := range etudiants {
		if err := db.Create(&etudiant).Error; err != nil {
			fmt.Println("Erreur lors de l'insertion de l'étudiant:", err)
			return err
		}
		fmt.Println("Étudiant inséré avec succès:", etudiant.Nom)
	}

	// Insérer des enregistrements pour la table Professeur
	professeurs := []model.Professeur{
		{Nom: "Professeur 1", Identifiant: GenerateRandomString(8), UtilisateurID: 1, Cours: []model.Cours{cours[0]}}, // Associer un cours au professeur
		{Nom: "Professeur 2", Identifiant: GenerateRandomString(8), UtilisateurID: 1, Cours: []model.Cours{cours[1]}}, // Associer un autre cours à un autre professeur
	}
	for _, professeur := range professeurs {
		if err := db.Create(&professeur).Error; err != nil {
			fmt.Println("Erreur lors de l'insertion du professeur:", err)
			return err
		}
		fmt.Println("Professeur inséré avec succès:", professeur.Nom)
	}

	fmt.Println("Tous les enregistrements ont été insérés avec succès.")
	return nil
}
