package entities

import (
	"time"

	"github.com/gofiber/fiber/v2/internal/uuid"
)

type Interface interface {
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	TableName()
	GetMap() map[string]interface{}
	GetFilterId() map[string]interface{}
}

type Base struct {
	ID        uuid.UUID `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (b *Base) GenerateID() {
	b.ID = uuid.New()

}

func (b *Base) SetCreatedAt() {
	b.CreatedAt = time.Now()
}

func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}

func GetTimeFormat() string {
	return "2010-01-02T15:04:05-0708"
}
