CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    show_id UUID NOT NULL REFERENCES shows(id) ON DELETE CASCADE,

    seat_id UUID NOT NULL REFERENCES show_seats(id) ON DELETE CASCADE,

    user_id TEXT NOT NULL,

    status TEXT NOT NULL CHECK (
        status IN ('confirmed', 'cancelled')
    ),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(show_id, seat_id)
);

CREATE INDEX idx_bookings_show_id
ON bookings(show_id);

CREATE INDEX idx_bookings_user_id
ON bookings(user_id);