package server

import (
	"AvitoInternship/config"
	"AvitoInternship/internal/handlers/users_handler"
	"AvitoInternship/internal/middleware"
	"AvitoInternship/internal/tools"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	c *config.ServerConfig
	l *logrus.Entry
}

func CreateServer(c *config.ServerConfig, l *logrus.Entry) *Server {
	return &Server{c: c, l: l}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	router := r.PathPrefix("/api/v1/").Subrouter()

	usersDB := tools.PostgreConnect(s.c.RepConfig)
	ordersDB := tools.PostgreConnect(s.c.RepConfig)
	transactionsDB := tools.PostgreConnect(s.c.RepConfig)

	usersRepo := users_repo.CreateUsersRepo(usersDB)
	ordersRepo := orders_repo.CreateOrdersRepo(ordersDB)
	transactionsRepo := transactions_repo.CreateTransactionsRepo(transactionsDB)

	usersController := users_controller.CreateUsersController(usersRepo)
	ordersController := orders_controller.CreateOrdersController(ordersRepo)
	transactionsController := transactions_controller.CreateTransactionController(transactionsRepo)

	usersManager := users_manager.CreateUsersManager(usersController, transactionsController)
	ordersManager := orders_manager.CreateOrdersManager(usersController, ordersController, transactionsController)
	reportsManager := reports_manager.CreateReportsManager(usersController, ordersController, transactionsController)

	usersHandler := users_handler.CreateUsersHandler(usersManager)
	servicesHandler := services_handler.CreateServiceHandler(ordersManager)
	reportsHandler := reports_handler.CreateReportHandler(reportsManager)

	router.HandleFunc("/users", usersHandler.GetBalance).Methods("GET")
	router.HandleFunc("/users/add", usersHandler.AddBalance).Methods("POST")
	router.HandleFunc("/transfer", usersHandler.Transfer).Methods("POST")

	router.HandleFunc("/services/reserve", servicesHandler.ReserveService).Methods("POST")
	router.HandleFunc("/services/accept", servicesHandler.AcceptService).Methods("POST")
	router.HandleFunc("/services/cancel", servicesHandler.CancelService).Methods("POST")

	router.HandleFunc("/reports/users", reportsHandler.GetUserReport).Methods("GET")
	router.HandleFunc("/reports/finances", reportsHandler.GetFinanceReport).Methods("GET")

	newRouter := middleware.Log(s.l, router)
	newRouter = middleware.Panic(router)

	return http.ListenAndServe(s.c.StartPort, newRouter)
}
