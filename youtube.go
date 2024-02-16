package youtube_downloader

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func TÃ©lÃ©charger(videoURL string) error {
	client := youtube.Client{}
	videoID := extractVideoID(videoURL)
	video, err := client.GetVideo(videoID)
	if err != nil {
		return fmt.Errorf("erreur lors de la rÃ©cupÃ©ration des dÃ©tails de la vidÃ©o: %v", err)
	}
	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		return fmt.Errorf("erreur : %v", err)
	}
	file, err := os.Create("video.mp4")
	if err != nil {
		return fmt.Errorf("erreur : %v", err)
	}
	defer file.Close()
	_, err = stream.WriteTo(file)
	if err != nil {
		return fmt.Errorf("erreur lors de l'enregistrement: %v", err)
	}
	fmt.Println("TÃ©lÃ©chargement terminÃ©!")
	return nil
}

func extractVideoID(url string) string {
	parts := strings.Split(url, "=")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
