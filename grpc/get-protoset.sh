#!/bin/bash
# requires https://github.com/fullstorydev/grpcurl

grpcurl -plaintext 192.168.100.1:9200 list
grpcurl -plaintext -d {\"get_status\":{}} 192.168.100.1:9200 SpaceX.API.Device.Device/Handle | grep softwareVersion | awk '{print $2}'

# dish grpc
grpcurl -plaintext -protoset-out dish.protoset 192.168.100.1:9200 describe SpaceX.API.Device.Device

# mesh
grpcurl -plaintext -protoset-out mesh.protoset 192.168.100.1:9200 describe SpaceX.API.Device.Mesh

# router
grpcurl -plaintext -protoset-out router.protoset 192.168.1.1:9000 describe SpaceX.API.Device.Device
