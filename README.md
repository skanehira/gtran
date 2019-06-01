# gtran
gtran is Google translate CLI tool it written in Go.  
this using the translate api endpoint maked by Google Apps Script LanguageApp.  
Also, this used in vim plugin [translate.vim](https://github.com/skanehira/translate.vim)

# About translate api endpoint
Translate logic.

```js
function doPost(e) {
  var p = JSON.parse(e.postData.getDataAsString());
  if (p.text == "") {
    return ContentService.createTextOutput("text is empty");
  }
  if (p.source == "") {
    return ContentService.createTextOutput("source is empty");
  }
  if (p.target == "") {
    return ContentService.createTextOutput("target is empty");
  }
  
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
$ go get github.com/skanehira/gtran
```

# Usage
The language code is bellow.  
https://cloud.google.com/translate/docs/languages

```sh
Usage of gtran:
  -endpoint string
        translate endpoint (default "https://script.google.com/macros/s/AKfycbywwDmlmQrNPYoxL90NCZYjoEzuzRcnRuUmFCPzEqG7VdWBAhU/exec")
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
