package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
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
		"https://pm1.narvii.com/6525/bdf76097499dd104090cc1398f8a3d28d36c5c7c_hq.jpg",
		"https://pm1.narvii.com/6525/cb3734b8c063b13fe9640fab7f58d2dc137f037e_hq.jpg",
		"https://pm1.narvii.com/6525/2db527c913fe3be5d7b34a3a520b7cf2e08505c0_hq.jpg",
		"https://pm1.narvii.com/6474/ac64930de94f131c20927abedd1e7b6a2395cd02_hq.jpg",
		"https://pm1.narvii.com/6474/f48c687672e4f3ffe7c7c7db8cf35a552ae8dd48_hq.jpg",
		"https://media1.tenor.com/images/f54e07d2285bfea75a29b7e9ac82e70a/tenor.gif",
		"https://pbs.twimg.com/media/C8_NeHKXkAAP31D.jpg",
		"https://i.pinimg.com/474x/2c/7c/b6/2c7cb632e7f6d1b23483d17fbc8d598e.jpg",
	}
	if message.Content == "... esqueleto" {
		discord.ChannelMessageSend(message.ChannelID, imageURLs[rand.Intn(len(imageURLs))])
	}
}

func init() {
	gotenv.Load()
}

func main() {
	rand.Seed(time.Now().Unix())
	token := os.Getenv("SKELETOKEN")
	discord, err := discordgo.New(fmt.Sprintf("Bot %v", token))
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
