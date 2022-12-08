package main

import (
	"github.com/IvanVojnic/cpGo.git"
	"github.com/IvanVojnic/cpGo.git/pkg/handler"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
	"github.com/IvanVojnic/cpGo.git/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	/*conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost", "quickstart-events", 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 20))
	message, _ := conn.ReadMessage(1e6)
	fmt.Println(message.Value)*/
	/*reader := kafka.NewKafkaReader()
	writer := kafka.NewKafkaWriter()

	ctx := context.Background()
	messages := make(chan kafkago.Message, 1000)
	messageCommitChan := make(chan kafkago.Message, 1000)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() (string, error) {
		return reader.FetchMessage(ctx, messages)
	})

	g.Go(func() error {
		return writer.WriteMessages(ctx, messages, messageCommitChan)
	})

	g.Go(func() error {
		return reader.CommitMessages(ctx, messageCommitChan)
	})

	err := g.Wait()
	if err != nil {
		log.Fatalln(err)
	}*/
	if err := initConfig(); err != nil {
		log.Fatalf("error init config", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to init db", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(cpGo.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
