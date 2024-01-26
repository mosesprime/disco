package main

import (
    discord "github.com/mosesprime/disco/internal/discord"
    docker "github.com/mosesprime/disco/internal/docker"
)

func main()  {
    docker.X()
    discord.BotToken = "DISCORD TOKEN ID"
    discord.RunDiscordBot()
}

