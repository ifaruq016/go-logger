package utils

import (
	"fmt"
	"time"
	"math/rand"
)

func ApiCallGenerator() string{
	now := time.Now()
	formatedDatetime := now.Format("19980130_141210")
	randInt := rand.Intn(900) + 100
	
	generatedApiIdentifier := fmt.Sprintf("API_CALL_%s%d", formatedDatetime, randInt)

	return generatedApiIdentifier
}