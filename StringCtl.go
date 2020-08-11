package main

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var (
	maxLengthDesc = 100
	regex         = map[string]string{
		"filterList": "<div id=\"menu_detalle_buscador\">(.|\n)*?</div></div>",

		"getLink":          "<a class=\"titulo_menu_izq\" href=\"(.)*?\">",
		"getDesc":          "<div id=\"buscador_detalle_sub\">(.)*?</div>",
		"getCountry":       "src=\"/pais/(.)*?.gif",
		"getScore":         "src=\"img/calif(.)*?.gif",
		"getFormat":        "<b>Formato:</b> (.)*? <b>",
		"getDate":          "<b>el</b> (.)*? </div>",
		"getDownloads":     "<b>Downloads:</b> (.)*? <b>",
		"getUploaderStep1": "<b>Subido por:</b> <a class=(.)*?\">(.)*?</a>",
		"getUploaderStep2": "\">(.)*?</",

		"getDownloadLink":   "<a class=\"link1\" href=\"(.)*?\">Bajar",
		"getDownloadLinkId": "?id=(.)*?&",
	}
)

func getDownloadLink(line []byte) string {
	return extract(toUtf8(line), "getDownloadLink")
}

func getLink(line []byte) string {
	return extract(toUtf8(line), "getLink")
}

func getDesc(line []byte) string {
	return extract(toUtf8(line), "getDesc")
}

func getCountry(line []byte) string {
	return extract(toUtf8(line), "getCountry")
}

func getScore(line []byte) string {
	return extract(toUtf8(line), "getScore")
}

func getFormat(line []byte) string {
	return extract(toUtf8(line), "getFormat")
}

func getDate(line []byte) string {
	return extract(toUtf8(line), "getDate")
}

func getDownloads(line []byte) string {
	return extract(toUtf8(line), "getDownloads")
}

func getUploader(line []byte) string {
	return extract(extract(toUtf8(line), "getUploaderStep1"), "getUploaderStep2")
}

func extract(line string, field string) string {
	re := regexp.MustCompile(regex[field])
	strParts := strings.Split(regex[field], "(.)*?")
	raw := re.FindString(line)
	parsedValue := raw[len(strParts[0]) : len(raw)-len(strParts[1])]
	return parsedValue
}

func getList(page []byte) [][]byte {
	re := regexp.MustCompile(regex["filterList"])
	lines := re.FindAll(page, -1)
	return lines
}

func toUtf8(iso8859_1buf []byte) string {
	buf := make([]rune, len(iso8859_1buf))
	for i, b := range iso8859_1buf {
		buf[i] = rune(b)
	}
	return string(buf)
}

func createTable() table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Descripción", "Descargas", "Usuario", "Calif.")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

func trimString(value string, length int) string {
	if len(value) > length {
		value = value[:length]
	}
	return value
}
