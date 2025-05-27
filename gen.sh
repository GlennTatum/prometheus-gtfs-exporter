#!/bin/bash

SRC_DIR=./proto/
DST_DIR=.

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/gtfs-realtime-NYCT.proto

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/gtfs-realtime.proto

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/transit.proto