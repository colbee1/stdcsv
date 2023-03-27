package stdcsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/colbee1/assertor"
	"github.com/colbee1/stdcsv/pkg/stdcsv/filters"
)

func Transform(fin io.Reader, fout io.Writer, config Config) (*Stats, error) {
	v := assertor.New()
	v.Assert(fin != nil, "input reader is missing")
	v.Assert(fout != nil, "output writer is missing")
	if err := v.Validate(); err != nil {
		return nil, err
	}

	// Skip the file BOM ?
	//
	if config.SkipBOM {
		if f, err := filters.BOMSkipper(fin); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrSkipBOM, err)
		} else {
			fin = f
		}
	}

	// Use a charset decoder ?
	//
	if charset := strings.ToLower(config.Charset); charset != "" {
		if charset == "utf-8" || charset == "utf8" {
			// no need to decode charset
		} else {
			if decoder, found := filters.TextDecoder[charset]; !found {
				return nil, fmt.Errorf("%w: %s", ErrCharsetDecoderNotFound, charset)
			} else {
				fin = decoder.NewDecoder().Reader(fin)
			}
		}
	}

	// Skip trailing comma ?
	//
	if config.SkipTrailingComma {
		fin = filters.SkipTrailingComma(fin, config.CommaIn)
	}

	// Create CSV reader
	//
	lineNo := 0
	csvReader := csv.NewReader(fin)
	if config.CommaIn != 0 {
		csvReader.Comma = config.CommaIn
	}
	if config.Comment != 0 {
		csvReader.Comment = config.Comment
	}
	csvReader.FieldsPerRecord = config.NbColumn
	csvReader.LazyQuotes = config.LazyQuotes

	nextRecords := func() ([]string, error) {
		lineNo++
		return csvReader.Read()
	}

	// Prepare stats
	//
	stats := &Stats{
		StartedAt: time.Now(),
	}

	// Is headers passed by configuration ?
	//
	headers := config.Headers
	if len(headers) == 0 {
		if h, err := nextRecords(); err != nil {
			return nil, fmt.Errorf("%w: (line #%d) %w", ErrReadHeaders, lineNo, err)
		} else {
			headers = h
		}
	}

	// Create CSV writer and write header line
	//
	csvWriter := csv.NewWriter(fout)
	if config.CommaOut != 0 {
		csvWriter.Comma = config.CommaOut
	}
	defer csvWriter.Flush()

	if err := csvWriter.Write(headers); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrWriteHeaders, err)
	}

	// Loop over all rows
	//
	offset := config.Offset
	for {
		cells, err := nextRecords()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("%w: (line #%d) %w", ErrReadRow, lineNo, err)
		}

		if offset > 0 {
			offset--
			continue
		}

		stats.ReadRows++

		// Handle unexpected number of columns ?
		//
		if len(cells) != len(headers) {
			// being here means that INPUT_NB_COLUMN == -1
			if len(headers) > len(cells) {
				add := len(headers) - len(cells)
				for i := 0; i < add; i++ {
					cells = append(cells, config.ColPad)
				}
				stats.PaddedOrTruncatedRows++
			} else {
				cells = cells[:len(headers)]
				stats.PaddedOrTruncatedRows++
			}
		}

		// Trim spaces around cells ?
		//
		if config.TrimSpaces {
			for i := range headers {
				val := strings.TrimSpace(cells[i])
				cells[i] = val
			}
		}

		// write records
		//
		if err = csvWriter.Write(cells); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrWriteRow, err)
		} else {
			stats.WrittenRows++
		}

		if config.Limit != 0 && config.Limit == stats.WrittenRows {
			break
		}
	}

	// Update stats
	//
	stats.StoppedAt = time.Now()

	return stats, nil
}
