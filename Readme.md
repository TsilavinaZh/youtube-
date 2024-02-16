# youtube

Ce module charger des vidéos YouTube.

## Installation

Pour utiliser ce module, vous pouvez l'installer via la commande suivante :go get github.com/TsilavinaZh/youtube

Voici un exemple d'utilisation du module pour télécharger une vidéo YouTube :

```go
package main

import (
	"fmt"
	"github.com/TsilavinaZh/youtube"
)

func main() {
	videoURL := "https://www.youtube.com/watch?v=YOUR_VIDEO_ID"
	err := youtube_downloader.Télécharger(videoURL)
	if err != nil {
		fmt.Printf("Erreur lors du téléchargement de la vidéo: %v\n", err)
		return
	}
}Assurez-vous de remplacer YOUR_VIDEO_ID par l'ID de la vidéo YouTube que vous souhaitez télécharger.LicenseCe projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.Ce fichier README.md comprend toutes les informations nécessaires sur le module youtube_downloader, y compris l'installation, un exemple d'utilisation et des informations de licence.
