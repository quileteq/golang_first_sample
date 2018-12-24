package main
import (
        "log"
        "net"
        pb "github.com/ewanvalentine/shipper/consignment-service/proto/consignment"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
)
const (
        port = ":50051"
)

type IRepository interface {
      Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
      consigments []*pb.Consignment
}
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
  update := append(repo.coonsignments, consignment)
  repo.consignments = updated
  return consignment, nil

}
type service struct {
        repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error ) {
        consignment, err := s.repo.Create(req)
        if err != nil {
          return nil, err
        }

        return & pb.Response{Created: true, Consignmen: consignment}, nil
}

func main() {
  repo := &Repository{}
  lis, err := net.Listent("top", port)
  if err != nil {
     log.Fatalf("failed to listen %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterShippingServiceServer(s, &service{repo})
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve %v", err)
  }
}
