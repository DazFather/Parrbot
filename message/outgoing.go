package message

import (
	"github.com/NicoNex/echotron/v3"
)

// Any single message
type Any interface {
	Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error)
}

// Animation message type
type Animation struct {
	File echotron.InputFile
	Opts *echotron.AnimationOptions
}

func (message Animation) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendAnimation(message.File, chatID, message.Opts)
}

func (message *Animation) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Animation {
	return message.editMarkup(kbd)
}

func (message *Animation) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Animation {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Animation) ForceReply(placeholder string, selective bool) *Animation {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Animation) editMarkup(value echotron.ReplyMarkup) *Animation {
	if message.Opts == nil {
		message.Opts = new(echotron.AnimationOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Audio message type
type Audio struct {
	File echotron.InputFile
	Opts *echotron.AudioOptions
}

func (message Audio) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendAudio(message.File, chatID, message.Opts)
}

func (message *Audio) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Audio {
	return message.editMarkup(kbd)
}

func (message *Audio) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Audio {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Audio) ForceReply(placeholder string, selective bool) *Audio {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Audio) editMarkup(value echotron.ReplyMarkup) *Audio {
	if message.Opts == nil {
		message.Opts = new(echotron.AudioOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Contact message type
type Contact struct {
	PhoneNumber, FirstName string
	Opts                   *echotron.ContactOptions
}

func (message Contact) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendContact(message.PhoneNumber, message.FirstName, chatID, message.Opts)
}

func (message *Contact) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Contact {
	return message.editMarkup(kbd)
}

func (message *Contact) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Contact {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Contact) ForceReply(placeholder string, selective bool) *Contact {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Contact) editMarkup(value echotron.ReplyMarkup) *Contact {
	if message.Opts == nil {
		message.Opts = new(echotron.ContactOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Dice message type
type Dice struct {
	Emoji echotron.DiceEmoji
	Opts  *echotron.BaseOptions
}

func (message Dice) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendDice(chatID, message.Emoji, message.Opts)
}

func (message *Dice) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Dice {
	return message.editMarkup(kbd)
}

func (message *Dice) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Dice {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Dice) ForceReply(placeholder string, selective bool) *Dice {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Dice) editMarkup(value echotron.ReplyMarkup) *Dice {
	if message.Opts == nil {
		message.Opts = new(echotron.BaseOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Document message type
type Document struct {
	File echotron.InputFile
	Opts *echotron.DocumentOptions
}

func (message Document) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendDocument(message.File, chatID, message.Opts)
}

func (message *Document) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Document {
	return message.editMarkup(kbd)
}

func (message *Document) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Document {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Document) ForceReply(placeholder string, selective bool) *Document {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Document) editMarkup(value echotron.ReplyMarkup) *Document {
	if message.Opts == nil {
		message.Opts = new(echotron.DocumentOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Game message type
type Game struct {
	GameShortName string
	Opts          *echotron.BaseOptions
}

func (message Game) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendGame(message.GameShortName, chatID, message.Opts)
}

func (message *Game) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Game {
	return message.editMarkup(kbd)
}

func (message *Game) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Game {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Game) ForceReply(placeholder string, selective bool) *Game {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Game) editMarkup(value echotron.ReplyMarkup) *Game {
	if message.Opts == nil {
		message.Opts = new(echotron.BaseOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Location message type
type Location struct {
	Latitude, Longitude float64
	Opts                *echotron.LocationOptions
}

func (message Location) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendLocation(chatID, message.Latitude, message.Longitude, message.Opts)
}

func (message *Location) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Location {
	return message.editMarkup(kbd)
}

func (message *Location) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Location {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Location) ForceReply(placeholder string, selective bool) *Location {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Location) editMarkup(value echotron.ReplyMarkup) *Location {
	if message.Opts == nil {
		message.Opts = new(echotron.LocationOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Text message type
type Text struct {
	Text string
	Opts *echotron.MessageOptions
}

func (message Text) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendMessage(message.Text, chatID, message.Opts)
}

func (message *Text) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Text {
	return message.editMarkup(kbd)
}

func (message *Text) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Text {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Text) ForceReply(placeholder string, selective bool) *Text {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Text) editMarkup(value echotron.ReplyMarkup) *Text {
	if message.Opts == nil {
		message.Opts = new(echotron.MessageOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Photo message type
type Photo struct {
	File echotron.InputFile
	Opts *echotron.PhotoOptions
}

func (message Photo) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendPhoto(message.File, chatID, message.Opts)
}

func (message *Photo) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Photo {
	return message.editMarkup(kbd)
}

func (message *Photo) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Photo {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Photo) ForceReply(placeholder string, selective bool) *Photo {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Photo) editMarkup(value echotron.ReplyMarkup) *Photo {
	if message.Opts == nil {
		message.Opts = new(echotron.PhotoOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Poll message type
type Poll struct {
	Question string
	Options  []string
	Opts     *echotron.PollOptions
}

func (message Poll) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendPoll(chatID, message.Question, message.Options, message.Opts)
}

func (message *Poll) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Poll {
	return message.editMarkup(kbd)
}

func (message *Poll) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Poll {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Poll) ForceReply(placeholder string, selective bool) *Poll {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Poll) editMarkup(value echotron.ReplyMarkup) *Poll {
	if message.Opts == nil {
		message.Opts = new(echotron.PollOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Sticker message type
type Sticker struct {
	StickerID string
	Opts      *echotron.BaseOptions
}

func (message Sticker) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendSticker(message.StickerID, chatID, message.Opts)
}

func (message *Sticker) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Sticker {
	return message.editMarkup(kbd)
}

func (message *Sticker) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Sticker {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Sticker) ForceReply(placeholder string, selective bool) *Sticker {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Sticker) editMarkup(value echotron.ReplyMarkup) *Sticker {
	if message.Opts == nil {
		message.Opts = new(echotron.BaseOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Venue message type
type Venue struct {
	Latitude, Longitude float64
	Title, Address      string
	Opts                *echotron.VenueOptions
}

func (message Venue) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendVenue(chatID, message.Latitude, message.Longitude, message.Title, message.Address, message.Opts)
}

func (message *Venue) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Venue {
	return message.editMarkup(kbd)
}

func (message *Venue) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Venue {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Venue) ForceReply(placeholder string, selective bool) *Venue {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Venue) editMarkup(value echotron.ReplyMarkup) *Venue {
	if message.Opts == nil {
		message.Opts = new(echotron.VenueOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Video message type
type Video struct {
	File echotron.InputFile
	Opts *echotron.VideoOptions
}

func (message Video) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendVideo(message.File, chatID, message.Opts)
}

func (message *Video) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Video {
	return message.editMarkup(kbd)
}

func (message *Video) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Video {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Video) ForceReply(placeholder string, selective bool) *Video {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Video) editMarkup(value echotron.ReplyMarkup) *Video {
	if message.Opts == nil {
		message.Opts = new(echotron.VideoOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// VideoNote message type
type VideoNote struct {
	File echotron.InputFile
	Opts *echotron.VideoNoteOptions
}

func (message VideoNote) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendVideoNote(message.File, chatID, message.Opts)
}

func (message *VideoNote) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *VideoNote {
	return message.editMarkup(kbd)
}

func (message *VideoNote) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *VideoNote {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *VideoNote) ForceReply(placeholder string, selective bool) *VideoNote {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *VideoNote) editMarkup(value echotron.ReplyMarkup) *VideoNote {
	if message.Opts == nil {
		message.Opts = new(echotron.VideoNoteOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// Voice message type
type Voice struct {
	File echotron.InputFile
	Opts *echotron.VoiceOptions
}

func (message Voice) Send(api echotron.API, chatID int64) (res echotron.APIResponseMessage, err error) {
	return api.SendVoice(message.File, chatID, message.Opts)
}

func (message *Voice) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Voice {
	return message.editMarkup(kbd)
}

func (message *Voice) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Voice {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

func (message *Voice) ForceReply(placeholder string, selective bool) *Voice {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

func (message *Voice) editMarkup(value echotron.ReplyMarkup) *Voice {
	if message.Opts == nil {
		message.Opts = new(echotron.VoiceOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

// -----------------------------------------------------
type ChatAction struct {
	Action echotron.ChatAction
} // APIResponseBool

type MediaGroup struct {
	Media []echotron.GroupableInputMedia
	Opts  *echotron.MediaGroupOptions
} // APIResponseArray
