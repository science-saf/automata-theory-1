package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

func RenderRegisterForm(ctx *gin.Context, result *RegisterResult, user *SiteUser) {
	tplData := gin.H{
		"title":              "Automata Theory - Lab 1, form validation",
		"alertMessage":       "",
		"showAlertName":      false,
		"showAlertEmail":     false,
		"showAlertPassword1": false,
		"showAlertPassword2": false,
	}
	if result != nil && len(result.errors) > 0 {
		if len(result.errors) > 0 {
			if ContainsInt32(result.errors, UserInvalidNickname) {
				tplData["showAlertName"] = true
				tplData["alertName"] = "Invalid nickname"
			}
			if ContainsInt32(result.errors, UserInvalidEmail) {
				tplData["showAlertEmail"] = true
				tplData["alertEmail"] = "Invalid email"
			}
			if ContainsInt32(result.errors, UserWeakPassword) {
				tplData["showAlertPassword1"] = true
				tplData["alertPassword1"] = "Invalid password"
			}
			if ContainsInt32(result.errors, UserPasswordMismatch) {
				tplData["showAlertPassword2"] = true
				tplData["alertPassword2"] = "Passwords don't match"
			}
			tplData["oldNickname"] = user.nickname
			tplData["oldEmail"] = user.email
			tplData["oldPassword1"] = user.password1
			tplData["oldPassword2"] = user.password2
		}
	}
	ctx.HTML(http.StatusOK, "reg-form.tpl", tplData)
}

func RenderUserPage(ctx *gin.Context, user *SiteUser) {
	ctx.HTML(http.StatusOK, "reg-results.tpl", gin.H{
		"title":        "Automata Theory - Lab 1, form validation",
		"userNickname": user.nickname,
		"userEmail":    user.email,
	})
}

func RenderLifeGame(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "life.tpl", gin.H{
		"title": "Automata Theory - Lab2, Game of life",
	})
}

func RenderCalc(ctx *gin.Context, expression string, isAjax string) {
	// calcTests := new(CalcTests)
	// calcTests.run()
	WriteToLogStr("isAjax: +" + isAjax + "+")
	if expression != "" {
		calc := new(Calc)
		calc.Init(expression)
		result := calc.ParseExpr()
		errors := calc.errors
		test := *ctx
		if isAjax == "true" {
			ctx.JSON(http.StatusOK, gin.H{
				"errors": errors,
				"result": result,
			})
		} else {
			ctx.HTML(http.StatusOK, "calc.tpl", gin.H{
				"title":  "Automata Theory - Lab3, Calculator",
				"errors": errors,
				"result": result,
				"test":   test,
			})
		}
	} else {
		ctx.HTML(http.StatusOK, "calc.tpl", gin.H{
			"title": "Automata Theory - Lab3, Calculator",
		})
	}
}

func main() {
	cache := NewSiteUsersCache()
	validator := new(RegisterFormValidator)

	router := gin.Default()
	router.Static("/css", "../site-content/css")
	router.Static("/js", "../site-content/js")
	router.LoadHTMLGlob("../site-content/tpl/*.tpl")
	router.GET("/life", func(ctx *gin.Context) {
		RenderLifeGame(ctx)
	})
	router.GET("/form", func(ctx *gin.Context) {
		RenderRegisterForm(ctx, nil, nil)
	})
	router.GET("/calc", func(ctx *gin.Context) {
		RenderCalc(ctx, "", "")
	})
	router.POST("/calc", func(ctx *gin.Context) {
		RenderCalc(ctx, ctx.PostForm("expression"), ctx.PostForm("is_ajax"))
	})
	router.POST("/form", func(ctx *gin.Context) {
		user := &SiteUser{
			nickname:  ctx.PostForm("userNickname"),
			email:     ctx.PostForm("userEmail"),
			password1: ctx.PostForm("userPassword"),
			password2: ctx.PostForm("userPasswordRepeat"),
		}
		checkResult := validator.Check(user)
		if len(checkResult.errors) == 0 {
			cache.AddUser(user)
			RenderUserPage(ctx, user)
		} else {
			RenderRegisterForm(ctx, &checkResult, user)
		}
	})
	router.Run(":8080")
}

func ContainsInt32(int32Haystack []int32, needle int32) bool {
	for _, value := range int32Haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func ContainsStr(strHaystack []string, needle string) bool {
	for _, value := range strHaystack {
		if value == needle {
			return true
		}
	}
	return false
}

func WriteArrayToLogStr(strArr []string) {
	for _, value := range strArr {
		WriteToLogStr(value)
	}
}

func WriteArrayToLogInt32(int32Arr []int32) {
	for _, value := range int32Arr {
		WriteToLogInt32(value)
	}
}

func WriteToLogInt32(int32var int32) {
	WriteToLogStr(fmt.Sprintf("%d", int32var))
}

func WriteToLogBool(boolVar bool) {
	WriteToLogStr(fmt.Sprintf("%t", boolVar))
}

func WriteToLogFloat32(float32var float32) {
	WriteToLogStr(fmt.Sprintf("%f", float32var))
}

func WriteToLogStr(str string) {
	fmt.Println("===== Log: " + str)
}
