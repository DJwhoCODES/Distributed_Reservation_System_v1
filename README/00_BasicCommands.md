1. To enter into docker-postgres : docker exec -it postgres psql -U postgres -d ticket-reservation

2. To migrate schema into db : Get-Content migrations\005_create_idempotency_keys.sql | docker exec -i postgres psql -U postgres -d ticket-reservation

3. List all tables inside postgres-docker : \dt

4. To close docker-postgres session: \q
