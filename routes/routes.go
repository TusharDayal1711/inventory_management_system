package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"inventory_management_system/handler"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"net/http"
)

func GetRoutes() *mux.Router {
	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "connection established...")
	}).Methods("GET")

	publicRoutes := mainRouter.PathPrefix("/api").Subrouter()

	//public route
	publicRoutes.HandleFunc("/user/register", handler.PublicRegister).Methods("POST")
	publicRoutes.HandleFunc("/user/login", handler.UserLogin).Methods("POST")

	//protected
	protectedRoutes := publicRoutes.NewRoute().Subrouter()
	protectedRoutes.Use(middlewares.JWTAuthMiddleware)
	protectedRoutes.HandleFunc("/users/dashboard", handler.GetUserDashboard).Methods("GET")

	//------------------asset manager and admin route-------------------
	inventoryRoutes := protectedRoutes.PathPrefix("/inventory").Subrouter()
	inventoryRoutes.Use(middlewares.RequireRole(models.AssetManagerRole, models.AdminRole))
	//post methods
	inventoryRoutes.HandleFunc("/asset", handler.AddNewAssetWithConfig).Methods("POST")
	inventoryRoutes.HandleFunc("/asset/assign", handler.AssignAssetToUser).Methods("POST")
	inventoryRoutes.HandleFunc("/asset/unassign", handler.RetrieveAsset).Methods("POST")
	inventoryRoutes.HandleFunc("/asset/service/send", handler.SendAssetToService).Methods("POST")
	inventoryRoutes.HandleFunc("/asset/service/received", handler.ReceivedFromService).Methods("POST")

	//put methods
	inventoryRoutes.HandleFunc("/asset/update", handler.UpdateAssetWithConfigHandler).Methods("PUT")

	//get methods
	inventoryRoutes.HandleFunc("/assets", handler.GetAllAssetsWithFilters).Methods("GET")
	inventoryRoutes.HandleFunc("/asset/timeline", handler.GetAssetTimeline).Methods("GET")

	//delete methods
	inventoryRoutes.HandleFunc("/asset/remove", handler.DeleteAsset).Methods("DELETE")

	//------------------Employee manager and admin routes-------------------
	employeeManagerRoutes := protectedRoutes.PathPrefix("/employee").Subrouter()
	employeeManagerRoutes.Use(middlewares.RequireRole(models.EmployeeMangerRole, models.AdminRole))

	//post methods
	employeeManagerRoutes.HandleFunc("/register", handler.RegisterEmployeeByManager).Methods("POST")
	employeeManagerRoutes.HandleFunc("/update", handler.UpdateEmployee).Methods("PUT")

	//get methods
	employeeManagerRoutes.HandleFunc("/employees", handler.GetEmployeesWithFilters).Methods("GET")
	employeeManagerRoutes.HandleFunc("/timeline", handler.GetUserTimeline).Methods("GET")

	//delete methods
	employeeManagerRoutes.HandleFunc("/remove", handler.DeleteUser).Methods("DELETE")

	//---------------------Admin-only routes-------------------
	adminRoutes := protectedRoutes.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middlewares.RequireRole(models.AdminRole))

	//post methods
	adminRoutes.HandleFunc("/employee/change-permissions", handler.ChangeUserRole).Methods("POST")

	return mainRouter
}
