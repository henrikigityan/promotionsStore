package service

import (
	"encoding/csv"
	"io"
	"net/http"
	"strconv"
	"test/simpleApi/model"
	"test/simpleApi/dao"
	"sync"
	"time"
)

const (
	maxMemory       = 10 * 1024 * 1024 * 10000 // Max file size: 10MB
	concurrentTasks = 2               // Number of concurrent processing goroutines
)

func HandleUpload(r *http.Request, promotionDao dao.PromotionDao) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		return
	}

	// Parse the multipart form data
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		return
	}

	// Get the file from the form data
	file, _, err := r.FormFile("testCsv")
	if err != nil {
		return
	}
	// defer file.Close()
	// fmt.Println(handler)

	// Read the uploaded CSV file in chunks
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to communicate the chunks of data
	chunkChannel := make(chan []string)

	// Start concurrent processing goroutines
	for i := 0; i < concurrentTasks; i++ {
		wg.Add(1)
		go processCSVChunks(&wg, chunkChannel, promotionDao)
	}

	// Read the file and send chunks to the processing goroutines
	for {
		chunk, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		chunkChannel <- chunk
	}

	close(chunkChannel) // Close the channel to signal that all chunks have been sent

	// Wait for all goroutines to finish processing
	wg.Wait()
}

func processCSVChunks(wg *sync.WaitGroup, chunkChannel <-chan []string,
	promotionDao dao.PromotionDao) {
	defer wg.Done()

	var promotions [] model.Promotion
	for chunk := range chunkChannel {
		// Process each chunk of data
		// You can perform any desired operations on the chunk here
		layout := "2006-01-02 15:04:05 -0700 MST"
		id := chunk[0]
		price, err := strconv.ParseFloat(chunk[1], 10)
		expirationDate, err := time.Parse(layout, chunk[2])
		if err == io.EOF {
            break
        }
        if err != nil {
            return
        }
		data := model.Promotion{  
			Id: id,
			Price: price,
			ExpirationDate : expirationDate,
		}
		promotions = append(promotions, data)
	}
	promotionDao.BatchInsert(promotions)
}