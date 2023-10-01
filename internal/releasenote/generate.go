package releasenote

import (
	"context"
	"strings"

	"releaseros/internal/config"
	"releaseros/internal/markdown"

	"github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
)

type Generator struct {
	gitTagFinder GitTagFinder
	gitLogFinder GitLogFinder
}

func NewGenerator() Generator {
	return Generator{
		gitTagFinder: gitTagFinder{},
		gitLogFinder: gitLogFinder{},
	}
}

func (releaseNoteGenerator Generator) Generate(ctx context.Context, config config.Config) (string, error) {
	latestTag, err := releaseNoteGenerator.gitTagFinder.LatestTag(ctx)
	if err != nil {
		return "", err
	}
	previousTag, err := releaseNoteGenerator.gitTagFinder.PreviousTag(ctx, latestTag)
	if err != nil {
		logger.Warn().Msg("Previous tag not found, using empty string.")
	}
	logger.Debug().Str("latestTag", latestTag).Str("previousTag", previousTag).Msg("")

	var log string
	if previousTag == "" {
		if config.InitialReleaseMessage != "" {
			return config.InitialReleaseMessage, nil
		}

		log, err = releaseNoteGenerator.gitLogFinder.LogTo(ctx, latestTag)
	} else {
		log, err = releaseNoteGenerator.gitLogFinder.Log(ctx, previousTag, latestTag)
	}
	if err != nil {
		return "", err
	}
	logger.Debug().Str("log", log).Msg("")

	var footerString string
	if config.Footer != "" {
		footer := Footer{
			LatestTag:   latestTag,
			PreviousTag: previousTag,
		}
		footerString, err = footer.Generate(config)
		if err != nil {
			return "", err
		}
		footerString += "\n"
	}
	logger.Debug().Str("footerString", footerString).Msg("")

	rawCommits := strings.Split(log, "\n")
	if lastLine := rawCommits[len(rawCommits)-1]; strings.TrimSpace(lastLine) == "" {
		rawCommits = rawCommits[0 : len(rawCommits)-1]
	}
	logger.Debug().Strs("rawCommits", rawCommits).Msg("")

	records := make(Records, len(rawCommits))
	for key, rawCommit := range rawCommits {
		commitHash := strings.Split(rawCommit, " ")[0]
		message := strings.Join(strings.Split(rawCommit, " ")[1:], " ")
		records[key] = Record{CommitHash: commitHash, Message: message}
	}
	logger.Debug().Array("records", records).Msg("")

	records, err = records.Filter(config)
	if err != nil {
		return "", err
	}
	logger.Debug().Array("filtered records", records).Msg("")

	records = records.Sort(config)
	logger.Debug().Array("sorted records", records).Msg("")

	releaseNote := releaseNote{
		title:   "Release Note",
		records: records,
	}
	if len(config.Categories) == 0 {
		return releaseNote.String() + footerString, nil
	}

	categories, err := categories(config, records)
	if err != nil {
		return "", err
	}
	logger.Debug().Array("categories", categories).Msg("")

	categories = categories.Sort()
	logger.Debug().Array("sorted categories", categories).Msg("")

	categorisedReleaseNote := categorisedReleaseNote{
		releaseNote: releaseNote,
		categories:  categories,
	}
	logger.Debug().Object("categorisedReleaseNote", categorisedReleaseNote).Msg("")

	return categorisedReleaseNote.String() + footerString, nil
}

type releaseNote struct {
	title   string
	records Records
}

func (r releaseNote) String() string {
	doc := markdown.NewDocument().With(markdown.H2(r.title))

	ul := markdown.UL()

	if len(r.records) > 0 {
		for _, record := range r.records {
			ul = ul.With(markdown.LI(record.String()))
		}
		doc = doc.With(ul)
	}

	return doc.String()
}

func (r releaseNote) MarshalZerologObject(e *zerolog.Event) {
	e.Str("title", r.title)
	e.Array("records", r.records)
}

type categorisedReleaseNote struct {
	releaseNote
	categories releaseNoteCategories
}

func (r categorisedReleaseNote) String() string {
	doc := markdown.NewDocument().With(markdown.H2(r.title))

	for _, category := range r.categories {
		if len(category.records) == 0 {
			continue
		}

		doc = doc.With(markdown.H3(category.title))

		ul := markdown.UL()
		for _, record := range category.records {
			ul = ul.With(markdown.LI(record.String()))
		}
		doc = doc.With(ul)
	}

	return doc.String()
}

func (r categorisedReleaseNote) MarshalZerologObject(e *zerolog.Event) {
	e.Object("releaseNote", r.releaseNote)
	e.Array("categories", r.categories)
}
