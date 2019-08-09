package lang

import (
	"fmt"
	"io/ioutil"

	"misteraladin.com/jasmine/go-boilerplate/config"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

var bundle *i18n.Bundle

func Translate(key string, data interface{}) string {
	locale := config.Config.App.Locale

	localizer := i18n.NewLocalizer(bundle, locale)
	config := i18n.LocalizeConfig{}
	config.MessageID = key

	if data != nil {
		config.TemplateData = data
	}

	msg, _ := localizer.Localize(
		&config,
	)
	return msg
}

func LoadLanguage() {
	langFiles := []string{"lang/en.yaml", "lang/id.yaml"}
	var err error

	bundle, err = CreateLocalizerBundle(langFiles)
	if err != nil {
		fmt.Printf("Error initialising localization, %v", err)
		panic(err)
	}
}

func CreateLocalizerBundle(langFiles []string) (*i18n.Bundle, error) {
	// Bundle stores a set of messages
	bundle := i18n.NewBundle(language.English)

	// Enable bundle to understand yaml
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var translations []byte
	var err error
	for _, file := range langFiles {
		// Read our language yaml file
		translations, err = ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Unable to read translation file %s", file)
			return nil, err
		}

		// It parses the bytes in buffer to add translations to the bundle
		bundle.MustParseMessageFileBytes(translations, file)
	}

	return bundle, nil
}
