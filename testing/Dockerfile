FROM python:3.11.5-slim-bookworm@sha256:edaf703dce209d774af3ff768fc92b1e3b60261e7602126276f9ceb0e3a96874 as python-base
RUN rm -f /etc/apt/apt.conf.d/docker-clean \
    && echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' | tee /etc/apt/apt.conf.d/keep-cache
ARG PYTHON_REQUIREMENTS_FILE=aiohttp/requirements.txt
ADD "$PYTHON_REQUIREMENTS_FILE" /tmp/requirements.txt


FROM python-base as python-grpc-client
WORKDIR /client
RUN pip install --require-hashes -qr /tmp/requirements.txt
# Copy the sources, including the stubs
COPY --chmod=777 client.py /client/grpc-kv-client.py
COPY kv /client/kv
CMD tail -f /dev/null


FROM python-base as aiohttp-service
ARG SERVICE_PORT=8080
ENV DEBIAN_FRONTEND=noninteractive \
    SERVICE_PORT=$SERVICE_PORT
ADD "$PYTHON_REQUIREMENTS_FILE" /tmp/requirements.txt
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt/lists,sharing=locked \
    pip3 install --require-hashes -qr /tmp/requirements.txt \
    && apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends netcat-traditional \
    && mkdir /code
HEALTHCHECK \
    --interval=1s \
    --timeout=1s \
    --start-period=1s \
    --retries=3 \
    CMD nc -zv localhost "$SERVICE_PORT"
ENTRYPOINT ["python3", "/code/service.py"]


FROM aiohttp-service as aiohttp-hello-service
ADD service.py /code/service.py
