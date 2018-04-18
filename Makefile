workspace:
	@docker run -it --rm --name Golang_Workspace -v ${PWD}:/usr/src/app -w /usr/src/app golang /bin/bash

workspace_80_mapped:
	@docker run -it --rm --name Golang_Workspace -p 80:80 -v ${PWD}:/usr/src/app -w /usr/src/app golang /bin/bash