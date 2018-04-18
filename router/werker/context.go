package werker

import (
	//"encoding/json"
	//"net/http"
	//"time"

	consulet "bitbucket.org/level11consulting/go-til/consul"
	ocelog "bitbucket.org/level11consulting/go-til/log"
	"bitbucket.org/level11consulting/ocelot/build/streamer"
	"bitbucket.org/level11consulting/ocelot/models"
	"bitbucket.org/level11consulting/ocelot/models/pb"
	"bitbucket.org/level11consulting/ocelot/storage"

)
type WerkerContext struct {
	BuildContexts map[string]*models.BuildContext
	*pb.WerkerFacts
	consul        *consulet.Consulet
	store         storage.OcelotStorage
	streamPack    *streamer.StreamPack
}
//
//func (w *WerkerContext) dumpData(wr http.ResponseWriter, r *http.Request) {
//	ocelog.Log().Info("writing out data for buildInfo")
//	wr.Header().Set("content-type", "application/json")
//	dataMap := make(map[string]int)
//	dataMap["time"] = int(time.Now().Unix())
//	wr.WriteHeader(http.StatusOK)
//	for hash, bytearray := range w.buildInfo {
//		dataMap[hash] = len(bytearray.GetData())
//	}
//	bit, err := json.Marshal(dataMap)
//	if err != nil {
//		ocelog.IncludeErrField(err).Error("couldn't marshal for dump")
//		wr.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	wr.Write(bit)
//}

func getWerkerContext(conf *pb.WerkerFacts, store storage.OcelotStorage) *WerkerContext {
	werkerConsul, err := consulet.Default()
	if err != nil {
		ocelog.IncludeErrField(err)
	}
	werkerCtx := &WerkerContext{
		BuildContexts: make(map[string]*models.BuildContext),
		WerkerFacts:   conf,
		consul:        werkerConsul,
		store:         store,
	}
	return werkerCtx
}
