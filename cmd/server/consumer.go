package main

import (
	"log"
	"notes-api/database"
	"notes-api/models"
	rb "notes-api/rabit_mq"
	r "notes-api/repositories"
	"os"
)

func SaveMsgToDB(logRepo *r.LogRepository, message string) {
	newLog := models.LogMsg{
		Message: message,
	}
	logRepo.Save(&newLog)
	log.Printf("Save %s done", message)
}

func main() {

	db := database.ConnectToDB(os.Getenv("DB_PATH"))

	err := db.AutoMigrate(&models.LogMsg{})
	if err != nil {
		panic("migrate log table error")
	}
	logRepository := r.NewLogRepository(db)

	channel, queue := rb.ConnectChannel()
	defer channel.Close()
	rb.StartConsume(channel, queue, logRepository, SaveMsgToDB)
}
