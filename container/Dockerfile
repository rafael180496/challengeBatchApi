#Ambiente para el api coreauth 
FROM rafael180496/gocore:v1
LABEL maintainer="rafael180496@gmail.com"

ENV WORKAPI $GOPATH/src/challengeapi

RUN mkdir $WORKAPI
RUN mkdir $WORKAPI/config
RUN mkdir $WORKAPI/config/log
RUN mkdir $WORKAPI/docs
COPY dist/ $WORKAPI
RUN mv $WORKAPI/challenge.ini $WORKAPI/config/challenge.ini
WORKDIR $WORKAPI
ENTRYPOINT $WORKAPI/apichallenge -mode="api"