package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func RunDiscordBot() {
    // create session
    discord, err := discordgo.New("Bot " + BotToken)
    if err != nil {
        log.Fatal("Failed to set up Discord session.")
    }
    // add event handler
    discord.AddHandler(newMessage)
    // open session
    discord.Open()
    defer discord.Close()
    // keep running untill interupt
    fmt.Println("Discord bot running ...")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
    // prevent bot from answering its own messages
    if message.Author.ID == discord.State.User.ID {
        return
    }
    //respond to user message 
    switch {
        case strings.Contains(message.Content, "!hi"):
            discord.ChannelMessageSend(message.ChannelID, "Hello!")
        case strings.Contains(message.Content, "!bye"):
            discord.ChannelMessageSend(message.ChannelID, "Goodbye!")
    }
}
