/*
	Copyright 2012-2013 1620469 Ontario Limited.

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package game

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

import (
	"gk/game/config"
	"gk/game/ses"
	"gk/gkerr"
	"gk/gklog"
	"gk/gknet"
	"gk/gkrand"
)

const _tokenRequest = "/gk/tokenServer"

// how many seconds is the token valid
// for now, way too many
const _tokenTimeoutSeconds = 60 * 60

type tokenContextDef struct {
	randContext    *gkrand.GkRandContextDef
	sessionContext *ses.SessionContextDef
	gameConfig     *config.GameConfigDef
	tokenMap       map[string]*tokenEntryDef
	tokenMutex     sync.Mutex
}

type tokenEntryDef struct {
	tokenId     string
	createdDate time.Time
	userName    string
}

func NewTokenContext(gameConfig *config.GameConfigDef, randContext *gkrand.GkRandContextDef, sessionContext *ses.SessionContextDef) *tokenContextDef {
	var tokenContext *tokenContextDef = new(tokenContextDef)

	tokenContext.gameConfig = gameConfig
	tokenContext.randContext = randContext
	tokenContext.sessionContext = sessionContext

	return tokenContext
}

func (tokenContext *tokenContextDef) gameInit() *gkerr.GkErrDef {
	//var gkErr *gkerr.GkErrDef

	tokenContext.tokenMap = make(map[string]*tokenEntryDef)

	return nil
}

func (tokenContext *tokenContextDef) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	gklog.LogTrace(req.Method)
	gklog.LogTrace(path)

	if req.Method == _methodGet || req.Method == _methodPost {
		if gknet.RequestMatches(path, _tokenRequest) {
			tokenContext.handleTokenRequest(res, req)
		} else {
			http.NotFound(res, req)
		}
	} else {
		http.NotFound(res, req)
	}
}

func (tokenContext *tokenContextDef) handleTokenRequest(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	var tokenEntry *tokenEntryDef = new(tokenEntryDef)

	var userName string = req.Form.Get(_userNameParam)

	if len(userName) > 2 {
		tokenEntry.tokenId = tokenContext.getSessionToken()
		tokenEntry.createdDate = time.Now()
		tokenEntry.userName = userName
		tokenContext.tokenMutex.Lock()
		tokenContext.tokenMap[tokenEntry.tokenId] = tokenEntry
		tokenContext.tokenMutex.Unlock()
		gklog.LogTrace(fmt.Sprintf("adding token entry: %+v", tokenEntry))
	} else {
		tokenEntry.tokenId = ""
	}

	res.Write([]byte(tokenEntry.tokenId))
}

func (tokenContext *tokenContextDef) getUserFromToken(token string) string {

	var ok bool
	var tokenEntry *tokenEntryDef

	tokenContext.purgeOldTokenEntries()

	tokenContext.tokenMutex.Lock()
	defer tokenContext.tokenMutex.Unlock()

	gklog.LogTrace(fmt.Sprintf("getting token entry: %+v", token))
	tokenEntry, ok = tokenContext.tokenMap[token]
	if !ok {
		gklog.LogTrace("did not find")
		return ""
	}

	var userName string

	userName = tokenEntry.userName
	gklog.LogTrace("found " + userName)

	// token cannot be reused
	// but for now we allow it to be reused :)
	//delete(tokenContext.tokenMap,tokenEntry.tokenId)

	return userName
}

func (tokenContext *tokenContextDef) purgeOldTokenEntries() {
	tokenContext.tokenMutex.Lock()
	defer tokenContext.tokenMutex.Unlock()

	for tokenId, tokenEntry := range tokenContext.tokenMap {
		if tokenEntry.createdDate.Add(time.Second * _tokenTimeoutSeconds).Before(time.Now()) {
			gklog.LogTrace(fmt.Sprintf("purge token entry: %+v", tokenEntry))
			delete(tokenContext.tokenMap, tokenId)
		}
	}
}

func (tokenContext *tokenContextDef) getSessionToken() string {
	return tokenContext.randContext.GetRandomString(12)
}
