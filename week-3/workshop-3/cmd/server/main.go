package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"route256.ru/workshop/postgres/internal/repository/animals"
)

const databaseDSN = "postgresql://postgres:password@localhost:5432/postgres"

type Animal struct {
	ID       int       `json:"id,omitempty"`
	Nickname string    `json:"nickname,omitempty"`
	Birthday time.Time `json:"birthday,omitempty"`
	Weight   int       `json:"weight,omitempty"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	r := gin.Default()

	dbpool, err := pgxpool.New(ctx, databaseDSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	if err := dbpool.Ping(ctx); err != nil {
		panic(err)
	}
	defer dbpool.Close()

	// Get a single animal
	r.GET("/animal", func(c *gin.Context) {
		animalsRepo := animals.New(dbpool)

		animalOne, err := animalsRepo.GetAnimal(c, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"animal": animalOne,
		})
	})

	// Get all animals
	r.GET("/animals", func(c *gin.Context) {
		tx, err := dbpool.Begin(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer func() {
			if err := tx.Rollback(c); err != nil {
				log.Println(err)
			}
		}()

		animalsRepo := animals.New(dbpool)

		animalsRepo = animalsRepo.WithTx(tx)

		animals, err := animalsRepo.ListAnimals(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"animals": animals,
		})
	})

	// Get all animals
	r.POST("/animal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.PUT("/animal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.DELETE("/animal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Default().Println(err)
		}
	}()

	<-ctx.Done()

	log.Default().Println("Gracefully shutting down...")
	srv.Shutdown(ctx)
	time.Sleep(time.Second * 5)
	log.Default().Println("Gracefully shut down.")
}
