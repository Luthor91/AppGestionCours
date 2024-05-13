package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"main_module/database"
	"main_module/model"
	"time"

	"github.com/redis/go-redis/v9"
)

// getAllProfesseurs est le gestionnaire HTTP qui affiche tous les professeurs
func GetAllProfesseur() ([]model.Professeur, error) {
	// Connexion à Redis
	rdb := database.ConnectToRedis()
	defer rdb.Close()

	ctx := context.Background()

	// Récupérer les étudiants depuis Redis s'ils existent
	professeurJSON, err := rdb.Get(ctx, "professeur:list").Bytes()
	if err == nil {
		var professeurs []model.Professeur
		if err := json.Unmarshal(professeurJSON, &professeurs); err != nil {
			return nil, err
		}
		return professeurs, nil
	} else if err != redis.Nil {
		// Une erreur autre que la clé n'existe pas
		return nil, err
	}

	// Si les étudiants n'existent pas dans Redis, les récupérer depuis la base de données PostgreSQL
	db := database.ConnectToPostgres()
	var professeurs []model.Professeur
	result := db.Preload("Cours").Find(&professeurs) // Charger les cours associés à chaque étudiant
	if result.Error != nil {
		return nil, result.Error
	}

	// Si les étudiants ont été récupérés depuis PostgreSQL, les stocker dans Redis
	professeurJSON, err = json.Marshal(professeurs)
	if err != nil {
		return nil, err
	}
	err = rdb.Set(ctx, "professeur:list", professeurJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return professeurs, nil
}

// GetProfesseurById récupère un professeur par son ID depuis Redis s'il existe, sinon depuis la base de données PostgreSQL
func GetProfesseurById(id int) (*model.Professeur, error) {
	if id <= 0 {
		return nil, errors.New("ID must be strictly greater than 0")
	}
	// Connexion à Redis
	rdb := database.ConnectToRedis()
	defer rdb.Close()

	ctx := context.Background()

	// Récupérer le professeur depuis Redis s'il existe
	professeurJSON, err := rdb.Get(ctx, fmt.Sprintf("professeur:%d", id)).Bytes()
	if err == nil {
		var professeur model.Professeur
		if err := json.Unmarshal(professeurJSON, &professeur); err != nil {
			return nil, err
		}
		return &professeur, nil
	} else if err != redis.Nil {
		// Une erreur autre que la clé n'existe pas
		return nil, err
	}

	// Si le professeur n'existe pas dans Redis, le récupérer depuis la base de données PostgreSQL avec les cours associés
	db := database.ConnectToPostgres()
	var professeur model.Professeur
	if err := db.Preload("Cours").First(&professeur, id).Error; err != nil {
		return nil, err
	}

	// Si le professeur a été récupéré depuis PostgreSQL, le stocker dans Redis
	professeurJSON, err = json.Marshal(professeur)
	if err != nil {
		return nil, err
	}
	err = rdb.Set(ctx, fmt.Sprintf("professeur:%d", id), professeurJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return &professeur, nil
}

// InsertProfesseur insère un étudiant dans la base de données Postgresql et dans le cache Redis
func InsertProfesseur(professeur *model.Professeur) error {

	rd := database.ConnectToRedis()
	db := database.ConnectToPostgres()

	// Insérer le professeur dans la base de données PostgreSQL
	if err := db.Create(professeur).Error; err != nil {
		return err
	}

	// Convertir le professeur en JSON
	professeurJSON, err := json.Marshal(professeur)
	if err != nil {
		return err
	}

	// Insérer le cours dans Redis
	ctx := context.Background()
	key := fmt.Sprintf("professeur:%d", professeur.ID) // Utiliser l'id du cours comme clé Redis
	expiration := 24 * time.Hour
	err = rd.Set(ctx, key, professeurJSON, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// UpdateProfesseur met à jour un professeur dans la base de données PostgreSQL et dans Redis
func UpdateProfesseur(professeur *model.Professeur) error {
	db := database.ConnectToPostgres()

	// Mettre à jour le professeur dans la base de données PostgreSQL
	if err := db.Model(&model.Professeur{}).Where("id = ?", professeur.ID).Updates(professeur).Error; err != nil {
		return err
	}

	// Mettre à jour le professeur dans Redis
	rd := database.ConnectToRedis()
	defer rd.Close()

	ctx := context.Background()

	// Convertir le professeur en JSON
	professeurJSON, err := json.Marshal(professeur)
	if err != nil {
		return err
	}

	// Mettre à jour le professeur dans Redis avec la clé correspondant à son ID
	err = rd.Set(ctx, fmt.Sprintf("professeur:%d", professeur.ID), professeurJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
