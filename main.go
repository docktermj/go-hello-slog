/*
 */
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

func main() {

	ctx := context.TODO()
	logger := slog.New(slog.NewJSONHandler(os.Stderr))

	// Test JSON messages

	aJsonMessage := `
    {
        "date": "2023-03-23",
        "time": "18:53:15.623837476",
        "level": "ERROR",
        "id": "error-60044001",
        "text": "Call to G2_addRecord(CUSTOMERS, 1002, {\"DATA_SOURCE\": \"BOB\", \"RECORD_ID\": \"1002\", \"RECORD_TYPE\": \"PERSON\", \"PRIMARY_NAME_LAST\": \"Smith\", \"PRIMARY_NAME_FIRST\": \"Bob\", \"DATE_OF_BIRTH\": \"11/12/1978\", \"ADDR_TYPE\": \"HOME\", \"ADDR_LINE1\": \"1515 Adela Lane\", \"ADDR_CITY\": \"Las Vegas\", \"ADDR_STATE\": \"NV\", \"ADDR_POSTAL_CODE\": \"89111\", \"PHONE_TYPE\": \"MOBILE\", \"PHONE_NUMBER\": \"702-919-1300\", \"DATE\": \"3/10/17\", \"STATUS\": \"Inactive\", \"AMOUNT\": \"200\"}, G2Engine_test) failed. Return code: -2",
        "duration": 505312,
        "location": "In AddRecord() at g2engineserver.go:66",
        "errors": [{
            "text": "0023E|Conflicting DATA_SOURCE values 'CUSTOMERS' and 'BOB'"
        }],
        "details": {
            "1": "CUSTOMERS",
            "2": 1002,
            "3": {
                "DATA_SOURCE": "BOB",
                "RECORD_ID": "1002",
                "RECORD_TYPE": "PERSON",
                "PRIMARY_NAME_LAST": "Smith",
                "PRIMARY_NAME_FIRST": "Bob",
                "DATE_OF_BIRTH": "11/12/1978",
                "ADDR_TYPE": "HOME",
                "ADDR_LINE1": "1515 Adela Lane",
                "ADDR_CITY": "Las Vegas",
                "ADDR_STATE": "NV",
                "ADDR_POSTAL_CODE": "89111",
                "PHONE_TYPE": "MOBILE",
                "PHONE_NUMBER": "702-919-1300",
                "DATE": "3/10/17",
                "STATUS": "Inactive",
                "AMOUNT": "200"
            },
            "4": "G2Engine_test",
            "5": -2,
            "6": 505312
        }
    }
    `

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(aJsonMessage), &jsonMap)
	if err != nil {
		panic(err)
	}

	logger.Info("hello", "PriorMessage", jsonMap)

	// Log levels and guards

	fmt.Printf("\nLog Levels: Debug: %d; Info: %d; Warn: %d; Error: %d\n\n", slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError)

	fmt.Printf("Log Level    -5: %v\n", logger.Enabled(ctx, -5))
	fmt.Printf("Log Level Debug: %v\n", logger.Enabled(ctx, slog.LevelDebug))
	fmt.Printf("Log Level    -1: %v\n", logger.Enabled(ctx, -1))
	fmt.Printf("Log Level  Info: %v\n", logger.Enabled(ctx, slog.LevelInfo))
	fmt.Printf("Log Level  Warn: %v\n", logger.Enabled(ctx, slog.LevelWarn))
	fmt.Printf("Log Level Error: %v\n", logger.Enabled(ctx, slog.LevelError))
	fmt.Printf("Log Level     5: %v\n", logger.Enabled(ctx, 5))
	fmt.Printf("Log Level     9: %v\n", logger.Enabled(ctx, 9))

}
