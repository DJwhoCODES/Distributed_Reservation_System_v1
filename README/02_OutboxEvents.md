Outbox events are a way to reliably publish events when your application updates the database.

The main problem they solve is:

How do you guarantee that both the database write and the event publication happen exactly once (or at least reliably), even if the application crashes?

The problem

Suppose a user books a movie ticket.

Your service does:

1. Insert booking into PostgreSQL ✅
2. Publish "BookingCreated" to Kafka ❌

What if the app crashes after step 1 but before step 2?

Result:

Booking exists in DB ✅
No event sent ❌
Payment service, notification service, analytics, etc. never know about it.

Now reverse the order:

1. Publish "BookingCreated" ✅
2. Insert booking into PostgreSQL ❌

If the DB insert fails:

Event exists ❌
Booking doesn't exist ❌

Now your system is inconsistent.

Outbox Pattern

Instead of publishing directly:

                 +----------------+
                 |  PostgreSQL    |
                 +----------------+
                        |
         +--------------+----------------+
         |                               |
         | bookings table                |
         | outbox table                  |
         +-------------------------------+

Inside one database transaction:

BEGIN;

INSERT INTO bookings (...);

INSERT INTO outbox (
event_id,
event_type,
payload,
status
);

COMMIT;

Both succeed together or both fail together.

No inconsistency.

Example
bookings
id movie seat

---

101 Avatar A10
outbox
event_id type payload processed

---

abc123 BookingCreated {...json...} false

After commit:

booking exists
event exists

No event is lost.

Then who publishes the event?

A separate background worker.

                 PostgreSQL
              +----------------+
              | bookings        |
              | outbox          |
              +--------+--------+
                       |
              polls unread events
                       |
                       v
              Outbox Publisher
                       |
          +------------+------------+
          |                         |
          v                         v

       Kafka                  RabbitMQ

Worker:

for {
events := GetUnprocessedEvents()

    for _, e := range events {
        publishToKafka(e)

        markProcessed(e.ID)
    }

}
Better approach: CDC (Change Data Capture)

Instead of polling:

PostgreSQL
|
| WAL
v
Debezium
|
v
Kafka
|
v
Consumers

Debezium reads PostgreSQL's Write-Ahead Log (WAL) and streams outbox rows automatically.

Advantages:

No polling
Lower latency
More scalable
Production-friendly

Many large systems use:

PostgreSQL
Outbox table
Debezium
Kafka
Outbox table schema
CREATE TABLE outbox (
id UUID PRIMARY KEY,
aggregate_id UUID NOT NULL,
event_type TEXT NOT NULL,
payload JSONB NOT NULL,
created_at TIMESTAMP NOT NULL,
processed BOOLEAN DEFAULT FALSE
);

Sometimes instead of processed, people use:

status

PENDING
PROCESSING
PUBLISHED
FAILED

or

retry_count
last_attempt_at
error

to support retries.

Real-world movie reservation flow
User
|
v

Reservation Service

BEGIN TX

Reserve Seat
Insert Reservation

Insert Outbox:
{
type: SeatReserved,
reservationId: 123,
seat: A10
}

COMMIT

        |
        v

Outbox Publisher

        |
        +------> Kafka

                     |
      +--------------+--------------+
      |              |              |

Payment Notification Analytics
Service Service Service
Key idea

The critical guarantee is:

✅ Business data and event record are written in the same PostgreSQL transaction
✅ If the transaction commits, both exist
✅ If it rolls back, neither exists
✅ A background publisher retries until the event is successfully delivered
✅ Consumers should be idempotent because duplicate deliveries can still occur in distributed systems

For production-grade distributed systems (ticketing, reservations, payments, order processing), the Outbox Pattern is one of the standard approaches to achieve reliable event-driven communication without relying on fragile dual writes.
