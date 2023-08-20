#!/bin/bash
# requires https://github.com/fullstorydev/grpcurl

grpcurl -plaintext 192.168.100.1:9200 list

# dish grpc
grpcurl -plaintext -protoset-out protoset/dish.protoset 192.168.100.1:9200 describe SpaceX.API.Device.Device

# mesh
#grpcurl -plaintext -protoset-out mesh.protoset 192.168.100.1:9200 describe SpaceX.API.Device.Mesh

# router
#grpcurl -plaintext -protoset-out router.protoset 192.168.1.1:9000 describe SpaceX.API.Device.Device
