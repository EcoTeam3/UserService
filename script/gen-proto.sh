#!/bin/zsh
CURRENT_DIR=$1

# Protobuf fayllarni saqlash uchun `genproto` katalogini o'chirib qayta yaratish
rm -rf ${CURRENT_DIR}/genproto
mkdir -p ${CURRENT_DIR}/genproto

# Protobuf fayllarni generatsiya qilish
for x in $(find ${CURRENT_DIR}/proto -type d); do
  echo "Processing directory: $x"
  protoc -I=${x} -I=${CURRENT_DIR}/proto -I /usr/local/include \
    --go_out=${CURRENT_DIR}/genproto --go-grpc_out=${CURRENT_DIR}/genproto \
    ${x}/*.proto
done

echo "Protobuf generation completed."
