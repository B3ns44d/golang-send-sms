build:
	docker build -t sendsms .

run:
	docker run --name sendsms -p 3006:3006 sendsms:latest