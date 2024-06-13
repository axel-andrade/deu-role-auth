package middlewares

import (
	"log"
	"net/http"
	"strings"

	common_ptr "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/presenters/common"
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

const (
	RequestBody    = "body"
	RequestHeaders = "headers"
	RequestQuery   = "query"
)

func errorsToStructs(errors []gojsonschema.ResultError) []common_ptr.ValidateDetail {
	var validationErrors []common_ptr.ValidateDetail

	for _, err := range errors {
		validationErrors = append(validationErrors, common_ptr.ValidateDetail{
			Namespace: err.Field(),
			Tag:       err.Type(),
			Param:     err.Description(),
		})
	}

	return validationErrors
}

func ValidateRequest(schemaPath string) gin.HandlerFunc {
	schemaLoader := gojsonschema.NewReferenceLoader("file://internal/adapters/primary/http/schemas/" + schemaPath + "_schema.json")
	jsonSchemaPtr := common_ptr.JsonSchemaPresenter{}
	schemaJson, err := schemaLoader.LoadJSON()

	if err != nil {
		log.Fatal("Failed to load schema json")
	}

	return func(c *gin.Context) {
		var input = make(map[string]any)
		var validationResult *gojsonschema.Result

		properties := schemaJson.(map[string]any)["properties"].(map[string]any)
		for paramName := range properties {
			if paramName == RequestBody {
				var bodyMap = make(map[string]any)
				if err := c.BindJSON(&bodyMap); err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, jsonSchemaPtr.Format(nil))
					return
				}
				input[RequestBody] = bodyMap
			} else if paramName == RequestHeaders {
				var headersMap = make(map[string][]string)
				for key, values := range c.Request.Header {
					headersMap[strings.ToLower(key)] = values
				}
				input[RequestHeaders] = headersMap
			} else if paramName == RequestQuery {
				var queryMap = make(map[string]string)
				for key := range c.Request.URL.Query() {
					queryMap[key] = c.Query(key)
				}
				input[RequestQuery] = queryMap
			}
		}

		validationResult, err = gojsonschema.Validate(
			schemaLoader,
			gojsonschema.NewGoLoader(input),
		)

		if err != nil || !validationResult.Valid() {
			c.AbortWithStatusJSON(http.StatusBadRequest, jsonSchemaPtr.Format(errorsToStructs(validationResult.Errors())))
			return
		}

		if value, ok := input[RequestBody]; ok {
			c.Set(RequestBody, value)
		}

		c.Next()
	}
}
