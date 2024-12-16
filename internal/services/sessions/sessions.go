package sessions

import (
	"server/internal/config"
	"server/internal/entities"
	"server/internal/pkg/sessions"
	"server/pkg/custom_error"
)

type Sessions struct {
	cfg           *config.Config
	sessionsStore sessionsStore
}

func New(cfg *config.Config, sessionsStore sessionsStore) *Sessions {
	return &Sessions{
		cfg:           cfg,
		sessionsStore: sessionsStore,
	}
}

func (s *Sessions) Create(entity entities.SessionsCreate) (sessionData *entities.SessionData, err error) {
	sessionData, err = sessions.NewCreator(s.cfg).Create(entity)
	if err != nil {
		return nil, err
	}

	err = s.sessionsStore.Add(*sessionData)
	if err != nil {
		return nil, err
	}

	return sessionData, nil
}

func (s *Sessions) HasSessionInStore(entity entities.SessionsHas) error {
	storeRefreshJWT, err := s.sessionsStore.Get(entity.UserID, entity.SessionID)
	if err != nil || storeRefreshJWT == "" {
		return custom_error.ErrNotAuthorized
	}

	if entity.RefreshJWT != storeRefreshJWT {
		return custom_error.ErrNotAuthorized
	}
	return nil
}

func (s *Sessions) Delete(entity entities.Logout) error {
	return s.sessionsStore.Delete(entity)
}

func (s *Sessions) DeleteAll(userID string) error {
	return s.sessionsStore.DeleteAll(userID)
}
