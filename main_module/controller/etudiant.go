package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"main_module/database"
	"main_module/model"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// GetAllEtudiant récupère tous les étudiants depuis Redis s'ils existent, sinon depuis la base de données PostgreSQL
func GetAllEtudiant() ([]model.Etudiant, error) {
	// Connexion à Redis
	rdb := database.ConnectToRedis()
	defer rdb.Close()

	ctx := context.Background()

	// Récupérer les étudiants depuis Redis s'ils existent
	etudiantJSON, err := rdb.Get(ctx, "etudiant:list").Bytes()
	if err == nil {
		var etudiants []model.Etudiant
		if err := json.Unmarshal(etudiantJSON, &etudiants); err != nil {
			return nil, err
		}
		return etudiants, nil
	} else if err != redis.Nil {
		// Une erreur autre que la clé n'existe pas
		return nil, err
	}

	// Si les étudiants n'existent pas dans Redis, les récupérer depuis la base de données PostgreSQL
	db := database.ConnectToPostgres()
	var etudiants []model.Etudiant
	result := db.Preload("Etudiant").Find(&etudiants) // Charger les etudiant associés à chaque étudiant
	if result.Error != nil {
		return nil, result.Error
	}

	// Si les étudiants ont été récupérés depuis PostgreSQL, les stocker dans Redis
	etudiantJSON, err = json.Marshal(etudiants)
	if err != nil {
		return nil, err
	}
	err = rdb.Set(ctx, "etudiant:list", etudiantJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return etudiants, nil
}

// GetEtudiantById récupère un etudiant par son ID depuis Redis s'il existe, sinon depuis la base de données PostgreSQL
func GetEtudiantById(id int) (*model.Etudiant, error) {
	if id <= 0 {
		return nil, errors.New("ID must be strictly greater than 0")
	}
	// Connexion à Redis
	rdb := database.ConnectToRedis()
	defer rdb.Close()

	ctx := context.Background()

	// Récupérer le etudiant depuis Redis s'il existe
	etudiantJSON, err := rdb.Get(ctx, fmt.Sprintf("etudiant:%d", id)).Bytes()
	if err == nil {
		var etudiant model.Etudiant
		if err := json.Unmarshal(etudiantJSON, &etudiant); err != nil {
			return nil, err
		}
		return &etudiant, nil
	} else if err != redis.Nil {
		// Une erreur autre que la clé n'existe pas
		return nil, err
	}

	// Si le etudiant n'existe pas dans Redis, le récupérer depuis la base de données PostgreSQL avec les cours associés
	db := database.ConnectToPostgres()
	var etudiant model.Etudiant
	if err := db.Preload("Cours").First(&etudiant, id).Error; err != nil {
		return nil, err
	}

	// Si le etudiant a été récupéré depuis PostgreSQL, le stocker dans Redis
	etudiantJSON, err = json.Marshal(etudiant)
	if err != nil {
		return nil, err
	}
	err = rdb.Set(ctx, fmt.Sprintf("etudiant:%d", id), etudiantJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return &etudiant, nil
}

// InsertEtudiant insère un étudiant dans la base de données Postgresql et dans le cache Redis
func InsertEtudiant(etudiant *model.Etudiant) error {

	rd := database.ConnectToRedis()
	db := database.ConnectToPostgres()

	// Insérer le etudiant dans la base de données PostgreSQL
	if err := db.Create(etudiant).Error; err != nil {
		return err
	}

	// Convertir le etudiant en JSON
	etudiantJSON, err := json.Marshal(etudiant)
	if err != nil {
		return err
	}

	// Insérer le cours dans Redis
	ctx := context.Background()
	key := fmt.Sprintf("etudiant:%d", etudiant.ID) // Utiliser l'id du cours comme clé Redis
	expiration := 24 * time.Hour
	err = rd.Set(ctx, key, etudiantJSON, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// UpdateInscriptionEtudiant met à jour les données de l'étudiant dans la base de données et dans Redis
func UpdateInscriptionEtudiant(etudiantID, coursID string) error {
	// Convertir les identifiants en entiers
	etudiantIDInt, err := strconv.Atoi(etudiantID)
	if err != nil {
		return err
	}
	coursIDInt, err := strconv.Atoi(coursID)
	if err != nil {
		return err
	}

	// Récupérer l'étudiant depuis Redis ou PostgreSQL
	etudiant, err := GetEtudiantById(etudiantIDInt)
	if err != nil {
		return err
	}

	// Récupérer le cours depuis Redis ou PostgreSQL
	cours, err := GetCoursById(coursIDInt)
	if err != nil {
		return err
	}

	// Ajouter le cours à la liste des cours de l'étudiant
	etudiant.Cours = append(etudiant.Cours, *cours)

	// Mettre à jour l'étudiant dans la base de données PostgreSQL
	db := database.ConnectToPostgres()
	if err := db.Save(&etudiant).Error; err != nil {
		return err
	}

	// Mettre à jour l'étudiant dans Redis
	rd := database.ConnectToRedis()
	defer rd.Close()

	ctx := context.Background()

	etudiantJSON, err := json.Marshal(etudiant)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("etudiant:%d", etudiantIDInt)
	expiration := 24 * time.Hour
	err = rd.Set(ctx, key, etudiantJSON, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// UpdateEtudiant met à jour un etudiant dans la base de données PostgreSQL et dans Redis
func UpdateEtudiant(etudiant *model.Etudiant) error {
	db := database.ConnectToPostgres()

	// Mettre à jour le etudiant dans la base de données PostgreSQL
	if err := db.Model(&model.Etudiant{}).Where("id = ?", etudiant.ID).Updates(etudiant).Error; err != nil {
		return err
	}

	// Mettre à jour le etudiant dans Redis
	rd := database.ConnectToRedis()
	defer rd.Close()

	ctx := context.Background()

	// Convertir le etudiant en JSON
	etudiantJSON, err := json.Marshal(etudiant)
	if err != nil {
		return err
	}

	// Mettre à jour le etudiant dans Redis avec la clé correspondant à son ID
	err = rd.Set(ctx, fmt.Sprintf("etudiant:%d", etudiant.ID), etudiantJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
