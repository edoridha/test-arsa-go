package db

import (
	"log"
	"test-ara/models"
	"time"
)

func Seed () {
	database := GetDBInstance()

	var filmCount int64
	var productionHouseCount int64

	database.Model(&models.Film{}).Count(&filmCount)
	database.Model(&models.ProductionHouse{}).Count(&productionHouseCount)

	if filmCount == 0 && productionHouseCount == 0 {
		log.Println("Seeding production house to database...")
		productionHouse := []models.ProductionHouse{
			{Name: "Warner Bros", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "20th Century Fox", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "Universal Pictures", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		for _, ph := range productionHouse{
			if err := database.Create(&ph).Error; err != nil {
				log.Printf("Failed to seed production house: %v", err)
			}
		}

		log.Println("Seeding film to database...")
		film := []models.Film{
			{
				Title: "Inception",
				ReleaseYear: "2010",
				Director: "Christopher Nolan",
				Sinopsis: "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
				Thumbnail: "https://m.media-amazon.com/images/M/MV5BMjExMjkwNTQ0Nl5BMl5BanBnXkFtZTcwNTY0OTk1Mw@@._V1_.jpg",
				ProductionHouseID: 1,
				CreatedAt: time.Now(), 
				UpdatedAt: time.Now(),
			},
			{
				Title: "Avatar",
				ReleaseYear: "2009",
				Director: "James Cameron",
				Sinopsis: "A paraplegic Marine dispatched to the moon Pandora on a unique mission becomes torn between following his orders and protecting the world he feels is his home.",
				Thumbnail: "https://upload.wikimedia.org/wikipedia/id/b/b0/Avatar-Teaser-Poster.jpg",
				ProductionHouseID: 2,
				CreatedAt: time.Now(), 
				UpdatedAt: time.Now(),
			},
			{
				Title: "The Dark Knight",
				ReleaseYear: "2008",
				Director: "Christopher Nolan",
				Sinopsis: "When the menace known as the Joker emerges from his mysterious past, he wreaks havoc and chaos on the people of Gotham. The Dark Knight must accept one of the greatest psychological and physical tests of his ability to fight injustice.",
				Thumbnail: "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcTfE_qrYMBZ_JB8om-34WGaZARhpX26yWRttqIDvn4_7l--UzX8mxKcPrc59IcvTpEA_G8gPA",
				ProductionHouseID: 1,
				CreatedAt: time.Now(), 
				UpdatedAt: time.Now(),
			},
			{
				Title: "Despicable Me",
				ReleaseYear: "2010",
				Director: "Pierre Coffin, Chris Renaud",
				Sinopsis: "Gru, a criminal mastermind, adopts three orphans as pawns to carry out the biggest heist in history. His life takes an unexpected turn when the little girls see the evildoer as their potential father.",
				Thumbnail: "https://upload.wikimedia.org/wikipedia/en/c/c0/Despicable_Me_%282010_animated_feature_film%29.jpg",
				ProductionHouseID: 3,
				CreatedAt: time.Now(), 
				UpdatedAt: time.Now(),
			},
			{
				Title: "Bruce Almighty",
				ReleaseYear: "2003",
				Director: "Tom Shadyac",
				Sinopsis: "Bruce, a TV reporter, is dissatisfied with his life despite his professional popularity and the love of his girlfriend. He denigrates God, who then gives Bruce divine powers to run the world.",
				Thumbnail: "https://upload.wikimedia.org/wikipedia/en/6/60/BruceAlmighty_poster.jpg",
				ProductionHouseID: 3,
				CreatedAt: time.Now(), 
				UpdatedAt: time.Now(),
			},
			{
				Title: "Lucy",
				ReleaseYear: "2014",
				Director: "Luc Besson",
				Sinopsis: "A woman, accidentally caught in a dark deal, turns the tables on her captors and transforms into a merciless warrior evolved beyond human logic.",
				Thumbnail: "https://resizing.flixster.com/yDWA3f7xiPk6iDbqe8FXacQptl8=/fit-in/352x330/v2/https://resizing.flixster.com/-XZAfHZM39UwaGJIFWKAE8fS0ak=/v3/t/assets/p10479310_p_v8_an.jpg",
				ProductionHouseID: 3,
				CreatedAt: time.Now(), 
				UpdatedAt: time.Now(),
			},
		}

		for _, f := range film{
			if err := database.Create(&f).Error; err != nil {
				log.Printf("Failed to seed film: %v", err)
			}
		}
	}
}