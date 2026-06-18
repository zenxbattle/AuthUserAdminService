package main

import (
	"net"
	"time"

	"xcode/cache"
	"xcode/configs"
	"xcode/db"
	"xcode/logutil"
	"xcode/repository"
	"xcode/service"

	authUserAdminProto "github.com/lijuuu/GlobalProtoXcode/AuthUserAdminService"
	"go.uber.org/zap"
	"go.uber.org/zapcore"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func mustConnect(fn func() error, name string, logShipper *logutil.LogShipper, logger *zap.Logger) {
	for {
		err := fn()
		if err == nil {
			logShipper.Log(zapcore.InfoLevel, "GENESISTRACEID", "Connected to "+name, nil, "INIT", nil)
			return
		}
		logShipper.Log(zapcore.WarnLevel, "GENESISTRACEID", "Retrying connection to "+name, map[string]any{"error": err.Error()}, "INIT", nil)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	// Load configuration
	config := configs.LoadConfig()

	// Initialize Zap logger based on environment
	var logger *zap.Logger
	var err error
	if config.Environment == "development" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic("Failed to initialize Zap logger: " + err.Error())
	}
	defer logger.Sync()

	// Initialize LokiLogShipper
	logShipper := logutil.New("auth-user-admin-service")

	// Initialize PostgreSQL connection with retry
	var dbConn *gorm.DB
	mustConnect(func() error {
		var dbErr error
		dbConn, dbErr = db.InitDB(config.PostgresDSN)
		return dbErr
	}, "PostgreSQL", logShipper, logger)
	defer db.Close(dbConn)

	// Initialize Redis cache
	redisCache := cache.NewRedisCache(config.RedisURL, "", 0)

	// Initialize repository and service
	userRepo := repository.NewUserRepository(dbConn, &config, logShipper)
	authUserAdminService := service.NewAuthUserAdminService(userRepo, *redisCache, &config, config.JWTSecretKey, logShipper)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":"+config.UserGRPCPort)
	if err != nil {
		logShipper.Log(zapcore.ErrorLevel, "GENESISTRACEID", "Failed to listen on port", map[string]any{
			"port":  config.UserGRPCPort,
			"error": err.Error(),
		}, "GRPC INIT", nil)
		// logger.Fatal("Failed to listen on port", zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	authUserAdminProto.RegisterAuthUserAdminServiceServer(grpcServer, authUserAdminService)

	// Log server startup
	logShipper.Log(zapcore.InfoLevel, "GENESISTRACEID", "AuthUserAdminService gRPC server running", map[string]any{
		"port": config.UserGRPCPort,
	}, "SERVICE INIT", nil)

	// Start gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		logShipper.Log(zapcore.ErrorLevel, "GENESISTRACEID", "Failed to serve gRPC server", map[string]any{
			"error": err.Error(),
		}, "GRPC SERVE", nil)
		// logger.Fatal("Failed to serve gRPC server", zap.Error(err))
	}
}
