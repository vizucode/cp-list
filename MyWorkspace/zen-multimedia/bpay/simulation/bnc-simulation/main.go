package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RequestToken struct {
	GrantType string `json:"grantType"`
}

func main() {

	app := fiber.New()

	app.Post("/open/bi/v1.0/get/token", func(ctx *fiber.Ctx) error {

		var payload RequestToken
		err := ctx.BodyParser(&payload)
		if err != nil {
			return err
		}

		signature := ctx.Get("X-SIGNATURE")
		string2sign, err := String2Sign("POST", "http://localhost:6677/open/bi/v1.0/get/token", "2019-07-03T12:08:56-07:00", `{"grantType":"client_credentials"}`)
		log.Println("server: ", string2sign)
		if err != nil {
			return ctx.JSON(map[string]interface{}{
				"accessToken": "",
				"expiresIn":   "",
				"code":        "4010200",
				"msg":         "Forbidden Access",
				"tokenType":   "",
			})
		}

		err = VerifySignature(signature, string2sign)
		if err != nil {
			log.Println(err)
			return ctx.JSON(map[string]interface{}{
				"accessToken": "",
				"expiresIn":   "",
				"code":        "4010201",
				"msg":         "Forbidden Access",
				"tokenType":   "",
			})
		}

		if !strings.EqualFold(payload.GrantType, "client_credentials") {
			return ctx.Status(400).JSON(map[string]interface{}{
				"accessToken": "",
				"expiresIn":   "",
				"code":        "4001101",
				"msg":         "Bad Request",
				"tokenType":   "Bearer",
			})
		}

		return ctx.Status(200).JSON(map[string]interface{}{
			"accessToken": "eyJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJCWUIiLCJzdWIiOiIwMDA1MTAwMDAyMzMiLCJpYXQiOjE3MDE4NDgyMzMsImV4cCI6MTcwMjQ1MzAzM30.3UNMyNge7iANifz_fzU7qR15nM98QTiIRLD_mQw-_30",
			"expiresIn":   "900",
			"code":        "2007300",
			"msg":         "Successful",
			"tokenType":   "Bearer",
		})
	})

	app.Post("/open/bi/private/v1.0/transfer-interbank", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(map[string]interface{}{
			"responseCode":    "4001800",
			"responseMessage": "Successful",
			"amount": map[string]interface{}{
				"value":                "1000.00",
				"currency":             "IDR",
				"beneficiaryAccountNo": "888801000003301",
			},
		})
	})

	log.Fatal(app.Listen(":6677"))
}
