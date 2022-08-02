// This file is part of MinIO DirectPV
// Copyright (c) 2021, 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package metrics

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"k8s.io/klog/v2"
)

// ServeMetrics starts metrics service.
func ServeMetrics(ctx context.Context, nodeID string, port int) {
	server := &http.Server{
		Handler: metricsHandler(nodeID),
	}

	lc := net.ListenConfig{}
	listener, lErr := lc.Listen(ctx, "tcp", fmt.Sprintf(":%v", port))
	if lErr != nil {
		panic(lErr)
	}

	klog.V(2).Infof("Starting metrics exporter in port: %d", port)
	if err := server.Serve(listener); err != nil {
		klog.Errorf("Failed to listen and serve metrics server: %v", err)
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
