CREATE TABLE show_seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    show_id UUID NOT NULL REFERENCES shows(id) ON DELETE CASCADE,

    row_label TEXT NOT NULL,
    seat_number INT NOT NULL CHECK (seat_number > 0),
    seat_label TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(show_id, seat_label),
    UNIQUE(show_id, row_label, seat_number)
);

CREATE INDEX idx_show_seats_show_id
ON show_seats(show_id);