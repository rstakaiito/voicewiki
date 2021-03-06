Voicepedia
=========

[![API Documentation](http://img.shields.io/badge/api-Godoc-blue.svg?style=flat)](http://godoc.org/github.com/tamnd/voicewiki)
[![BSD License](http://img.shields.io/badge/license-BSD-red.svg?style=flat)](http://opensource.org/licenses/MIT)

Voicepedia is an open source web application written in Go in the server-side and HTML5 technologies in the client side. You can use your voice to tell which Wikipedia entry to search for. No need to use your eyse, the entry's content will be reading to you!

Voicepedia is also a social web app, when you can record your voice and hear voices by  people around the world!

#### Voicepedia can run on desktop...

![desktop](https://raw.githubusercontent.com/tamnd/voicewiki/master/screenshots/desktop.png)

#### Or on mobile phone...

![mobile](https://raw.githubusercontent.com/tamnd/voicewiki/master/screenshots/mobile.png)

#### with voice commands and fancy gestures...
- Draw **R** to listen recent articles read by real people
- Draw **S** to start searching
- Draw **?** to listen how to use Voicepedia

![gesture](https://raw.githubusercontent.com/tamnd/voicewiki/master/screenshots/gestures.png)

#### or if you want to read this wikipedia and contribute, just say "Record".
![record](https://raw.githubusercontent.com/tamnd/voicewiki/master/screenshots/record.png)


### How to install

#### Open source library used

* [Wikipedia extractor](https://github.com/bwbaugh/wikipedia-extractor)
* [RethinkDB](http://rethinkdb.com/)
* [Redis](http://redis.io/)
* [Goji](https://goji.io/)
* [Annyang](https://github.com/TalAter/annyang)
* [Recordmp3js](https://github.com/nusofthq/Recordmp3js)
* [Mouse Gesture Recognition](http://www.bytearray.org/?p=91)

#### Database

* Start RethinkDB and Redis
* Download lastest wikipedia database at [Wikipedia database dump](http://download.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2)
* Uncompress using bzip2 (44GB uncompressed)
* Clean database using Wikipedia extractor (about 5 - 6 hours)
* Import database using scripts/import.go (about 4 hours)
* Index database using scripts/index.go (about 4.5 hours)

#### Running

```
go get github.com/tamnd/voicewiki
cd $GOPATH/src/github.com/tamnd/voicewiki
go build && ./voicewiki
```


### What includes in the source code?

* A index/search engine write from scratch (scripts/index.go, model/article/search.go)
* A light HTTP middleware (middleware)
* A gesture engine (/public/public/js/gesture.js, /public/public/js/pad.js)

And much more...

### License
Under BSD license.
