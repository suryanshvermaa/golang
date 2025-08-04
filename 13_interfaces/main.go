package main

type dbContract interface {
	Close()
	InsertUser(username string) error
	SelectSingleUser(username string) (string, error)
}

type Application struct {
	db dbContract
}

func NewApplication(db dbContract) *Application {
	return &Application{db: db}
}

func main() {

}
