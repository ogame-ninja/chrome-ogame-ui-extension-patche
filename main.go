package main

import (
	ep "github.com/ogame-ninja/extension-patcher"
)

func main() {
	const (
		webstoreURL               = "https://chromewebstore.google.com/detail/ogame-ui++/nhbgpipnadhelnecpcjcikbnedilhddf"
		ui_2024_7_31_17002_sha256 = "c47eced2b8eea6891f547141706af5387bd25a7013a001a0d56b7c2c9d32e4c2"
	)

	files := []ep.FileAndProcessors{
		ep.NewFile("manifest.json", processManifest),
		ep.NewFile("src/utils/load-universe-api.js", srcUtilsLoadUniverseApiJs),
		ep.NewFile("_locales/de/messages.json", messagesJson),
		ep.NewFile("_locales/en/messages.json", messagesJson),
		ep.NewFile("_locales/es/messages.json", messagesJson),
		ep.NewFile("_locales/fr/messages.json", messagesJson),
		ep.NewFile("_locales/pl/messages.json", messagesJson),
		ep.NewFile("_locales/tr/messages.json", messagesJson),
	}

	ep.MustNew(ep.Params{
		ExpectedSha256: ui_2024_7_31_17002_sha256,
		WebstoreURL:    webstoreURL,
		Files:          files,
	}).Start()
}

var replN = ep.MustReplaceN

func processManifest(by []byte) []byte {
	by = replN(by, `"*://*.ogame.gameforge.com/game/*"`,
		`"*://*.ogame.gameforge.com/game/*",
      "http://127.0.0.1:*/bots/*/browser/html*",
      "https://*.ogame.ninja/bots/*/browser/html*"`, 1)
	return by
}

func messagesJson(by []byte) []byte {
	return replN(by, `"message": "OGame UI++",`, `"message": "OGame UI++ Ninja",`, 1)
}

func srcUtilsLoadUniverseApiJs(by []byte) []byte {
	by = replN(by, `_loadUniverseApi(cb) {`,
		`_loadUniverseApi(cb) {
	var l = new URL(document.location.href);
    var tmp = l.pathname.split('/').slice(-1)[0].split('-');
    var server = tmp[0];
    var lang = tmp[1];
`, 1)
	by = replN(by, `/api/players.xml`, `/api/'+server+'/'+lang+'/players.xml`, 1)
	by = replN(by, `/api/universe.xml`, `/api/'+server+'/'+lang+'/universe.xml`, 1)
	by = replN(by, `/api/highscore.xml?category=1&type=1`, `/api/'+server+'/'+lang+'/highscore.xml?category=1&type=1`, 1)
	by = replN(by, `/api/highscore.xml?category=1&type=0`, `/api/'+server+'/'+lang+'/highscore.xml?category=1&type=0`, 1)
	by = replN(by, `/api/highscore.xml?category=1&type=3`, `/api/'+server+'/'+lang+'/highscore.xml?category=1&type=3`, 1)
	by = replN(by, `/api/highscore.xml?category=1&type=2`, `/api/'+server+'/'+lang+'/highscore.xml?category=1&type=2`, 1)
	by = replN(by, `/api/highscore.xml?category=1&type=7`, `/api/'+server+'/'+lang+'/highscore.xml?category=1&type=7`, 1)
	by = replN(by, `/api/alliances.xml`, `/api/'+server+'/'+lang+'/alliances.xml`, 1)
	by = replN(by, `/api/serverData.xml`, `/api/'+server+'/'+lang+'/serverData.xml`, 1)
	by = replN(by, `/api/localization.xml`, `/api/'+server+'/'+lang+'/localization.xml`, 1)
	return by
}
