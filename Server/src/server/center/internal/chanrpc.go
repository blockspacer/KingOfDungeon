package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"proto"
	"server/conf"
)

func init() {
	skeleton.RegisterChanRPC("DisconnectFromServer", rpcDisconnectFromServer)
	skeleton.RegisterChanRPC("MapLoaded", rpcMapLoaded)

	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

//客户端连接
func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("client connected %v", a.RemoteAddr())
}

//客户端断开
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("client disconnected %v", a.RemoteAddr())
	mapRole := a.UserData().(*MapRole)
	mapRole.m.ChanRPCServer.Go("RoleDisconnect", mapRole)
}

//====================================================================================
//当和center server断开
func rpcDisconnectFromServer(args []interface{}) {
	a := args[0].(*TCPAgent)
	_ = a

	log.Release("disconnect from center server")
}

//当地图协程加载完毕后通知
func rpcMapLoaded(args []interface{}) {
	mapId := args[0].(uint32)
	log.Debug("map loaded %v", mapId)

	loadedMapCount++
	if loadedMapCount == len(Maps) {

		log.Debug("all map loaded, count = %v", loadedMapCount)

		ctAgent.WriteMsg(&proto.NotifyServerInited{
			LoadedMaps: conf.Server.MapLoad,
		})
	}
}
