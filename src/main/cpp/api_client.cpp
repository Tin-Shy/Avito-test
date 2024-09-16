#include "crow_all.h"
#include <pqxx/pqxx> // Для работы с PostgreSQL

std::string postgres_conn_str = "host=localhost port=5432 dbname=mydb user=myuser password=mypassword";

// Эндпоинт для проверки работы сервера (ping)
void setup_routes(crow::SimpleApp& app) {
    // GET /api/ping
    CROW_ROUTE(app, "/api/ping")
    ([](){
        return "ok";
    });

    // GET /api/tenders
    CROW_ROUTE(app, "/api/tenders")
    ([](){
        // Логика для получения списка тендеров из БД
        return crow::response(200, "[{\"id\":1,\"name\":\"Tender 1\"}]");
    });

    // POST /api/tenders/new
    CROW_ROUTE(app, "/api/tenders/new").methods(crow::HTTPMethod::Post)
    ([](const crow::request& req){
        auto json_data = crow::json::load(req.body);
        if (!json_data) {
            return crow::response(400, "Invalid JSON");
        }

        try {
            pqxx::connection conn(postgres_conn_str);
            pqxx::work txn(conn);
            std::string query = "INSERT INTO tenders (name, description, service_type, status, organization_id) "
                                "VALUES (" +
                                txn.quote(json_data["name"].s()) + ", " +
                                txn.quote(json_data["description"].s()) + ", " +
                                txn.quote(json_data["serviceType"].s()) + ", 'CREATED', " +
                                txn.quote(json_data["organizationId"].i()) + ") RETURNING id";

            pqxx::result r = txn.exec(query);
            txn.commit();

            return crow::response(200, "{\"id\":" + std::to_string(r[0][0].as<int>()) + "}");
        } catch (const std::exception& e) {
            return crow::response(500, e.what());
        }
    });

    // Пример маршрута для тендеров текущего пользователя
    CROW_ROUTE(app, "/api/tenders/my").methods(crow::HTTPMethod::Get)
    ([](const crow::request& req){
        std::string username = req.url_params.get("username");

        // Логика для получения тендеров пользователя из базы данных
        return crow::response(200, "[{\"id\":1,\"name\":\"Tender for User\"}]");
    });
}

int main() {
    crow::SimpleApp app;

    // Настройка маршрутов
    setup_routes(app);

    // Запуск сервера с настройкой адреса и порта
    app.port(8080).multithreaded().run();
}
