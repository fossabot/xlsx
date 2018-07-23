package format_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/stretchr/testify/require"
	"testing"
	"github.com/plandem/xlsx/internal/ml/styles"
)

func TestFontScheme(t *testing.T) {
	type Element struct {
		Property styles.FontSchemeType `xml:"property,omitempty"`
	}

	list := map[string]styles.FontSchemeType{
		"none":     format.FontSchemeNone,
		"major":    format.FontSchemeMajor,
		"minor":    format.FontSchemeMinor,
		"schema-a": styles.FontSchemeType("schema-a"),
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property></Element>`, s), string(encoded))

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
