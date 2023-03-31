package releasenote

import (
	"fmt"
	"regexp"
	"sort"

	"releaseros/internal/context"

	"github.com/rs/zerolog"
)

type releaseNoteCategory struct {
	title   string
	records Records
	weight  int
}

func (r releaseNoteCategory) MarshalZerologObject(e *zerolog.Event) {
	e.Str("title", r.title)
	e.Array("records", r.records)
	e.Int("weight", r.weight)
}

type releaseNoteCategories []releaseNoteCategory

func (r releaseNoteCategories) Sort() releaseNoteCategories {
	sort.Slice(r, func(i, j int) bool {
		return r[i].weight < r[j].weight
	})
	return r
}

func (r releaseNoteCategories) MarshalZerologArray(a *zerolog.Array) {
	for _, v := range r {
		a.Object(v)
	}
}

func categories(ctx *context.Context, records Records) (releaseNoteCategories, error) {
	categories := releaseNoteCategories{}
	for _, categoryFromConfig := range ctx.Config.Categories {
		category := releaseNoteCategory{
			title:   categoryFromConfig.Title,
			records: Records{},
			weight:  categoryFromConfig.Weight,
		}

		if categoryFromConfig.Regexp == "" {
			category.records = records
			categories = append(categories, category)
			break
		}

		regexp, err := regexp.Compile(categoryFromConfig.Regexp)
		if err != nil {
			return categories, fmt.Errorf("failed to compile regexp for category %q: %w", categoryFromConfig.Title, err)
		}

		i := 0
		for _, record := range records {
			if !regexp.MatchString(record.Message) {
				// Keep unmatched entry.
				records[i] = record
				i++
				continue
			}

			category.records = append(category.records, record)
		}
		records = records[:i]

		categories = append(categories, category)

		if len(records) == 0 {
			break
		}
	}
	return categories, nil
}
