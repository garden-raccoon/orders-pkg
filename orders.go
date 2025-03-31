package orders

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/garden-raccoon/orders-pkg/models"
	proto "github.com/garden-raccoon/orders-pkg/protocols/orders-pkg"
)

// TODO need to set timeout via lib initialisation
// timeOut is  hardcoded GRPC requests timeout value

// Debug on/off
var Debug = true

type OrdersPkgAPI interface {
	CreateOrUpdateOrder(s *models.Order) (uuid.UUID, error)
	GetOrders() ([]*models.Order, error)
	DeleteOrder(orderUuid uuid.UUID) error
	OrdersByUserUuid(userUuid uuid.UUID) ([]*models.Order, error)
	OrderByOrderUuid(orderUuid uuid.UUID) (*models.Order, error)

	HealthCheck() error
	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	*grpc.ClientConn
	mu sync.Mutex
	proto.OrdersServiceClient
	grpc_health_v1.HealthClient
}

// New create new Battles Api instance
func New(addr string, timeout time.Duration) (OrdersPkgAPI, error) {
	api := &Api{timeout: timeout}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create OrdersApi:  %w", err)
	}
	api.HealthClient = grpc_health_v1.NewHealthClient(api.ClientConn)

	api.OrdersServiceClient = proto.NewOrdersServiceClient(api.ClientConn)
	return api, nil
}

func (api *Api) DeleteOrder(orderUuid uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	req := &proto.OrderDeleteReq{
		OrderUuid: orderUuid.Bytes(),
	}
	_, err := api.OrdersServiceClient.DeleteOrder(ctx, req)
	if err != nil {
		return fmt.Errorf("DeleteOrder api request: %w", err)
	}
	return nil
}

func (api *Api) GetOrders() ([]*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var resp *proto.Orders
	empty := new(proto.OrdersEmpty)
	resp, err := api.OrdersServiceClient.GetOrders(ctx, empty)
	if err != nil {
		return nil, fmt.Errorf("GetOrders api request: %w", err)
	}

	orders, err := models.OrdersFromProto(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to GetOrders %w", err)
	}
	return orders, nil
}

func (api *Api) CreateOrUpdateOrder(s *models.Order) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	resp, err := api.OrdersServiceClient.CreateOrUpdateOrder(ctx, models.Proto(s))
	if err != nil {
		return uuid.Nil, fmt.Errorf("create orders api request: %w", err)
	}
	resp.OrderUuid = s.OrderUuid.Bytes()
	return uuid.FromBytesOrNil(resp.OrderUuid), nil
}

// initConn initialize connection to Grpc servers
func (api *Api) initConn(addr string) (err error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	api.ClientConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}
	return
}

func (api *Api) OrdersByUserUuid(userUuid uuid.UUID) ([]*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	req := &proto.ByUserUuidReq{UserUuid: userUuid.Bytes()}
	resp, err := api.OrdersServiceClient.OrdersByUserUuid(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("OrderAPI OrdersByUserUuid request failed: %w", err)
	}
	orders, err := models.OrdersFromProto(resp)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return orders, nil
}
func (api *Api) OrderByOrderUuid(orderUuid uuid.UUID) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	req := &proto.ByOrderUuidReq{OrderUuid: orderUuid.Bytes()}
	resp, err := api.OrdersServiceClient.OrderByOrderUuid(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("OrderAPI OrderByOrderUuid request failed: %w", err)
	}
	return models.OrderFromProto(resp), nil
}

func (api *Api) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	api.mu.Lock()
	defer api.mu.Unlock()

	resp, err := api.HealthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{Service: "ordersapi"})
	if err != nil {
		return fmt.Errorf("healthcheck error: %w", err)
	}

	if resp.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		return fmt.Errorf("node is %s", errors.New("service is unhealthy"))
	}
	//api.status = service.StatusHealthy
	return nil
}
