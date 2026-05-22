# Teto Word of the Day: Generator
![The program in "action"](./screenshot.png?raw=true)

This is a simple program written in Golang for generating a random word at the bottom of popular UTAU voicebank Kasane Teto.
# How It Works
The program is simply, really.  It uses [Ebitengine](https://github.com/hajimehoshi/ebiten) to render a picture of a Kasane Teto plush (and the text), [etxt](https://github.com/tinne26/etxt) to render the image macro text you would see in memes, and [GopherLua](https://github.com/yuin/gopher-lua) to deliver the word "of the day" to the program.

You can run the program with any word you want, but be careful; pass more than one word, and your "Teto Word of the Day" will be "overflow".

If you do not pass any word, then that's okay; the program will use its Lua library to call a random word from its default list (for example, Teto's word of the day will be "Github", as seen in the screenshot.)
# How to Run
First off, you may need to have [Golang](https://go.dev/dl/) installed (version 1.25.0 or newer), and that is that.

Then you can execute the program in three different ways.
## via Terminal
As usual, you can always run this simple command (if you have internet access, of course):
```bash
go run https://github.com/EbenezerChicago/teto-word@latest
```
## via Source
If you want to run using the source code, you can always execute these commands:
```bash
git clone https://github.com/EbenezerChicago/teto-word
cd teto-word
go run . [any word you want (optional)]
```
## via Startup
Want Teto to greet you with a daily word?  You can always copy the script used in the "via Terminal" method, and paste it in a new `.sh` file.

Just make sure you use `crontab -e` and add `@reboot /path/where/you/put/the/script.sh` at the bottom of it.

That way, you can feel machine love without needing a Mac.
# Credits
This program was authored by me, Ebenezer Chicagom and is licensed under the MIT License.

All rights reserved.  Batteries not included.
## Additional Credits
Ebitengine was made by [Hajime Hoshi](https://github.com/hajimehoshi).

etxt was made by [tinne26](https://github.com/tinne26).

GopherLua was made by [Yusuke Inuzuka](https://github.com/yuin).

Kasane Teto was made by TWINDRILL.
