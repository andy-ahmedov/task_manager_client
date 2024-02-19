all: run


run:
	sudo docker-compose run client sh

build_up:
	sudo docker-compose up --build client 

up:
	sudo docker-compose up client 
# By default, the make up command starts the client 
# and calls the getall command. To call another command,
# change the command CMD[""] in the dockerfile

local_run: 
	go run client/cmd/main.go

client_run: 
ifeq ($(CMD),update)
	go run cmd/main.go $(CMD) $(ID) --name $(NAME) --description $(DESCRIPTION) --status $(STATUS)
else
	go run cmd/main.go $(CMD) $(filter-out $@,$(MAKECMDGOALS))
endif

# Exampe:
#	-$> make client_run create "Name" "Description" "Status"

# Example for using update command: 
#	-$> CMD=update ID=7 NAME="SomeName" DESCRIPTION='SomeDescriptio' STATUS='SomeStatus' make client_run

#	Important! When using the update command via make, 
#	only 1 word can be passed in parameters. If you want 
#	to use a command and pass several words to a specific 
#	field, then use without make

#	-$> go run client/cmd/main.go update 7 --status "Some Status"
