package main

import (
	"context"
	"errors"
	pb "github.com/mauamy/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/mauamy/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"log"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all the methods to satisfy the service
// we defined in our protobuf definition.
type consignmentService struct {
	repo         repository
	vesselClient vesselProto.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if vesselResponse == nil {
		return errors.New("error fetching vessel, returned nil")
	}

	if err != nil {
		return err
	}

	// we set the vesselId as the vessel we got back from our vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	service := micro.NewService(
		// this name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	service.Init()

	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo: repo, vesselClient: vesselClient}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
