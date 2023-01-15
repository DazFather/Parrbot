package message

import (
	"fmt"

	"github.com/NicoNex/echotron/v3"
)

// api is used by echotron library to send request to Telegram
var api echotron.API

// LoadAPI resets the api to a new value using given token.
// Keep in mind that both api and token will be already set douring robot.Start and
// if you don't want to use program's argument, you can use robot.Config.SetAPIToken
// intead. You are probably NOT going to need this function
func LoadAPI(token string) {
	api = echotron.NewAPI(token)
}

// API return the current api, useful for compatibility with not yet supported
// echotron functions calls
func API() echotron.API {
	return api
}

// ResponseError is an error generated by a echotron / Telegram resonse
type ResponseError struct {
	From        string
	ErrorCode   int
	Description string
}

// Error returns a complete error description (by creating this method ResponseError is a error interface)
func (err ResponseError) Error() string {
	return fmt.Sprint("[", err.ErrorCode, "] ", err.From, ": ", err.Description)
}

func parseResponseError(res echotron.APIResponse, err error) error {
	if err != nil {
		return &ResponseError{"Echotron", 1, err.Error()}
	}
	if base := res.Base(); !base.Ok {
		return &ResponseError{"Telegram", base.ErrorCode, base.Description}
	}
	return nil
}

// clearResponse it clears the echotron.APIResponseMessage and returns the actual
// message of type *UpdateMessage (casting it from Result),
// and an error by checking both, the APIResponseBase and the echotron err
func clearResponse(res echotron.APIResponseMessage, err error) (*UpdateMessage, error) {
	if e := parseResponseError(res, err); e != nil {
		return nil, e
	}

	return castMessage(res.Result), nil
}

// Any rapresent any single message type with the exeption of MediaGroup
type Any interface {
	// Send the message to the specified user and return a pointer to the messa sent and an error
	Send(chatID int64) (*UpdateMessage, error)
}

// Animation message type
type Animation struct {
	File echotron.InputFile
	Opts *echotron.AnimationOptions
}

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Animation) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendAnimation(message.File, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Animation) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Animation {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Animation) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Animation {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Animation) ForceReply(placeholder string, selective bool) *Animation {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Audio) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendAudio(message.File, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Audio) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Audio {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Audio) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Audio {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Audio) ForceReply(placeholder string, selective bool) *Audio {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Contact) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendContact(message.PhoneNumber, message.FirstName, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Contact) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Contact {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Contact) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Contact {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Contact) ForceReply(placeholder string, selective bool) *Contact {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Dice) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendDice(chatID, message.Emoji, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Dice) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Dice {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Dice) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Dice {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Dice) ForceReply(placeholder string, selective bool) *Dice {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Document) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendDocument(message.File, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Document) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Document {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Document) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Document {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Document) ForceReply(placeholder string, selective bool) *Document {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Game) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendGame(message.GameShortName, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Game) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Game {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Game) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Game {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Game) ForceReply(placeholder string, selective bool) *Game {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Location) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendLocation(chatID, message.Latitude, message.Longitude, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Location) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Location {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Location) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Location {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Location) ForceReply(placeholder string, selective bool) *Location {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Text) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendMessage(message.Text, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Text) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Text {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Text) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Text {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Text) ForceReply(placeholder string, selective bool) *Text {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Photo) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendPhoto(message.File, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Photo) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Photo {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Photo) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Photo {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Photo) ForceReply(placeholder string, selective bool) *Photo {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Poll) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendPoll(chatID, message.Question, message.Options, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Poll) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Poll {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Poll) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Poll {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Poll) ForceReply(placeholder string, selective bool) *Poll {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Sticker) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendSticker(message.StickerID, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Sticker) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Sticker {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Sticker) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Sticker {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Sticker) ForceReply(placeholder string, selective bool) *Sticker {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Venue) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendVenue(chatID, message.Latitude, message.Longitude, message.Title, message.Address, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Venue) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Venue {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Venue) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Venue {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Venue) ForceReply(placeholder string, selective bool) *Venue {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Video) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendVideo(message.File, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Video) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Video {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Video) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Video {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Video) ForceReply(placeholder string, selective bool) *Video {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message VideoNote) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendVideoNote(message.File, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *VideoNote) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *VideoNote {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *VideoNote) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *VideoNote {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *VideoNote) ForceReply(placeholder string, selective bool) *VideoNote {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
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

// Send the message to the specified user (by this method the stuct can be used a Any interface)
func (message Voice) Send(chatID int64) (res *UpdateMessage, err error) {
	return clearResponse(api.SendVoice(message.File, chatID, message.Opts))
}

// ClipKeyboard allows to quickly add or change the Opts.ReplyMarkup of the current message
func (message *Voice) ClipKeyboard(kbd echotron.ReplyKeyboardMarkup) *Voice {
	return message.editMarkup(kbd)
}

// ClipInlineKeyboard allows to quickly add or change an inline keyboard to the message Opts
func (message *Voice) ClipInlineKeyboard(kbd [][]echotron.InlineKeyboardButton) *Voice {
	return message.editMarkup(echotron.InlineKeyboardMarkup{InlineKeyboard: kbd})
}

// ForceReply allows to quickly force the user to reply to the current message when sent
func (message *Voice) ForceReply(placeholder string, selective bool) *Voice {
	return message.editMarkup(echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	})
}

// editMarkup is a helper method to change or add a ReplyMarkup on the message Opts
func (message *Voice) editMarkup(value echotron.ReplyMarkup) *Voice {
	if message.Opts == nil {
		message.Opts = new(echotron.VoiceOptions)
	}
	message.Opts.ReplyMarkup = value

	return message
}

/* -----------------------------------------------------
   NOT YET IMPLEMENTED
   -----------------------------------------------------
type ChatAction struct {
	Action echotron.ChatAction
} // APIResponseBool

type MediaGroup struct {
	Media []echotron.GroupableInputMedia
	Opts  *echotron.MediaGroupOptions
} // APIResponseArray
*/
