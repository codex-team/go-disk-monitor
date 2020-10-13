go-disk-monitor
==========

Disk monitoring and alerts via CodeX Bot.

Usage
-----

Run binary file with the following arguments

```
Usage of ./diskMonitor-linux:
  -alert int
    	disk usage percentage to trigger notification (default 90)
  -debug
    	show debug information
  -path string
    	path to the disk volume (default "/")
  -webhook string
    	notification URI from CodeX Bot
```

If you specify `-webhook` argument, you will be notified in Telegram via CodeX Bot.
You can get `webhook` from [@codex-bot](https://t.me/@codex-bot) as written in [CodeX Bot Docs](https://github.com/codex-bot/notify#getting-started).

Output
-----

Sample output of `./diskMonitor-linux -debug -alert 45`
```
2020/10/13 14:20:13 All: 46.90 GB
2020/10/13 14:20:13 Used: 23.49 GB
2020/10/13 14:20:13 Free: 23.42 GB
2020/10/13 14:20:13 Percent: 50%
2020/10/13 14:20:13 🔥🚒 Running out of space `23.42GB(50%)` on server server-1263
```

Installation
-----

```
go get github.com/n0str/go-disk-monitor
```