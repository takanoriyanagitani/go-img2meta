package img2meta

import (
	"encoding/json"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func ModelToString(m color.Model, unknown string) string {
	switch m {
	case color.RGBAModel:
		return "rgba"
	case color.RGBA64Model:
		return "rgba64"
	case color.NRGBAModel:
		return "nrgba"
	case color.NRGBA64Model:
		return "nrgba64"
	case color.AlphaModel:
		return "alpha"
	case color.Alpha16Model:
		return "alpha16"
	case color.GrayModel:
		return "gray"
	case color.Gray16Model:
		return "gray16"
	default:
		return unknown
	}
}

type Metadata struct {
	FormatName string `json:"format_name"`
	ColorModel string `json:"color_model"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

func (m Metadata) ToJson() (jdata []byte, e error) {
	return json.Marshal(&m)
}

func ReaderToMetadata(rdr io.Reader) (Metadata, error) {
	cfg, f, e := image.DecodeConfig(rdr)
	return Metadata{
		FormatName: f,
		ColorModel: ModelToString(cfg.ColorModel, "unknown-model"),
		Width:      cfg.Width,
		Height:     cfg.Height,
	}, e
}
