package stdcsv

import "errors"

var (
	ErrPackage                = errors.New("stdcsv")
	ErrSkipBOM                = errors.New("skip BOM")
	ErrCharsetDecoderNotFound = errors.New("no decoder found for charset")
	ErrSkipTrailingComma      = errors.New("skip trailing comma")
	ErrReadHeaders            = errors.New("read headers")
	ErrWriteHeaders           = errors.New("write headers")
	ErrReadRow                = errors.New("read row")
	ErrWriteRow               = errors.New("write row")
)
