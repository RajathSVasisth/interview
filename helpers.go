package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	models "github.com/RajathSVasisth/interview/models"
)

func getDataFromFile(filename string) ([]models.IndexInfo, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var rawIndexes []models.RawIndex
	err = json.Unmarshal(raw, &rawIndexes)
	if err != nil {
		return nil, err
	}

	return parseRawIndexes(rawIndexes)
}

func getDataFromServer(endpoint string, days int) ([]models.IndexInfo, error) {
	var allData []models.RawIndex
	for i := 1; i <= days; i++ {
		date := time.Now().AddDate(0, 0, -i)
		url := fmt.Sprintf("https://%s/_cat/indices/*%d*%02d*%02d?v&h=index,pri.store.size,pri&format=json&bytes=b",
			endpoint, date.Year(), date.Month(), date.Day())

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var dailyData []models.RawIndex
		if err := json.Unmarshal(body, &dailyData); err != nil {
			return nil, err
		}

		allData = append(allData, dailyData...)
	}
	return parseRawIndexes(allData)
}

func parseRawIndexes(raw []models.RawIndex) ([]models.IndexInfo, error) {
	var data []models.IndexInfo
	for _, r := range raw {
		size, err := strconv.ParseInt(r.SizeStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing size for index %s: %v", r.Name, err)
		}
		shards, err := strconv.Atoi(r.ShardsStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing shards for index %s: %v", r.Name, err)
		}
		gb := float64(size) / (1e9)
		data = append(data, models.IndexInfo{
			Name:              r.Name,
			SizeGB:            gb,
			Shards:            shards,
			BalanceRatio:      int(math.Round(gb / float64(shards))),
			RecommendedShards: int(math.Round(gb / 30)),
		})
	}
	return data, nil
}
