## Tech Stack

Language: Go
HTTP Router: net/http
Database: PostgreSQL
Cache & Temporary Holds: Redis
Real-time Updates: WebSockets
SQL: Raw SQL + pgx
Migrations: golang-migrate
Config: .env
Logging: log/slog
Docker Compose
Graceful shutdown
Context propagation
Repository pattern only where it adds value

---

domain/ => Pure business entities. Just business models.
handler/ => HTTP Request -> Decode JSON -> Validate -> Call Service -> Encode Response
service/ => Entire business logic lives here.
repository/postgres => Only database interaction.
repository/redis => Responsible for: SET NX EX / GET / DEL / TTL / Pub/Sub
websocket/ => connection -> register -> unregister -> broadcast -> cleanup
