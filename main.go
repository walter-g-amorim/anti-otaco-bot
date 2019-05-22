package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix string
	botID         string
)

func onReady(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(0, "Esmurrando uns otaco")
	if err != nil {
		fmt.Println("Deu ruim no status change painho")
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == botID || message.Author.Bot {
		return
	}
	imageURLs := [...]string{
		"https://cdn.discordapp.com/attachments/84109753864687616/445391728212705290/C4gp0JsWYAA3zza.png",
		"https://pbs.twimg.com/media/C8RPiHoWsAER2Gq.jpg",
		"https://cdn.discordapp.com/attachments/84109753864687616/445391821187579914/C8_MNshXYAAgI38.png",
	}
	if message.Content == "... esqueleto" {
		discord.ChannelMessageSend(message.ChannelID, imageURLs[rand.Intn(len(imageURLs))])
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	discord, err := discordgo.New("Bot token")
	if err != nil {
		fmt.Println("Deu ruim painho")
	}
	user, err := discord.User("@me")
	if err != nil {
		fmt.Println("Cadê minha conta porra?")
	}
	botID = user.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(onReady)

	err = discord.Open()
	if err != nil {
		fmt.Println("Quebrou a merda da conexão")
	}
	defer discord.Close()

	<-make(chan struct{})
}
