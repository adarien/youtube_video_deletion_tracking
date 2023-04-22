package localization

import (
	"github.com/Masterminds/sprig"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type Bundle struct {
	b *i18n.Bundle
}

type Lang struct {
	l          *i18n.Localizer
	botButtons map[Button]string
}

type Message string

const (
	MsgBye               Message = "msgBye"
	MsgHello             Message = "msgHello"
	MsgInitEnd           Message = "msgInitEnd"
	MsgInitLang          Message = "msgInitLang"
	MsgLangSelect        Message = "msgLangSelect"
	MsgSessionConflict   Message = "msgSessionConflict"
	MsgSettings          Message = "msgSettings"
	MsgUserNotRegistered Message = "msgUserNotRegistered"
)

func (m Message) String() string {
	return string(m)
}

type Button string

const (
	ButtonBack         Button = "buttonBack"
	ButtonEN           Button = "buttonEN"
	ButtonRU           Button = "buttonRU"
	ButtonSettingsLang Button = "buttonSettingsLang"
)

var buttons = []Button{
	ButtonBack,
	ButtonEN,
	ButtonRU,
	ButtonSettingsLang,
}

func (b Button) String() string {
	return string(b)
}

func Init() (Bundle, error) {

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	if _, err := bundle.LoadMessageFile("./modules/localization/locales/en.yaml"); err != nil {
		return Bundle{}, err
	}
	if _, err := bundle.LoadMessageFile("./modules/localization/locales/ru.yaml"); err != nil {
		return Bundle{}, err
	}

	bundle.LanguageTags()

	return Bundle{
		b: bundle,
	}, nil
}

func (b *Bundle) LangSwitch(tag string) (Lang, error) {

	l := i18n.NewLocalizer(b.b, tag)

	bb := make(map[Button]string)

	for _, btn := range buttons {
		button, err := l.Localize(&i18n.LocalizeConfig{
			MessageID: btn.String(),
		})
		if err != nil {
			return Lang{}, err
		}

		bb[btn] = button
	}

	return Lang{
		l:          l,
		botButtons: bb,
	}, nil
}

func (l *Lang) BotButton(button Button) string {
	return l.botButtons[button]
}

func (l *Lang) MessageCreate(tag string, in any) (string, error) {
	return l.l.Localize(
		&i18n.LocalizeConfig{
			MessageID:    tag,
			Funcs:        sprig.TxtFuncMap(),
			TemplateData: in,
		},
	)
}
