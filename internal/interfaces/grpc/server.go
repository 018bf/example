package grpc

import (
	"context"
	"errors"
	"net"

	"github.com/018bf/example/internal/configs"
	"github.com/018bf/example/internal/domain/errs"
	"github.com/018bf/example/pkg/log"
	examplepb "github.com/018bf/example/pkg/examplepb/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server	*grpc.Server
	config	*configs.Config
}

func NewServer(
	logger log.Logger,
	config *configs.Config,
	authMiddleware *AuthMiddleware,
	authHandler examplepb.AuthServiceServer,
	userHandler examplepb.UserServiceServer, sessionHandler examplepb.SessionServiceServer, equipmentHandler examplepb.EquipmentServiceServer, planHandler examplepb.PlanServiceServer, dayHandler examplepb.DayServiceServer, archHandler examplepb.ArchServiceServer,
) *Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
				authMiddleware.UnaryServerInterceptorAuth,
				RequestIDUnaryServerInterceptor,
				grpcZap.UnaryServerInterceptor(
					logger.Logger(),
					grpcZap.WithMessageProducer(DefaultMessageProducer),
				),
			),
		),
	)
	reflection.Register(server)
	healthServer := health.NewServer()
	for service := range server.GetServiceInfo() {
		healthServer.SetServingStatus(service, grpc_health_v1.HealthCheckResponse_SERVING)
	}
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	examplepb.RegisterAuthServiceServer(server, authHandler)
	examplepb.RegisterUserServiceServer(server, userHandler)
	examplepb.RegisterSessionServiceServer(server, sessionHandler)
	examplepb.RegisterEquipmentServiceServer(server, equipmentHandler)
	examplepb.RegisterPlanServiceServer(server, planHandler)
	examplepb.RegisterDayServiceServer(server, dayHandler)
	examplepb.RegisterArchServiceServer(server, archHandler)
	return &Server{
		server:	server,
		config:	config,
	}
}

func (s *Server) Start(_ context.Context) error {
	listener, err := net.Listen("tcp", s.config.BindAddr)
	if err != nil {
		return err
	}
	return s.server.Serve(listener)
}

func (s *Server) Stop(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}

func DefaultMessageProducer(
	ctx context.Context,
	msg string,
	level zapcore.Level,
	code codes.Code,
	err error,
	duration zapcore.Field,
) {
	logger := ctxzap.Extract(ctx)
	params := []zap.Field{
		zap.String("grpc.code", code.String()),
		duration,
		zap.Any("request_id", ctx.Value(log.RequestIDKey)),
	}
	if err != nil {
		sts := status.Convert(err)
		msg = sts.Message()
		for _, v := range sts.Details() {
			errParams := errs.Params{}
			badRequest, ok := v.(*errdetails.BadRequest)
			if ok {
				for _, violation := range badRequest.GetFieldViolations() {
					errParams[violation.GetField()] = violation.GetDescription()
				}
			}
			errorInfo, ok := v.(*errdetails.ErrorInfo)
			if ok {
				for key, value := range errorInfo.GetMetadata() {
					errParams[key] = value
				}
			}
			params = append(params, zap.Object("params", errParams))
		}
	}
	logger.Check(level, msg).Write(params...)
}

func decodeError(err error) error {
	var domainError *errs.Error
	if errors.As(err, &domainError) {
		stat := status.New(codes.Code(domainError.Code), domainError.Message)
		var withDetails *status.Status
		switch domainError.Code {
		case errs.ErrorCodeInvalidArgument:
			d := &errdetails.BadRequest{}
			for key, value := range domainError.Params {
				d.FieldViolations = append(d.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:		key,
					Description:	value,
				})
			}
			withDetails, err = stat.WithDetails(d)
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		default:
			d := &errdetails.ErrorInfo{
				Reason:		domainError.Message,
				Domain:		"",
				Metadata:	domainError.Params,
			}
			withDetails, err = stat.WithDetails(d)
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		}
		return withDetails.Err()
	}
	return status.Error(codes.Internal, err.Error())
}