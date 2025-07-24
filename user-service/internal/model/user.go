package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	AvatarURL string    `db:"avatar_url"`
	XP        int       `db:"xp"`
	CreatedAt time.Time `db:"created_at"`
}
