package models

import (
	"time"

	"github.com/google/uuid"
)

// TestModel is a minimal model for testing GORM migration
type TestModel struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User represents a user in the system
type User struct {
	ID                 uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email              string    `json:"email" gorm:"uniqueIndex;not null"`
	Password           string    `json:"-" gorm:"not null"`
	DiscordID          string    `json:"discord_id"`
	IsActive           bool      `json:"is_active" gorm:"default:true"`
	SubscriptionTier   string    `json:"subscription_tier" gorm:"default:'basic'"`
	SubscriptionStatus string    `json:"subscription_status" gorm:"default:'active'"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// Portfolio represents a user's portfolio
type Portfolio struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Addresses []Address `json:"addresses" gorm:"foreignKey:PortfolioID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Address represents a blockchain address in a portfolio
type Address struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	PortfolioID uuid.UUID `json:"portfolio_id" gorm:"type:uuid;not null"`
	Portfolio   Portfolio `json:"portfolio" gorm:"foreignKey:PortfolioID"`
	Address     string    `json:"address" gorm:"not null"`
	Network     string    `json:"network" gorm:"not null"`
	Label       string    `json:"label"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	Balances    []Balance `json:"balances" gorm:"foreignKey:AddressID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Transaction represents a blockchain transaction
type Transaction struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	PortfolioID  uuid.UUID `json:"portfolio_id" gorm:"type:uuid;not null"`
	TxHash       string    `json:"tx_hash" gorm:"not null"`
	Network      string    `json:"network" gorm:"not null"`
	TokenAddress string    `json:"token_address"`
	Amount       string    `json:"amount" gorm:"type:decimal(65,18)"`
	BlockNumber  uint64    `json:"block_number"`
	Timestamp    time.Time `json:"timestamp"`
	CreatedAt    time.Time `json:"created_at"`
}

// Alert represents a user's alert
type Alert struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Type       string    `json:"type" gorm:"not null"`
	Name       string    `json:"name" gorm:"not null"`
	Conditions string    `json:"conditions" gorm:"type:jsonb;not null"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Balance represents a token balance
type Balance struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	AddressID    uuid.UUID `json:"address_id" gorm:"type:uuid;not null"`
	Address      Address   `json:"address" gorm:"foreignKey:AddressID"`
	TokenAddress string    `json:"token_address"`
	Symbol       string    `json:"symbol"`
	Name         string    `json:"name"`
	Amount       string    `json:"amount" gorm:"type:decimal(65,18)"`
	Decimals     uint8     `json:"decimals"`
	Price        string    `json:"price" gorm:"type:decimal(20,8)"`
	Value        string    `json:"value" gorm:"type:decimal(20,8)"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Forum models

type Question struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"type:text;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Tags      []Tag     `json:"tags" gorm:"many2many:question_tags;"`
	Answers   []Answer  `json:"answers" gorm:"foreignKey:QuestionID"`
	Votes     []Vote    `json:"votes" gorm:"polymorphic:Votable;polymorphicValue:question"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Answer struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Body       string    `json:"body" gorm:"type:text;not null"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	QuestionID uuid.UUID `json:"question_id" gorm:"type:uuid;not null"`
	Question   Question  `json:"question" gorm:"foreignKey:QuestionID"`
	IsAccepted bool      `json:"is_accepted" gorm:"default:false"`
	Votes      []Vote    `json:"votes" gorm:"polymorphic:Votable;polymorphicValue:answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Comment struct {
	ID         uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Body       string     `json:"body" gorm:"type:text;not null"`
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	User       User       `json:"user" gorm:"foreignKey:UserID"`
	QuestionID *uuid.UUID `json:"question_id" gorm:"type:uuid"`
	AnswerID   *uuid.UUID `json:"answer_id" gorm:"type:uuid"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type Tag struct {
	ID          uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string     `json:"name" gorm:"uniqueIndex;not null"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions" gorm:"many2many:question_tags;"`
}

type Vote struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Value       int       `json:"value" gorm:"not null"` // +1 or -1
	VotableID   uuid.UUID `json:"votable_id" gorm:"type:uuid;not null"`
	VotableType string    `json:"votable_type" gorm:"not null"` // "question" or "answer"
	CreatedAt   time.Time `json:"created_at"`
}

type Reputation struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;uniqueIndex"`
	Points    int       `json:"points" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;uniqueIndex"`
	Role      string    `json:"role" gorm:"not null"` // user, moderator, admin
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
