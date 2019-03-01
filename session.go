package auth

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Sessionable stores session information
type Sessionable interface {
	SessionID() string
	SessionType() int
	UserID() string
	UserAgent() Agentable
	Permissions() []string
	OriginID() string
	Origin() string
	Expiry() time.Duration // Used by scheduler
	CreatedAt() time.Time
}

// Session implements the Sessionable interface
type Session struct {
	sessionID   string
	sessionType int
	userID      string
	userAgent   Agentable
	permissions []string
	originID    string
	origin      string
	expiry      time.Duration
	createdAt   time.Time
}

// NewSession returns a new Session
func NewSession(sessionType int, userID string, userAgent Agentable, permissions []string, originID string, origin string, expiry time.Duration) *Session {
	return &Session{
		sessionID:   uuid.NewV4().String(),
		sessionType: sessionType,
		userID:      userID,
		userAgent:   userAgent,
		permissions: permissions,
		originID:    originID,
		origin:      origin,
		expiry:      expiry,
		createdAt:   time.Now(),
	}
}

// SessionID implements the Sessionable interface
func (s Session) SessionID() string {
	return s.sessionID
}

// SessionType implements the Sessionable interface
func (s Session) SessionType() int {
	return s.sessionType
}

// UserID implements the Sessionable interface
func (s Session) UserID() string {
	return s.userID
}

// UserAgent implements the Sessionable interface
func (s Session) UserAgent() Agentable {
	return s.userAgent
}

// Permissions implements the Sessionable interface
func (s Session) Permissions() []string {
	return s.permissions
}

// OriginID implements the Sessionable interface
func (s Session) OriginID() string {
	return s.originID
}

// Origin implements the Sessionable interface
func (s Session) Origin() string {
	return s.origin
}

// Expiry implements the Sessionable interface
func (s Session) Expiry() time.Duration {
	return s.expiry
}

// CreatedAt implements the Sessionable interface
func (s Session) CreatedAt() time.Time {
	return s.createdAt
}

// Sessionables represents multiple Sessionable
type Sessionables interface {
	Length() int
	Get(i int) Sessionable
}

// Sessions implements the Sessionables interface
type Sessions []Session

// Length implements the Sessionables interface
func (s Sessions) Length() int {
	return len(s)
}

// Get implements the Sessionables interface
func (s Sessions) Get(i int) Sessionable {
	return s[i]
}
