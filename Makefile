.PHONY=generate
generate:
	@echo "Hye"
	@docker run --rm -v "$(PWD):/nsq-messages" -w "/nsq-messages" bufbuild/buf generate --template "./configs/grpc.gen.yaml" --path ./common/proto