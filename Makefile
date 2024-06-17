# Go parameters
GOCMD=go
APP=api/app.go
GORUN=$(GOCMD) run $(APP) #for clean-architecture structure
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=api-csm
BINARY_UNIX=$(BINARY_NAME)
BINARY_WIN=$(BINARY_NAME).exe
BINARY_MAC=$(BINARY_NAME).macos
PORT=8080
IP=127.0.0.1
URL=https://$(IP):$(PORT)
EXPOSE=$(IP):$(PORT):$(PORT) 
REPO_NAME=tobias0406/$(BINARY_NAME)
TAG=0.1.17

# muestra el log de github (no añadir | head -n 10 | tail -n 5 | sed 's/^/  /')
log:
	git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative

st:
	git status -sb --untracked-files --ignore-submodules

ust:
	git restore --staged .
	
push:
	git push && git rebase dev cert && git push && git switch dev

pull:
	git pull --autostash --rebase

module:
	cd tickets && mkdir $$(MODULE) && cd $$(MODULE) && mkdir ./infrastructure/ && mkdir ./domain/ ./application/ ./infrastructure/handler/ ./infrastructure/storage/ ./infrastructure/presenter/ && cd ../..

ex:
	export DB_USER=postgres && export DB_PASSWORD=lEy9gfGEqbdYxl1fWcqd && export DB_HOST=127.0.0.1 && export DB_NAME=ticket && export DB_PORT=5432 && export ENABLE_SUBSTITUTION=true
	export DB_USER=postgres && export DB_PASSWORD=lEy9gfGEqbdYxl1fWcqd && export DB_HOST=containers-us-west-62.railway.app && export DB_NAME=ticket && export DB_PORT=5962 && export SECURE=true
	export DB_USER=postgres && export DB_PASSWORD=xdSMD27LjfX2zR502Gds && export DB_HOST=34.176.235.236 && export DB_NAME=ticket && export DB_PORT=80 && export URL=https://api-edg-atlbqiwvnq-tl.a.run.app && export SECURE=true

build:
	CGO_ENABLED=0 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_NAME) -v $(APP)

buildb:
	CGO_ENABLED=0 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o api-block -v block/app.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_UNIX) -v $(APP)

build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_MAC) -v $(APP)

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -trimpath -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -tags netgo -o $(BINARY_WIN) -v $(APP)

test:
	$(GOTEST) -v -shuffle=on -count=1 -race -timeout=10m ./... -coverprofile=coverage.out

mod:
	$(GOMOD) tidy

run-bin:
	./$(BINARY_NAME)

run-binb:
	DB_USER=postgres DB_PASSWORD=lEy9gfGEqbdYxl1fWcqd DB_HOST=$(IP) DB_NAME=vault DB_PORT=5432 BLOB_URL=http://127.0.0.1:9500 ./api-block

#limpia los binarios compilados
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_WIN)

# ejecuta el main.go
run:
	DB_USER=postgres DB_PASSWORD=lEy9gfGEqbdYxl1fWcqd DB_HOST=$(IP) DB_NAME=vault DB_PORT=5432 $(GORUN)

brun:
	DB_USER=postgres DB_PASSWORD=lEy9gfGEqbdYxl1fWcqd DB_HOST=$(IP) DB_NAME=vault DB_PORT=5432 go run block/app.go

#hace el build de la imagen Docker
dbuild:
	DOCKER_BUILDKIT=1 docker build --target dev -t $(REPO_NAME):$(TAG) .

dbuildt:
	DOCKER_BUILDKIT=1 docker build --target test -t $(REPO_NAME):$(TAG) .

dbuildcert:
	DOCKER_BUILDKIT=1 docker build --target cert -t $(REPO_NAME):$(TAG)-cert .

dbuildpre:
	DOCKER_BUILDKIT=1 docker build --target pre -t $(REPO_NAME):$(TAG)-pre .

dbuildprod:
	DOCKER_BUILDKIT=1 docker build --target prod -t $(REPO_NAME):$(TAG)-prod .

# hace el build y el push de la imagen Docker en arquitecturas diferentes
dbuildpush:
#	docker buildx build --platform linux/amd64 -t $(REPO_NAME):$(TAG) --push .
	docker buildx build --platform linux/amd64 -t tobias0406/blob:$(TAG) -f ./Dockerfile.blob --push .
#	docker buildx build --platform ,linux/arm64 -t $(REPO_NAME) -t $(REPO_NAME):$(TAG) --push .

drun:
	docker run -p $(EXPOSE) --name $(BINARY_NAME) $(REPO_NAME):$(TAG)

drund:
	docker run -dp $(EXPOSE) --name $(BINARY_NAME) $(REPO_NAME):$(TAG)

dclean:
	docker rm -f $(BINARY_NAME)
	docker rmi $(REPO_NAME):$(TAG)

dbuildb:
	docker build -t blob:$(TAG) -f ./Dockerfile.blob .

drunb:
	docker run -dp $(IP):9500:9500 --name=blob blob:$(TAG)

dcleans:
	docker rm -f socket
	docker rmi tobias0406/skt-tsm:$(TAG)

dpush:
# docker push $(REPO_NAME):latest
	docker push $(REPO_NAME):$(TAG)

gpush:
	docker build -t "us-central1-docker.pkg.dev/proyecto-egx/dev/edg:latest" .
	docker push "us-central1-docker.pkg.dev/proyecto-egx/dev/edg:latest"

grun:
	docker build -t "us-central1-docker.pkg.dev/proyecto-egx/dev/edg:latest" .
	docker run -dp $(EXPOSE) us-central1-docker.pkg.dev/proyecto-egx/dev/edg:latest

#kubernetes
kube-deploy:
	kubectl apply -f ./$(BINARY_NAME).yaml
	
kube-service:
	kubectl expose deployment $(BINARY_NAME) --port=$(EXPOSE) --type=NodePort

certs:
	chmod +x makecert.sh
	./makecert.sh bgutierrez@datec.com.bo