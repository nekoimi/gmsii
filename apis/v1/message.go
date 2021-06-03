package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nekoimi/gmsii/message"
	"io"
	"regexp"
	"strings"
)

var (
	oldPushRegexMap = map[string]string{
		"title":       `####(.*?)\n`,
		"date_time":   `time:(.*?)\n`,
		"err_code":    `Code:.+\[\s*([0-9]+?)\s*\]\n`,
		"err_message": `Message:(.*?)\n`,
		"err_clazz":   `Type:(.*?)\n`,
		"err_trace":   `File:(.*?)\n`,
	}
)

type TxtStruct struct {
	Content string `json:"content"`
}

func OldPush(ctx *gin.Context) {
	content := ctx.Request.FormValue("content")
	if len(content) <= 0 {
		ctx.JSON(200, gin.H{
			"code":    0,
			"data":    "",
			"message": "ok",
		})
	}

	var regexResult = make(map[string]string)
	for field, fieldRegex := range oldPushRegexMap {
		compile, err := regexp.Compile(fieldRegex)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code":    0,
				"data":    "",
				"message": err.Error(),
			})
			return
		}
		findResult := compile.FindStringSubmatch(content)
		if len(findResult) > 1 {
			regexResult[field] = strings.ReplaceAll(strings.ReplaceAll(findResult[1], "`", ""), "  ", "")
		} else {
			regexResult[field] = ""
			fmt.Println(findResult)
		}
	}

	go func() {
		message.Pipeline <- message.NewMarkdown(fmt.Sprintf(`
<font color="warning">%s</font>
> 时间: <font color="comment">%s</font>
> 错误码: <font color="comment">%s</font>
> 异常消息: <font color="comment">%s</font>
> 异常类型: <font color="comment">%s</font>
> Trace: <font color="comment">%s</font>
`,
			regexResult["title"],
			regexResult["date_time"],
			regexResult["err_code"],
			regexResult["err_message"],
			regexResult["err_clazz"],
			regexResult["err_trace"]))
	}()

	ctx.JSON(200, gin.H{
		"code":    0,
		"data":    "",
		"message": "ok",
	})
}

func Text(ctx *gin.Context) {
	var txt = TxtStruct{}
	err := ctx.BindJSON(&txt)
	if err != nil {
		errMsg := "Invalid request."
		if err != io.EOF {
			errMsg = err.Error()
		}
		ctx.JSON(200, gin.H{
			"code":    -1,
			"data":    "",
			"message": errMsg,
		})
		return
	}

	go func() {
		message.Pipeline <- message.NewText(txt.Content)
	}()

	ctx.JSON(200, gin.H{
		"code":    0,
		"data":    "",
		"message": "ok",
	})
}

type ErrStruct struct {
	Title      string `json:"title"`
	DateTime   string `json:"date_time"`
	ErrCode    string `json:"err_code"`
	ErrMessage string `json:"err_message"`
	ErrClazz   string `json:"err_clazz"`
	ErrTrace   string `json:"err_trace"`
}

func Error(ctx *gin.Context) {
	var errRequest = ErrStruct{}
	err := ctx.BindJSON(&errRequest)
	if err != nil {
		errMsg := "Invalid request."
		if err != io.EOF {
			errMsg = err.Error()
		}
		ctx.JSON(200, gin.H{
			"code":    -1,
			"data":    "",
			"message": errMsg,
		})
		return
	}

	go func() {
		message.Pipeline <- message.NewMarkdown(fmt.Sprintf(`
<font color="warning">%s</font>
> 时间: <font color="comment">%s</font>
> 错误码: <font color="comment">%s</font>
> 异常消息: <font color="comment">%s</font>
> 异常类型: <font color="comment">%s</font>
> Trace: <font color="comment">%s</font>
`,
			errRequest.Title,
			errRequest.DateTime,
			errRequest.ErrCode,
			errRequest.ErrMessage,
			errRequest.ErrClazz,
			errRequest.ErrTrace))
	}()

	ctx.JSON(200, gin.H{
		"code":    0,
		"data":    "",
		"message": "ok",
	})
}
