package server

import (
	"fmt"
	"html/template"
	"main_module/controller"
	"main_module/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetIndexHandler(c *fiber.Ctx) error {
	tmpl := template.Must(template.ParseFiles("../public/html/index.html"))
	err := tmpl.Execute(c.Response().BodyWriter(), nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

// GetAllCoursHandler est le gestionnaire HTTP qui affiche tous les cours
func GetAllCoursHandler(c *fiber.Ctx) error {
	all_cours, err := controller.GetAllCours()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des cours")
	}

	// Passer les données au template et générer la réponse HTTP
	page_liste_cours := template.Must(template.ParseFiles("../public/html/liste-cours.html"))
	err = page_liste_cours.Execute(c.Response().BodyWriter(), all_cours)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

// GetAllEtudiantHandler est le gestionnaire HTTP qui affiche tous les étudiants
func GetAllEtudiantHandler(c *fiber.Ctx) error {
	all_etudiant, err := controller.GetAllEtudiant()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des cours")
	}

	// Passer les données au template et générer la réponse HTTP
	page_liste_etudiant := template.Must(template.ParseFiles("../public/html/liste-etudiant.html"))
	err = page_liste_etudiant.Execute(c.Response().BodyWriter(), all_etudiant)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

// GetAllEtudiantHandler est le gestionnaire HTTP qui affiche tous les étudiants
func GetAllProfesseurHandler(c *fiber.Ctx) error {
	all_professeur, err := controller.GetAllProfesseur()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des professeur")
	}

	// Passer les données au template et générer la réponse HTTP
	page_liste_professeur := template.Must(template.ParseFiles("../public/html/liste-professeur.html"))
	err = page_liste_professeur.Execute(c.Response().BodyWriter(), all_professeur)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func GetCoursByIdHandler(c *fiber.Ctx) error {
	coursID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID de cours invalide")
	}

	cours, err := controller.GetCoursById(coursID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération du cours")
	}

	// Passer les données au template et générer la réponse HTTP
	page_cours := template.Must(template.ParseFiles("../public/html/cours.html"))
	err = page_cours.Execute(c.Response().BodyWriter(), cours)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func GetProfesseurByIdHandler(c *fiber.Ctx) error {
	professeurID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID du professeur invalide")
	}

	professeur, err := controller.GetProfesseurById(professeurID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des professeur")
	}

	// Passer les données au template et générer la réponse HTTP
	page_professeur := template.Must(template.ParseFiles("../public/html/professeur.html"))
	err = page_professeur.Execute(c.Response().BodyWriter(), professeur)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func GetEtudiantByIdHandler(c *fiber.Ctx) error {
	etudiantID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID de l'étudiant invalide")
	}

	etudiant, err := controller.GetEtudiantById(etudiantID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la récupération des etudiant")
	}

	// Passer les données au template et générer la réponse HTTP
	page_etudiant := template.Must(template.ParseFiles("../public/html/etudiant.html"))
	err = page_etudiant.Execute(c.Response().BodyWriter(), etudiant)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func CoursInsertFormHandler(c *fiber.Ctx) error {

	page_cours := template.Must(template.ParseFiles("../public/html/add-cours.html"))

	err := page_cours.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func EtudiantInsertFormHandler(c *fiber.Ctx) error {
	page_etudiant := template.Must(template.ParseFiles("../public/html/add-etudiant.html"))

	err := page_etudiant.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func ProfesseurInsertFormHandler(c *fiber.Ctx) error {
	page_professeur := template.Must(template.ParseFiles("../public/html/add-professeur.html"))

	err := page_professeur.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func CoursUpdateFormHandler(c *fiber.Ctx) error {

	page_cours := template.Must(template.ParseFiles("../public/html/update-cours.html"))

	err := page_cours.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func EtudiantUpdateFormHandler(c *fiber.Ctx) error {
	page_etudiant := template.Must(template.ParseFiles("../public/html/update-etudiant.html"))

	err := page_etudiant.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func ProfesseurUpdateFormHandler(c *fiber.Ctx) error {
	page_professeur := template.Must(template.ParseFiles("../public/html/update-professeur.html"))

	err := page_professeur.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func InscriptionFormHandler(c *fiber.Ctx) error {
	page_professeur := template.Must(template.ParseFiles("../public/html/add-inscription.html"))

	err := page_professeur.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

// CoursFormHandler est le gestionnaire HTTP pour soumettre le formulaire d'ajout de produit
func CoursInsertHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	titre := c.FormValue("titre")
	niveau := c.FormValue("niveau")
	resume := c.FormValue("resume")
	places_str := c.FormValue("places")
	expiration_str := c.FormValue("expiration")

	// Convertir le prix en uint
	places, err := strconv.Atoi(places_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Le poids doit être un nombre valide")
	}
	expiration, err := strconv.Atoi(expiration_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Le poids doit être un nombre valide")
	}

	// Créer un nouveau produit avec les données du formulaire
	new_cours := model.Cours{
		Identifiant: controller.GenerateRandomString(8), Titre: titre, Niveau: niveau, Resume: resume, PlaceDisponible: uint(places), TempsExpiration: uint(expiration)}

	// Ajouter le nouveau produit à la base de données
	if err := controller.InsertCours(&new_cours); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'ajout du cours")
	}

	// Rediriger l'utilisateur vers la page des produits après l'ajout réussi
	return c.Redirect("/cours/list")
}

// EtudiantFormHandler est le gestionnaire HTTP pour soumettre le formulaire d'ajout de produit
func EtudiantInsertHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	nom := c.FormValue("nom")

	// Créer un nouveau produit avec les données du formulaire
	new_etudiant := model.Etudiant{
		Identifiant: controller.GenerateRandomString(8), Nom: nom}

	// Ajouter le nouveau produit à la base de données
	if err := controller.InsertEtudiant(&new_etudiant); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'ajout de l'étudiant")
	}

	// Rediriger l'utilisateur vers la page des produits après l'ajout réussi
	return c.Redirect("/etudiant/list")
}

// ProfesseurFormHandler est le gestionnaire HTTP pour soumettre le formulaire d'ajout de produit
func ProfesseurInsertHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	nom := c.FormValue("nom")

	// Créer un nouveau produit avec les données du formulaire
	new_professeur := model.Professeur{
		Identifiant: controller.GenerateRandomString(8), Nom: nom}

	// Ajouter le nouveau produit à la base de données
	if err := controller.InsertProfesseur(&new_professeur); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'ajout de l'étudiant")
	}

	// Rediriger l'utilisateur vers la page des produits après l'ajout réussi
	return c.Redirect("/professeur/list")
}

// CoursUpdateHandler est le gestionnaire HTTP pour soumettre le formulaire de modification de cours
func CoursUpdateHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	coursId_str := c.FormValue("coursId")
	titre := c.FormValue("titre")
	niveau := c.FormValue("niveau")
	resume := c.FormValue("resume")
	places_str := c.FormValue("places")
	expiration_str := c.FormValue("expiration")

	// Convertir en int
	coursId, err := strconv.Atoi(coursId_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("L'identifiant du cours doit être un nombre valide")
	}
	places, err := strconv.Atoi(places_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Le nombre de places doit être un nombre valide")
	}
	expiration, err := strconv.Atoi(expiration_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("La durée d'expiration doit être un nombre valide")
	}

	// Mettre à jour le cours dans la base de données PostgreSQL
	updatedCours := model.Cours{
		ID:              uint(coursId),
		Titre:           titre,
		Niveau:          niveau,
		Resume:          resume,
		PlaceDisponible: uint(places),
		TempsExpiration: uint(expiration),
	}

	if err := controller.UpdateCours(&updatedCours); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la mise à jour du cours")
	}

	// Rediriger l'utilisateur vers la page des cours après la modification réussie
	return c.Redirect(fmt.Sprintf("/cours/%d", coursId))
}

// EtudiantFormHandler est le gestionnaire HTTP pour soumettre le formulaire d'ajout de produit
func EtudiantUpdateHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	etudiantId_str := c.FormValue("etudiantId")
	nom := c.FormValue("nom")

	// Convertir en int
	etudiantId, err := strconv.Atoi(etudiantId_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("L'id du cours doit être un nombre valide")
	}

	// Mettre à jour le cours dans la base de données PostgreSQL
	updatedEtudiant := model.Etudiant{
		ID:  int(etudiantId),
		Nom: nom,
	}

	if err := controller.UpdateEtudiant(&updatedEtudiant); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la mise à jour du cours")
	}

	// Rediriger l'utilisateur vers la page des cours après la modification réussie
	return c.Redirect(fmt.Sprintf("/etudiant/%d", etudiantId))
}

// ProfesseurFormHandler est le gestionnaire HTTP pour soumettre le formulaire d'ajout de produit
func ProfesseurUpdateHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	professeurId_str := c.FormValue("professeurId")
	nom := c.FormValue("nom")

	// Convertir en int
	professeurId, err := strconv.Atoi(professeurId_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("L'identifiant du cours doit être un nombre valide")
	}

	// Mettre à jour le cours dans la base de données PostgreSQL
	updatedProfesseur := model.Professeur{
		ID:  int(professeurId),
		Nom: nom,
	}

	if err := controller.UpdateProfesseur(&updatedProfesseur); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la mise à jour du cours")
	}

	// Rediriger l'utilisateur vers la page des cours après la modification réussie
	return c.Redirect(fmt.Sprintf("/professeur/%d", professeurId))

}

// InscriptionInsertHandler est le gestionnaire HTTP pour soumettre le formulaire d'inscription à un cours
func InscriptionInsertHandler(c *fiber.Ctx) error {
	// Vérifier si la méthode de requête est POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Méthode non autorisée")
	}

	// Récupérer les données du formulaire
	etudiantID := c.FormValue("etudiantId")
	coursID := c.FormValue("coursId")

	// Vérifier si les identifiants sont vides
	if etudiantID == "" || coursID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Identifiant de l'étudiant ou du cours manquant")
	}

	// Mettre à jour l'étudiant en ajoutant le cours
	if err := controller.UpdateInscriptionEtudiant(etudiantID, coursID); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'ajout du cours à l'étudiant")
	}

	// Rediriger l'utilisateur vers une page de confirmation
	return c.Redirect("/index")
}

func UpdateCoursHandler(c *fiber.Ctx) error {

	page_cours := template.Must(template.ParseFiles("../public/html/update-cours.html"))

	err := page_cours.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func UpdateEtudiantHandler(c *fiber.Ctx) error {

	page_cours := template.Must(template.ParseFiles("../public/html/update-etudiant.html"))

	err := page_cours.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}

func UpdateProfesseurHandler(c *fiber.Ctx) error {

	page_cours := template.Must(template.ParseFiles("../public/html/update-professeur.html"))

	err := page_cours.Execute(c.Response().BodyWriter(), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la génération de la réponse HTTP")
	}

	c.Type("html")

	return nil
}
