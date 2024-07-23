package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed music.mp3
var musicData []byte

// readCloser is a wrapper for *bytes.Reader to implement io.ReadCloser.
type readCloser struct {
	*bytes.Reader
}

// Close implements the io.Closer interface.
func (rc *readCloser) Close() error {
	// No resource to close
	return nil
}

func main() {
	// MP3 dosyasını byte dilimi olarak kullanın
	musicReader := &readCloser{Reader: bytes.NewReader(musicData)}

	// MP3 dosyasını decode et
	streamer, format, err := mp3.Decode(musicReader)
	if err != nil {
		fmt.Println("Error decoding mp3:", err)
		return
	}
	defer streamer.Close()

	// Beep ile müzik çalmak için hoparlörü başlat
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// MP3 dosyasının toplam uzunluğunu hesapla
	duration := float64(streamer.Len()) / float64(format.SampleRate.N(1)) // in seconds

	// Müzik çalmasını başlat
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		// Müzik bittiğinde çalışacak kod
		go func() {
			command := `rundll32.exe powrprof.dll,SetSuspendState 0,1,0`

			// Komutu çalıştırın
			cmd := exec.Command("cmd", "/C", command)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error turning off screen:", err)
			}
			// Programı kapat
			os.Exit(0)
		}()
	})))

	// Müzik son saniyelerinde bir işlem başlatacak bir zamanlayıcı
	go func() {
		startTime := time.Now()
		for {
			// Şu anki süreyi hesapla
			elapsed := time.Since(startTime).Seconds()
			if float64(duration)-elapsed <= 5 { // Son 5 saniye
				// Komutu çalıştır
				go func() {
					command := `rundll32.exe powrprof.dll,SetSuspendState 0,1,0`

					// Komutu çalıştırın
					cmd := exec.Command("cmd", "/C", command)
					err := cmd.Run()
					if err != nil {
						fmt.Println("Error turning off screen:", err)
					}
				}()
				// İşlemi yalnızca bir kez yapın
				break
			}
			time.Sleep(time.Second) // Her saniye kontrol et
		}
	}()

	// Programın bitmesini bekleyin
	select {} // Keeps the program running until manually stopped
}
