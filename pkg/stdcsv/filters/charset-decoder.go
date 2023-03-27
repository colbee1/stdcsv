package filters

import "golang.org/x/text/encoding/charmap"

var TextDecoder = map[string]*charmap.Charmap{
	"codepage037":       charmap.CodePage037,
	"codepage437":       charmap.CodePage437,
	"codepage850":       charmap.CodePage850,
	"codepage852":       charmap.CodePage852,
	"codepage855":       charmap.CodePage855,
	"codepage858":       charmap.CodePage858,
	"codepage860":       charmap.CodePage860,
	"codepage862":       charmap.CodePage862,
	"codepage863":       charmap.CodePage863,
	"codepage865":       charmap.CodePage865,
	"codepage866":       charmap.CodePage866,
	"codepage1047":      charmap.CodePage1047,
	"codepage1140":      charmap.CodePage1140,
	"iso8859_1":         charmap.ISO8859_1,
	"iso8859_2":         charmap.ISO8859_2,
	"iso8859_3":         charmap.ISO8859_3,
	"iso8859_4":         charmap.ISO8859_4,
	"iso8859_5":         charmap.ISO8859_5,
	"iso8859_6":         charmap.ISO8859_6,
	"iso8859_7":         charmap.ISO8859_7,
	"iso8859_8":         charmap.ISO8859_8,
	"iso8859_9":         charmap.ISO8859_9,
	"iso8859_10":        charmap.ISO8859_10,
	"iso8859_13":        charmap.ISO8859_13,
	"iso8859_14":        charmap.ISO8859_14,
	"iso8859_15":        charmap.ISO8859_15,
	"iso8859_16":        charmap.ISO8859_16,
	"koi8r":             charmap.KOI8R,
	"koi8u":             charmap.KOI8U,
	"macintosh":         charmap.Macintosh,
	"macintoshcyrillic": charmap.MacintoshCyrillic,
	"windows874":        charmap.Windows874,
	"windows1250":       charmap.Windows1250,
	"windows1251":       charmap.Windows1251,
	"windows1252":       charmap.Windows1252,
	"windows1253":       charmap.Windows1253,
	"windows1254":       charmap.Windows1254,
	"windows1255":       charmap.Windows1255,
	"windows1256":       charmap.Windows1256,
	"windows1257":       charmap.Windows1257,
	"windows1258":       charmap.Windows1258,
}

var Decoders []string = make([]string, 0, len(TextDecoder))

func init() {
	for k := range TextDecoder {
		Decoders = append(Decoders, k)
	}
}
