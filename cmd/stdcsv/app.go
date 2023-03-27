package stdcsv

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	pkg "github.com/colbee1/stdcsv/pkg/stdcsv"
	"github.com/colbee1/stdcsv/pkg/stdcsv/filters"
)

type App struct {
	CliOutput io.Writer
	Flag      struct {
		fs                *flag.FlagSet
		FileIn            string
		FileOut           string
		SkipBOM           bool
		Comma             string
		CommaOutput       string
		SkipTrailingComma bool
		Charset           string
		Comment           string
		LazyQuotes        bool
		NbColumn          int
		ColPad            string
		Headers           string
		TrimSpaces        bool
		Offset            int64
		Limit             int64
		ListCharset       bool
		Quiet             bool
	}
}

func (app *App) Close() error {
	return nil
}

func (app *App) Logf(format string, args ...any) {
	if !app.Flag.Quiet {
		fmt.Fprintf(app.CliOutput, format, args...)
	}
}

func (app *App) Setup() error {
	app.Flag.fs = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	app.CliOutput = app.Flag.fs.Output()

	app.Flag.fs.StringVar(&app.Flag.FileIn, "in", "-", "input file")
	app.Flag.fs.StringVar(&app.Flag.FileOut, "out", "-", "output file")
	app.Flag.fs.BoolVar(&app.Flag.SkipBOM, "skip-bom", false, "skip the BOM")
	app.Flag.fs.StringVar(&app.Flag.Comma, "comma-in", ",", "comma used in input")
	app.Flag.fs.StringVar(&app.Flag.CommaOutput, "comma-out", ",", "comma to use for output")
	app.Flag.fs.BoolVar(&app.Flag.SkipTrailingComma, "skip-trailing-comma", false, "skip the trailing comma if present")
	app.Flag.fs.StringVar(&app.Flag.Charset, "charset", "utf-8", "charset of input")
	app.Flag.fs.StringVar(&app.Flag.Comment, "comment", "#", "char used as comment")
	app.Flag.fs.BoolVar(&app.Flag.LazyQuotes, "lazy", false, "lazy quoting is used in input")
	app.Flag.fs.IntVar(&app.Flag.NbColumn, "columns", -1, "number of columns per rows")
	app.Flag.fs.StringVar(&app.Flag.ColPad, "colpad", "", "column content when padding")
	app.Flag.fs.StringVar(&app.Flag.Headers, "headers", "", "headers separated by comma")
	app.Flag.fs.BoolVar(&app.Flag.TrimSpaces, "trim-spaces", false, "trim cells spaces")
	app.Flag.fs.Int64Var(&app.Flag.Offset, "offset", 0, "number of rows to skip before start")
	app.Flag.fs.Int64Var(&app.Flag.Limit, "limit", 0, "maximum number of rows in output")
	app.Flag.fs.BoolVar(&app.Flag.ListCharset, "charset-list", false, "list recognized charset")
	app.Flag.fs.BoolVar(&app.Flag.Quiet, "quiet", false, "don't be to verbose")

	app.Flag.fs.Usage = func() {
		out := app.Flag.fs.Output()
		fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
		app.Flag.fs.PrintDefaults()
		fmt.Fprintf(out, ``)
	}

	app.Flag.fs.Parse(os.Args[1:])

	return nil
}

func (app *App) Run() error {
	if app.Flag.ListCharset {
		names := make([]string, 0, len(filters.TextDecoder))
		for name := range filters.TextDecoder {
			names = append(names, name)
		}
		fmt.Fprintln(app.CliOutput, strings.Join(names, ", "))

		return nil
	}

	var fin io.ReadCloser = os.Stdin
	if name := app.Flag.FileIn; name != "-" {
		if f, err := os.Open(name); err != nil {
			return err
		} else {
			fin = f
		}
	}
	defer fin.Close()

	var fout io.WriteCloser = os.Stdout
	if name := app.Flag.FileOut; name != "-" {
		if f, err := os.Create(name); err != nil {
			return err
		} else {
			fout = f
		}
	}
	defer fout.Close()

	config := pkg.Config{
		SkipBOM:           app.Flag.SkipBOM,
		Comma:             []rune(app.Flag.Comma)[0],
		CommaOutput:       []rune(app.Flag.CommaOutput)[0],
		SkipTrailingComma: app.Flag.SkipTrailingComma,
		Charset:           app.Flag.Charset,
		Comment:           []rune(app.Flag.Comment)[0],
		LazyQuotes:        app.Flag.LazyQuotes,
		NbColumn:          app.Flag.NbColumn,
		ColPad:            app.Flag.ColPad,
		Headers:           strings.Split(app.Flag.Headers, ","),
		TrimSpaces:        app.Flag.TrimSpaces,
		Offset:            app.Flag.Offset,
		Limit:             app.Flag.Limit,
	}

	stats, err := pkg.Transform(fin, fout, config)
	if err != nil {
		return err
	}

	data, err := json.Marshal(stats)
	if err != nil {
		return err
	}

	app.Logf("%s\n", data)

	return nil
}
