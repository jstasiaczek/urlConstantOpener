# urlConstantOpener

This application is part of bigger set with fakeBrowser.

Basically it is used to open url on host machine from virtual machine.
It is checking for new file in predefined directory and  it is opening url found in file.
It can also listen on some port as http server and open urls sended to it.

## Configuration

```
{
	"DirectoryToScan" 	: "/where/to/read/UrlSync",
	"LogFile" 			: "/where/to/save/.urlSync.log",
	"UseFiles" 			: true,
	"UseHttp" 			: false,
	"UseFileLog"		: false,
	"HttpPort"			: 4040,
	"SecretKey"			: "some-secret-key"
}

```

Description:
* DirectoryToScan **[existing directory path]** _directory where files to read are saved, be aware that opened files are deleted._
* LogFile **[existing file path]** _path of file to save logs_
* UseFiles **[true|false]** _do you want to scan directory__
* UseHttp **[true|false]** _do you want to listen on http__
* UseFileLog **[true|false]** _do you want to save logs in file__
* HttpPort **[int]** _port where app will listen__
* SecretKey **[some randor string]** _secter key used only for http__

Command line parameters:
* -c config file location

## How to use

Copy Copy **config.json.dist** to **config.json** and fill it with correct values. Now run it as background task. On ubuntu you can configure app to run on login, run **gnome-session-properties** in terminal.