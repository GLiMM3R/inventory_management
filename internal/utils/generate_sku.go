package utils

import (
	"regexp"
	"strings"
)

func GenerateSKU(name string, length int, prefix string, characters int) string {
	// Convert the name to uppercase
	upperName := strings.ToUpper(name)

	// Remove special characters using regex but keep spaces for word splitting
	reg, _ := regexp.Compile("[^A-Z0-9 ]+")
	cleanedName := reg.ReplaceAllString(upperName, "")

	// Split the cleaned name by spaces
	words := strings.Fields(cleanedName)

	// Take the first 4 characters of each word
	var skuParts []string
	for _, word := range words {
		if len(word) > characters {
			skuParts = append(skuParts, word[:characters])
		} else {
			skuParts = append(skuParts, word)
		}
	}

	// Join the parts together and truncate to the desired total length
	sku := strings.Join(skuParts, "")

	// Limit the SKU to the desired length
	if len(sku) > length {
		sku = sku[:length]
	}

	// Return the SKU with the optional prefix
	return prefix + sku
}
