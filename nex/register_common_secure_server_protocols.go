package nex

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	commonmatchmaking "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	commonmatchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	commonmatchmakeextension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	commonnattraversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	commonrankinglegacy "github.com/PretendoNetwork/nex-protocols-common-go/v2/ranking-legacy"
	commonsecure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	commonstoragemanager "github.com/PretendoNetwork/nex-protocols-common-go/v2/storage-manager"
	matchmaking "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	matchmakeextension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	nattraversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	rankinglegacy "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/legacy"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	storagemanager "github.com/PretendoNetwork/nex-protocols-go/v2/storage-manager"
	"github.com/PretendoNetwork/puyo-puyo-20th/globals"
)

func CreateReportDBRecord(_ types.PID, _ types.UInt32, _ types.QBuffer) error {
	return nil
}

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	commonSecureProtocol := commonsecure.NewCommonProtocol(secureProtocol)
	commonSecureProtocol.EnableInsecureRegister()
	commonSecureProtocol.CreateReportDBRecord = CreateReportDBRecord

	natTraversalProtocol := nattraversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	commonnattraversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := matchmaking.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	commonMatchMakingProtocol := commonmatchmaking.NewCommonProtocol(matchMakingProtocol)
	commonMatchMakingProtocol.SetManager(globals.MatchmakingManager)

	matchMakingExtProtocol := matchmakingext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol := commonmatchmakingext.NewCommonProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol.SetManager(globals.MatchmakingManager)

	matchmakeExtensionProtocol := matchmakeextension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := commonmatchmakeextension.NewCommonProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol.SetManager(globals.MatchmakingManager)

	rankingProtocol := rankinglegacy.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(rankingProtocol)
	commonRankingProtocol := commonrankinglegacy.NewCommonProtocol(rankingProtocol)
	commonRankingProtocol.SetManager(globals.RankingManager)

	storageManagerProtocol := storagemanager.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(storageManagerProtocol)
	commonStorageManagerProtocol := commonstoragemanager.NewCommonProtocol(storageManagerProtocol)
	commonStorageManagerProtocol.SetManager(globals.StorageManagerManager)
}
