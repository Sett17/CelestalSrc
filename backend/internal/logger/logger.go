package logger

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

var Formatter = func(param gin.LogFormatterParams) string {
	var statusStyle, methodStyle lipgloss.Style
	if param.IsOutputColor() {
		methodStyle = styleForMethod(param.Method)
		statusStyle = styleForStatus(param.StatusCode)
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	return fmt.Sprintf("✨ %v | %3s | %11v | %7s | %s | %15s | %#v \n",
		param.TimeStamp.Format("2006/01/02 15:04:05"),
		statusStyle.Render(strconv.Itoa(param.StatusCode)),
		param.Latency,
		humanize.Bytes(uint64(param.BodySize)),
		methodStyle.Render(param.Method),
		param.ClientIP,
		param.Path,
	)
}

func styleForStatus(status int) lipgloss.Style {
	switch {
	case status >= http.StatusOK && status < http.StatusMultipleChoices:
		return lipgloss.NewStyle().Background(lipgloss.Color("86")).Foreground(lipgloss.Color("0"))
    case status < http.StatusInternalServerError:
        return lipgloss.NewStyle().Background(lipgloss.Color("192")).Foreground(lipgloss.Color("0"))
	default:
        return lipgloss.NewStyle().Background(lipgloss.Color("204")).Foreground(lipgloss.Color("0"))
	}
}

func styleForMethod(method string) lipgloss.Style {
	switch method {
	case http.MethodGet:
		return lipgloss.NewStyle().Background(lipgloss.Color("86")).Foreground(lipgloss.Color("0"))
	case http.MethodPost:
        return lipgloss.NewStyle().Background(lipgloss.Color("63")).Foreground(lipgloss.Color("17"))
    case http.MethodPut:
        return lipgloss.NewStyle().Background(lipgloss.Color("192")).Foreground(lipgloss.Color("17"))
    case http.MethodDelete:
        return lipgloss.NewStyle().Background(lipgloss.Color("204")).Foreground(lipgloss.Color("0"))
    case http.MethodPatch:
        return lipgloss.NewStyle().Background(lipgloss.Color("218")).Foreground(lipgloss.Color("17"))
    case http.MethodHead:
        return lipgloss.NewStyle().Background(lipgloss.Color("131")).Foreground(lipgloss.Color("0"))
    case http.MethodOptions:
        return lipgloss.NewStyle().Background(lipgloss.Color("153")).Foreground(lipgloss.Color("0"))
	default:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	}
}
