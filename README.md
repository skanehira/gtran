# gtran
gtran is Google translate CLI tool it written in Go.  
The translate api endpoint maked by Google Apps Script LanguageApp.
Also, this used in vim plugin [translate.vim](https://github.com/skanehira/translate.vim)

# About translate api endpoint
Translate logic.

```js
function doGet(e) {
  var p = e.parameter;
  var translatedText = LanguageApp.translate(p.text, p.source, p.target);
  return ContentService.createTextOutput(translatedText);
}
```

You can set env `GTRAN_ENDPOINT` or use the flag `-endpoint` to set yourself GAS endpoint.  
Env takes precedence.

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
  -endpoint string
        translate endpoint (default "https://script.google.com/macros/s/AKfycbzi15QCo0IsjutiMnI5FYf43
-TKqfrUDiaM03x5C5IcH7-setg/exec?")
  -source string
        translate source (default "en")
  -target string
        translate traget (default "ja")
  -text string
        translate source text

$ gtran -text "海賊王におれはなる"
Become a Pirate King

$ gtran -source=en -target=ja -text="Become a Pirate King"
海賊王になる
```
