package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// StartWebsite initialise le serveur web avec Fiber
func DefineRoutes() {
	// Initialisation de l'application Fiber
	app := fiber.New()

	app.Static("/", "public")

	// Gestionnaire pour l'URL racine "/"
	app.Get("/index", GetIndexHandler)

	app.Get("/cours/list", GetAllCoursHandler)
	app.Get("/etudiant/list", GetAllEtudiantHandler)
	app.Get("/professeur/list", GetAllProfesseurHandler)

	app.Get("/cours/insert", CoursInsertFormHandler)
	app.Get("/etudiant/insert", EtudiantInsertFormHandler)
	app.Get("/professeur/insert", ProfesseurInsertFormHandler)
	app.Get("/inscription/insert", InscriptionFormHandler)

	app.Get("/cours/update", CoursUpdateFormHandler)
	app.Get("/etudiant/update", EtudiantUpdateFormHandler)
	app.Get("/professeur/update", ProfesseurUpdateFormHandler)

	app.Get("/cours/:id", GetCoursByIdHandler)
	app.Get("/etudiant/:id", GetEtudiantByIdHandler)
	app.Get("/professeur/:id", GetProfesseurByIdHandler)

	app.Post("/cours/insert", CoursInsertHandler)
	app.Post("/etudiant/insert", EtudiantInsertHandler)
	app.Post("/professeur/insert", ProfesseurInsertHandler)
	app.Post("/inscription/insert", InscriptionInsertHandler)

	app.Post("/cours/update", CoursUpdateHandler)
	app.Post("/etudiant/update", EtudiantUpdateHandler)
	app.Post("/professeur/update", ProfesseurUpdateHandler)

	/*
		// route pour récupérer/afficher quelque chose
		app.Get("/object/list", getAll...)
	*/

	/*
		// route pour rajouter quelque chose dans la base de données
		app.Get("/object/add", func(c *fiber.Ctx) error {
			tmpl := template.Must(template.ParseFiles("../../public/colis/add_colis.html"))
			c.Type("html")
			return tmpl.Execute(c.Response().BodyWriter(), nil)
		})
	*/

	/*
		// route pour récupérer un objet spécifique
		app.Get("/colis/seek", func(c *fiber.Ctx) error {
			tmpl := template.Must(template.ParseFiles("../../public/colis/recherche_colis.html"))
			c.Type("html")
			return tmpl.Execute(c.Response().BodyWriter(), nil)
		})
	*/

	/*
		// route pour effectuer une action quelconque
		app.Post("/colis/affect", affectColis)
	*/

	/*
		// route pour supprimer quelque chose
		app.Delete("/colis/:id", deleteColisHandler)
	*/

	// Route de fallback pour les URL non trouvées, doit se trouver tout à la fin du code
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Page non trouvée")
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
