package L10n

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

func Configure(dirPath string, currentLocale string) error {

	currentTag, err := language.Default.Parse(currentLocale)

	if err != nil {
		return fmt.Errorf("can't parse current locale %s to language tag due to -> %w", currentLocale, err)
	}

	cat, err := _configure(dirPath)

	if err != nil {
		return err
	}

	Print = message.NewPrinter(currentTag, message.Catalog(cat))

	return nil
}

func _configure(dirPath string) (catalog.Catalog, error) {
	fileInfo, err := ioutil.ReadDir(dirPath)
	builder := catalog.NewBuilder()

	if err != nil {
		return nil, err
	}

	for _, it := range fileInfo {
		split := strings.Split(it.Name(), ".")

		if len(split) == 0 {
			log.Printf("[INFO] File with name %s wasn't parsed as language config", split[0])
			continue
		}

		if split[1] != "yaml" {
			log.Printf("[INFO] File with name %s wasn't parsed as language config", split[0])
			continue
		}

		tag, err := language.Default.Parse(split[0])

		if err != nil {
			return nil, fmt.Errorf("cauldn;t parse file name %s as a BCP47 language tag -> %w", split[0], err)
		}

		path := fmt.Sprintf("%s/%s", dirPath, it.Name())

		if err = parseFile(path, tag, builder); err != nil {
			return nil, err
		}
	}

	return builder, nil
}

func parseFile(filePath string, languageTag language.Tag, builder *catalog.Builder) error {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return fmt.Errorf("can't read file at path %s -> %w", filePath, err)
	}

	var dict map[string]string

	if err = yaml.Unmarshal(data, &dict); err != nil {
		return fmt.Errorf("can't parse file at path %s into yaml due to %w", filePath, err)
	}

	for key, value := range dict {
		if err = builder.Set(languageTag, key, catalog.String(value)); err != nil {
			return fmt.Errorf("couldn't set key %s and value %s as localized locales due to %w", key, value, err)
		}
	}

	return nil
}
