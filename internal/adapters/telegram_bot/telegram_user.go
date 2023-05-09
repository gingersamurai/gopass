package telegram_bot

type TelegramUserStorage interface {
	AddTelegramUser(tgUserId, userId int64) error
	GetTelegramUser(tgUserId int64) (int64, error)
	DeleteTelegramUser(tgUseId int64) error
}

type TelegramUser struct {
	Id     int64
	UserId int64
}
