package twitch

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
)

type commandHandler func(twitch.PrivateMessage)

type command struct {
	Trigger     string   `json:"trigger"`
	Response    string   `json:"response"`
	Permissions []string `json:"permissions"`
}

type alias struct {
	Alias   string `json:"alias"`
	Trigger string `json:"trigger"`
}

type commands struct {
	Commands []command `json:"commands"`
	Aliases  []alias   `json:"aliases"`
}

var (
	errorCommandAlreadyExists = errors.New("command already exists")
	errorCommandDoesNotExist  = errors.New("command does not exist")
	errorAliasAlreadyExists   = errors.New("alias already exists")
)

func (s *ChatService) loadCommands() error {
	log.Printf("Loading command file: %s\n", s.config.Commands)

	rawFile, err := os.ReadFile(s.config.Commands)

	if err != nil {
		return err
	}

	commands := commands{}

	err = json.Unmarshal(rawFile, &commands)

	if err != nil {
		return err
	}

	for _, v := range commands.Commands {
		permissionsMap := toPermissionMap(v.Permissions)
		err := s.registerCommand(v.Trigger, s.defaultCommandHandler(v.Response, getPermissionMask(permissionsMap)))

		if err != nil {
			log.Printf("Error registering command '%s': %s", v.Trigger, err.Error())
		}
	}

	for _, v := range commands.Aliases {
		s.registerAlias(v.Alias, v.Trigger)

		if err != nil {
			log.Printf("Error registering alias '%s': %s", v.Trigger, err.Error())
		}
	}

	return nil
}

func (s *ChatService) registerCommands() {
	log.Println("Registering commands")

	s.registerCommand("!reload", s.reloadCommandsHandler)
	s.registerCommand("!dc", s.discordInviteHandler)
	s.registerAlias("!discord", "!dc")

	s.loadCommands()
}

func (s *ChatService) registerCommand(trigger string, handler commandHandler) error {
	if _, ok := s.commands[trigger]; ok {
		return errorCommandAlreadyExists
	}

	s.commands[trigger] = handler
	return nil
}

func (s *ChatService) registerAlias(alias string, trigger string) error {
	if _, ok := s.commands[alias]; ok {
		return errorAliasAlreadyExists
	}

	if h, ok := s.commands[trigger]; ok {
		s.commands[alias] = h
		return nil
	}

	return errorCommandDoesNotExist
}

func (s *ChatService) executeCommand(message twitch.PrivateMessage) {
	fields := strings.Fields(message.Message)

	if len(fields) == 0 {
		return
	}

	if h, ok := s.commands[fields[0]]; ok {
		h(message)
	}
}
