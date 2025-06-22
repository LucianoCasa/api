package grpcapi

import (
	"context"

	pb "github.com/lucianocasa/api/internal/api/grpcapi/pb"
	"github.com/lucianocasa/api/internal/order"
	"google.golang.org/grpc"
)

type OrderGRPCHandler struct {
	pb.UnimplementedOrderServiceServer
	Service order.Service
}

func Register(server *grpc.Server, service order.Service) {
	pb.RegisterOrderServiceServer(server, &OrderGRPCHandler{Service: service})
}

func (h *OrderGRPCHandler) ListOrders(ctx context.Context, _ *pb.Empty) (*pb.OrderList, error) {
	orders, err := h.Service.ListOrders()
	if err != nil {
		return nil, err
	}
	var result []*pb.Order
	for _, o := range orders {
		result = append(result, &pb.Order{
			Id:           int32(o.ID),
			CustomerName: o.CustomerName,
			Total:        o.Total,
			CreatedAt:    o.CreatedAt,
		})
	}
	return &pb.OrderList{Orders: result}, nil
}
