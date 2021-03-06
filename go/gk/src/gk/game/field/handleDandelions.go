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

package field

import (
	"fmt"
	"math/rand"
)

import (
	"gk/game/message"
	//	"gk/game/ses"
	"gk/gkerr"
)

func (fieldContext *FieldContextDef) handleDandelions() *gkerr.GkErrDef {
	if fieldContext.rainContext.rainCurrentlyOn {
		if rand.Int31n(5) == 10 {
			fieldContext.addDandelion()
		}
	} else {
		if rand.Int31n(4) == 10 {
			fieldContext.removeDandelion()
		}
	}

	return nil
}

func (fieldContext *FieldContextDef) addDandelion() {

	var messageToClient *message.MessageToClientDef = new(message.MessageToClientDef)
	var svgJsonData *message.SvgJsonDataDef = new(message.SvgJsonDataDef)
	var fileName = "dandelion"

	svgJsonData.Id = fieldContext.getNextObjectId()
	//	svgJsonData.IsoXYZ.X = int16(rand.Int31n(50))
	//	svgJsonData.IsoXYZ.Y = int16(rand.Int31n(50))

	var podId int32 = firstPodId // dandelions are only in the first pod

	for _, websocketConnectionContext := range fieldContext.podMap[podId].websocketConnectionMap {
		//		var singleSession *ses.SingleSessionDef
		//		singleSession = fieldContext.sessionContext.GetSessionFromId(websocketConnectionContext.sessionId)
		var terrainJson *terrainJsonDef
		terrainJson = fieldContext.podMap[podId].terrainJson

		var index = rand.Int31n(int32(len(terrainJson.jsonMapData.TileList)))

		if terrainJson.jsonMapData.TileList[index].Terrain == "grass" {

			svgJsonData.IsoXYZ.X = int16(terrainJson.jsonMapData.TileList[index].X)
			svgJsonData.IsoXYZ.Y = int16(terrainJson.jsonMapData.TileList[index].Y)
			//svgJsonData.IsoXYZ.Z = int16(terrainJson.jsonMapData.TileList[index].Z)
			svgJsonData.IsoXYZ.Z = 0

			messageToClient.BuildSvgMessageToClient(fieldContext.terrainSvgDir, message.AddSvgReq, fileName, svgJsonData)

			fieldContext.queueMessageToClient(websocketConnectionContext.sessionId, messageToClient)

			var fieldObject *fieldObjectDef = new(fieldObjectDef)
			fieldObject.id = svgJsonData.Id
			fieldObject.fileName = fileName
			fieldObject.isoXYZ = svgJsonData.IsoXYZ
			fieldContext.addTerrainObject(fieldObject, podId)
		}
	}
}

func (fieldContext *FieldContextDef) removeDandelion() {

	var messageToClient *message.MessageToClientDef = new(message.MessageToClientDef)
	var fileName = "dandelion"

	messageToClient.Command = message.DelSvgReq

	for podId, podEntry := range fieldContext.podMap {
		for _, fieldObject := range podEntry.objectMap {
			if fieldObject.fileName == fileName {
				messageToClient.JsonData = []byte(fmt.Sprintf("{ \"id\": \"%s\"}", fieldObject.id))
				for _, websocketConnectionContext := range podEntry.websocketConnectionMap {
					fieldContext.queueMessageToClient(websocketConnectionContext.sessionId, messageToClient)
				}
				fieldContext.delTerrainObject(podId, fieldObject)
				break
			}
		}
	}
}
