package routes

import (
	"context"

	"github.com/JulianElisii/Crud-Go-/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UseMoviesRoute(router fiber.Router) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/GO-MOvies-DB"))

	if err != nil {
		panic(err)
	}

	router.Post("/", func(c *fiber.Ctx) error {
		var movie models.Movie
		c.BodyParser(&movie)

		coll := client.Database("GO-MOvies-DB").Collection("movies")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: movie.Name,
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"Movie":  movie.Name,
			"status": "save mobie",
			"data":   result,
		})

	})

	router.Get("/", func(c *fiber.Ctx) error {
		var movies []models.Movie

		coll := client.Database("GO-MOvies-DB").Collection("movies")
		results, error := coll.Find(context.TODO(), bson.M{})

		if error != nil {
			panic(error)
		}

		for results.Next(context.TODO()) {
			var movie models.Movie
			results.Decode(&movie)         //por cada item(movie) en la base de datos vas a convertirlo a un objeto movie.
			movies = append(movies, movie) //por cada objeto que se convierte a movie lo insertas al arreglo de movies.
		}

		return c.JSON(&fiber.Map{
			"Movie": movies,
		})
	})
}
