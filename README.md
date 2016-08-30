##### disqusScraper

disqusScraper is a console-based go-routine oriented Go application that lets you extract all popular threads from any given disqus forum name that is not private. For private disqus forums, disqusScraper accumulates a list of all top commenters and individually parses their activities to locate all threads that relate to the given disqus forum name.

> This project was created in contract with removeyourmedia.com in order to 
> make the disqus platform more accessible to scrutiny and thereby to help fight piracy and thriving pirated content on the aforementioned platform.

### Tech

disqusScraper uses the following components to work properly:

* [https://github.com/go-playground/pool] - Package pool implements a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.

### Installation

disqusScraper requires [Golang](https://golang.org/) v1.7+ to run.
Install the dependencies and devDependencies and compile the project using following command.
```sh
$ go build
```

### Usage Instructions
```sh
$ c:\myEpicFolder\disqusScraper.exe -forum=kissanime -worker=10 -debug=true -nonoise=true
```
where;
* **-forum**	switch takes the name of the forum that we will scrape users off and get links; default “fiestaonline”
* **-worker**	switch defines how many simultaneous “threads” are doing the job; default 10
* **-debug**	switch defines whether we want to be verbose; default true
* **-nonoise**	switch decides whether to only pick up links that stem from given forum name, or ignore that and pick up everything, considered true if forum is public; defaults true
* **-parseusers**	switch decides to parse users even if forum is public, considered true if forum 		is private; defaults to false

### Todos

 - Filter User comments to only show links that belong to provided forum name. (Done)
 - Get all links from a given disqus forum name without getting their users if the forum isn’t private. (Probably done)
 - Write Tests
 - Add Web Interface

N.B. More discussions in source-code

### Example output text-file contents:
```html
https://disqus.com/home/discussion/kissmanga/read_manga_happy_if_you_died_ch014_online_in_high_quality
https://disqus.com/home/discussion/kissmanga/revival_man_manga_read_revival_man_manga_online_in_high_quality
https://disqus.com/home/discussion/kissmanga/read_manga_young_gun_ch003_online_in_high_quality
https://disqus.com/home/discussion/kissmanga/read_manga_i_am_my_wife_ch002_online_in_high_quality
```
 ####  Contact:
 *  Manish Prakash Singh
 *  contact@kryptodev.com
 *  Skype: kryptodev

License
----

This work is licensed under a Creative Commons Attribution-ShareAlike 4.0 International License.