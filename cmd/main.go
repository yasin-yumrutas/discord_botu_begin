package main

import (
	"fmt"
	"os"
	"os/signal"

	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Discord botunuzun token'ını buraya girin
	token := "BOT_TOKEN"

	// Discord botunu başlat
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Bot başlatılamadı:", err)
		return
	}

	// Botun "messageCreate" olayına yanıt vermesini sağlama
	dg.AddHandler(messageCreate)

	// Botun çalıştırılması
	err = dg.Open()
	if err != nil {
		fmt.Println("Discord oturumu açılamadı:", err)
		return
	}

	fmt.Println("Bot başarıyla çalışıyor. CTRL+C ile çıkış yapabilirsiniz.")

	// Botun kapatılması
	defer dg.Close()

	// CTRL+C sinyali alındığında kapatma işlemi
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Botun kendi mesajlarını işlememesi için kontrol
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Mesajın içeriği
	content := m.Content

	fmt.Println(content)

	// Gerçek kişi tarafından yazıldıysa belirli bir mesajı ekleyin
	var response string
	if m.Author.Bot {
		response = "Bu mesaj bot tarafından gönderildi."
	} else {
		response = "Bu mesaj gerçek bir kişi tarafından gönderildi."
	}

	// Kaç kişiye iletildiğini gösterme
	numRecipients := len(m.Mentions) + 1 // Kendi mesajımızı da ekleyin
	response += fmt.Sprintf(" Toplam %d kişiye iletilmiştir.", numRecipients)

	// Botun girdiği kanala yanıt verme
	_, _ = s.ChannelMessageSend(m.ChannelID, response)
}
