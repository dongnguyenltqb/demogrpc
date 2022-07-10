FROM ubuntu
WORKDIR /app
COPY demogrpc /app/demogrpc
CMD /app/demogrpc