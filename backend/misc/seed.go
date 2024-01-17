package misc

import (
	"backend/internal/database"
	"backend/internal/models"
	"context"
	"math/rand"
)

func Seed(dbService database.Service, numAdmins, numEvents, maxImagesPerEvent int) error {

	// Generate and insert Admins
	var adminIDs []string
	for i := 0; i < numAdmins; i++ {
		admin := createRandomAdmin()
		adminID, err := dbService.CreateAdmin(context.Background(), admin)
		if err != nil {
			return err
		}
		adminIDs = append(adminIDs, adminID)
	}

	//TODO: Generate and insert Events
	return nil
}

// func createRandomEventRequest(adminIDs []string, maxImagesPerEvent int) models.CreateEventRequest {

// 	isDraft := rand.Intn(2) == 0 // 50% chance of being a draft
// 	var published_at *time.Time

// 	// a draft event has no published_at, otherwise it needds one
// 	if !isDraft {
// 		date := randomDate()
// 		published_at = &date
// 	}

// 	return models.CreateEventRequest{
// 		EventName:   fmt.Sprintf("Event %d", rand.Intn(100)),
// 		Date:        randomDate(),
// 		Description: fmt.Sprintf("Description for event %d", rand.Intn(100)),
// 		Content:     fmt.Sprintf("Content for event %d", rand.Intn(100)),
// 		IsDraft:     rand.Intn(2) == 0,
// 		PublishedAt: published_at,
// 		Images:      createRandomImages(maxImagesPerEvent),
// 		AuthorIDs:   selectRandomAdmins(adminIDs),
// 	}
// }

func createRandomAdmin() models.Admin {
	return models.Admin{
		Name:     randomString(10),
		Email:    randomString(10) + "@example.com",
		Position: randomString(10),
		Status:   "active",
	}
}

func randomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// func selectRandomAdmins(adminIDs []string) []string {
// 	rand.Shuffle(len(adminIDs), func(i, j int) {
// 		adminIDs[i], adminIDs[j] = adminIDs[j], adminIDs[i]
// 	})

// 	numAdmins := rand.Intn(len(adminIDs)) + 1 // Ensure at least one admin is selected
// 	return adminIDs[:numAdmins]
// }

// func createRandomImages(max int) []models.EventImage {
// 	var images []models.EventImage
// 	numImages := rand.Intn(max + 1) // Number of images can be from 0 to max

// 	for i := 0; i < numImages; i++ {
// 		images = append(images, models.EventImage{
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 			ImageURL:  fmt.Sprintf("https://example.com/image%d.jpg", rand.Intn(100)),
// 			IsDisplay: i == 0, // Make the first image the display image
// 		})
// 	}

// 	return images
// }

// func randomDate() time.Time {
// 	// Define the range for the random date
// 	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC) // Start date (e.g., Jan 1, 2020)
// 	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC) // End date (e.g., Dec 31, 2023)

// 	// Calculate total duration between start and end
// 	totalDays := end.Sub(start).Hours() / 24

// 	// Generate a random number of days to add to start
// 	daysToAdd := rand.Intn(int(totalDays))

// 	// Add the random number of days to the start date
// 	randomDate := start.AddDate(0, 0, daysToAdd)

// 	// Return the random date as a time.Time object
// 	return randomDate
// }
