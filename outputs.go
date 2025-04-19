package main

import (
	"fmt"
	"sort"

	models "github.com/RajathSVasisth/interview/models"
)

func printLargestIndexes(data []models.IndexInfo) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].SizeGB > data[j].SizeGB
	})

	fmt.Println("\nPrinting largest indexes by storage size")
	for i := range min(5, len(data)) {
		fmt.Printf("Index: %s\nSize: %.2f GB\n", data[i].Name, data[i].SizeGB)
	}
}

func printMostShards(data []models.IndexInfo) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Shards > data[j].Shards
	})

	fmt.Println("\nPrinting largest indexes by shard count")
	for i := range min(5, len(data)) {
		fmt.Printf("Index: %s\nShards: %d\n", data[i].Name, data[i].Shards)
	}
}

func printLeastBalanced(data []models.IndexInfo) {

	sort.Slice(data, func(i, j int) bool {
		return data[i].BalanceRatio > data[j].BalanceRatio
	})

	fmt.Println("\nPrinting least balanced indexes")
	for i := range min(5, len(data)) {
		fmt.Printf("Index: %s\n", data[i].Name)
		fmt.Printf("Size: %.2f GB\n", data[i].SizeGB)
		fmt.Printf("Shards: %d\n", data[i].Shards)
		fmt.Printf("Balance Ratio: %d\n", data[i].BalanceRatio)
		fmt.Printf("Recommended shard count is %d\n", max(1, data[i].RecommendedShards))
	}
}
