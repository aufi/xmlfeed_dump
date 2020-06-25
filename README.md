# xmlfeed_dump

An RSS/Atom or other xml-based feed dump tool.

**This is just a project for golang programming practise and is not considered as stable/completed, however could be useful.**

## Installation

```
$ wget https://github.com/aufi/xmlfeed_dump/raw/master/xmlfeed_dump
$ chmod +x xmlfeed_dump
```

## Usage

Possible options

```
$ ./xmlfeed_dump -h
Usage of xmlfeed_dump:
  -destination string
        Local directory path where items should
  -url string
        RSS or Atom feed URL
```

Example usage

```
$ ./xmlfeed_dump -url http://cvut2tut.blogspot.com/feeds/posts/default
2020/06/25 20:56:38 Starting RSS/Atom feed website dumper for http://cvut2tut.blogspot.com/feeds/posts/default
2020/06/25 20:56:38 Preparing destination directory..
2020/06/25 20:56:38 Fetching the feed..
2020/06/25 20:56:39 Have valid feed with name: Ze ČVUTu na TUT Erasmus
2020/06/25 20:56:39 Fetching items..
2020/06/25 20:56:39 Storing Item from URL:http://cvut2tut.blogspot.com/2010/08/doprava-v-tampere.html
2020/06/25 20:56:39 Stored to: /home/user/cvut2tut.blogspot.com/http:--cvut2tut.blogspot.com-2010-08-doprava-v-tampere.html
2020/06/25 20:56:39 Storing Item from URL:http://cvut2tut.blogspot.com/2010/08/mobily.html
2020/06/25 20:56:40 Stored to: /home/user/cvut2tut.blogspot.com/http:--cvut2tut.blogspot.com-2010-08-mobily.html
2020/06/25 20:56:40 Storing Item from URL:http://cvut2tut.blogspot.com/2010/08/jidla.html
2020/06/25 20:56:40 Stored to: /home/user/cvut2tut.blogspot.com/http:--cvut2tut.blogspot.com-2010-08-jidla.html
2020/06/25 20:56:40 Storing Item from URL:http://cvut2tut.blogspot.com/2010/08/ruzne-fotky.html
2020/06/25 20:56:41 Stored to: /home/user/cvut2tut.blogspot.com/http:--cvut2tut.blogspot.com-2010-08-ruzne-fotky.html
2020/06/25 20:56:41 Storing Item from URL:http://cvut2tut.blogspot.com/2010/08/cesta-do-tampere.html
2020/06/25 20:56:41 Stored to: /home/user/cvut2tut.blogspot.com/http:--cvut2tut.blogspot.com-2010-08-cesta-do-tampere.html
2020/06/25 20:56:41 Storing Item from URL:http://cvut2tut.blogspot.com/2010/08/formality-ii.html
2020/06/25 20:56:42 Stored to: /home/user/cvut2tut.blogspot.com/http:--cvut2tut.blogspot.com-2010-08-formality-ii.html
2020/06/25 20:56:42 Storing Item from URL:http://cvut2tut.blogspot.com/2010/08/formality-i.html
2020/06/25 20:56:42 Stored to: /home/user/cvut2tut.blogspot.com/http:--cvut2tut.blogspot.com-2010-08-formality-i.html
2020/06/25 20:56:42 Dump completed.
```

Tip: There is no pagination, consider adding parameter extending amount of records in the feed, e.g. ```?max-results=99```
