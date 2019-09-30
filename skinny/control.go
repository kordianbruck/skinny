package skinny

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/danrl/skinny/proto/control"
)

// Status exposes internal state information of an instance
func (in *Instance) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	in.mu.Lock()
	defer in.mu.Unlock()

	status := pb.StatusResponse{
		Name:      in.name,
		Increment: in.increment,
		Timeout:   in.timeout.String(),
		Promised:  in.promised,
		ID:        in.id,
		Holder:    in.holder,
	}

	for _, peer := range in.peers {
		status.Peers = append(status.Peers, &pb.StatusResponse_Peer{
			Name:    peer.name,
			Address: peer.address,
		})
	}

	return &status, nil
}

func (in *Instance) StatusHttp(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	status, _ := in.Status(request.Context(), &pb.StatusRequest{})
	jsonString, _ := json.Marshal(status)
	fmt.Fprintf(w, "%s", string(jsonString))
}
