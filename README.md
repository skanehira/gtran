# gtran
gtran is Google translate CLI tool it written in Go.
The translate api endpoint maked by Google Apps Script LanguageApp.

# About translate api endpoint
Translate logic.

```js
function doGet(e) {
  var p = e.parameter;
  var translatedText = LanguageApp.translate(p.text, p.source, p.target);
  return ContentService.createTextOutput(translatedText);
}
```

# Requirement
- Go 1.12 or higher

# Install
```sh
$ git clone https://github.com/skanehira/gtran.git
$ cd gtran
$ go install
```

# Usage
```sh
$ gtran -h
Usage of gtran:
  -source string
        translate source (default "ja")
  -target string
        translate traget (default "en")
  -text string
        translate source text

$ gtran -text "海賊王におれはなる"
Become a Pirate King

$ gtran -source=en -target=ja -text="Become a Pirate King"
海賊王になる
```
