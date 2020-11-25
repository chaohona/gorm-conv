APPS = gorm-conv
BASE = $(PWD)/src/lib

all:install

install:
	@rm -rf make.log;\
	export GOPATH=$(BASE):$(PWD)\
	&& for ser in $(APPS);\
	do \
		echo "\033[32mmake $$ser start... \033[0m";\
		go install -x $$ser...;\
		if [ "$$?" != "0" ]; then\
			echo "\033[31mmake $$ser failed... \033[0m";\
            exit 1;\
        else\
        	echo "\033[32mmake $$ser success... \033[0m";\
        fi;\
	done
	cp -rf bin/gorm-conv /root/gorm/tools/bin/
	cp -rf bin/gorm-conv ./src/gorm-conv/test/gorm_release/bin
	cp -rf bin/gorm-conv /root/gorm/cpp/bin
	cp -rf bin/gorm-conv /root/gorm_cpp/bin
	cp -rf bin/gorm-conv /root/hongchao/gorm/tools/bin
	cp -rf bin/gorm-conv /root/daobatu/gorm/tools/bin
	cp -rf bin/gorm-conv /root/github.com/gorm_make/tools/bin
	cp -rf bin/gorm-conv /root/database/gorm_client_cpp/bin
	cp -rf bin/gorm-conv /root/database/gorm/tools/bin
	cp -rf bin/gorm-conv /root/github.com/async/gorm/tools/bin

	
clean:
	export GOPATH=$(PWD)\
	&& go clean -i ./...\
	&& go clean -i ./src/lib\
	
gz:
	@rm -rf release
	@mkdir -pv release/
	@mkdir -pv release/bin
	@cp -avf run.sh.example release/
	@cp -avf bin/* release/bin
	@svn info>release/ver.txt
	@tar zcvf `date +%Y%m%d%H%M`.tar.gz release/
