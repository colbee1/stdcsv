package stdcsv

import "time"

type (
	Config struct {
		// SkipBOM skip the BOM if present.
		// Default is false
		SkipBOM bool

		// CommaIn defines the comma used by input file.
		// Default is ','.
		CommaIn rune

		// CommaOut define comma to use in output file.
		// Default is ','.
		CommaOut rune

		// SkipTrailingComma skip trailing comma if present.
		// Default is false
		SkipTrailingComma bool

		// Encoding define the input encoding.
		// Default is utf-8
		Charset string

		// Skip row beginning with this rune.
		// Default is '' (no comment)
		Comment rune

		// LazyQuotes define lazy quoting for input file
		// Default is false
		LazyQuotes bool

		// NbColumn define the expected number of column par row
		// Default is -1: Pad/truncate according to number of columns in header.
		// 0: fails if column do not contains same number of columns as in headers.
		// >0: expects this number of columns
		NbColumn int

		// When NbColumn == -1 and column are missing, pad with ColPad
		ColPad string

		// Headers set the headers line. If empty, headers line is read from input.
		Headers []string

		// TrimSpaces trim spaces around each cells
		// Default: false
		TrimSpaces bool

		// Offset defines the number of rows to skip before start to write output.
		// Default 0
		Offset int64

		// Limit defines the maximum number of rows to write in output.
		Limit int64
	}

	Stats struct {
		ReadRows              int64
		WrittenRows           int64
		PaddedOrTruncatedRows uint64    `json:",omitempty"`
		StartedAt             time.Time `json:",omitempty"`
		StoppedAt             time.Time `json:",omitempty"`
		Duration              float64   `json:",omitempty"`
	}
)
