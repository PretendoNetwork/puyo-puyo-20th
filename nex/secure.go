package nex

import (
	"os"
	"strconv"

	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/puyo-puyo-20th/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()
	globals.SecureServer.ByteStreamSettings.UseStructureHeader = false

	globals.SecureEndpoint = nex.NewPRUDPEndPoint(1)
	globals.SecureEndpoint.IsSecureEndPoint = true
	globals.SecureEndpoint.ServerAccount = globals.SecureServerAccount
	globals.SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)

	globals.SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(2, 3, 2))
	globals.SecureServer.AccessKey = "9ccfebcf"

	globals.SecureEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()
		globals.Logger.Infof("PP20 Secure - Protocol %d, Method %d", request.ProtocolID, request.MethodID)
	})

	globals.SecureEndpoint.OnError(func(err *nex.Error) {
		globals.Logger.Errorf("Secure: %v", err)
	})

	globals.MatchmakingManager = common_globals.NewMatchmakingManager(globals.SecureEndpoint, globals.Postgres)
	globals.MatchmakingManager.GetUserFriendPIDs = globals.GetUserFriendPIDs

	globals.StorageManagerManager = common_globals.NewStorageManagerManager(globals.SecureEndpoint, globals.Postgres)
	globals.RankingManager = common_globals.NewRankingManager(globals.SecureEndpoint, globals.Postgres)
	globals.RankingManager.GetUserFriendPIDs = globals.GetUserFriendPIDs

	registerCommonSecureServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_PP20_SECURE_SERVER_PORT"))

	globals.SecureServer.Listen(port)
}
