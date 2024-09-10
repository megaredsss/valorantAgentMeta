<h1 align="center">Valorant agent meta</h1>
If you play Valorant, you know that agent selection is a big part of your potential victory. Your choice depends on the card and role, but the stats sites have numbers that don't correlate with each other and seem to be random. To me the real stats are from professional games, the best site about them is https://www.vlr.gg/, but it has an unreadable table. This program allows you to simply enter the name of a map and get the 5 agents with the highest winrate on that map. The program generates data based on data from the site and from the last official tournament.

# Installation and usage
This section all about installation. There two diffent path of installation: run native(Classic version) or run with docker
## Classic version
Installation is very simple
- Clone the repository or download ZIP

- After that check your Go version in console by this
```console
go version
```
if you dont have Go visit this https://go.dev/dl/
- Jump to the folder there you donwload program, for example like this in console
```console
cd path\to\folder
```
- Open this folder in terminal or cmd and type this
```console
go run .
```
Command above runs main func in valorantAgentMeta.go 
- Result should be like that
```console
Scrapping Completed
Connect: https://www.vlr.gg/event/agents/2097/valorant-champions-2024
Type map name(like this: Bind)
```
- After that you should type map and program shows the best pick for this map. Correct output for Bind
```console
Scrapping Completed
Connect: https://www.vlr.gg/event/agents/2097/valorant-champions-2024
Type map name(like this: Bind)
Bind
viper raze gekko brimstone fade
```
## Docker version
- If you have docker, you can run this application with it. Download the program, navigate to the folder as in the classic version, and enter
```console
 docker build . -t valorantagentmeta
```
- After build you should run the container
```console
 docker run -it valorantagentmeta
```
And follow classic version
