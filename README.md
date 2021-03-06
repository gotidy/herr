# Herr

### Errors

`Herr` implements the error model of the well established
[Google API Design Guide](https://cloud.google.com/apis/design/errors). Additionally, it makes the fields `request` and `reason` available. A complete `Herr` error response looks like this:

```json
{
    "error": {
        "code": 400,
        "status": "Bad Request",
        "request": "8512585f-587a-40bf-9ec7-2937f10d1d97",
        "reason": "INVALID_PARAMETER",
        "items": [
            {
                "name": "user",
                "code": "EMPTY",
                "description": "User is not defined"
            }
        ],
        "details": [
            {"additional": "info"}
        ],
        "message":"foo"
    }
}
```

### Open API Error description

[Open API file](openapi.yml)

```yaml
components:
  schemas:
    ErrorItem:
      type: object
      properties:
        name:
          type: string
          description: Name of parameter or field.
          example: error message
        code:
          type: string
          description: Error code.
          example: "Empty"
        Description:
          type: string
          description: Description.
          example: The field is empty
      required:
        - "name"

    Error:
      type: object
      additionalProperties: false
      properties:
        code:
          type: integer
          description: HTTP status code.
          example: 400
        status:
          type: string
          description: Status.
          example: "Bad Request"
        reason:
          type: string
          description: Error reason.
          example: INVALID_PARAMETERS
        request_id:
          type: string
          description: Request ID.
          example: c4b9df35-1716-4099-b24f-f844b5491e22
        description:
          type: string
          description: Description.
          example: Something went wrong.
        debug:
          type: string
          description: Debug information.
          example: S0011.
        details:
          type: object
        items:
          type: array
          items:
            $ref: "#/components/schemas/ErrorItem"
      required:
        - "status"
        - "code"
```
